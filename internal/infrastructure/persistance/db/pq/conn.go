package pq

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Blank import without comment
)

const driverName = "postgres"

type DB interface {
	Conn() *sql.DB
}

type Pq struct {
	config Config
	db     *sql.DB
	err    error
}

var _ DB = (*Pq)(nil)

func New(config Config) *Pq {
	return &Pq{config: config}
}

func (p *Pq) ConnectTo() error {
	cnn := fmt.Sprintf("host=%s port=%s userentity=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		p.config.Host, p.config.Port, p.config.Username, p.config.Password,
		p.config.Database, p.config.SSLMode)

	p.db, p.err = sql.Open(driverName, cnn)
	if p.err != nil {
		return fmt.Errorf("can`t open postgres connection: %w", p.err)
	}

	// See "Important config" section
	p.db.SetMaxIdleConns(p.config.MaxIdleConns)
	p.db.SetMaxOpenConns(p.config.MaxOpenConns)
	p.db.SetConnMaxLifetime(p.config.ConnMaxLiftTime)

	return nil
}

func (p *Pq) Conn() *sql.DB {
	return p.db
}

func (p *Pq) Error() error {
	return p.err
}
