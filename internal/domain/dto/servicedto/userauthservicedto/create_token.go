package userauthservicedto

import (
	"go-backend-clean-arch/internal/domain/entity"
)

type CreateTokenRequest struct {
	User entity.User
}

type CreateTokenResponse struct {
	Token string
}
