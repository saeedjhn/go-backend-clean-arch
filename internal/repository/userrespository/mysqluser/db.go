package mysqluser

import (
	"context"
	"database/sql"
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

type DB struct {
	conn *mysql.Mysql
}

//var _ userservice.Repository = (*DB)(nil)

func New(conn *mysql.Mysql) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create(ctx context.Context, u entity.User) (entity.User, error) {
	query := "INSERT INTO users (name, mobile, email, password) values(?, ?, ?, ?)"

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
	var exists bool

	err := r.conn.Conn().
		QueryRow(`select exists(select 1 from users where mobile = ?)`, mobile).Scan(&exists)

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
	query := "SELECT * FROM users WHERE mobile = ?"
	row := r.conn.Conn().QueryRow(query, mobile)

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
	query := "SELECT * FROM users WHERE id = ?"
	row := r.conn.Conn().QueryRow(query, id)

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
