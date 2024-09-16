package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const driverName = "mysql"

type DB interface {
	Conn() *sql.DB
}

type MySqlDB struct {
	config Config
	db     *sql.DB
	err    error
}

var _ DB = (*MySqlDB)(nil)

func New(config Config) *MySqlDB {
	return &MySqlDB{config: config}
}

func (m *MySqlDB) ConnectTo() error {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.Database)

	m.db, m.err = sql.Open(driverName, conn)
	if m.err != nil {
		//log.Fatalf("can't open mysql db: %v", err)
		return fmt.Errorf("can`t open mysql db: %w", m.err)
	}

	// See "Important settings" section.
	m.db.SetMaxIdleConns(m.config.MaxIdleConns)
	m.db.SetMaxOpenConns(m.config.MaxOpenConns)
	m.db.SetConnMaxLifetime(m.config.ConnMaxLiftTime * time.Second)

	return nil
}

func (m *MySqlDB) Conn() *sql.DB {
	return m.db
}

func (m *MySqlDB) Error() error {
	return m.err
}
