package persistent

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/prok05/gophermart/internal/entity"
	"github.com/prok05/gophermart/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) Create(ctx context.Context, u entity.User) (err error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Pool.Begin: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `
		INSERT INTO "user" (login, password) VALUES ($1, $2) RETURNING id
	`

	if err := tx.QueryRow(ctx, query, u.Login, u.Password.Hash).Scan(&u.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return fmt.Errorf("UserRepo - Create - tx.QueryRow: %w", entity.ErrDuplicateLogin)
			default:
				return fmt.Errorf("UserRepo - Create - tx.QueryRow: %w", err)
			}
		}
	}

	balanceQuery := `
		INSERT INTO user_balance (user_id) VALUES ($1)
	`

	_, err = tx.Exec(ctx, balanceQuery, u.ID)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - tx.Exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("UserRepo - Create - tx.Commit: %w", err)
	}

	return nil
}

func (r *UserRepo) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	query := `
		SELECT id, login, password, created_at 
		FROM "user" 
		WHERE login = $1
	`

	u := &entity.User{}

	if err := r.Pool.QueryRow(ctx, query, login).Scan(&u.ID, &u.Login, &u.Password.Hash, &u.CreatedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case errors.Is(pgErr, pgx.ErrNoRows):
				return nil, fmt.Errorf("UserRepo - GetByLogin - r.Pool.QueryRow: %w", entity.ErrNotFound)
			default:
				return nil, fmt.Errorf("UserRepo - GetByLogin - r.Pool.QueryRow: %w", err)
			}
		}
	}

	return u, nil
}

func (r *UserRepo) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	query := `
		SELECT id, login, password, created_at 
		FROM "user" 
		WHERE id = $1
	`

	u := &entity.User{}

	if err := r.Pool.QueryRow(ctx, query, userID).Scan(&u.ID, &u.Login, &u.Password.Hash, &u.CreatedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case errors.Is(pgErr, pgx.ErrNoRows):
				return nil, fmt.Errorf("UserRepo - GetByID - r.Pool.QueryRow: %w", entity.ErrNotFound)
			default:
				return nil, fmt.Errorf("UserRepo - GetByID - r.Pool.QueryRow: %w", err)
			}
		}
	}

	return u, nil
}

func (r *UserRepo) GetOrders(ctx context.Context, userID string) (*[]entity.UserOrder, error) {
	query := `
		SELECT number, status, accrual, uploaded_at 
		FROM user_order
		WHERE user_id = $1
		ORDER BY uploaded_at DESC
	`

	rows, err := r.Pool.Query(ctx, query, userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case errors.Is(pgErr, pgx.ErrNoRows):
				return nil, fmt.Errorf("UserRepo - GetOrders - r.Pool.QueryRow: %w", entity.ErrNoContent)
			default:
				return nil, fmt.Errorf("UserRepo - GetOrders - r.Pool.QueryRow: %w", err)
			}
		}
	}
	defer rows.Close()

	orders := make([]entity.UserOrder, 0)

	for rows.Next() {
		var order entity.UserOrder

		err = rows.Scan(&order.Number, &order.Status, &order.Accrual, &order.UploadedAt)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetOrders - rows.Scan: %w", err)
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (r *UserRepo) CreateOrder(ctx context.Context, userID string, orderNumber string) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo - CreateOrder - r.Pool.Begin: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	getQuery := `
		SELECT user_id, status, number
		FROM user_order
		WHERE number = $1
	`

	order := &entity.UserOrder{}

	err = tx.QueryRow(ctx, getQuery, orderNumber).Scan(
		&order.UserID,
		&order.Status,
		&order.Number,
	)

	// TODO: выглядит как жопа?
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("UserRepo - CreateOrder - Select: %w", err)
		}
	} else {
		if order.UserID == userID {
			return fmt.Errorf("UserRepo - CreateOrder - order.UserID == userID: %w", entity.ErrOrderAlreadyLoaded)
		}
		return fmt.Errorf("UserRepo - CreateOrder - order.UserID != userID: %w", entity.ErrOrderLoadedByAnotherUser)
	}

	newOrderQuery := `
		INSERT INTO user_order (user_id, number) 
		VALUES ($1, $2)
	`

	_, err = tx.Exec(ctx, newOrderQuery, userID, orderNumber)
	if err != nil {
		return fmt.Errorf("UserRepo - CreateOrder - tx.Exec: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("UserRepo - CreateOrder - tx.Commit: %w", err)
	}

	return nil
}

func (r *UserRepo) GetBalance(ctx context.Context, userID string) (*entity.UserBalance, error) {
	query := `
		SELECT "current", withdrawn 
		FROM user_balance
		WHERE user_id = $1
	`

	var balance entity.UserBalance

	if err := r.Pool.QueryRow(ctx, query, userID).Scan(&balance.Current, &balance.Withdrawn); err != nil {
		return nil, fmt.Errorf("UserRepo - GetBalance - r.Pool.QueryRow: %w", err)
	}

	return &balance, nil
}

func (r *UserRepo) GetWithdrawals(ctx context.Context, userID string) (*[]entity.UserWithdrawal, error) {
	query := `
		SELECT "order", "sum", processed_at 
		FROM user_withdrawal
		WHERE user_id = $1
	`

	rows, err := r.Pool.Query(ctx, query, userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case errors.Is(pgErr, pgx.ErrNoRows):
				return nil, fmt.Errorf("UserRepo - GetWithdrawals - r.Pool.Query: %w", entity.ErrNoContent)
			default:
				return nil, fmt.Errorf("UserRepo - GetWithdrawals - r.Pool.Query: %w", err)
			}
		}
	}
	defer rows.Close()

	withdrawals := make([]entity.UserWithdrawal, 0)

	for rows.Next() {
		var withdrawal entity.UserWithdrawal

		err = rows.Scan(&withdrawal.Order, &withdrawal.Sum, &withdrawal.ProcessedAt)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetWithdrawals - rows.Scan: %w", err)
		}
		withdrawals = append(withdrawals, withdrawal)
	}

	return &withdrawals, nil
}
