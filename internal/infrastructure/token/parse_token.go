package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Token) ParseToken(requestToken string, secret string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
	})

	token, err = jwt.ParseWithClaims(requestToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
