package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Cost int

const (
	MinCost     Cost = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     Cost = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost Cost = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func Generate(str string, cost Cost) (string, error) {
	if cost > MaxCost {
		return "", fmt.Errorf("bcrypt: unsupported cost value %d (maximum allowed is %d)", cost, MaxCost)
	}

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(str), int(cost))
	if err != nil {
		return "", fmt.Errorf("bcrypt: failed to generate hash: %w", err)
	}

	return string(encryptedPass), nil
}

func CompareHashAndSTR(hashed string, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str))
}
