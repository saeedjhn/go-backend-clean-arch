package user

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data   UserInfo `json:"data"`
	Tokens Tokens   `json:"tokens"`
}
