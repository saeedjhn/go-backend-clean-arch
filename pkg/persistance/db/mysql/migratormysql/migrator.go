package migratormysql

import (
	"fmt"
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

const dialect = "mysql"

type Config struct {
	MigrationPath   string
	MigrationDBName string
}

type Migrator struct {
	config     Config
	conn       *mysql.Mysql
	dialect    string
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(conn *mysql.Mysql, config Config) Migrator {
	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: config.MigrationPath,
	}

	return Migrator{
		conn:       conn,
		dialect:    dialect,
		migrations: migrations,
		config:     config,
	}
}

func (m Migrator) Up() error {
	migrate.SetTable(m.config.MigrationDBName)

	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("can't apply migrations: %w", err)
	}

	log.Printf("Applied %d migrations!\n", n)

	return nil
}

func (m Migrator) Down() error {
	migrate.SetTable(m.config.MigrationDBName)

	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Down)
	if err != nil {
		return fmt.Errorf("can't rollback migrations: %w", err)
	}

	log.Printf("Rollbacked %d migrations!\n", n)

	return nil
}

func (m Migrator) Status() {
	// TODO - add status
	panic("IMPLEMENT ME")
}
