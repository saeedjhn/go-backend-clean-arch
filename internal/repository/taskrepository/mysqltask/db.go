package mysqltask

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

//var _ taskservice.Repository = (*DB)(nil)

func New(conn *mysql.Mysql) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create(ctx context.Context, t entity.Task) (entity.Task, error) {
	query := "INSERT INTO tasks (user_id, title, description, status)  values(?, ?, ?, ?)"

	res, err := r.conn.Conn().Exec(query, t.UserID, t.Title, t.Description, t.Status)
	if err != nil {
		return entity.Task{},
			richerror.New(_opMysqlTaskCreate).WithErr(err).
				WithMessage(message.ErrorMsg500InternalServerError).
				WithKind(kind.KindStatusInternalServerError)
	}

	id, _ := res.LastInsertId()
	t.ID = uint64(id) // #nosec G115 // integer overflow conversion int64 -> uint64

	return t, nil
}

func (r *DB) IsExistsUser(ctx context.Context, id uint64) (bool, error) {
	var exists bool

	err := r.conn.Conn().
		QueryRow(`select exists(select 1 from users where id = ?)`, id).Scan(&exists) // TODO - IsExistsUser

	if err != nil {
		return false, richerror.New(_opMysqlTaskIsExistsUser).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (r *DB) GetByID(ctx context.Context, id uint64) (entity.Task, error) {
	query := "SELECT * FROM users WHERE id = ?" // TODO - GetByID
	row := r.conn.Conn().QueryRow(query, id)

	task, err := scanTask(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return entity.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return task, nil
}

func (r *DB) GetAllByUserID(ctx context.Context, userID uint64) ([]entity.Task, error) {
	query := "SELECT * FROM tasks WHERE user_id = ? ORDER BY id DESC"

	rows, err := r.conn.Conn().Query(query, userID)
	if err != nil || rows.Err() != nil {
		return nil, richerror.New(_opMysqlTaskGetAllByUserID).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	tasks, err := scanTasks(rows)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return nil, richerror.New(_opMysqlTaskGetAllByUserID).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return tasks, nil
}

func (r *DB) GetAll(ctx context.Context) ([]entity.Task, error) {
	rows, err := r.conn.Conn().Query("SELECT * FROM tasks ORDER BY id DESC ")
	if err != nil || rows.Err() != nil {
		return nil, richerror.New(_opMysqlTaskGetAll).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	tasks, err := scanTasks(rows)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return nil, richerror.New(_opMysqlTaskGetAll).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
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
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
