package mysqltask

import (
	"database/sql"
	"errors"
	"go-backend-clean-arch/internal/domain"
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

func (r *DB) Create(t domain.Task) (domain.Task, error) {
	const op = message.OpMysqlTaskCreate

	query := "INSERT INTO tasks (user_id, title, description, status)  values(?, ?, ?, ?)"

	res, err := r.conn.Conn().Exec(query, t.UserID, t.Title, t.Description, t.Status)
	if err != nil {
		return domain.Task{},
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

func (r *DB) GetByID(id uint) (domain.Task, error) {
	const op = message.OpMysqlTaskGetByID

	query := "SELECT * FROM users WHERE id = ?" // TODO - GetByID
	row := r.conn.Conn().QueryRow(query, id)
	task, err := scanTask(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBRecordNotFound).WithKind(kind.KindStatusNotFound)
		}

		return domain.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
	}

	return task, nil
}

func scanTask(scanner Scanner) (domain.Task, error) {
	var task domain.Task

	err := scanner.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	return task, err
}

func (r *DB) GetAllByUserID(userID uint) ([]domain.Task, error) {
	const op = message.OpMysqlTaskGetAllByUserID

	query := "SELECT * FROM tasks WHERE user_id = ? ORDER BY id DESC"
	rows, err := r.conn.Conn().Query(query, userID)
	defer rows.Close()

	if err != nil {
		return []domain.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []domain.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *DB) GetAll() ([]domain.Task, error) {
	const op = message.OpMysqlTaskGetAll

	var tasks []domain.Task

	rows, err := r.conn.Conn().Query("SELECT * FROM tasks ORDER BY DESC ")
	defer rows.Close()

	if err != nil {
		return []domain.Task{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []domain.Task{}, richerror.New(op).WithErr(err).
				WithMessage(message.ErrorMsgDBCantScanQueryResult).WithKind(kind.KindStatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
