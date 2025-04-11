package models

// Permission describe the rules for accessing resources.
type Permission struct {
	Allow RWXD // Allowed actions
	Deny  RWXD // Denied actions
}

type Action string

const (
	ReadAction    Action = "r"
	WriteAction   Action = "w"
	ExecuteAction Action = "x"
	DeleteAction  Action = "d"
)

// RWXD specifies a set of permissions for resources.
type RWXD struct {
	R bool // ReadAction permission
	W bool // WriteAction permission
	X bool // ExecuteAction permission
	D bool // DeleteAction permission
}
