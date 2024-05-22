package taskgateway

type TaskInteractor interface {
	List()
}

type TaskGateway struct {
	taskInteractor TaskInteractor
}

func New(taskInteractor TaskInteractor) *TaskGateway {
	return &TaskGateway{taskInteractor: taskInteractor}
}
