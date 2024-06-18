package authservice

import (
	"fmt"
	"go-backend-clean-arch/internal/domain/entity"
	"go-backend-clean-arch/internal/infrastructure/token"
	"testing"
	"time"
)

var config = Config{
	AccessTokenSecret:      "TOKENSECRET",
	RefreshTokenSecret:     "REFRESHSECRET",
	AccessTokenSubject:     "as",
	RefreshTokenSubject:    "rs",
	AccessTokenExpiryTime:  7 * time.Hour,
	RefreshTokenExpiryTime: 120 * time.Hour,
}

func TestCreateToken(t *testing.T) {
	u := entity.User{
		ID:        7,
		Name:      "",
		Mobile:    "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	ai := New(config, token.New())

	at, _ := ai.CreateAccessToken(u)
	fmt.Println(at)
	fmt.Println(ai.ExtractIDFromAccessToken(at))
	fmt.Println(ai.ParseAccessToken(at))

	rt, _ := ai.CreateRefreshToken(u)
	fmt.Println(rt)
	fmt.Println(ai.ExtractIDFromRefreshToken(rt))
	fmt.Println(ai.ParseRefreshToken(rt))
}
