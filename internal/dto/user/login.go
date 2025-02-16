package user

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserInfo    Info              `json:"user"`
	Tokens      Tokens            `json:"tokens"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
