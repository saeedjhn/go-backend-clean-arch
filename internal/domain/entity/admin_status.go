package entity

type AdminStatus string

const (
	AdminActiveStatus   = AdminStatus("active")
	AdminInactiveStatus = AdminStatus("inactive")
)

var _adminStatusStrings = map[AdminStatus]string{
	AdminActiveStatus:   "active",
	AdminInactiveStatus: "inactive",
}

func (a AdminStatus) IsValid() bool {
	_, ok := _adminStatusStrings[a]

	return ok
}
