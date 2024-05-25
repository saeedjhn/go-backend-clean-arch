package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
)

func (u *UserInteractor) Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	const op = "userInteractor - Register"

	u.repository.Create()

	//err0 := richerror.New(op).WithErr(errors.New("ERROR")).
	//	WithKind(kind.KindStatusForbidden).
	//	WithMeta(map[string]interface{}{"req": req})

	//err := richerror.New("ANY").WithErr(err0).WithMessage("MSGGG")
	//WithMeta(map[string]interface{}{"req": req})

	//log.Print("UserInteractor -> Register - IMPL ME")
	return userdto.RegisterResponse{
		User: userdto.UserInfo{
			ID:     0,
			Mobile: "0123456789",
			Name:   "John Doe",
		},
	}, nil
}
