package bcrypt_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
)

func TestBcrypt(t *testing.T) {
	t.Log(bcrypt.Generate("pass", 10))
}
