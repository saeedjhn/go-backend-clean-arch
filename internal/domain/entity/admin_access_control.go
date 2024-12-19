package entity

type AdminAccessControl struct {
	ID         uint64
	ActorID    uint64
	ActorType  AdminActorType
	Permission AdminPermission
}

type AdminActorType string

const (
	AdminRoleActorType  = AdminActorType("role")
	AdminAdminActorType = AdminActorType("admin")
)
