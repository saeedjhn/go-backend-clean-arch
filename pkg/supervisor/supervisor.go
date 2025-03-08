package supervisor

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"
)

type Process struct {
	name    string
	handler ProcessFunc
	options ProcessOption
	state   ProcessState
}

type ProcessState struct {
	// RecoveredNum count number of time the process recovered
	RecoveredNum int
}

type ProcessOption struct {
	Recover         bool
	RecoverInterval time.Duration
	RecoverCount    int
	RetryCount      int
	RetryInterval   time.Duration
	IsFatal         bool
}

// ProcessFunc is a long-running process which listens on finishSignal
// It notifies the supervisor by terminate channel when it terminates.
type ProcessFunc func(finishSignal context.Context, processName string, terminateChannel chan<- string) error

var noopProcessFunc = func(_ context.Context, _ string, _ chan<- string) error { //nolint:gochecknoglobals // nothing
	return nil
}

// Supervisor is responsible to manage long-running processes
// Supervisor is not for concurrent use and should be used as the main goroutine of application.
type Supervisor struct {
	ctx              context.Context
	ctxCancel        context.CancelFunc
	logger           contract.Logger
	lock             *sync.Mutex
	processes        map[string]Process
	shutdownSignal   chan os.Signal
	shutdownTimeout  time.Duration
	terminateChannel chan string // terminateChannel should be used to notify supervisor when a process terminates
}

func New(shutdownTimeout time.Duration, l contract.Logger) *Supervisor {
	ctx, cancel := context.WithCancel(context.Background())

	if shutdownTimeout == 0 {
		shutdownTimeout = DefaultGracefulShutdownTimeout
	}

	return &Supervisor{
		lock:            &sync.Mutex{},
		logger:          l,
		processes:       make(map[string]Process),
		shutdownSignal:  make(chan os.Signal, 1),
		ctx:             ctx,
		ctxCancel:       cancel,
		shutdownTimeout: shutdownTimeout,
		// TODO : how to set terminateChannel buffer?
		terminateChannel: make(chan string, 10), //nolint:mnd // nothing
	}
}

// Register registers a new process to supervisor.
func (s *Supervisor) Register(name string, process ProcessFunc, options *ProcessOption) {
	// TODO : don't allow any registration after Start is called using a mutex

	s.warnIfNameAlreadyInUse(name)

	// TODO : validate name
	p := Process{
		name:    name,
		handler: process,
		options: defaultOptions,
		state:   ProcessState{RecoveredNum: 0},
	}

	if options != nil {
		p.options = *options
	}

	s.lock.Lock()
	s.processes[name] = p
	s.lock.Unlock()
}

// Start spawns a new goroutine for each process
// Spawned goroutine is responsible to handle the panics and restart the process.
func (s *Supervisor) Start() {
	// TODO : is it viable to use a goroutine pool such as Ants ?
	for name := range s.processes {
		go s.executeProcessWithRetryPolicy(name)
	}
}

// WaitOnShutdownSignal wait to receive shutdown signal.
// WaitOnShutdownSignal should not be called in other goroutines except main goroutine of application.
func (s *Supervisor) WaitOnShutdownSignal() {
	// TODO : is it necessary to add os.Interrupt to supervisor config?
	signal.Notify(s.shutdownSignal, os.Interrupt)
	<-s.shutdownSignal

	s.gracefulShutdown()
}

func (s *Supervisor) executeProcessWithRetryPolicy(name string) {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Errorw("recover from panic", "process_name", name, "panic", r)

			if s.isRecoverable(name) {
				s.incRecover(name)
				s.waitFoRecover(name)
				s.logger.Infow("restart the process", "process_name", name)

				// spawn new goroutine to avoid heap/stack memory leak when the recover count is big
				go s.executeProcessWithRetryPolicy(name)

				return
			}

			s.logger.Infow("don't try any more to restart the process", "process_name", name)
			s.removeProcess(name)

			if s.isFatal(name) {
				s.logger.Errorw("can't recover important process. exit..", "process_name", name)
				s.shutdownSignal <- os.Interrupt
			}
		}
	}()

	for i := 1; i <= s.retryCount(name); i++ {
		s.logger.Infow("execute process", "process_name", name)
		f := s.handler(name)
		err := f(s.ctx, name, s.terminateChannel)
		if err != nil {
			s.logger.Errorw("failed to execute process", "process_name", name,
				"attempt", i, "error", err.Error())
			s.waitFoRetry(name)

			continue
		}

		// don't expect handler return if it hasn't any error because it's long-running process
		// it should return when receives shutdown signal
		s.logger.Infow("process terminates with no error", "process_name", name)

		if s.isFatal(name) {
			s.logger.Errorw("can't recover important process. exit..", "process_name", name)
			s.shutdownSignal <- os.Interrupt
		}

		return
	}

	s.logger.Infow("don't try any more to execute process", "process_name", name)
	s.removeProcess(name)
}

func (s *Supervisor) gracefulShutdown() {
	s.logger.Info("shutdown all processes gracefully")

	s.logger.Infow(
		"notify all processes (goroutines) to finish their jobs",
		"shutdown_timeout", s.shutdownTimeout)
	s.ctxCancel()

	forceExitCtx, forceExitCancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer forceExitCancel()

	for {
		select {
		case name := <-s.terminateChannel:
			s.logger.Infow("process terminates gracefully", "process_name", name)
			s.removeProcess(name)

		case <-forceExitCtx.Done():
			s.logger.Infow("supervisor terminates its job.",
				"number_of_unfinished_processes", len(s.processes),
			)

			return
		}
	}
}

func (s *Supervisor) removeProcess(name string) {
	s.lock.Lock()
	delete(s.processes, name)
	s.lock.Unlock()
}

func (s *Supervisor) isRecoverable(name string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return false
	}

	if v.options.Recover && v.state.RecoveredNum < v.options.RecoverCount {
		return true
	}

	return false
}

func (s *Supervisor) isFatal(name string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return false
	}

	return v.options.IsFatal
}

func (s *Supervisor) incRecover(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return
	}

	v.state.RecoveredNum++
	s.processes[name] = v
}

func (s *Supervisor) retryCount(name string) int {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return -1
	}

	return v.options.RetryCount
}

func (s *Supervisor) retryInterval(name string) time.Duration {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return -1
	}

	return v.options.RetryInterval
}

func (s *Supervisor) waitFoRecover(name string) {
	s.lock.Lock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return
	}

	t := v.options.RecoverInterval

	// free lock before sleep
	s.lock.Unlock()

	time.Sleep(t)
}

func (s *Supervisor) waitFoRetry(name string) {
	s.lock.Lock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return
	}

	t := v.options.RetryInterval

	// free lock before sleep
	s.lock.Unlock()

	s.logger.Infow("wait to retry execute process after sleep interval",
		"process_name", name, "interval", t,
	)

	time.Sleep(t)
}

func (s *Supervisor) handler(name string) ProcessFunc {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, ok := s.processes[name]
	if !ok {
		s.logger.Warnw("process doesn't exist", "process_name", name)

		return noopProcessFunc
	}

	return v.handler
}

func (s *Supervisor) warnIfNameAlreadyInUse(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.processes[name]; ok {
		s.logger.Warnw("process name already in use", "process_name", name)
	}
}
