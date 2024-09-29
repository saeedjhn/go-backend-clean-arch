package mysqluser

/*
stubbing is a way to simulate the behavior of some components in your application, typically for the purpose of testing.
Stubs are often used to replace real implementations with mock implementations that return predefined results.
This allows you to test your code in isolation and control the environment more precisely.
*/

import (
	"errors"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

const (
	_FirstRecordID  = 1001
	_SecondRecordID = 1002
)

type DBStub struct {
	conn map[uint64]entity.User
}

func NewDBStub() *DBStub {
	conn := map[uint64]entity.User{
		1001: {
			ID:        _FirstRecordID,
			Name:      "n",
			Mobile:    "m",
			Email:     "e",
			Password:  "p",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		1002: {
			ID:        _SecondRecordID,
			Name:      "nn",
			Mobile:    "mm",
			Email:     "ee",
			Password:  "pp",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return &DBStub{conn: conn}
}

func (s DBStub) Create(u entity.User) (uint64, error) {
	s.conn[u.ID] = u
	exists := s.IsByID(u.ID)
	if exists {
		return 0, errors.New("user already")
	}

	return u.ID, nil
}

func (s DBStub) GetByID(id uint64) (entity.User, error) {
	user, exists := s.conn[id]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}

func (s DBStub) IsByID(id uint64) bool {
	_, exists := s.conn[id]

	return exists
}
