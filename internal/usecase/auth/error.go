package auth

import "errors"

var (
	ErrMissingAccessTokenSecret  = errors.New("missing access token secret: ensure the secret is configured correctly")
	ErrMissingRefreshTokenSecret = errors.New("missing refresh token secret: ensure the secret is configured correctly")
	ErrInvalidExpireTime         = errors.New("invalid expire time: must be greater than zero")
)
