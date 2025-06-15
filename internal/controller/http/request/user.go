package request

type RegisterUser struct {
	Login    string `json:"login" validate:"required,max=100" example:"login"`
	Password string `json:"password" validate:"required,min=8,max=72" example:"password"`
}

type LoginUser struct {
	Login    string `json:"login" validate:"required,max=100" example:"login"`
	Password string `json:"password" validate:"required,min=8,max=72" example:"password"`
}

type WithdrawBalance struct {
	Order string  `json:"order" validate:"required,max=100" example:"2377225624"`
	Sum   float64 `json:"sum" validate:"required" example:"751"`
}
