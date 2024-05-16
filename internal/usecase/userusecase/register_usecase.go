package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"log"
)

func (u *UserInteractor) Register(req userdto.UserRequest) {
	u.userGateway.SaveUserToDB()

	log.Print("UserInteractor -> Register - IMPL ME")
}
