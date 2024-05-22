package authusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/token"
	"time"
)

type Config struct {
	AccessTokenSecret      string        `mapstructure:"secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}

type AuthInteractor struct {
	token token.Token
}

func New(token token.Token) *AuthInteractor {
	return &AuthInteractor{token: token}
}

func (a AuthInteractor) CreateAccessToken(id uint) (string, error) {
	return a.token.CreateAccessToken(id)
}

func (a AuthInteractor) RefreshAccessToken(id uint) (string, error) {
	return a.token.CreateRefreshToken(id)
}

func (a AuthInteractor) IsAuthorized(requestToken string, secret string) (bool, error) {
	return a.token.IsAuthorized(requestToken, secret)
}

func (a AuthInteractor) ExtractIdFromToken(requestToken string, secret string) (string, error) {
	return a.token.ExtractIdFromToken(requestToken, secret)
}

func (a AuthInteractor) ParseToken(requestToken string, secret string) (Claims, error) {
	claims, err := a.token.ParseToken(requestToken, secret)

	return Claims(*claims), err
}
