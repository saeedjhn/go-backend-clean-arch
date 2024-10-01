package entity

import (
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Mobile    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (u User) ToUserInfoDTO() userdto.UserInfo {
//	return userdto.UserInfo{
//		ID:        u.ID,
//		Name:      u.Name,
//		Mobile:    u.Mobile,
//		Email:     u.Email,
//		CreatedAt: u.CreatedAt,
//		UpdatedAt: u.UpdatedAt,
//	}
// }
