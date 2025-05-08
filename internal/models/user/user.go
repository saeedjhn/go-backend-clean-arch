package user

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type User struct {
	ID        types.ID
	Name      string
	Mobile    string
	Email     string
	Password  string
	Events    []contract.DomainEvent
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (u *User) AddEvent(event contract.DomainEvent) {
// 	u.Events = append(u.Events, event)
// }
//

func (u *User) AddEvents(events ...contract.DomainEvent) {
	u.Events = append(u.Events, events...)
}

func (u *User) GetEvents() []contract.DomainEvent {
	return u.Events
}

func (u *User) ClearEvents() {
	u.Events = nil
}

func (u *User) PullEvents() []contract.DomainEvent {
	events := u.Events
	u.Events = nil

	return events
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
