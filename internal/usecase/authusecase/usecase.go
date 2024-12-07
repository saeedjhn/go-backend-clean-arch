package authusecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
)

type Config struct {
	AccessTokenSecret      string        `mapstructure:"access_secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenSubject     string        `mapstructure:"access_subject"`
	RefreshTokenSubject    string        `mapstructure:"refresh_subject"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}

type Interactor struct {
	Config Config
}

// var _ userservice.AuthGenerator = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(config Config) *Interactor {
	return &Interactor{Config: config}
}

func (i Interactor) CreateAccessToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   i.Config.AccessTokenSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.Config.AccessTokenExpiryTime)),
		},
		UserID: req.User.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(i.Config.AccessTokenSecret))
	if err != nil {
		return userauthservicedto.CreateTokenResponse{}, err
	}

	return userauthservicedto.CreateTokenResponse{Token: tokenString}, err
}

func (i Interactor) CreateRefreshToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   i.Config.RefreshTokenSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.Config.RefreshTokenExpiryTime)),
		},
		UserID: req.User.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(i.Config.RefreshTokenSecret))
	if err != nil {
		return userauthservicedto.CreateTokenResponse{}, err
	}

	return userauthservicedto.CreateTokenResponse{Token: tokenString}, err
}

func (i Interactor) IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (i Interactor) ParseToken(
	req userauthservicedto.ParseTokenRequest,
) (userauthservicedto.ParseTokenResponse[*Claims], error) {
	token, err := jwt.ParseWithClaims(
		req.Token,
		&Claims{},
		func(_ *jwt.Token) (interface{}, error) {
			return []byte(req.Secret), nil // secret is accessTokenSecret or refreshTokenSecret_
		},
	)
	if err != nil {
		return userauthservicedto.ParseTokenResponse[*Claims]{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return userauthservicedto.ParseTokenResponse[*Claims]{Claims: claims}, nil
	}

	return userauthservicedto.ParseTokenResponse[*Claims]{}, err
}
