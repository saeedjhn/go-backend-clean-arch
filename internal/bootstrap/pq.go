package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

func NewPostgresConnection(c pq.Config) (*pq.Postgres, error) {
	db := pq.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func ClosePostgresConnection(db *pq.Postgres) error {
	if err := db.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close pq connection: %w", err)
	}

	return nil
}
