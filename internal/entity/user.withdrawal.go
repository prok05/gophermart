package entity

import "time"

type UserWithdrawal struct {
	ID          string    `json:"id,omitempty" example:"UUID"`
	UserID      string    `json:"user_id,omitempty" example:"UUID"`
	Order       string    `json:"order,omitempty" example:"2377225624"`
	Sum         float64   `json:"sum,omitempty" example:"500"`
	ProcessedAt time.Time `json:"processed_at,omitempty" example:"2020-12-09T16:09:57+03:00"`
}
