package userdto

type ProfileRequest struct {
	ID uint `json:"id"`
}

type ProfileResponse struct {
	User UserInfo `json:"user_info"`
}
