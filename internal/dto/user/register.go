package user

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserInfo    UserInfo          `json:"user"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
