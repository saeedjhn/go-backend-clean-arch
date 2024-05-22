package taskgateway

import "log"

func (g TaskGateway) TaskList() {
	g.taskInteractor.List()

	// Any impl codes

	log.Print("TaskGateway -> TaskList - IMPL ME")
}
