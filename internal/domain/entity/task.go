package entity

import "time"

type Task struct {
	ID          uint64
	UserID      uint64
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
