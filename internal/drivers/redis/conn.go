package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	cfg Config
	db  *redis.Client
}

func New(cfg Config) *RedisConnection {
	return (&RedisConnection{cfg: cfg}).conn()
}

func (r *RedisConnection) conn() *RedisConnection {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.Port),
		Password: r.cfg.Password,
		DB:       r.cfg.DB,
	})
	r.db = rdb

	return r
}

func (r *RedisConnection) Client() *redis.Client {
	return r.db
}
