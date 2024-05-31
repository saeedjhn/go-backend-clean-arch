package taskinggateway

import "log"

func (g TaskingGateway) TaskList() {
	g.taskInteractor.List()

	// Any impl codes

	log.Print("TaskingGateway -> TaskList - IMPL ME")
}
