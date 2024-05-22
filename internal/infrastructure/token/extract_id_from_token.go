package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func (a *Token) ExtractIdFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	//claims, fiberok := token.Claims.(jwt.MapClaims)
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return strconv.FormatFloat(claims["id"].(float64), 'f', -1, 64), nil
}
