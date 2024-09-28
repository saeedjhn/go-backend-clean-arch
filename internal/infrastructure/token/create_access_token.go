package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (a *Token) CreateAccessToken(id uint64, secret string, subject string, expire time.Duration) (string, error) {
	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: subject,
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.AccessTokenExpiryTime * time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		},
		UserID: id,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := accessToken.SignedString([]byte(a.config.AccessTokenSecret))
	tokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
