package mysqluser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

type DB struct {
	conn mysql.DB
}

func New(conn mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Register(u domain.User) (domain.User, error) {
	const op = message.OpMysqlUserRegister

	query := `insert into users(name, mobile, email, password) values(?, ?, ?, ?)` //for mysql

	res, err := r.conn.Conn().Exec(query, u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return domain.User{},
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
