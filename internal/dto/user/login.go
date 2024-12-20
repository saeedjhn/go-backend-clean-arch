package user

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data        Data              `json:"data"`
	Tokens      Tokens            `json:"tokens"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
