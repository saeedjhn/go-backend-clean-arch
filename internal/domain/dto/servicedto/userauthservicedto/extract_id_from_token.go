package userauthservicedto

type ExtractIDFromTokenRequest struct {
	Token string
}

type ExtractIDFromTokenResponse struct {
	UserID uint64
}
