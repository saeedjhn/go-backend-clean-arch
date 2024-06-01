package postgresqluser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

type DB struct {
	conn postgresql.DB
}

func New(conn postgresql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Register(u domain.User) (domain.User, error) {
	const op = "postgresqluser.register"

	res, err := r.conn.Conn().Exec(`insert into users(name, mobile, email, password) values(?, ?, ?, ?)`,
		u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return domain.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsgSomethingWentWrong).
				WithKind(kind.KindStatusInternalServerError)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}
