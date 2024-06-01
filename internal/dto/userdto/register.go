package userdto

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	User UserInfo `json:"user_info"`
}
