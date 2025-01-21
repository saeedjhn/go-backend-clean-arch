package user

type ProfileRequest struct {
	ID uint64 `json:"id"`
}

type ProfileResponse struct {
	UserInfo    UserInfo          `json:"user"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
