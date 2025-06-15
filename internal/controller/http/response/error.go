package response

type LoginResponse struct {
	Token string `json:"token" example:"token"`
}

type JSON struct {
	Data any `json:"data"`
}

type Error struct {
	Error string `json:"error"`
}
