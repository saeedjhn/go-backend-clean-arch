package authusecase_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

func TestCreateToken(t *testing.T) {
	config := authusecase.Config{
		AccessTokenSecret:      "TOKENSECRET",
		RefreshTokenSecret:     "REFRESHSECRET",
		AccessTokenSubject:     "as",
		RefreshTokenSubject:    "rs",
		AccessTokenExpiryTime:  7 * time.Hour,
		RefreshTokenExpiryTime: 120 * time.Hour,
	}

	user := entity.User{
		ID:        7,
		Name:      "John",
		Mobile:    "09111111111",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	auth := authusecase.New(config)

	accessToken, _ := auth.CreateAccessToken(entity.Authenticable{ID: user.ID})
	t.Log(accessToken)

	// refreshToken, _ := auth.CreateRefreshToken(userauthservicedto.CreateTokenRequest{Data: user})
	// t.Log(refreshToken)

	isAuthorized, err := auth.IsAuthorized(accessToken, config.AccessTokenSecret)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isAuthorized)
}
