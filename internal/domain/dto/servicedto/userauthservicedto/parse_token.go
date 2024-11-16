package userauthservicedto

type ParseTokenRequest struct {
	Secret string
	Token  string
}

type ParseTokenResponse[T interface{}] struct {
	Claims T
}
