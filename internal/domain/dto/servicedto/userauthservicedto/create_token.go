package userauthservicedto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type CreateTokenRequest struct {
	User entity.User
}

type CreateTokenResponse struct {
	Token string
}
