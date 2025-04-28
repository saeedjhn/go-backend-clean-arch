package circuitbreaker

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sony/gobreaker/v2"
)

type State struct {
	Name              string
	CurrentState      string
	TotalRequests     uint32
	TotalSuccesses    uint32
	TotalFailures     uint32
	ConsecutiveErrors uint32
	LastError         error
	OpenedAt          time.Time
}

type Manager[T any] struct {
	breakers    sync.Map
	timeout     time.Duration
	threshold   uint32
	globalStats map[string]State
	statsMutex  sync.RWMutex
}

func NewManager[T any](
	timeout time.Duration,
	failureThreshold uint32,
) *Manager[T] {
	return &Manager[T]{
		timeout:     timeout,
		threshold:   failureThreshold,
		globalStats: make(map[string]State),
	}
}

func (m *Manager[T]) GetBreaker(name string) *gobreaker.CircuitBreaker[T] {
	cb, loaded := m.breakers.Load(name)
	if loaded {
		return cb.(*gobreaker.CircuitBreaker[T]) //nolint:errcheck // noproblem
	}

	settings := gobreaker.Settings{
		Name:    name,
		Timeout: m.timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= m.threshold
		},
		OnStateChange: m.handleStateChange(name),
	}

	newCb := gobreaker.NewCircuitBreaker[T](settings)
	cb, _ = m.breakers.LoadOrStore(name, newCb)
	return cb.(*gobreaker.CircuitBreaker[T]) //nolint:errcheck // noproblem
}

func (m *Manager[T]) Execute(name string, action func() (T, error)) (T, error) {
	cb := m.GetBreaker(name)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("circuit breaker panic: %v", r)
			log.Printf("panic in circuit breaker %s: %v", name, r)
			m.recordError(name, err)
		}
	}()

	res, execErr := cb.Execute(func() (T, error) {
		return action()
	})

	if execErr != nil {
		m.recordError(name, execErr)
	} else {
		m.recordSuccess(name)
	}

	return res, execErr
}

func (m *Manager[T]) State(name string) (State, error) {
	cb, exists := m.breakers.Load(name)
	if !exists {
		return State{}, ErrBreakerNotFound
	}

	breaker := cb.(*gobreaker.CircuitBreaker[T]) //nolint:errcheck // noproblem
	counts := breaker.Counts()

	m.statsMutex.RLock()
	defer m.statsMutex.RUnlock()

	return State{
		Name:              name,
		CurrentState:      breaker.State().String(),
		TotalRequests:     counts.Requests,
		TotalSuccesses:    counts.TotalSuccesses,
		TotalFailures:     counts.TotalFailures,
		ConsecutiveErrors: counts.ConsecutiveFailures,
		LastError:         m.globalStats[name].LastError,
		OpenedAt:          m.globalStats[name].OpenedAt,
	}, nil
}

func (m *Manager[T]) Reset(name string) {
	m.breakers.Delete(name)
	m.statsMutex.Lock()
	delete(m.globalStats, name)
	m.statsMutex.Unlock()
}

func (m *Manager[T]) handleStateChange(name string) func(string, gobreaker.State, gobreaker.State) {
	return func(_ string, from gobreaker.State, to gobreaker.State) {
		m.statsMutex.Lock()
		defer m.statsMutex.Unlock()

		stats := m.globalStats[name]
		stats.Name = name
		stats.CurrentState = to.String()

		if to == gobreaker.StateOpen {
			stats.OpenedAt = time.Now()
		}

		m.globalStats[name] = stats

		log.Printf("Circuit Breaker '%s' changed state: %s → %s",
			name,
			from.String(),
			to.String(),
		)
	}
}

func (m *Manager[T]) recordError(name string, err error) {
	m.statsMutex.Lock()
	defer m.statsMutex.Unlock()

	stats := m.globalStats[name]
	stats.LastError = err
	m.globalStats[name] = stats
}

func (m *Manager[T]) recordSuccess(name string) {
	m.statsMutex.Lock()
	defer m.statsMutex.Unlock()

	stats := m.globalStats[name]
	stats.LastError = nil
	m.globalStats[name] = stats
}
