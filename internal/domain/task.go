package domain

import "time"

type Task struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
