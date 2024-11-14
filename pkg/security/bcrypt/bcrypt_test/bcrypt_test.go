package bcrypt_test

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	t.Log(bcrypt.Generate("pass", 32))
}
