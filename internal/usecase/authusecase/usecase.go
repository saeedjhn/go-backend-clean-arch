package authusecase

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/token"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
)

type Config struct {
	AccessTokenSecret      string        `mapstructure:"secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenSubject     string        `mapstructure:"access_subject"`
	RefreshTokenSubject    string        `mapstructure:"refresh_subject"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}

type Interactor struct {
	config Config
	token  *token.Token
}

// var _ userservice.AuthGenerator = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(config Config, token *token.Token) *Interactor {
	return &Interactor{config: config, token: token}
}

func (i Interactor) CreateAccessToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	at, err := i.token.CreateAccessToken(
		req.User.ID,
		i.config.AccessTokenSecret,
		i.config.AccessTokenSubject,
		i.config.AccessTokenExpiryTime,
	)

	return userauthservicedto.CreateTokenResponse{Token: at}, err
}

func (i Interactor) CreateRefreshToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	rt, err := i.token.CreateRefreshToken(
		req.User.ID,
		i.config.RefreshTokenSecret,
		i.config.RefreshTokenSubject,
		i.config.RefreshTokenExpiryTime,
	)

	return userauthservicedto.CreateTokenResponse{Token: rt}, err
}

func (i Interactor) IsAuthorized(requestToken string, secret string) (bool, error) {
	return i.token.IsAuthorized(requestToken, secret)
}

func (i Interactor) ExtractIDFromAccessToken(
	req userauthservicedto.ExtractIDFromTokenRequest,
) (userauthservicedto.ExtractIDFromTokenResponse, error) {
	et, err := i.token.ParseToken(req.Token, i.config.AccessTokenSecret)

	return userauthservicedto.ExtractIDFromTokenResponse{
		UserID: et.UserID,
	}, err
}

func (i Interactor) ExtractIDFromRefreshToken(
	req userauthservicedto.ExtractIDFromTokenRequest,
) (userauthservicedto.ExtractIDFromTokenResponse, error) {
	et, err := i.token.ParseToken(req.Token, i.config.RefreshTokenSecret)

	return userauthservicedto.ExtractIDFromTokenResponse{
		UserID: et.UserID,
	}, err
}

func (i Interactor) ParseAccessToken(requestToken string) (Claims, error) {
	claims, err := i.token.ParseToken(requestToken, i.config.AccessTokenSecret)

	return Claims(*claims), err
}

func (i Interactor) ParseRefreshToken(requestToken string) (Claims, error) {
	claims, err := i.token.ParseToken(requestToken, i.config.RefreshTokenSecret)

	return Claims(*claims), err
}
