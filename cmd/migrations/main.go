package main

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq/migratorpq"
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
