package sms

type Status string

const (
	Active   = Status("active")
	Inactive = Status("inactive")
)

var statusStrings = map[Status]string{ //nolint:gochecknoglobals // nothing
	Active:   "active",
	Inactive: "inactive",
}

func (a Status) IsValid() bool {
	_, ok := statusStrings[a]

	return ok
}
