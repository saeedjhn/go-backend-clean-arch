package role

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, r entity.Role) (entity.Role, error)
	GetByID(ctx context.Context, id uint64) (entity.Role, error)
	GetAll(
		ctx context.Context,
		filter dto.FilterRequest,
		pagination dto.PaginationRequest,
		sort dto.SortRequest,
		searchParams *dto.QuerySearch,
	) ([]entity.Role, uint, error)
	Update(ctx context.Context, r entity.Role) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteAll(ctx context.Context) error
}

type Interactor struct {
	repository Repository
}

func New(repository Repository) *Interactor {
	return &Interactor{repository: repository}
}
