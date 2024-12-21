package entity

// Permissions describe the rules for accessing resources.
type Permissions struct {
	Allow RWXD // Allowed actions
	Deny  RWXD // Denied actions
}
