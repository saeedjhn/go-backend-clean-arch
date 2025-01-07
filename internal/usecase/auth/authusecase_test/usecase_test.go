package authusecase_test

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _config = auth.Config{ //nolint:gochecknoglobals // nothing
	AccessTokenSecret:     "secret123",
	AccessTokenSubject:    "access-subject",
	AccessTokenExpiryTime: time.Minute,
}

type request struct {
	Secret     string
	Data       entity.Authenticable
	ExpireTime time.Duration
}

//go:generate go test -count=1 -v ./...

//go:generate go test -v -race -count=1 -run Test_AuthInteractor_CreateAccessToken
func Test_AuthInteractor_CreateAccessToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		req           request
		expectedError bool
	}{
		{
			name: "NegativeExpireTime_TokenNotGenerated",
			req: request{
				Secret:     "secret",
				Data:       entity.Authenticable{ID: 1},
				ExpireTime: -time.Second,
			},
			expectedError: true,
		},
		{
			name: "ZeroExpireTime_TokenNotGenerated",
			req: request{
				Secret:     "secret",
				Data:       entity.Authenticable{ID: 1},
				ExpireTime: 0,
			},
			expectedError: true,
		},
		{
			name: "EmptySecret_TokenNotGenerated",
			req: request{
				Secret:     "",
				Data:       entity.Authenticable{ID: 1},
				ExpireTime: 5 * time.Second,
			},
			expectedError: true,
		},
		{
			name: "ValidRequest_TokenGenerated",
			req: request{
				Secret:     "secret",
				Data:       entity.Authenticable{ID: 1},
				ExpireTime: 5 * time.Minute,
			},
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			interactor := auth.New(auth.Config{
				AccessTokenSecret:     tc.req.Secret,
				AccessTokenExpiryTime: tc.req.ExpireTime,
			})
			token, err := interactor.CreateAccessToken(tc.req.Data)

			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func Test_AuthInteractor_IsAuthorized(t *testing.T) {
	t.Parallel()
	interactor := auth.New(_config)

	tests := []struct {
		name          string
		token         string
		secret        string
		expectedAuth  bool
		expectedError bool
	}{
		{
			name:          "ValidToken_IsAuthorized",
			token:         generateTestToken("secret123", time.Minute),
			secret:        "secret123",
			expectedAuth:  true,
			expectedError: false,
		},
		{
			name:          "InValidSecret_IsNotAuthorized",
			token:         generateTestToken("wrongSecret", time.Minute),
			secret:        "secret123",
			expectedAuth:  false,
			expectedError: true,
		},
		{
			name:          "ExpiredToken_IsNotAuthorized",
			token:         generateTestToken("secret123", -time.Minute),
			secret:        "secret123",
			expectedAuth:  false,
			expectedError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			isAuthorized, err := interactor.IsAuthorized(tc.token, tc.secret)

			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.expectedAuth, isAuthorized)
		})
	}
}

func Test_AuthInteractor_ParseToken(t *testing.T) {
	t.Parallel()

	interactor := auth.New(_config)

	tests := []struct {
		name           string
		token          string
		secret         string
		expectedUserID uint64
		expectedError  bool
	}{
		{
			name:           "ValidToken_ClaimsReturned",
			token:          generateTestTokenWithClaims("secret123", time.Minute, 1),
			secret:         "secret123",
			expectedUserID: 1,
			expectedError:  false,
		},
		{
			name:           "InValidSecret_ClaimsNotReturned",
			token:          generateTestTokenWithClaims("wrongSecret", time.Minute, 1),
			secret:         "secret123",
			expectedUserID: 0,
			expectedError:  true,
		},
		{
			name:           "ExpireToken_ClaimsNotReturned",
			token:          generateTestTokenWithClaims("secret123", -time.Minute, 1),
			secret:         "secret123",
			expectedUserID: 0,
			expectedError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			claims, err := interactor.ParseToken(tc.secret, tc.token)

			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, claims)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, claims)
				assert.Equal(t, tc.expectedUserID, claims.UserID)
			}
		})
	}
}

func generateTestToken(secret string, duration time.Duration) string {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func generateTestTokenWithClaims(secret string, duration time.Duration, userID uint64) string {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}
