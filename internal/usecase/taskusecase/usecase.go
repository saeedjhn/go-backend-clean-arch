package taskusecase

type Repository interface {
	List()
}

type TaskInteractor struct {
	repository Repository
}

func New(taskGateway Repository) *TaskInteractor {
	return &TaskInteractor{repository: taskGateway}
}
