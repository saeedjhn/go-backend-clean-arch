package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type MysqlConnection struct {
	cfg Config
	db  *sql.DB
}

func New(config Config) *MysqlConnection {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatalf("can't open mysql db: %v", err)
	}

	// See "Important settings" section.
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLiftTime * time.Second)

	return &MysqlConnection{cfg: config, db: db}
}

func (m *MysqlConnection) Conn() *sql.DB {
	return m.db
}
