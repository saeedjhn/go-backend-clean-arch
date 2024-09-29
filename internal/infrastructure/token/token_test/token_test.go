package token_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
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
	as := token.New()

	t.Log("Access Token is:")
	accessToken, _ := as.CreateAccessToken(ID, AccessTokenSecret, AccessSubject, AccessTokenExpiryTime)
	t.Log(accessToken)

	t.Log("Parse Token, Access Token")
	pt, _ := as.ParseToken(accessToken, AccessTokenSecret)
	t.Log(pt)

	t.Log("Is authorized, access Token")
	t.Log(as.IsAuthorized(accessToken, AccessTokenSecret))

	t.Log("Refresh Token is")
	refreshToken, _ := as.CreateRefreshToken(ID, RefreshTokenSecret, RefreshSubject, RefreshTokenExpiryTime)
	t.Log(refreshToken)

	t.Log("Parse Token, Refresh Token")
	t.Log(as.ParseToken(refreshToken, RefreshTokenSecret))

	t.Log("Is authorized, refresh Token")
	t.Log(as.IsAuthorized(refreshToken, RefreshTokenSecret))
}
