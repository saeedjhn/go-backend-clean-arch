package usecase

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
)

//go:generate mockery --name AuthInteractor
type AuthInteractor interface {
	CreateAccessToken(req models.Authenticable) (string, error)
	CreateRefreshToken(req models.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authentication.Claims, error)
}
