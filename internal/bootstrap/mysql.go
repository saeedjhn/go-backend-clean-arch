package bootstrap

import (
	"fmt"
	mysql2 "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

func NewMysqlConnection(c mysql2.Config) (mysql2.DB, error) {
	db := mysql2.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseMysqlConnection(db mysql2.DB) error {
	if err := db.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close mysql connection: %w", err)
	}

	return nil
}
