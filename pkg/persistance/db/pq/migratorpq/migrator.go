package migratorpq

import (
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

const dialect = "postgres"

type Migrator struct {
	conn       *pq.Postgres
	dialect    string
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(conn *pq.Postgres, absolutePath string) Migrator {
	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: absolutePath,
	}

	return Migrator{conn: conn, dialect: dialect, migrations: migrations}
}

func (m Migrator) Up() {
	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %w", err))
	}
	log.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %w", err))
	}
	log.Printf("Rollbacked %d migrations!\n", n)
}

func (m Migrator) Status() {
	// TODO - add status
}
