package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/docs"
	_ "github.com/prok05/gophermart/docs" // Swagger docs.
	v1 "github.com/prok05/gophermart/internal/controller/http/v1"
	"github.com/prok05/gophermart/internal/usecase"
	"github.com/prok05/gophermart/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

// NewRouter -.
// Swagger spec:
//
//	@title						Gophermart
//	@description				Service for orders
//	@version					1.0
//	@host						localhost:80
//	@BasePath					/api
//
//	@securityDefinitions.apiKey	AuthToken
//	@in							header
//	@name						Authorization
func NewRouter(cfg *config.Config, u usecase.User, l logger.Interface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(":80/swagger/doc.json")))

	docs.SwaggerInfo.Version = cfg.App.Version

	// main routes
	r.Route("/api", func(r chi.Router) {
		{
			v1.NewUserRoutes(*cfg, r, u, l)
		}
	})

	return r
}
