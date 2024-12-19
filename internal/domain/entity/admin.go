package entity

type Admin struct {
	ID          uint64
	FirstName   string
	LastName    string
	Password    string
	PhoneNumber string
	Role        AdminRole
	Description string
	Email       string
	Gender      Gender
	Status      AdminStatus
}
