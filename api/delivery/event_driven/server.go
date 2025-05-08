package eventdriven

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

const _bufferSize = 1024

type Server struct {
	app *bootstrap.Application
	ed  *event.C
}

func New(app *bootstrap.Application) Server {
	return Server{app: app}
}

func (s Server) Run() error {
	router := event.NewRouter()

	for t, h := range s.app.EventRegister {
		router.Register(t, h)
	}

	s.ed = event.NewEventConsumer(
		_bufferSize,
		router,
		s.app.Rabbitmq,
	).WithLogger(s.app.Logger)

	s.ed.Start()

	return nil
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.ed.Shutdown(ctx)
}
