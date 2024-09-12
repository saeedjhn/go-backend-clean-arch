package pq

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const driverName = "postgres"

type DB interface {
	Conn() *sql.DB
}

type PostgresDB struct {
	config Config
	db     *sql.DB
}

var _ DB = (*PostgresDB)(nil)

func New(config Config) *PostgresDB {
	cnn := fmt.Sprintf("host=%s port=%s userentity=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		config.Host, config.Port, config.Username, config.Password,
		config.Database, config.SSLMode)

	db, err := sql.Open(driverName, cnn)
	if err != nil {
		log.Fatalf("can`t open postgres connection: %v", err)
	}

	// See "Important config" section
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLiftTime * time.Second)

	return &PostgresDB{config: config, db: db}
}

func (m *PostgresDB) Conn() *sql.DB {
	return m.db
}
