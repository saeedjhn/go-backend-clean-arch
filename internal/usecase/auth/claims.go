package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"user_id"`
}

// func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
// }
