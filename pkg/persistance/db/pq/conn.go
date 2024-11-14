package pq

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Blank import without comment
)

const _driverName = "postgres"

type DB interface {
	Conn() *sql.DB
}

type Postgres struct {
	config Config
	db     *sql.DB
}

var _ DB = (*Postgres)(nil)

func New(config Config) *Postgres {
	return &Postgres{config: config}
}

func (p *Postgres) ConnectTo() error {
	var err error

	uri := fmt.Sprintf("host=%s port=%s userentity=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		p.config.Host, p.config.Port, p.config.Username, p.config.Password,
		p.config.Database, p.config.SSLMode)

	p.db, err = sql.Open(_driverName, uri)
	if err != nil {
		return fmt.Errorf("can`t open postgres connection: %w", err)
	}

	// See "Important config" section
	p.db.SetMaxIdleConns(p.config.MaxIdleConns)
	p.db.SetMaxOpenConns(p.config.MaxOpenConns)
	p.db.SetConnMaxLifetime(p.config.ConnMaxLiftTime)

	return nil
}

func (p *Postgres) Conn() *sql.DB {
	return p.db
}
