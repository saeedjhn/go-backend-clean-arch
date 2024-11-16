package pquser

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

type DB struct {
	conn *pq.Postgres
}

func New(conn *pq.Postgres) *DB {
	return &DB{
		conn: conn,
	}
}
