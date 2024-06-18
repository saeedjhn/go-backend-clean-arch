package userdto

import "time"

type UserInfo struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
