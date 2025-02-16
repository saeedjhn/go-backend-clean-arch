package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql" // Blank import without comment
)

const _driverName = "mysql"

type DB struct {
	config     Config
	db         *sql.DB
	mu         sync.Mutex
	statements map[uint]*sql.Stmt
}

func New(config Config) *DB {
	return &DB{config: config, statements: make(map[uint]*sql.Stmt)}
}

func (db *DB) ConnectTo() error {
	var err error
	uri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",

		db.config.Username, db.config.Password, db.config.Host, db.config.Port, db.config.Database)

	db.db, err = sql.Open(_driverName, uri)
	if err != nil {
		return fmt.Errorf("can`t open mysql db: %w", err)
	}

	// See "Important settings" section.
	db.db.SetMaxIdleConns(db.config.MaxIdleConns)
	db.db.SetMaxOpenConns(db.config.MaxOpenConns)
	db.db.SetConnMaxLifetime(db.config.ConnMaxLiftTime)

	return nil
}

func (db *DB) PrepareStatement(ctx context.Context, key uint, query string) (*sql.Stmt, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if stmt, ok := db.statements[key]; ok {
		return stmt, nil
	}

	stmt, err := db.db.PrepareContext(ctx, query) //nolint:sqlclosecheck // nothing
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %w", err)
	}

	db.statements[key] = stmt

	return stmt, nil
}

func (db *DB) Conn() *sql.DB {
	return db.db
}

func (db *DB) CloseStatements() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	var errs []error
	for k, stmt := range db.statements {
		if err := stmt.Close(); err != nil {
			errs = append(errs, err)
		}

		delete(db.statements, k)
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to close MySQL statements: %v", errs)
	}

	return nil
}
