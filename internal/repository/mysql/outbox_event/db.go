package outboxevent

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) DB {
	return DB{conn: conn}
}

func (d DB) Create(_ context.Context, _ models.OutboxEvent) error {
	// TODO implement me
	panic("implement me")
}
