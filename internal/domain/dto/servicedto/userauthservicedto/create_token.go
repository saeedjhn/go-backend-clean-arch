package userauthservicedto

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
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
