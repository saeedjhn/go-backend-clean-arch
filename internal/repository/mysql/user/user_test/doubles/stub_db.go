package doubles

/*
stubbing is a way to simulate the behavior of some components in your application, typically for the purpose of testing.
Stubs are often used to replace real implementations with mock implementations that return predefined results.
This allows you to test your code in isolation and control the environment more precisely.
*/

import (
	"errors"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/types"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
)

const (
	_FirstRecordID  = 1001
	_SecondRecordID = 1002
)

type StubDB struct {
	conn map[types.ID]entity.User
}

func NewStubDB() *StubDB {
	conn := map[types.ID]entity.User{
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

	return &StubDB{conn: conn}
}

func (s StubDB) Create(u entity.User) (types.ID, error) {
	s.conn[u.ID] = u
	exists := s.IsByID(u.ID)
	if exists {
		return 0, errors.New("user already")
	}

	return u.ID, nil
}

func (s StubDB) GetByID(id types.ID) (entity.User, error) {
	user, exists := s.conn[id]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}

func (s StubDB) IsByID(id types.ID) bool {
	_, exists := s.conn[id]

	return exists
}
