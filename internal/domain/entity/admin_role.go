package entity

type AdminRole uint

const (
	AdminSuperAdminRole AdminRole = iota + 1
	AdminManagerRole
	AdminModeratorRole
	AdminSupportRole
	AdminDeveloperRole
	AdminViewerRole
	AdminEditorRole
	AdminAnalystRole
	AdminTesterRole
)

var _adminRoleStrings = map[AdminRole]string{ //nolint:gochecknoglobals // nothing
	AdminSuperAdminRole: "super-admin",
	AdminManagerRole:    "manager",
	AdminModeratorRole:  "moderator",
	AdminSupportRole:    "support",
	AdminDeveloperRole:  "developer",
	AdminViewerRole:     "viewer",
	AdminEditorRole:     "editor",
	AdminAnalystRole:    "analyst",
	AdminTesterRole:     "tester",
}

func (s AdminRole) String() string {
	return _adminRoleStrings[s]
}

func (s AdminRole) IsValid() bool {
	// Which Method is More Efficient?
	// If the role values are always sequential and start from zero (like in the code you've written),
	// the first method:

	// return s < AdminRole(len(_adminRoleStrings))

	// is more efficient and faster.
	//
	// However, if there is a possibility of introducing non-sequential or indirect values in the future,
	// the second method (checking existence in the map) is safer and more reliable.

	// Recommendation:
	// For greater stability, it is better to use the map-based method unless you are certain that the roles
	// will always be sequential.

	return s <= AdminRole(len(_adminRoleStrings))

	// _, ok := _adminRoleStrings[s]
	// return ok
}

func MapToAdminRole(roleStr string) AdminRole {
	for role, str := range _adminRoleStrings {
		if str == roleStr {
			return role
		}
	}

	return AdminRole(0)
}
