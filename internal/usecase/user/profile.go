package user

import (
	"context"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto/user"
)

func (i *Interactor) Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error) {
	ctx, span := i.trc.Span(ctx, "Profile")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Profile",
		"usecase.request": req,
	})
	defer span.End()

	u, err := i.repository.GetByID(ctx, req.ID.Uint64())
	if err != nil {
		return user.ProfileResponse{}, err
	}

	return user.ProfileResponse{UserInfo: user.Info{
		ID:        u.ID,
		Name:      u.Name,
		Mobile:    u.Mobile,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}}, nil // Or
	// return userdto.ProfileResponse{
	//	Info: user.ToUserInfoDTO(),
	// }, nil
}
