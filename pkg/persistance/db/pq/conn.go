package pq

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq" // Blank import without comment
)

const _driverName = "postgres"

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

	uri := fmt.Sprintf("host=%s port=%s userentity=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		db.config.Host, db.config.Port, db.config.Username, db.config.Password,
		db.config.Database, db.config.SSLMode)

	db.db, err = sql.Open(_driverName, uri)
	if err != nil {
		return fmt.Errorf("can`t open postgres connection: %w", err)
	}

	// See "Important config" section
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

	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
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

	for _, stmt := range db.statements {
		err := stmt.Close()
		if err != nil {
			return fmt.Errorf("failed to close SQL statement: %v", err)
		}
	}

	return nil
}
