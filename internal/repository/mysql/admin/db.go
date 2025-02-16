package admin

import "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"

type DB struct {
	conn *mysql.DB
}

// var _ adminusecase.Repository = (*DB)(nil)

func New(conn *mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}
