package entity

import "time"

type UserOrder struct {
	ID         string      `json:"id,omitempty"`
	UserID     string      `json:"user_id,omitempty"`
	Status     OrderStatus `json:"status,omitempty"`
	Number     string      `json:"number,omitempty"`
	Accrual    float64     `json:"accrual,omitempty"`
	UploadedAt time.Time   `json:"uploaded_at,omitempty"`
}

type OrderStatus string

const (
	OrderNewStatus        OrderStatus = "NEW"
	OrderProcessingStatus OrderStatus = "PROCESSING"
	OrderInvalidStatus    OrderStatus = "INVALID"
	OrderProcessedStatus  OrderStatus = "PROCESSED"
)
