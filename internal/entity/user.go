package entity

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

type User struct {
	ID        types.ID
	Name      string
	Mobile    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (u UserInfo) ToUserInfoDTO() userdto.UserInfo {
//	return userdto.UserInfo{
//		ID:        u.ID,
//		Name:      u.Name,
//		Mobile:    u.Mobile,
//		Email:     u.Email,
//		CreatedAt: u.CreatedAt,
//		UpdatedAt: u.UpdatedAt,
//	}
// }
