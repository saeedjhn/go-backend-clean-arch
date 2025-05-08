package user

import (
	"context"
	"errors"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/models/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	ctx, span := i.trc.Span(ctx, "Register")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Register",
		"usecase.request": req,
	})
	defer span.End()

	if fieldsErrs, err := i.vld.ValidateRegisterRequest(req); err != nil {
		return userdto.RegisterResponse{FieldErrors: fieldsErrs}, err
	}

	isUnique, err := i.repository.IsExistsByMobile(ctx, req.Mobile)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	if !isUnique {
		return userdto.RegisterResponse{},
			richerror.New(_opUserServiceRegister).
				WithErr(errors.New(errMsgMobileIsNotUnique)).
				WithMessage(errMsgMobileIsNotUnique).
				WithKind(richerror.KindStatusBadRequest)
	}

	u := user.User{
		Name:   req.Name,
		Mobile: req.Mobile,
		Email:  req.Email,
	}

	encryptedPass, _ := GenerateHash(req.Password) // TODO: usecase>userusecase>register>CheckErrorGenerateHash
	u.Password = encryptedPass

	createdUser, err := i.repository.Create(ctx, u)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	u.AddEvents(events.NewUserRegisteredEvent(u.ID, "reason"))
	if err = i.outboxIntr.Create(ctx, u.PullEvents()); err != nil {
		return userdto.RegisterResponse{}, err
	}

	return userdto.RegisterResponse{
		UserInfo: userdto.Info{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Mobile:    createdUser.Mobile,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		}, // Or
		// Info: createdUser.ToUserInfoDTO(),
	}, nil
}
