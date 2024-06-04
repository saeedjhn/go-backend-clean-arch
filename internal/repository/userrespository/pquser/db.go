package pquser

import (
	"database/sql"
	"errors"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
	"log"
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
	const op = message.OpPqUserRegister

	//query := `insert into users(name, mobile, email, password) values(?, ?, ?, ?)` //for mysql
	//query := `insert into users(name, mobile, email, password) values($1, $2, $3, $4)` // dont lastInsertID -  for postgres
	query := `insert into users(name, mobile, email, password) values($1, $2, $3, $4) RETURNING id` // for postgres
	lastInsertID := 0
	err := r.conn.Conn().QueryRow(query, u.Name, u.Mobile, u.Email, u.Password).Scan(&lastInsertID)
	if err != nil {
		return domain.User{},
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	u.ID = uint(lastInsertID)

	return u, nil
}

func (r *DB) IsMobileUnique(mobile string) (bool, error) {
	const op = "mysql.GetUserByPhoneNumber"

	var exists bool
	// query := `insert into users(name, mobile, email, password) values($1, $2, $3, $4) RETURNING id` // for postgres
	//query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE mobile = %s);", mobile)
	//query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE mobile = %s);", mobile)
	query := "SELECT 1 FROM users WHERE mobile = $1"

	err := r.conn.Conn().QueryRow(query, mobile).Scan(&exists)

	//log.Println("**************************** ", op, err.Error())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no row **************************** ", op, err.Error())
			return true, nil
		}

		log.Println("ise **************************** ", op, err.Error())
		return false,
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}
	log.Println("ok **************************** ", op, err.Error())

	return exists, err
}

//func scanUser(scanner mysql.Scanner) (domain.User, error) {
//	var createdAt time.Time
//	var user entity.User
//
//	var roleStr string
//
//	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)
//
//	user.Role = entity.MapToRoleEntity(roleStr)
//
//	return user, err
//}
