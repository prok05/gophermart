package entity

import "errors"

var (
	ErrNotFound             = errors.New("resource was not found")
	ErrNoContent            = errors.New("data is empty")
	ErrDuplicateLogin       = errors.New("user with such login already exists")
	ErrWrongLoginOrPassword = errors.New("wrong login or password")
	ErrInvalidOrderNumber   = errors.New("invalid order number")
	ErrNotEnoughBalance     = errors.New("not enough balance")

	ErrOrderLoadedByAnotherUser = errors.New("order was already loaded by another user")
	ErrOrderAlreadyLoaded       = errors.New("order was already loaded by user")
)
