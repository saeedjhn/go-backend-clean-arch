package userdto

type RegisterRequest struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type RegisterResponse struct {
	User  UserInfo `json:"user_info"`
	Token Token    `json:"token"`
}
