package redisuser

import (
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
)

type DB struct {
	conn redis.DB
}

func New(conn redis.DB) *DB {
	return &DB{conn: conn}
}

func (d *DB) Set() {
	log.Print("redis-set - IMPL ME")
}
