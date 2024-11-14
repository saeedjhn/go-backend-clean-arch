package bootstrap

import (
	"fmt"
	pq2 "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

func NewPostgresConnection(c pq2.Config) (pq2.DB, error) {
	db := pq2.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func ClosePostgresConnection(db pq2.DB) error {
	if err := db.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close pq connection: %w", err)
	}

	return nil
}
