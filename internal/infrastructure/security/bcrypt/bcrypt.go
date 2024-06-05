package bcrypt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Cost int

const (
	MinCost     Cost = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     Cost = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost Cost = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func Generate(str string, cost Cost) (string, error) {
	if cost > 31 {
		return "", errors.New("don`t supported cost")
	}

	encryptedPass, _ := bcrypt.GenerateFromPassword(
		[]byte(str),
		int(cost),
	)

	return string(encryptedPass), nil
}
