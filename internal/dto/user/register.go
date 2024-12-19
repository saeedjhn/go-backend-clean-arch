package user

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Data UserInfo `json:"data"`
}
