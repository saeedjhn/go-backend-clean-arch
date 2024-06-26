package mysqltask

import (
	"database/sql"
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

type DB struct {
	conn mysql.DB
}

func New(conn mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create(t entity.Task) (entity.Task, error) {
	const op = message.OpMysqlTaskCreate

	query := "INSERT INTO tasks (user_id, title, description, status)  values(?, ?, ?, ?)"

	res, err := r.conn.Conn().Exec(query, t.UserID, t.Title, t.Description, t.Status)
	if err != nil {
		return entity.Task{},
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	id, _ := res.LastInsertId()
	t.ID = uint(id)

	return t, nil
}

func (r *DB) IsExistsUser(id uint) (bool, error) {
	const op = message.OpMysqlTaskIsExistsUser
	var exists bool

	err := r.conn.Conn().
		QueryRow(`select exists(select 1 from users where id = ?)`, id).Scan(&exists) // TODO - IsExistsUser

	if err != nil {
		return false,
			richerror.New(op).
				WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (r *DB) GetByID(id uint) (entity.Task, error) {
	const op = message.OpMysqlTaskGetByID

	query := "SELECT * FROM users WHERE id = ?" // TODO - GetByID
	row := r.conn.Conn().QueryRow(query, id)
	task, err := scanTask(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBRecordNotFound).WithKind(kind.KindStatusNotFound)
		}

		return entity.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return task, nil
}

func scanTask(scanner Scanner) (entity.Task, error) {
	var task entity.Task

	err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	return task, err
}

func (r *DB) GetAllByUserID(userID uint) ([]entity.Task, error) {
	const op = message.OpMysqlTaskGetAllByUserID

	query := "SELECT * FROM tasks WHERE user_id = ? ORDER BY id DESC"
	rows, err := r.conn.Conn().Query(query, userID)
	defer rows.Close()

	if err != nil {
		return []entity.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []entity.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *DB) GetAll() ([]entity.Task, error) {
	const op = message.OpMysqlTaskGetAll

	var tasks []entity.Task

	rows, err := r.conn.Conn().Query("SELECT * FROM tasks ORDER BY id DESC ")
	defer rows.Close()

	if err != nil {
		return []entity.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []entity.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
