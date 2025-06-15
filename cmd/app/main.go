package main

import (
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
