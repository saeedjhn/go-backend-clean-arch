package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	mysqluser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user/user_test/doubles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate go test -v -race -count=1 -run Test_MysqlUser_GetByMobile
func Test_MysqlUser_GetByMobile(t *testing.T) {
	t.Parallel()

	conn, err := setuptest.NewMySQLDB(_myDBConfig)
	if err != nil {
		t.Fatalf("failed to create database connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := mysqluser.New(doubles.NewDummyTracer(), conn)

	testCases := []struct {
		name         string
		mobile       string
		expectedUser entity.User
		expectedErr  error
	}{
		{
			name:         "GetByMobile_MobileNotExists_ReturnError",
			mobile:       "09120000000",
			expectedUser: entity.User{},
			expectedErr:  errUserNotFound,
		},
		{
			name:         "GetByMobile_DBUnexpectedError_ReturnError",
			mobile:       "09130000000",
			expectedUser: entity.User{},
			expectedErr:  errUnexpected,
		},
		{
			name:   "GetByMobile_MobileExists_ReturnUser",
			mobile: "09120000001",
			expectedUser: entity.User{
				ID:       1,
				Name:     "Bob Smith",
				Mobile:   "09120000001",
				Email:    "bob.smith@example.com",
				Password: "password123",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user, errDB := db.GetByMobile(ctx, tc.mobile)

			if tc.expectedErr != nil {
				require.Error(t, errDB)

				if errors.Is(tc.expectedErr, errUserNotFound) {
					assert.ErrorIs(t, tc.expectedErr, errUserNotFound)
				}
				if errors.Is(tc.expectedErr, errUnexpected) {
					assert.ErrorIs(t, tc.expectedErr, errUnexpected)
				}
			} else {
				require.NoError(t, errDB)
				assert.Equal(t, tc.expectedUser.ID, user.ID)
				assert.Equal(t, tc.expectedUser.Mobile, user.Mobile)
			}
		})
	}
}
