package main

import (
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/pq/migratorpq"
)

func main() {
	app := bootstrap.App(configs.Development)

	// PostgresQL
	pqDir := "./internal/repository/migrations/pqmigration"
	migratorPq := migratorpq.New(app.PostgresDB, pqDir)
	migratorPq.Down()
	migratorPq.Up()

	// MysQL
	//mysqlDir := "./internal/repository/migrations/mysqlmigration"
	//migratorMysql := migratormysql.New(app.MysqlDB, mysqlDir)
	//migratorMysql.Down()
	//migratorMysql.Up()

	// etc
}
