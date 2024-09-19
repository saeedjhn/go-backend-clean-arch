package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
)

func NewMysqlConnection(config mysql.Config) (mysql.DB, error) {
	myDB := mysql.New(config)

	if err := myDB.ConnectTo(); err != nil {
		return nil, err
	}

	return myDB, nil
}

func CloseMysqlConnection(mysqlDB mysql.DB) error {
	if err := mysqlDB.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close mysql connection: %w", err)
	}

	return nil
}
