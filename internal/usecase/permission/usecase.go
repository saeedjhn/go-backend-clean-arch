package permission

import (
	"context"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-domain-driven-design/configs"
	"github.com/saeedjhn/go-domain-driven-design/internal/dto"
	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
)

//go:generate mockery --name Validator
type Validator interface {
}

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, p entity.Permission) (entity.Permission, error)
	GetByID(ctx context.Context, id uint64) (entity.Permission, error)
	GetAll(
		ctx context.Context,
		filter dto.FilterRequest,
		pagination dto.PaginationRequest,
		sort dto.SortRequest,
		searchParams *dto.QuerySearch,
	) ([]entity.Permission, uint, error)
	Update(ctx context.Context, p entity.Permission) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteAll(ctx context.Context) error
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
