package aes_test

import (
	"testing"

	"github.com/saeedjhn/go-domain-driven-design/pkg/security/aes"
)

//go:generate go test -v -race -count=1 ./...

func TestCryptAES_Encrypt_ValidKey_ReturnsEncryptedText(t *testing.T) {
	t.Parallel()

	secretKey := "thisis32bitlongpassphraseimusing"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"

	encryptedText, err := crypt.Encrypt(plainText)

	if err != nil {
		t.Errorf("Encrypt failed with error: %v", err)
	}
	if encryptedText == "" {
		t.Error("Encrypted text should not be empty")
	}
}

func TestCryptAES_Encrypt_InvalidKey_ReturnsError(t *testing.T) {
	t.Parallel()

	secretKey := "shortkey"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"

	_, err := crypt.Encrypt(plainText)

	if err == nil {
		t.Error("Expected error for invalid key length, got nil")
	}
}

func TestCryptAES_Decrypt_ValidCipherText_ReturnsOriginalText(t *testing.T) {
	t.Parallel()

	secretKey := "thisis32bitlongpassphraseimusing"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"
	encryptedText, _ := crypt.Encrypt(plainText)

	decryptedText, err := crypt.Decrypt(encryptedText)

	if err != nil {
		t.Errorf("Decrypt failed with error: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be '%s', got '%s'", plainText, decryptedText)
	}
}

func TestCryptAES_Decrypt_InvalidCipherText_ReturnsError(t *testing.T) {
	t.Parallel()

	secretKey := "thisis32bitlongpassphraseimusing"
	crypt := aes.New(secretKey)
	invalidCipherText := "invalidbase64encodedtext"

	_, err := crypt.Decrypt(invalidCipherText)

	if err == nil {
		t.Error("Expected error for invalid cipher text, got nil")
	}
}

func TestCryptAES_SetSecret_GetSecret_ReturnsUpdatedSecret(t *testing.T) {
	t.Parallel()

	secretKey := "thisis32bitlongpassphraseimusing"
	newSecretKey := "new32bitlongpassphraseimusing"
	crypt := aes.New(secretKey)

	crypt.SetSecret(newSecretKey)
	retrievedSecret := crypt.GetSecret()

	if retrievedSecret != newSecretKey {
		t.Errorf("Expected secret key to be '%s', got '%s'", newSecretKey, retrievedSecret)
	}
}

func TestCryptAES_EncryptDecrypt_24ByteKey_ReturnsOriginalText(t *testing.T) {
	t.Parallel()

	secretKey := "24bytekeylongenough12322"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"

	encryptedText, err := crypt.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Encrypt failed with error: %v", err)
	}
	decryptedText, err := crypt.Decrypt(encryptedText)

	if err != nil {
		t.Errorf("Decrypt failed with error: %v", err)
	}
	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be '%s', got '%s'", plainText, decryptedText)
	}
}

func TestCryptAES_EncryptDecrypt_16ByteKey_ReturnsOriginalText(t *testing.T) {
	t.Parallel()

	secretKey := "16bytekeylong123"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"

	encryptedText, err := crypt.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Encrypt failed with error: %v", err)
	}
	decryptedText, err := crypt.Decrypt(encryptedText)

	if err != nil {
		t.Errorf("Decrypt failed with error: %v", err)
	}
	if decryptedText != plainText {
		t.Errorf("Expected decrypted text to be '%s', got '%s'", plainText, decryptedText)
	}
}

func TestCryptAES_Decrypt_TamperedCipherText_ReturnsError(t *testing.T) {
	t.Parallel()

	secretKey := "thisis32bitlongpassphraseimusing"
	crypt := aes.New(secretKey)
	plainText := "Hello, World!"
	encryptedText, _ := crypt.Encrypt(plainText)

	tamperedCipherText := encryptedText[:len(encryptedText)-1] + "x"

	_, err := crypt.Decrypt(tamperedCipherText)

	if err == nil {
		t.Error("Expected error for tampered cipher text, got nil")
	}
}
