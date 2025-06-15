package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/internal/usecase"
	"github.com/prok05/gophermart/pkg/logger"
)

func NewUserRoutes(cfg config.Config, r chi.Router, u usecase.User, l logger.Interface) {
	h := &V1{
		cfg: cfg,
		u:   u,
		l:   l,
		v:   validator.New(validator.WithRequiredStructEnabled()),
	}

	r.Route("/user", func(r chi.Router) {
		r.Post("/login", h.login)
		r.Post("/register", h.register)

		r.Route("/orders", func(r chi.Router) {
			r.Use(h.AuthTokenMiddleware)
			r.Post("/", h.loadOrder)
			r.Get("/", h.getUserOrders)
		})

		r.Route("/balance", func(r chi.Router) {
			r.Use(h.AuthTokenMiddleware)
			r.Get("/", h.getUserBalance)
			r.Post("/withdraw", h.withdrawUserBalance)
		})

		r.Route("/withdrawals", func(r chi.Router) {
			r.Use(h.AuthTokenMiddleware)
			r.Get("/", h.getUserWithdrawals)
		})
	})
}
