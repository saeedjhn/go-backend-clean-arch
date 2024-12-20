package user

import "time"

type Data struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
