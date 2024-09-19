package authservice

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

// func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
// }
