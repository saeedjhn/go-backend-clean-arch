package admin

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

//go:generate mockery --name Validator
type Validator interface {
}

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, a entity.Admin) (entity.Admin, error)
	GetByID(ctx context.Context, id uint64) (entity.Admin, error)
	GetAll(
		ctx context.Context,
		filter dto.FilterRequest,
		pagination dto.PaginationRequest,
		sort dto.SortRequest,
		searchParams *dto.QuerySearch,
	) ([]entity.Admin, uint, error)
	GetRolesIDsByID(ctx context.Context, id uint64) ([]uint64, error)
	MobileExists(ctx context.Context, mobile string) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	DeleteByID(ctx context.Context, id uint64) error
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
