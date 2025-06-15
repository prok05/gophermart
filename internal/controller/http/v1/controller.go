package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/internal/usecase"
	"github.com/prok05/gophermart/pkg/logger"
)

type V1 struct {
	cfg config.Config
	u   usecase.User
	l   logger.Interface
	v   *validator.Validate
}
