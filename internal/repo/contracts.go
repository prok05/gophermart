package repo

import (
	"context"
	"github.com/prok05/gophermart/internal/entity"
)

type (
	UserRepo interface {
		Create(ctx context.Context, user entity.User) error
		GetByLogin(ctx context.Context, login string) (*entity.User, error)
		GetByID(ctx context.Context, userID string) (*entity.User, error)

		GetOrders(ctx context.Context, userID string) (*[]entity.UserOrder, error)
		CreateOrder(ctx context.Context, userID string, orderNumber string) error

		GetBalance(ctx context.Context, userID string) (*entity.UserBalance, error)

		GetWithdrawals(ctx context.Context, userID string) (*[]entity.UserWithdrawal, error)
	}
)
