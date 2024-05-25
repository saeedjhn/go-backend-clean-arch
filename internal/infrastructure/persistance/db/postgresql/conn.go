package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DB interface {
	Conn() *sql.DB
}

type PostgresqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *PostgresqlDB {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		config.Host, config.Port, config.Username, config.Password,
		config.Database, config.SSLMode)

	db, err := sql.Open("postgres", cnn)
	if err != nil {
		log.Fatalf("can`t open postgres connection: %v", err)
	}

	// See "Important config" section
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLiftTime * time.Second)

	return &PostgresqlDB{config: config, db: db}
}

func (m *PostgresqlDB) Conn() *sql.DB {
	return m.db
}
