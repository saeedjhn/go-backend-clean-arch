package migratorpq

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
)

const dialect = "postgres"

type Migrator struct {
	conn       pq.DB
	dialect    string
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(conn pq.DB, absolutePath string) Migrator {
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
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {

	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %w", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}

func (m Migrator) Status() {
	// TODO - add status
}
