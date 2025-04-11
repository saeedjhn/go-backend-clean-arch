package roleresourcepermission

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto"
)

//go:generate mockery --name Validator
type Validator interface {
}

//go:generate mockery --name Repository
type Repository interface {
	Create(_ context.Context, r models.RoleResourcePermission) (models.RoleResourcePermission, error)
	GetByID(_ context.Context, id uint64) (models.RoleResourcePermission, error)
	GetByRoleIDAndResourceID(_ context.Context, roleID, resourceID uint64) (models.RoleResourcePermission, error)
	GetByRoleIDsAndResourceID(
		_ context.Context,
		roleIDs []uint64,
		resourceID uint64,
	) ([]models.RoleResourcePermission, error)
	GetAll(
		_ context.Context,
		filter dto.FilterRequest,
		pagination dto.PaginationRequest,
		sort dto.SortRequest,
		searchParams *dto.QuerySearch,
	) ([]models.RoleResourcePermission, uint, error)
	Update(_ context.Context, r models.RoleResourcePermission) error
	DeleteByID(_ context.Context, id uint64) error
	DeleteAllPermissionsForRole(_ context.Context, roleID uint64) error
	DeleteAll(_ context.Context) error
}

type Interactor struct {
	cfg        *configs.Config
	trc        contract.Tracer
	vld        Validator
	repository Repository
}

func New(
	cfg *configs.Config,
	trc contract.Tracer,
	vld Validator,
	repository Repository,
) *Interactor {
	return &Interactor{
		cfg:        cfg,
		trc:        trc,
		vld:        vld,
		repository: repository,
	}
}
