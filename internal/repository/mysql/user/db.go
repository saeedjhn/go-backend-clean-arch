package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

type DB struct {
	trc  contract.Tracer
	conn *mysql.Mysql
}

// var _ userservice.Repository = (*DB)(nil)

func New(trc contract.Tracer, conn *mysql.Mysql) *DB {
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

	res, err := r.conn.Conn().Exec(query, u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return entity.User{},
			richerror.New(_opMysqlUserCreate).WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	id, _ := res.LastInsertId()
	u.ID = uint64(id) // #nosec G115 // integer overflow conversion int64 -> uint64

	return u, nil
}

func (r *DB) IsMobileUnique(ctx context.Context, mobile string) (bool, error) {
	_, span := r.trc.Span(ctx, "DB IsMobileUnique")
	span.SetAttributes(map[string]interface{}{
		"db.system":    "MYSQL",  // MYSQL, MARIA, POSTGRES, MONGO
		"db.operation": "SELECT", // SELECT, INSERT, UPDATE, DELETE
	})
	defer span.End()

	var exists bool

	query := "select exists(select 1 from users where mobile = ?)"
	err := r.conn.Conn().
		QueryRow(query, mobile).Scan(&exists)

	span.SetAttribute("db.query", query)

	if err != nil {
		return false,
			richerror.New(_opMysqlUserIsMobileUnique).WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
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
	row := r.conn.Conn().QueryRow(query, mobile)

	span.SetAttribute("db.query", query)

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
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
	row := r.conn.Conn().QueryRow(query, id)

	span.SetAttribute("db.query", query)

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return user, nil
}

func scanUser(scanner Scanner) (entity.User, error) {
	var user entity.User

	err := scanner.Scan(&user.ID, &user.Name, &user.Mobile, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}
