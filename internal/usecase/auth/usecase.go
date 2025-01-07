package auth

import (
	"fmt"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/golang-jwt/jwt/v5"
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
	req entity.Authenticable,
) (string, error) {
	if len(i.Config.AccessTokenSecret) == 0 {
		return "", ErrMissingAccessTokenSecret
	}

	if i.Config.AccessTokenExpiryTime <= 0 {
		return "", ErrInvalidExpireTime
	}

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   i.Config.AccessTokenSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.Config.AccessTokenExpiryTime)),
		},
		UserID: req.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(i.Config.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (i Interactor) CreateRefreshToken(
	req entity.Authenticable,
) (string, error) {
	if len(i.Config.RefreshTokenSecret) == 0 {
		return "", ErrMissingRefreshTokenSecret
	}

	if i.Config.RefreshTokenExpiryTime <= 0 {
		return "", ErrInvalidExpireTime
	}

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   i.Config.RefreshTokenSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.Config.RefreshTokenExpiryTime)),
		},
		UserID: req.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(i.Config.RefreshTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, err
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

func (i Interactor) ParseToken(secret, requestToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		requestToken,
		&Claims{},
		func(_ *jwt.Token) (interface{}, error) {
			return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
