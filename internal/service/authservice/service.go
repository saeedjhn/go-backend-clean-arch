package authservice

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
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

// var _ userservice.AuthGenerator = (*AuthInteractor)(nil) // Commented, because it happens import cycle.

func New(config Config, token *token.Token) *AuthInteractor {
	return &AuthInteractor{config: config, token: token}
}

func (a AuthInteractor) CreateAccessToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	at, err := a.token.CreateAccessToken(
		req.User.ID,
		a.config.AccessTokenSecret,
		a.config.AccessTokenSubject,
		a.config.AccessTokenExpiryTime,
	)

	return userauthservicedto.CreateTokenResponse{Token: at}, err
}

func (a AuthInteractor) CreateRefreshToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	rt, err := a.token.CreateRefreshToken(
		req.User.ID,
		a.config.RefreshTokenSecret,
		a.config.RefreshTokenSubject,
		a.config.RefreshTokenExpiryTime,
	)

	return userauthservicedto.CreateTokenResponse{Token: rt}, err
}

func (a AuthInteractor) IsAuthorized(requestToken string, secret string) (bool, error) {
	return a.token.IsAuthorized(requestToken, secret)
}

func (a AuthInteractor) ExtractIDFromAccessToken(
	req userauthservicedto.ExtractIDFromTokenRequest,
) (userauthservicedto.ExtractIDFromTokenResponse, error) {
	et, err := a.token.ParseToken(req.Token, a.config.AccessTokenSecret)

	return userauthservicedto.ExtractIDFromTokenResponse{
		UserID: et.UserID,
	}, err
}

func (a AuthInteractor) ExtractIDFromRefreshToken(
	req userauthservicedto.ExtractIDFromTokenRequest,
) (userauthservicedto.ExtractIDFromTokenResponse, error) {
	et, err := a.token.ParseToken(req.Token, a.config.RefreshTokenSecret)

	return userauthservicedto.ExtractIDFromTokenResponse{
		UserID: et.UserID,
	}, err
}

func (a AuthInteractor) ParseAccessToken(requestToken string) (Claims, error) {
	claims, err := a.token.ParseToken(requestToken, a.config.AccessTokenSecret)

	return Claims(*claims), err
}

func (a AuthInteractor) ParseRefreshToken(requestToken string) (Claims, error) {
	claims, err := a.token.ParseToken(requestToken, a.config.RefreshTokenSecret)

	return Claims(*claims), err
}
