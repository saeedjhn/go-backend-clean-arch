package user_test

import (
	"testing"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	uservld "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ValidateLoginRequest(t *testing.T) {
	t.Parallel()

	entropyPassword := float64(80)
	validator := uservld.New(entropyPassword)

	t.Run("EmptyMobileField_ReturnsMobileFieldError", func(t *testing.T) {
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

	t.Run("EmptyPasswordField_ReturnsPasswordFieldError", func(t *testing.T) {
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

	t.Run("EmptyMobileAndPasswordFields_ReturnsMobileAndPasswordErrors", func(t *testing.T) {
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

	t.Run("InvalidMobileLength_ReturnsMobileFieldError", func(t *testing.T) {
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

	t.Run("ValidInput_ReturnsNoErrors", func(t *testing.T) {
		t.Parallel()

		req := userdto.LoginRequest{
			Mobile:   "09123456789",
			Password: "password123",
		}

		fieldsErr, err := validator.ValidateLoginRequest(req)

		require.NoError(t, err)
		assert.Nil(t, fieldsErr)
	})
}
