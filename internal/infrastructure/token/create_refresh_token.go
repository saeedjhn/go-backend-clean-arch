package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (a *Token) CreateRefreshToken(id uint) (string, error) {
	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.RefreshTokenExpiryTime * time.Second)),
		},
		UserId: id,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(a.config.RefreshTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
