package inmemory

import (
	"errors"
)

type DB struct {
	m map[string]interface{}
}

func New() *DB {
	return &DB{m: make(map[string]interface{})}
}

func (i *DB) Exists(key string) (bool, error) {
	_, ok := i.m[key]
	if !ok {
		return false, nil
	}

	return true, nil
}

func (i *DB) Set(key string, value interface{}) error {
	i.m[key] = value

	return nil
}

func (i *DB) Get(key string) (interface{}, error) {
	exists, err := i.Exists(key)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("not found")
	}

	return i.m[key], nil
}

func (i *DB) Del(key string) (bool, error) {
	exists, err := i.Exists(key)
	if err != nil {
		return false, err
	}

	if !exists {
		return false, errors.New("not found")
	}

	delete(i.m, key)

	return true, nil
}
