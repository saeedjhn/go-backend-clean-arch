package mysqltask

import (
	"database/sql"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

type DB struct {
	conn mysql.DB
}

var _ taskservice.Repository = (*DB)(nil)

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

func (r *DB) GetAllByUserID(userID uint) ([]entity.Task, error) {
	const op = message.OpMysqlTaskGetAllByUserID

	query := "SELECT * FROM tasks WHERE user_id = ? ORDER BY id DESC"
	rows, err := r.conn.Conn().Query(query, userID)
	defer rows.Close()

	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	tasks, err := scanTasks(rows)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return tasks, nil
}

func (r *DB) GetAll() ([]entity.Task, error) {
	const op = message.OpMysqlTaskGetAll

	rows, err := r.conn.Conn().Query("SELECT * FROM tasks ORDER BY id DESC ")
	defer rows.Close()

	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	tasks, err := scanTasks(rows)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return tasks, nil
}

func scanTask(scanner RowScanner) (entity.Task, error) {
	var (
		task   entity.Task
		status string
	)

	err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &status, &task.CreatedAt, &task.UpdatedAt)

	task.Status = entity.MapToTaskStatus(status)

	return task, err
}

func scanTasks(scanner RowsScanner) ([]entity.Task, error) {
	var (
		tasks  []entity.Task
		task   entity.Task
		status string
	)

	for scanner.Next() {
		err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &status, &task.CreatedAt, &task.UpdatedAt)
		task.Status = entity.MapToTaskStatus(status)

		if err != nil {
			return nil, nil
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
