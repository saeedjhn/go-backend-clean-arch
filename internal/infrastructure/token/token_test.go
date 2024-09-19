package token

import (
	"log"
	"testing"
	"time"
)

const (
	ID                     = 7
	AccessTokenSecret      = "access_token"
	RefreshTokenSecret     = "refresh_token"
	AccessSubject          = "as"
	RefreshSubject         = "rs"
	AccessTokenExpiryTime  = 7 * time.Hour
	RefreshTokenExpiryTime = 120 * time.Hour
)

func TestService(t *testing.T) {
	// as := New(Config{
	//	AccessTokenSecret:      AccessTokenSecret,
	//	RefreshTokenSecret:     RefreshTokenSecret,
	//	AccessTokenExpiryTime:  60,
	//	RefreshTokenExpiryTime: 60,
	// })
	as := New()

	log.Println("Access Token is:")
	accessToken, _ := as.CreateAccessToken(ID, AccessTokenSecret, AccessSubject, AccessTokenExpiryTime)
	log.Println(accessToken)

	log.Println("Parse Token, Access Token")
	pt, _ := as.ParseToken(accessToken, AccessTokenSecret)
	log.Println(pt)

	log.Println("Is authorized, access Token")
	log.Println(as.IsAuthorized(accessToken, AccessTokenSecret))

	log.Println("Extract Id from access token")
	log.Println(as.ExtractIdFromToken(accessToken, AccessTokenSecret))

	log.Println("Refresh Token is")
	refreshToken, _ := as.CreateRefreshToken(ID, RefreshTokenSecret, RefreshSubject, RefreshTokenExpiryTime)
	log.Println(refreshToken)

	log.Println("Parse Token, Refresh Token")
	log.Println(as.ParseToken(refreshToken, RefreshTokenSecret))

	log.Println("Is authorized, refresh Token")
	log.Println(as.IsAuthorized(refreshToken, RefreshTokenSecret))

	log.Println("Extract Id from refresh token")
	log.Println(as.ExtractIdFromToken(refreshToken, RefreshTokenSecret))
}
