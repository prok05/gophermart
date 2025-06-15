package user

import (
	"context"
	"fmt"
	"github.com/prok05/gophermart/config"
	"github.com/prok05/gophermart/internal/controller/http/request"
	"github.com/prok05/gophermart/internal/entity"
	"github.com/prok05/gophermart/internal/repo"
	"github.com/prok05/gophermart/pkg/jwt"
	"time"
)

type UseCase struct {
	cfg  *config.Config
	repo repo.UserRepo
}

func New(r repo.UserRepo, cfg *config.Config) *UseCase {
	return &UseCase{repo: r, cfg: cfg}
}

func (uc *UseCase) Register(ctx context.Context, input request.RegisterUser) error {
	u := entity.User{
		Login: input.Login,
	}

	// password set
	if err := u.Password.Set(input.Password); err != nil {
		return fmt.Errorf("UserUseCase - Register - u.Password.Set: %w", err)
	}

	// storing
	if err := uc.repo.Create(ctx, u); err != nil {
		return fmt.Errorf("UserUseCase - Register - uc.repo.Create: %w", err)
	}

	return nil
}

func (uc *UseCase) Login(ctx context.Context, input request.LoginUser) (string, error) {
	// получить пользователя по логину
	u, err := uc.repo.GetByLogin(ctx, input.Login)
	if err != nil {
		return "", fmt.Errorf("UserUseCase - Login - uc.repo.GetByLogin: %w", err)
	}

	// провалидировать пароль
	if err := u.Password.Compare(input.Password); err != nil {
		return "", fmt.Errorf("UserUseCase - Login - u.Password.Compare: %w", entity.ErrWrongLoginOrPassword)
	}

	// сформировать токен
	claims := jwt.GenerateJWTClaims(
		u.ID,
		uc.cfg.App.Name,
		uc.cfg.App.Name,
		time.Duration(uc.cfg.JWT.ExpDays),
	)

	token, err := jwt.GenerateToken(claims, uc.cfg.JWT.Secret)
	if err != nil {
		return "", fmt.Errorf("UserUseCase - Login - jwt.GenerateToken: %w", err)
	}

	return token, nil
}

func (uc *UseCase) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	u, err := uc.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetByID - uc.repo.GetByID: %w", err)
	}

	return u, nil
}

func (uc *UseCase) LoadOrder(ctx context.Context, orderNumber string) error {
	userID := ctx.Value(entity.ContextUserID).(string)

	if !uc.ValidateOrderNumber(orderNumber) {
		return fmt.Errorf("UserUseCase - LoadOrder - uc.ValidateOrderNumber: %w", entity.ErrInvalidOrderNumber)
	}

	if err := uc.repo.CreateOrder(ctx, userID, orderNumber); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetOrders(ctx context.Context) (*[]entity.UserOrder, error) {
	userID := ctx.Value(entity.ContextUserID).(string)
	orders, err := uc.repo.GetOrders(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetOrders - uc.repo.GetOrders: %w", err)
	}

	if len(*orders) == 0 {
		return nil, fmt.Errorf("UserUseCase - GetOrders - len(*orders) == 0: %w", entity.ErrNoContent)
	}

	return orders, nil
}

func (uc *UseCase) GetBalance(ctx context.Context, userID string) (*entity.UserBalance, error) {
	balance, err := uc.repo.GetBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetBalance - uc.repo.GetBalance: %w", err)
	}

	return balance, nil
}

func (uc *UseCase) WithdrawBalance(ctx context.Context, request request.WithdrawBalance) error {
	userID := ctx.Value(entity.ContextUserID).(string)

	if !uc.ValidateOrderNumber(request.Order) {
		return fmt.Errorf("UserUseCase - WithdrawBalance - uc.ValidateOrderNumber: %w", entity.ErrInvalidOrderNumber)
	}

	balance, err := uc.repo.GetBalance(ctx, userID)
	if err != nil {
		return fmt.Errorf("UserUseCase - WithdrawBalance - uc.repo.GetBalance: %w", err)
	}

	if balance.Current < request.Sum {
		return fmt.Errorf("UserUseCase - WithdrawBalance - balance.Current < request.Sum: %w", entity.ErrNotEnoughBalance)
	}

	// another service

	return nil
}

func (uc *UseCase) GetWithdrawals(ctx context.Context) (*[]entity.UserWithdrawal, error) {
	userID := ctx.Value(entity.ContextUserID).(string)

	withdrawals, err := uc.repo.GetWithdrawals(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetWithdrawals - uc.repo.GetWithdrawals: %w", err)
	}

	if len(*withdrawals) == 0 {
		return nil, fmt.Errorf("UserUseCase - GetOrders - len(*withdrawals) == 0: %w", entity.ErrNoContent)
	}

	return withdrawals, nil
}
