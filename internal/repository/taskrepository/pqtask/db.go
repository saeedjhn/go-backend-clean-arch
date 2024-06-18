package pqtask

import (
	"go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
)

type DB struct {
	conn pq.DB
}

func New(conn pq.DB) *DB {
	return &DB{
		conn: conn,
	}
}
