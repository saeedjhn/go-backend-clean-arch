package authservice_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
)

func TestCreateToken(t *testing.T) {
	config := authservice.Config{
		AccessTokenSecret:      "TOKENSECRET",
		RefreshTokenSecret:     "REFRESHSECRET",
		AccessTokenSubject:     "as",
		RefreshTokenSubject:    "rs",
		AccessTokenExpiryTime:  7 * time.Hour,
		RefreshTokenExpiryTime: 120 * time.Hour,
	}

	u := entity.User{
		ID:        7,
		Name:      "John",
		Mobile:    "09111111111",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	ai := authservice.New(config, token.New())

	at, _ := ai.CreateAccessToken(userauthservicedto.CreateTokenRequest{User: u})
	t.Log(at)

	t.Log(ai.ExtractIDFromAccessToken(struct{ Token string }{Token: at.Token}))

	t.Log(ai.ParseAccessToken(at.Token))

	rt, _ := ai.CreateRefreshToken(userauthservicedto.CreateTokenRequest{User: u})
	t.Log(rt)

	t.Log(ai.ExtractIDFromRefreshToken(struct{ Token string }{Token: rt.Token}))

	t.Log(ai.ParseRefreshToken(rt.Token))
}
