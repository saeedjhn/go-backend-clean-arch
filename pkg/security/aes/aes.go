package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
)

type Either int

// either 16, 24, or 32 bytes to select.
const (
	Either16 Either = 16
	Either24 Either = 24
	Either32 Either = 32
)

type CryptAES struct {
	secretKey string
}

func New(secretKey string) *CryptAES {
	return &CryptAES{secretKey: secretKey}
}

func (a *CryptAES) Encrypt(plainText string) (string, error) {
	s, err := a.checkSecretKeyLen()
	if err != nil {
		return s, err
	}

	newCipher, err := aes.NewCipher([]byte(a.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	// Generate a 12-byte nonce for GCM.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", fmt.Errorf("failed to generate random nonce: %w", err)
	}

	// Append nonce to ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return string(ciphertext), nil
}

func (a *CryptAES) Decrypt(cipherText string) (string, error) {
	newCipher, err := aes.NewCipher([]byte(a.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("ciphertext is too short to contain nonce and data")
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt ciphertext: %w", err)
	}

	return string(plaintext), nil
}

func (a *CryptAES) checkSecretKeyLen() (string, error) {
	switch Either(len(a.secretKey)) {
	case Either16, Either24, Either32:
	default:
		return "", fmt.Errorf(
			"invalid key length: got %d, expected one of %d, %d, or %d bytes",
			len(a.secretKey),
			Either16,
			Either24,
			Either32,
		)
	}

	return "", nil
}
