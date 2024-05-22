package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type DB interface {
	Client() *redis.Client
}

type RedisDB struct {
	cfg Config
	db  *redis.Client
}

func New(cfg Config) *RedisDB {
	return (&RedisDB{cfg: cfg}).conn()
}

func (r *RedisDB) conn() *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.Port),
		Password: r.cfg.Password,
		DB:       r.cfg.DB,
	})
	r.db = rdb

	return r
}

func (r *RedisDB) Client() *redis.Client {
	return r.db
}
