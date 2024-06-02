package pquser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

type DB struct {
	conn pq.DB
}

func New(conn pq.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Register(u domain.User) (domain.User, error) {
	const op = "postgresqluser.register"

	//query := `insert into users(name, mobile, email, password) values(?, ?, ?, ?)` //for mysql
	//query := `insert into users(name, mobile, email, password) values($1, $2, $3, $4)` // dont lastInsertID -  for postgres
	query := `insert into users(name, mobile, email, password) values($1, $2, $3, $4) RETURNING id` // for postgres
	lastInsertID := 0
	err := r.conn.Conn().QueryRow(query, u.Name, u.Mobile, u.Email, u.Password).Scan(&lastInsertID)
	if err != nil {
		return domain.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsgSomethingWentWrong).
				WithKind(kind.KindStatusInternalServerError)
	}

	u.ID = uint(lastInsertID)

	return u, nil
}
