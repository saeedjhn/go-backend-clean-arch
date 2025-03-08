package user

import (
	"context"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/pkg/persistance/cache/redis"
)

type DB struct {
	conn *redis.DB
}

func New(conn *redis.DB) *DB {
	return &DB{conn: conn}
}

func (d *DB) Exists(_ context.Context, _ string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DB) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
	// TODO implement me
	panic("implement me")
}

func (d *DB) Get(_ context.Context, _ string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DB) Del(_ context.Context, _ string) (bool, error) {
	// TODO implement me
	panic("implement me")
}
