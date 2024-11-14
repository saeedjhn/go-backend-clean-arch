package pquser

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

type DB struct {
	conn pq.DB
}

func New(conn pq.DB) *DB {
	return &DB{
		conn: conn,
	}
}
