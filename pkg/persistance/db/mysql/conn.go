package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Blank import without comment
)

const _driverName = "mysql"

type DB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *DB {
	return &DB{config: config}
}

func (m *DB) ConnectTo() error {
	var err error
	uri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",

		m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.Database)

	m.db, err = sql.Open(_driverName, uri)
	if err != nil {
		return fmt.Errorf("can`t open mysql db: %w", err)
	}

	// See "Important settings" section.
	m.db.SetMaxIdleConns(m.config.MaxIdleConns)
	m.db.SetMaxOpenConns(m.config.MaxOpenConns)
	m.db.SetConnMaxLifetime(m.config.ConnMaxLiftTime)

	return nil
}

func (m *DB) Conn() *sql.DB {
	return m.db
}
