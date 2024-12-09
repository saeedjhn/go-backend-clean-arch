package userdto

type ProfileRequest struct {
	ID uint64 `json:"id"`
}

type ProfileResponse struct {
	Data UserInfo `json:"data"`
}
