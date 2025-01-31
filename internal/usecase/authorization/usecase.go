package authorization

import (
	"context"
)

type Repository interface {
	// GetAdminPermissions(ctx context.Context, adminID uint, role entity.AdminRole) ([]entity.AdminPermission, error)
}

type Config struct {
}

type Interactor struct {
	Config Config
}

func New(config Config) *Interactor {
	return &Interactor{Config: config}
}

func (s *Interactor) CheckAccess(
	ctx context.Context,
	adminID uint,
	// role entity.AdminRole,
	// permissions ...entity.AdminPermission,
) (bool, error) {

	// AdminPermissions, err := s.repo.GetAdminPermissions(ctx, adminID, role)
	// if err != nil {
	// 	return false, richerror.New(op).WithErr(err)
	// }
	//
	// for _, p := range permissions {
	// 	for _, ap := range AdminPermissions {
	// 		if p == ap {
	// 			return true, nil
	// 		}
	// 	}
	// }
	//

	return false, nil
}
