package scheduler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	sch gocron.Scheduler
	mu  sync.Mutex
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Configure() error {
	sch, err := gocron.NewScheduler(gocron.WithLocation(time.UTC))
	if err != nil {
		return fmt.Errorf("failed to create scheduler: %w", err)
	}
	s.sch = sch

	return nil
}

func (s *Scheduler) RepeatTaskEvery(
	ctx context.Context,
	fn func(),
	t time.Duration,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.checkInitialized(); err != nil {
		return err
	}

	task := s.wrapTaskWithContext(ctx, fn)

	_, err := s.sch.NewJob(gocron.DurationJob(t), gocron.NewTask(task))
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (s *Scheduler) StartAt(
	ctx context.Context,
	fn func(),
	t time.Duration,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.checkInitialized(); err != nil {
		return err
	}

	task := s.wrapTaskWithContext(ctx, fn)

	_, err := s.sch.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(time.Now().Add(t))),
		gocron.NewTask(task),
	)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (s *Scheduler) Start() error {
	if err := s.checkInitialized(); err != nil {
		return err
	}

	s.sch.Start()

	return nil
}

func (s *Scheduler) StopJobs() error {
	if err := s.checkInitialized(); err != nil {
		return err
	}

	if err := s.sch.StopJobs(); err != nil {
		return fmt.Errorf("failed to stop scheduled jobs: %w", err)
	}

	return nil
}

func (s *Scheduler) Shutdown() error {
	if err := s.checkInitialized(); err != nil {
		return err
	}

	if err := s.sch.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown scheduler: %w", err)
	}

	return nil
}

func (s *Scheduler) wrapTaskWithContext(ctx context.Context, fn func()) func() {
	return func() {
		select {
		case <-ctx.Done():
			return
		default:
			fn()
		}
	}
}

func (s *Scheduler) checkInitialized() error {
	if s.sch == nil {
		return _errNotInitialized
	}

	return nil
}

// func (s *Scheduler) RunTaskEvery(
// 	t time.Duration,
// 	fn func() error,
// ) error {
// 	_, err := s.sch.NewJob(
// 		gocron.DurationJob(t),
// 		gocron.NewTask(func() {
// 			if err := fn(); err != nil {
// 				s.errorChannel <- err
// 			}
// 		}),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to create task: %w", err)
// 	}
//
// 	return nil
// }
//
// func (s *Scheduler) Start() {
// 	s.sch.Start()
// }
//
// func (s *Scheduler) Errors() <-chan error {
// 	return s.errorChannel
// }
//
// func (s *Scheduler) Shutdown() error {
// 	close(s.errorChannel)
//
// 	if err := s.sch.Shutdown(); err != nil {
// 		return fmt.Errorf("failed to shutdown scheduler: %w", err)
// 	}
//
// 	return nil
// }
