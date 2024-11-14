package token

import (
	"github.com/golang-jwt/jwt/v5"
)

func (a *Token) ParseToken(requestToken string, secret string) (*Claims, error) {
	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	// token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	//	}
	//	return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
	// })

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
