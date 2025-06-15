package app

import (
	"context"
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/internal/controller/http"
	"github.com/prok05/gophermart/internal/repo/persistent"
	"github.com/prok05/gophermart/internal/usecase/user"
	"github.com/prok05/gophermart/pkg/httpserver"
	"github.com/prok05/gophermart/pkg/logger"
	"github.com/prok05/gophermart/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	// logger
	l, err := logger.New(cfg.Log.Level, cfg.App.ENV)
	if err != nil {
		log.Fatalf("logger error: %s", err)
	}

	// db
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		l.Fatal("app - Run - postgres.New:", err)
	}

	if err := pg.Pool.Ping(context.Background()); err != nil {
		l.Fatal("app - Run - postgres.Ping:", err)
	}

	defer pg.Close()

	// usecase
	userUsecase := user.New(
		persistent.New(pg),
		cfg,
	)

	// HTTP Server
	router := http.NewRouter(cfg, userUsecase, l)
	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	httpServer.Start()
	l.Info("server started at",
		"port", cfg.HTTP.Port,
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run",
			"signal", s.String(),
		)
	case err = <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify",
			"err", err,
		)
	}

	err = httpServer.ShutDown(context.Background())
	if err != nil {
		l.Error("app - Run",
			"httpServer.Shutdown", err,
		)
	}
}
