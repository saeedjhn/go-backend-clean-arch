package authusecase

import (
	"fmt"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/token"
	"testing"
	"time"
)

const ID = 7

var config = Config{
	AccessTokenSecret:      "TOKENSECRET",
	RefreshTokenSecret:     "REFRESHSECRET",
	AccessTokenSubject:     "as",
	RefreshTokenSubject:    "rs",
	AccessTokenExpiryTime:  7 * time.Hour,
	RefreshTokenExpiryTime: 120 * time.Hour,
}

func TestCreateToken(t *testing.T) {
	ai := New(config, token.New())

	at, _ := ai.CreateAccessToken(7)
	fmt.Println(at)
	fmt.Println(ai.ExtractIdFromAccessToken(at))
	fmt.Println(ai.ParseAccessToken(at))

	rt, _ := ai.RefreshAccessToken(7)
	fmt.Println(rt)
	fmt.Println(ai.ExtractIdFromRefreshToken(rt))
	fmt.Println(ai.ParseRefreshToken(rt))
}
