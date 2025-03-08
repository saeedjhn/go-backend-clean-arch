package admin

import "github.com/saeedjhn/go-domain-driven-design/pkg/persistance/db/mysql"

type DB struct {
	conn *mysql.DB
}

// var _ adminusecase.Repository = (*DB)(nil)

func New(conn *mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}
