package authorization //nolint:cyclop // nothing

import (
	"context"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

//go:generate mockery --name RoleResourcePermissionInteractor
type RoleResourcePermissionInteractor interface {
	GetByRoleIDAndResourceID(ctx context.Context, roleID, resourceID uint64) (models.RoleResourcePermission, error)
	// GetByRoleIDsAndResourceID(
	// ctx context.Context,
	// roleIDs []uint64,
	// resourceID uint64,
	// ) ([]models.RoleResourcePermission, error)
}

//go:generate mockery --name ResourceInteractor
type ResourceInteractor interface {
	GetIDByName(ctx context.Context, name string) (uint64, error)
}

//go:generate mockery --name Cache
// type Cache interface {
// 	Get(ctx context.Context, key string) ([]byte, error)
// 	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
// }

type Config struct {
	// CacheTTL time.Duration `mapstructure:"ttl"`
}

type Interactor struct {
	Config                     Config
	resourceIntr               ResourceInteractor
	roleResourcePermissionIntr RoleResourcePermissionInteractor
	// Cache                      Cache
}

func New(
	config Config,
	resourceIntr ResourceInteractor,
	roleResourcePermissionIntr RoleResourcePermissionInteractor,
	// cache Cache,
) *Interactor {
	return &Interactor{
		Config:                     config,
		resourceIntr:               resourceIntr,
		roleResourcePermissionIntr: roleResourcePermissionIntr,
		// Cache:                      cache,
	}
}

func (a *Interactor) CheckAccess( //nolint:gocognit // nothing
	ctx context.Context,
	roleIDs []types.ID,
	resourceName string,
	actions ...models.Action,
) (bool, error) {
	if len(actions) == 0 || len(roleIDs) == 0 {
		return false, nil
	}

	resourceID, err := a.resourceIntr.GetIDByName(ctx, resourceName)
	if err != nil {
		return false, err
	}

	for _, roleID := range roleIDs {
		// TODO: [OPTIMIZATION] Refactor this function to improve performance.
		// cacheKey := fmt.Sprintf("perm:%d:%d", roleID, resourceID)
		// cachedPerm, err := a.Cache.Get(ctx, cacheKey)
		//
		// var perm models.RoleResourcePermission
		// if err == nil {
		// 	_ = decodeGob(cachedPerm, &perm)
		// } else {
		// 	perm, err = a.roleResourcePermissionIntr.GetByRoleIDAndResourceID(ctx, roleID, resourceID)
		// 	if err != nil {
		// 		continue
		// 	}
		//
		// 	encoded, _ := encodeGob(perm)
		// 	a.Cache.Set(ctx, cacheKey, encoded, a.Config.CacheTTL)
		// }

		perm, errRRP := a.roleResourcePermissionIntr.GetByRoleIDAndResourceID(ctx, roleID.Uint64(), resourceID)
		if errRRP != nil {
			continue
		}

		allowAccess := true
		for _, action := range actions {
			switch action {
			case models.ReadAction:
				if !(perm.Permissions.Allow.R && !perm.Permissions.Deny.R) {
					allowAccess = false
				}
			case models.WriteAction:
				if !(perm.Permissions.Allow.W && !perm.Permissions.Deny.W) {
					allowAccess = false
				}
			case models.ExecuteAction:
				if !(perm.Permissions.Allow.X && !perm.Permissions.Deny.X) {
					allowAccess = false
				}
			case models.DeleteAction:
				if !(perm.Permissions.Allow.D && !perm.Permissions.Deny.D) {
					allowAccess = false
				}
			default:
				return false, errors.New("invalid action type: " + string(action))
			}

			if !allowAccess {
				break
			}
		}

		if allowAccess {
			return true, nil
		}
	}

	return false, nil
}

// func encodeGob(value interface{}) ([]byte, error) {
// 	var buf bytes.Buffer
//
// 	enc := gob.NewEncoder(&buf)
// 	err := enc.Encode(value)
//
// 	return buf.Bytes(), err
// }
//
// func decodeGob(data []byte, to interface{}) error {
// 	buf := bytes.NewBuffer(data)
// 	dec := gob.NewDecoder(buf)
//
// 	return dec.Decode(to)
// }

// func (a *Interactor) CheckAccess(
// ctx context.Context,
// roleIDs []uint64,
// resourceName string,
// actions ...models.Action,
// ) (bool, error) {
// 	if len(roleIDs) == 0 || len(actions) == 0 {
// 		return false, nil
// 	}
//
// 	resourceID, err := a.resourceIntr.GetIDByName(ctx, resourceName)
// 	if err != nil {
// 		return false, err
// 	}
//
// 	allowAccess := true
//
// 	for _, roleID := range roleIDs {
// 		perm, errRRP := a.roleResourcePermissionIntr.GetByRoleIDAndResourceID(ctx, roleID, resourceID)
// 		if errRRP != nil {
// 			return false, err
// 		}
//
// 		for _, action := range actions {
// 			switch action {
// 			case models.ReadAction:
// 				if !(perm.Permissions.Allow.R && !perm.Permissions.Deny.R) {
// 					allowAccess = false
// 				}
// 			case models.WriteAction:
// 				if !(perm.Permissions.Allow.W && !perm.Permissions.Deny.W) {
// 					allowAccess = false
// 				}
// 			case models.ExecuteAction:
// 				if !(perm.Permissions.Allow.X && !perm.Permissions.Deny.X) {
// 					allowAccess = false
// 				}
// 			case models.DeleteAction:
// 				if !(perm.Permissions.Allow.D && !perm.Permissions.Deny.D) {
// 					allowAccess = false
// 				}
// 			default:
// 				return false, errors.New("invalid action type: " + string(action))
// 			}
//
// 			if !allowAccess {
// 				return false, nil
// 			}
// 		}
// 	}
//
// 	return allowAccess, nil
// }

// func (a *Interactor) CheckAccess(ctx context.Context, roleID, resourceID uint64, actions ...string) (bool, error) {
// 	perm, err := a.roleResourcePermissionIntr.GetByRoleIDAndResourceID(ctx, roleID, resourceID)
// 	if err != nil {
// 		return false, err
// 	}
//
// 	// if perm == nil {
// 	// 	return false, errors.New("permission not found")
// 	// }
//
// 	for _, action := range actions {
// 		switch action {
// 		case "read":
// 			if perm.Permissions.Allow.R && !perm.Permissions.Deny.R {
// 				return true, nil
// 			}
// 		case "write":
// 			if perm.Permissions.Allow.W && !perm.Permissions.Deny.W {
// 				return true, nil
// 			}
// 		case "execute":
// 			if perm.Permissions.Allow.X && !perm.Permissions.Deny.X {
// 				return true, nil
// 			}
// 		case "delete":
// 			if perm.Permissions.Allow.D && !perm.Permissions.Deny.D {
// 				return true, nil
// 			}
// 		default:
// 			return false, errors.New("invalid action type: " + action)
// 		}
// 	}
//
// 	return false, nil
// }

// func (s *Interactor) CheckAccess(
// 	ctx context.Context,
// 	adminID uint,
// role models.AdminRole,
// permissions ...models.AdminPermission,
// ) (bool, error) {

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

// return false, nil
// }
