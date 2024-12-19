package entity

type AdminAccessControl struct {
	ID         uint
	ActorID    uint
	ActorType  AdminActorType
	Permission AdminPermission
}

type AdminActorType string

const (
	AdminRoleActorType  = AdminActorType("role")
	AdminAdminActorType = AdminActorType("admin")
)
