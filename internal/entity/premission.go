package entity

// Permissions describe the rules for accessing resources.
type Permissions struct {
	Allow RWXD // Allowed actions
	Deny  RWXD // Denied actions
}

// RWXD specifies a set of permissions for resources.
type RWXD struct {
	R bool // Read permission
	W bool // Write permission
	X bool // Execute permission
	D bool // Delete permission
}
