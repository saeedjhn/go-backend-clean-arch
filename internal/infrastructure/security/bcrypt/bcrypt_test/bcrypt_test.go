package bcrypt_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/security/bcrypt"
)

func TestBcrypt(t *testing.T) {
	t.Log(bcrypt.Generate("pass", 32))
}
