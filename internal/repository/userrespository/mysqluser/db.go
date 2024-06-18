package mysqluser

import (
	"database/sql"
	"errors"
	"go-backend-clean-arch/internal/domain/entity"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

type DB struct {
	conn mysql.DB
}

func New(conn mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create(u entity.User) (entity.User, error) {
	const op = message.OpMysqlUserCreate

	query := "INSERT INTO users (name, mobile, email, password) values(?, ?, ?, ?)"

	res, err := r.conn.Conn().Exec(query, u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return entity.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (r *DB) IsMobileUnique(mobile string) (bool, error) {
	const op = message.OpMysqlUserIsMobileUnique
	var exists bool

	err := r.conn.Conn().
		QueryRow(`select exists(select 1 from users where mobile = ?)`, mobile).Scan(&exists)

	if err != nil {
		return false,
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	if !exists {
		return true, nil
	}

	return false, nil
}

func (r *DB) GetByMobile(mobile string) (entity.User, error) {
	const op = message.OpMysqlUserGetByMobile

	query := "SELECT * FROM users WHERE mobile = ?"
	row := r.conn.Conn().QueryRow(query, mobile)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBRecordNotFound).WithKind(kind.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return user, nil
}

func (r *DB) GetByID(id uint) (entity.User, error) {
	const op = message.OpMysqlUserGetByID

	query := "SELECT * FROM users WHERE id = ?"
	row := r.conn.Conn().QueryRow(query, id)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBRecordNotFound).WithKind(kind.KindStatusNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return user, nil
}

func scanUser(scanner Scanner) (entity.User, error) {
	var user entity.User

	err := scanner.Scan(&user.ID, &user.Name, &user.Mobile, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}
