package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
)

func NewMysqlConnection(c mysql.Config) (mysql.DB, error) {
	db := mysql.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseMysqlConnection(db mysql.DB) error {
	if err := db.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close mysql connection: %w", err)
	}

	return nil
}
