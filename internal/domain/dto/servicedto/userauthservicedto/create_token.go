package userauthservicedto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"time"
)

type CreateTokenRequest struct {
	User       entity.User
	Secret     string
	Subject    string
	ExpireTime time.Duration
}

type CreateTokenResponse struct {
	Token string
}
