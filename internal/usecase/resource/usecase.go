package resource

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
	Create(_ context.Context, r entity.Resource) (entity.Resource, error)
	GetByID(_ context.Context, id uint64) (entity.Resource, error)
	GetAll(
		_ context.Context,
		filter dto.FilterRequest,
		pagination dto.PaginationRequest,
		sort dto.SortRequest,
		searchParams *dto.QuerySearch,
	) ([]entity.Resource, uint, error)
	Update(_ context.Context, r entity.Resource) error
	DeleteByID(_ context.Context, id uint64) error
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
