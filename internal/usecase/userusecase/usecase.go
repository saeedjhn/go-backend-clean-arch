package userusecase

type Repository interface {
	Create()
}

type Gateway interface {
	TaskList()
}

type UserInteractor struct {
	repository      Repository
	taskListGateway Gateway
}

func New(taskListGateway Gateway, repository Repository) *UserInteractor {
	return &UserInteractor{taskListGateway: taskListGateway, repository: repository}
}
