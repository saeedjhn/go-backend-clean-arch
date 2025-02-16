package user

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserInfo    Info              `json:"user"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
