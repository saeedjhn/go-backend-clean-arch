package userdto

type ProfileRequest struct {
	ID uint64 `json:"id"`
}

type ProfileResponse struct {
	User UserInfo `json:"user_info"`
}
