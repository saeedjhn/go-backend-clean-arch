package entity

import "time"

type User struct {
	ID        uint
	Name      string
	Mobile    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
