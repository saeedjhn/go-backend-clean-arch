package token

import (
	"fmt"
	"testing"
)

const (
	ID                 = 7
	AccessTokenSecret  = "access_token"
	RefreshTokenSecret = "refresh_token"
)

func TestService(t *testing.T) {

	as := New(Config{
		AccessTokenSecret:      AccessTokenSecret,
		RefreshTokenSecret:     RefreshTokenSecret,
		AccessTokenExpiryTime:  60,
		RefreshTokenExpiryTime: 60,
	})
	fmt.Println("Access Token is:")
	accessToken, _ := as.CreateAccessToken(ID)
	fmt.Println(accessToken)

	fmt.Println("Parse Token, Access Token")
	pt, _ := as.ParseToken(accessToken, AccessTokenSecret)
	fmt.Println(pt)

	fmt.Println("Is authorized, access Token")
	fmt.Println(as.IsAuthorized(accessToken, AccessTokenSecret))

	fmt.Println("Extract Id from access token")
	fmt.Println(as.ExtractIdFromToken(accessToken, AccessTokenSecret))

	fmt.Println("Refresh Token is")
	refreshToken, _ := as.CreateRefreshToken(ID)
	fmt.Println(refreshToken)

	fmt.Println("Parse Token, Refresh Token")
	fmt.Println(as.ParseToken(refreshToken, RefreshTokenSecret))

	fmt.Println("Is authorized, refresh Token")
	fmt.Println(as.IsAuthorized(refreshToken, RefreshTokenSecret))

	fmt.Println("Extract Id from refresh token")
	fmt.Println(as.ExtractIdFromToken(refreshToken, RefreshTokenSecret))
}
