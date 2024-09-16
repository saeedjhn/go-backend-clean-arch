package bootstrap

import (
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
)

func NewPostgresConnection(config pq.Config) (pq.DB, error) {
	//return pq.New(config)

	myDB := pq.New(config)

	if err := myDB.ConnectTo(); err != nil {
		return nil, err
	}

	return myDB, nil
}

func ClosePostgresConnection(postgresDB pq.DB) error {
	if err := postgresDB.Conn().Close(); err != nil {
		return fmt.Errorf("don`t close pq connection: %w", err)
	}

	return nil
}
