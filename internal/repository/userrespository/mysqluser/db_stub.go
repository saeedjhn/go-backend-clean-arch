package mysqluser

/*
stubbing is a way to simulate the behavior of some components in your application, typically for the purpose of testing.
Stubs are often used to replace real implementations with mock implementations that return predefined results.
This allows you to test your code in isolation and control the environment more precisely.
*/

import (
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"time"
)

type DBStub struct {
	conn map[uint]entity.User
}

func NewDBStub() *DBStub {
	conn := map[uint]entity.User{
		1001: {1001, "s", "m", "e", "p", time.Now(), time.Now()},
		1002: {1002, "ss", "mm", "ee", "pp", time.Now(), time.Now()},
	}

	return &DBStub{conn: conn}
}

func (s DBStub) Create(u entity.User) (uint, error) {
	s.conn[u.ID] = u
	exists := s.IsByID(u.ID)
	if exists {
		return 0, errors.New("user already")
	}

	return u.ID, nil
}

func (s DBStub) GetByID(id uint) (entity.User, error) {
	user, exists := s.conn[id]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}

func (s DBStub) IsByID(id uint) bool {
	_, exists := s.conn[id]

	return exists
}
