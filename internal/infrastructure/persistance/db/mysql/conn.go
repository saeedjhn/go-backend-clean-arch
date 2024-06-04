package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const driverName = "mysql"

type DB interface {
	Conn() *sql.DB
}

type MySqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySqlDB {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open(driverName, conn)
	if err != nil {
		log.Fatalf("can't open mysql db: %v", err)
	}

	// See "Important settings" section.
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLiftTime * time.Second)

	return &MySqlDB{config: config, db: db}
}

func (m *MySqlDB) Conn() *sql.DB {
	return m.db
}
