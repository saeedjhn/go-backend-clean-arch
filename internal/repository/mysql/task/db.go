package task

import (
	"context"
	"database/sql"
	"errors"
	entity2 "github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

type DB struct {
	conn *mysql.Mysql
}

// var _ taskservice.Repository = (*DB)(nil)

func New(conn *mysql.Mysql) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create(ctx context.Context, t entity2.Task) (entity2.Task, error) {
	query := "INSERT INTO tasks (user_id, title, description, status)  values(?, ?, ?, ?)"

	res, err := r.conn.Conn().Exec(query, t.UserID, t.Title, t.Description, t.Status)
	if err != nil {
		log.Println(err)

		return entity2.Task{},
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

func (r *DB) GetByID(ctx context.Context, id uint64) (entity2.Task, error) {
	query := "SELECT * FROM users WHERE id = ?" // TODO - GetByID
	row := r.conn.Conn().QueryRow(query, id)

	task, err := scanTask(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity2.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return entity2.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return task, nil
}

func (r *DB) GetAllByUserID(ctx context.Context, userID uint64) ([]entity2.Task, error) {
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
			return []entity2.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return nil, richerror.New(_opMysqlTaskGetAllByUserID).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return tasks, nil
}

func (r *DB) GetAll(ctx context.Context) ([]entity2.Task, error) {
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
			return []entity2.Task{}, richerror.New(_opMysqlTaskGetByID).WithErr(err).
				WithMessage(_errMsgDBRecordNotFound).
				WithKind(kind.KindStatusNotFound)
		}

		return nil, richerror.New(_opMysqlTaskGetAll).WithErr(err).
			WithMessage(_errMsgDBCantScanQueryResult).
			WithKind(kind.KindStatusInternalServerError)
	}

	return tasks, nil
}

func scanTask(scanner RowScanner) (entity2.Task, error) {
	var (
		task   entity2.Task
		status string
	)

	err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &status, &task.CreatedAt, &task.UpdatedAt)

	task.Status = entity2.TaskStatus(status)

	return task, err
}

func scanTasks(scanner RowsScanner) ([]entity2.Task, error) {
	var (
		tasks  []entity2.Task
		task   entity2.Task
		status string
	)

	for scanner.Next() {
		err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &status, &task.CreatedAt, &task.UpdatedAt)
		task.Status = entity2.TaskStatus(status)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
