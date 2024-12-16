package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

type Either int

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

func (c *CryptAES) SetSecret(secret string) *CryptAES {
	c.secretKey = secret

	return c
}

func (c *CryptAES) GetSecret() string {
	return c.secretKey
}

func (c *CryptAES) Encrypt(plainText string) (string, error) {
	if err := c.checkSecretKeyLen(); err != nil {
		return "", err
	}

	// Create AES cipher
	newCipher, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Create GCM cipher mode
	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	// Generate c secure random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate random nonce: %w", err)
	}

	// Encrypt the plaintext and prepend the nonce
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (c *CryptAES) Decrypt(cipherText string) (string, error) {
	cipherData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("failed to decode Base64: %w", err)
	}

	// Create AES cipher
	newCipher, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Create GCM cipher mode
	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		return "", errors.New("ciphertext is too short to contain nonce and data")
	}

	// Extract nonce and ciphertext
	nonce, ciphertext := cipherData[:nonceSize], cipherData[nonceSize:]

	// Verify nonce length
	if len(nonce) != nonceSize {
		return "", fmt.Errorf("invalid nonce length: expected %d, got %d", nonceSize, len(nonce))
	}

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil) // #nosec G407
	if err != nil {
		return "", fmt.Errorf("failed to decrypt ciphertext: %w", err)
	}

	return string(plaintext), nil
}

func (c *CryptAES) checkSecretKeyLen() error {
	switch Either(len(c.secretKey)) {
	case Either16, Either24, Either32:
		// Valid key lengths
	default:
		return fmt.Errorf(
			"invalid key length: got %d, expected one of %d, %d, or %d bytes",
			len(c.secretKey),
			Either16,
			Either24,
			Either32,
		)
	}
	return nil
}
