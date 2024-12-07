package aes_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/aes"
)

const _secret = "LKJKGKJHLJKLNGRWKLoolhsl"

func TestAES(t *testing.T) {
	cryptAES := aes.New(_secret)

	encrypt, err := cryptAES.Encrypt("1234556")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Encrypted encrypt: %x \n", encrypt)

	decrypt, err := cryptAES.Decrypt(encrypt)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decrypted plaintext: %s \n", decrypt)
}
