package userdto

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token Token    `json:"token"`
}
