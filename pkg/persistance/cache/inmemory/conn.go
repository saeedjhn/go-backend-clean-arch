package inmemory

import (
	"errors"
)

type InMemory struct {
	m map[string]interface{}
}

func New() *InMemory {
	return &InMemory{m: make(map[string]interface{})}
}

func (i *InMemory) Exists(key string) (bool, error) {
	_, ok := i.m[key]
	if !ok {
		return false, nil
	}

	return true, nil
}

func (i *InMemory) Set(key string, value interface{}) error {
	i.m[key] = value

	return nil
}

func (i *InMemory) Get(key string) (interface{}, error) {
	exists, err := i.Exists(key)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("not found")
	}

	return i.m[key], nil
}

func (i *InMemory) Del(key string) (bool, error) {
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
