package user

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string            `json:"access_token"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
