package usecase

import (
	"context"
	"github.com/prok05/gophermart/internal/controller/http/request"
	"github.com/prok05/gophermart/internal/entity"
)

type (
	User interface {
		Register(ctx context.Context, input request.RegisterUser) error
		Login(ctx context.Context, input request.LoginUser) (string, error)
		GetByID(ctx context.Context, userID string) (*entity.User, error)

		LoadOrder(ctx context.Context, orderNumber string) error
		GetOrders(ctx context.Context) (*[]entity.UserOrder, error)

		GetBalance(ctx context.Context, userID string) (*entity.UserBalance, error)
		WithdrawBalance(ctx context.Context, input request.WithdrawBalance) error

		GetWithdrawals(ctx context.Context) (*[]entity.UserWithdrawal, error)
	}
)
