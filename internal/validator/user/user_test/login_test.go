package user_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	uservld "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
)

func TestValidateLoginRequest(t *testing.T) {
	t.Parallel()

	entropyPassword := float64(80)
	validator := uservld.New(entropyPassword)

	t.Run("Empty.Mobile", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "",
			Password: "password123",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		require.Error(t, err)
		assert.NotNil(t, fieldsErr)
		assert.Contains(t, fieldsErr, "mobile")
	})

	t.Run("Empty.Password", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "09123456789",
			Password: "",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		require.Error(t, err)
		assert.NotNil(t, fieldsErr)
		assert.Contains(t, fieldsErr, "password")
	})

	t.Run("Empty.Mobile.And.Password", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "",
			Password: "",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		require.Error(t, err)
		assert.NotNil(t, fieldsErr)
		assert.Contains(t, fieldsErr, "mobile")
		assert.Contains(t, fieldsErr, "password")
	})

	t.Run("Invalid.Mobile.Len", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "0912345678",
			Password: "password",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		require.Error(t, err)
		assert.NotNil(t, fieldsErr)
		assert.Contains(t, fieldsErr, "mobile")
	})

	t.Run("Valid.Input", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "09123456789",
			Password: "password123",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		assert.NoError(t, err)
		assert.Nil(t, fieldsErr)
		// assert.Equal(t, fieldsErr, map[string]string{})
	})
}
