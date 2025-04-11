package admin

type Status string

const (
	ActiveStatus   = Status("active")
	InactiveStatus = Status("inactive")
)

var _adminStatusStrings = map[Status]string{ //nolint:gochecknoglobals // nothing
	ActiveStatus:   "active",
	InactiveStatus: "inactive",
}

func (a Status) IsValid() bool {
	_, ok := _adminStatusStrings[a]

	return ok
}
