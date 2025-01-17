package migrations

import (
	"flag"
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql/migratormysql"
)

const (
	mysqlDIR = "./internal/repository/migrations/mysql"
	pqDir    = "./internal/repository/migrations/pq"
)

const (
	ArgUP       = "up"
	ArgDown     = "down"
	ArgRollback = "rollback"
)

func Start(app *bootstrap.Application) {
	// Define flags
	up := flag.Bool(ArgUP, false, "Example flag to demonstrate passing --up")
	down := flag.Bool(ArgDown, false, "Example flag to demonstrate passing --down")
	rollback := flag.Bool(ArgRollback, true, "Example flag to demonstrate passing --down")

	// Parse flags
	flag.Parse()

	log.Println("Startup migrations...")

	if err := HandleFlag(app, *up, *down, *rollback); err != nil {
		log.Fatal(err)
	}

	log.Println("finished migrations.")
}

func HandleFlag(app *bootstrap.Application, up, down, rollback bool) error {
	switch {
	case up:
		return Up(app)
	case down:
		return Down(app)
	case rollback:
		return Rollback(app)
	default:
		return Rollback(app)
	}
}

func Up(app *bootstrap.Application) error {
	// Mysql
	migratorMysql := migratormysql.New(app.DB.MySQL, migratormysql.Config{
		MigrationPath:   mysqlDIR,
		MigrationDBName: "gorp_migrations",
	})
	if err := migratorMysql.Up(); err != nil {
		return err
	}

	return nil

	// Pq
	// migratorPq := migratorpq.New(app.Postgres, pqDIR)
	// migratorPq.Up()

	// Etc
}

func Down(app *bootstrap.Application) error {
	// Mysql
	migratorMysql := migratormysql.New(app.DB.MySQL, migratormysql.Config{
		MigrationPath:   mysqlDIR,
		MigrationDBName: "gorp_migrations",
	})
	if err := migratorMysql.Down(); err != nil {
		return err
	}

	return nil

	// Pq
	// migratorPq := migratorpq.New(app.Postgres, pqDIR)
	// migratorPq.Down()

	// Etc
}

func Rollback(app *bootstrap.Application) error {
	// Mysql
	migratorMysql := migratormysql.New(app.DB.MySQL, migratormysql.Config{
		MigrationPath:   mysqlDIR,
		MigrationDBName: "gorp_migrations",
	})
	if err := migratorMysql.Down(); err != nil {
		return err
	}

	if err := migratorMysql.Up(); err != nil {
		return err
	}

	return nil

	// Pq
	// migratorPq := migratorpq.New(app.Postgres, pqDIR)
	// migratorPq.Down()
	// migratorPq.Up()

	// etc
}
