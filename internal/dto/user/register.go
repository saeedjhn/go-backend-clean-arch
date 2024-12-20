package user

type RegisterRequest struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Data        Data              `json:"data"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
