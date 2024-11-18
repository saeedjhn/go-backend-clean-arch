package inmemoryuser

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/inmemory"
)

type DB struct {
	conn *inmemory.InMemory
}

func New(conn *inmemory.InMemory) *DB {
	return &DB{conn: conn}
}

func (d *DB) Exists(ctx context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DB) Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (d *DB) Get(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DB) Del(ctx context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
