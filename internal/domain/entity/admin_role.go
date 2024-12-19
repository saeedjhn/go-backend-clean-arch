package entity

type AdminRole uint

const (
	AdminSuperAdminRole AdminRole = iota
	AdminAdminRole
	AdminAgentRole
	AdminManagerRole
	AdminModeratorRole
	AdminSupportRole
	AdminDeveloperRole
	AdminViewerRole
	AdminEditorRole
	AdminAnalystRole
	AdminTesterRole
)

var _adminRoleStrings = map[AdminRole]string{
	AdminSuperAdminRole: "super-admin",
	AdminAdminRole:      "admin",
	AdminAgentRole:      "agent",
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
	return s > 0 && int(s) <= len(_adminRoleStrings)
}

func MapToAdminRole(roleStr string) AdminRole {
	for role, str := range _adminRoleStrings {
		if str == roleStr {
			return role
		}
	}

	return AdminRole(0)
}
