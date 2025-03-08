package task

import (
	"context"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
	taskusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/task"
	"github.com/saeedjhn/go-domain-driven-design/pkg/persistance/db/mysql"
)

type DB struct {
	conn *mysql.DB
}

func (d *DB) Create(_ context.Context, _ entity.Task) (entity.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DB) GetByID(_ context.Context, _ uint64) (entity.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DB) GetAllByUserID(_ context.Context, _ uint64) ([]entity.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DB) IsExistsUser(_ context.Context, _ uint64) (bool, error) {
	// TODO implement me
	panic("implement me")
}

var _ taskusecase.Repository = (*DB)(nil)

func New(conn *mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}
