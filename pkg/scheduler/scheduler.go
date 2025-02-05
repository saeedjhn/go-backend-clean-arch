package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	sch gocron.Scheduler
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
	t time.Duration,
	fn func(),
) error {
	_, err := s.sch.NewJob(gocron.DurationJob(t), gocron.NewTask(fn))
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (s *Scheduler) StartAt(t time.Duration, fn func()) error {
	_, err := s.sch.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(time.Now().Add(t))),
		gocron.NewTask(fn),
	)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (s *Scheduler) Start() {
	s.sch.Start()
}

func (s *Scheduler) StopJobs() error {
	if err := s.sch.StopJobs(); err != nil {
		return fmt.Errorf("failed to stop scheduled jobs: %w", err)
	}

	return nil
}

func (s *Scheduler) Shutdown() error {
	if err := s.sch.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown scheduler: %w", err)
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
