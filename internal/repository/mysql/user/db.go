package user

import (
	"context"
	"database/sql"
	"errors"

	mysqlrepo "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"
)

type DB struct {
	trc  contract.Tracer
	conn *mysql.DB
}

// var _ userservice.Repository = (*DB)(nil)

func New(trc contract.Tracer, conn *mysql.DB) *DB {
	return &DB{
		trc:  trc,
		conn: conn,
	}
}

func (r *DB) Create(ctx context.Context, u entity.User) (entity.User, error) {
	_, span := r.trc.Span(ctx, "DB Create")
	span.SetAttributes(map[string]interface{}{
		"db.system":    "MYSQL",  // MYSQL, MARIA, POSTGRES, MONGO
		"db.operation": "INSERT", // SELECT, INSERT, UPDATE, DELETE
	})
	defer span.End()

	query := "INSERT INTO users (name, mobile, email, password) values(?, ?, ?, ?)"

	span.SetAttribute("db.query", query)

	stmt, err := r.conn.PrepareStatement(ctx, uint(mysqlrepo.StatementKeyUserCreate), query)
	if err != nil {
		return entity.User{}, richerror.New(_opMysqlUserCreate).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	res, err := stmt.ExecContext(ctx, u.Name, u.Mobile, u.Email, u.Password)
	// res, err := r.conn.Conn().Exec(query, u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return entity.User{},
			richerror.New(_opMysqlUserCreate).WithErr(err).
				WithMessage(msg.ErrorMsg500InternalServerError).
				WithKind(richerror.KindStatusInternalServerError)
	}

	id, _ := res.LastInsertId()
	u.ID = types.ID(id) // #nosec G115 // integer overflow conversion int64 -> uint64

	return u, nil
}

func (r *DB) IsExistsByMobile(ctx context.Context, mobile string) (bool, error) {
	_, span := r.trc.Span(ctx, "DB IsExistsByMobile")
	span.SetAttributes(map[string]interface{}{
		"db.system":    "MYSQL",  // MYSQL, MARIA, POSTGRES, MONGO
		"db.operation": "SELECT", // SELECT, INSERT, UPDATE, DELETE
	})
	defer span.End()

	var exists bool

	query := "select exists(select 1 from users where mobile = ?)"

	span.SetAttribute("db.query", query)

	stmt, err := r.conn.PrepareStatement(ctx, uint(mysqlrepo.StatementKeyUserIsExistsByMobile), query)
	if err != nil {
		return false, richerror.New(_opMysqlUserIsExistsByMobile).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	err = stmt.QueryRowContext(ctx, mobile).Scan(&exists)
	// err := r.conn.Conn().QueryRow(query, mobile).Scan(&exists)
	if err != nil {
		return false,
			richerror.New(_opMysqlUserIsExistsByMobile).WithErr(err).
				WithMessage(msg.ErrorMsg500InternalServerError).
				WithKind(richerror.KindStatusInternalServerError)
	}

	if !exists {
		return true, nil
	}

	return false, nil
}

func (r *DB) GetByMobile(ctx context.Context, mobile string) (entity.User, error) {
	_, span := r.trc.Span(ctx, "DB GetByMobile")
	span.SetAttributes(map[string]interface{}{
		"db.system":    "MYSQL",  // MYSQL, MARIA, POSTGRES, MONGO
		"db.operation": "SELECT", // SELECT, INSERT, UPDATE, DELETE
	})
	defer span.End()

	query := "SELECT * FROM users WHERE mobile = ?"

	span.SetAttribute("db.query", query)

	// row := r.conn.Conn().QueryRow(query, mobile)
	stmt, err := r.conn.PrepareStatement(ctx, uint(mysqlrepo.StatementKeyUserGetByMobile), query)
	if err != nil {
		return entity.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	row := stmt.QueryRowContext(ctx, mobile)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
				WithMessage(errMsgDBRecordNotFound).
				WithKind(richerror.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
			WithMessage(errMsgDBCantScanQueryResult).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user, nil
}

func (r *DB) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	_, span := r.trc.Span(ctx, "DB GetByID")
	span.SetAttributes(map[string]interface{}{
		"db.system":    "MYSQL",  // MYSQL, MARIA, POSTGRES, MONGO
		"db.operation": "SELECT", // SELECT, INSERT, UPDATE, DELETE
	})
	defer span.End()

	query := "SELECT * FROM users WHERE id = ?"

	// row := r.conn.Conn().QueryRow(query, id)
	stmt, err := r.conn.PrepareStatement(ctx, uint(mysqlrepo.StatementKeyUserGetByID), query)
	if err != nil {
		return entity.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	span.SetAttribute("db.query", query)

	row := stmt.QueryRowContext(ctx, id)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
				WithMessage(errMsgDBRecordNotFound).
				WithKind(richerror.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
			WithMessage(errMsgDBCantScanQueryResult).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user, nil
}

func scanUser(scanner Scanner) (entity.User, error) {
	var user entity.User

	err := scanner.Scan(&user.ID, &user.Name, &user.Mobile, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	// Convert something...

	return user, err
}
