package authservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
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

	at, _ := ai.CreateAccessToken(userauthservicedto.CreateTokenRequest{User: u})
	log.Println(at)
	log.Println(ai.ExtractIDFromAccessToken(userauthservicedto.ExtractIDFromTokenRequest{Token: at.Token}))
	log.Println(ai.ParseAccessToken(at.Token))

	rt, _ := ai.CreateRefreshToken(userauthservicedto.CreateTokenRequest{User: u})
	log.Println(rt)
	log.Println(ai.ExtractIDFromRefreshToken(userauthservicedto.ExtractIDFromTokenRequest{Token: rt.Token}))
	log.Println(ai.ParseRefreshToken(rt.Token))
}
