package authusecase

import (
	"go-backend-clean-arch/internal/infrastructure/token"
	"time"
)

type Config struct {
	AccessTokenSecret      string        `mapstructure:"secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenSubject     string        `mapstructure:"access_subject"`
	RefreshTokenSubject    string        `mapstructure:"refresh_subject"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}

type AuthInteractor struct {
	config Config
	token  *token.Token
}

func New(config Config, token *token.Token) *AuthInteractor {
	return &AuthInteractor{config: config, token: token}
}

func (a AuthInteractor) CreateAccessToken(id uint) (string, error) {
	return a.token.CreateAccessToken(
		id,
		a.config.AccessTokenSecret,
		a.config.AccessTokenSubject,
		a.config.AccessTokenExpiryTime,
	)
}

func (a AuthInteractor) RefreshAccessToken(id uint) (string, error) {
	return a.token.CreateRefreshToken(
		id,
		a.config.RefreshTokenSecret,
		a.config.RefreshTokenSubject,
		a.config.RefreshTokenExpiryTime,
	)
}

func (a AuthInteractor) IsAuthorized(requestToken string, secret string) (bool, error) {
	return a.token.IsAuthorized(requestToken, secret)
}

func (a AuthInteractor) ExtractIDFromAccessToken(requestToken string) (string, error) {
	return a.token.ExtractIdFromToken(requestToken, a.config.AccessTokenSecret)
}

func (a AuthInteractor) ExtractIDFromRefreshToken(requestToken string) (string, error) {
	return a.token.ExtractIdFromToken(requestToken, a.config.RefreshTokenSecret)
}

func (a AuthInteractor) ParseAccessToken(requestToken string) (Claims, error) {
	claims, err := a.token.ParseToken(requestToken, a.config.AccessTokenSecret)

	return Claims(*claims), err
}

func (a AuthInteractor) ParseRefreshToken(requestToken string) (Claims, error) {
	claims, err := a.token.ParseToken(requestToken, a.config.RefreshTokenSecret)

	return Claims(*claims), err
}
