package outboxevent

import (
	"context"

	mysqlrepo "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) DB {
	return DB{conn: conn}
}

func (d DB) Create(ctx context.Context, oe models.OutboxEvent) (types.ID, error) {
	var resultID types.ID

	query := `
		INSERT INTO outbox_events 
		(topic, payload, is_published, retried_count, last_retried_at, published_at) 
		VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := d.conn.PrepareStatement(ctx, uint(mysqlrepo.StatementKeyOutboxCreate), query)
	if err != nil {
		return 0, richerror.New(_opMysqlOutboxCreate).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	result, err := stmt.ExecContext(
		ctx,
		oe.Type,
		oe.Payload,
		oe.IsPublished,
		0,   // retried_count
		nil, // last_retried_at as NULL
		nil, // published_at as NULL
	)
	if err != nil {
		return 0, richerror.New(_opMysqlOutboxCreate).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, richerror.New(_opMysqlOutboxCreate).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	resultID = types.ID(lastID) // #nosec G115 // integer overflow conversion int64 -> uint64

	return resultID, nil
}
