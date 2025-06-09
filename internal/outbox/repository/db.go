package repository

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) DB {
	return DB{conn: conn}
}

func (d DB) UpdatePublished(_ context.Context, _ []types.ID, _ time.Time) error {
	// TODO implement me
	panic("implement me")
}

func (d DB) UpdateUnpublished(_ context.Context, _ []types.ID, _ time.Time) error {
	// TODO implement me
	panic("implement me")
}

func (d DB) UnpublishedCount(_ context.Context, _ int64) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (d DB) GetUnPublished(_ context.Context, _, _, _ int) ([]models.OutboxEvent, error) {
	// TODO implement me
	panic("implement me")
}
