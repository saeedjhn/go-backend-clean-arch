package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/mysql"
	"log"
)

func NewMysqlDB(config mysql.Config) mysql.DB {
	return mysql.New(config)
}

func CloseMysqlDB(mysqlDB mysql.DB) {
	err := mysqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close mysql connection: %s", err.Error())
	}
}
