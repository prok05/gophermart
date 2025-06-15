package entity

type UserBalance struct {
	UserID    string  `json:"user_id,omitempty" example:"UUID"`
	Current   float64 `json:"current" example:"500.5"`
	Withdrawn float64 `json:"withdrawn" example:"42"`
}
