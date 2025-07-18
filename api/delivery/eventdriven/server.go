package eventdriven

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/eventdriven/handler"

	"github.com/saeedjhn/go-backend-clean-arch/configs"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

type Server struct {
	ctx context.Context
	app *bootstrap.Application
	ed  *event.C
}

func New(app *bootstrap.Application) *Server {
	return &Server{ctx: context.Background(), app: app}
}

func (s *Server) WithContextConsumer(ctx context.Context) *Server {
	s.ctx = ctx

	return s
}

func (s *Server) Run() error {
	router := event.NewRouter()

	// for t, h := range s.app.EventRegister {
	// 	router.Register(t, h)
	// }

	for t, h := range handler.Setup(s.app) {
		router.Register(t, h)
	}

	// rmq, err := NewRabbitmq(s.app.Config.RabbitMQ, s.app.EventRegister)
	rmq, err := NewRabbitmq(s.app.Config.RabbitMQ, handler.Setup(s.app))
	if err != nil {
		return err
	}

	s.ed = event.NewEventConsumer(
		configs.EventBufferSize,
		router,
		rmq,
	).WithLogger(s.app.Logger).WithContext(s.ctx)

	s.ed.Start()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.ed.Shutdown(ctx)
}
