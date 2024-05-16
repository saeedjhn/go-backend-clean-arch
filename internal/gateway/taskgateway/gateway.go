package taskgateway

import "log"

type Repository interface {
	Find()
}

type TaskGateway struct {
	taskRepository Repository
}

func New(taskRepository Repository) *TaskGateway {
	return &TaskGateway{taskRepository: taskRepository}
}

func (t *TaskGateway) List() {
	t.taskRepository.Find()

	// Any impl codes

	log.Print("TaskGateway -> List - IMPL ME")
}
