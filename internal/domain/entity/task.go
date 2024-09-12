package entity

import "time"

type Task struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
