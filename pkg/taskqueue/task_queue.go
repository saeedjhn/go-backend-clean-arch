package taskqueue

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/hibiken/asynq"
)

type HandlerFunc func(context.Context, Task) error

type TaskQueue struct {
	client    *asynq.Client
	server    *asynq.Server
	scheduler *asynq.Scheduler
	mux       *asynq.ServeMux
}

func New(config Config) *TaskQueue {
	client := asynq.NewClient(getRedisConfig(config.RedisConfig))

	server := asynq.NewServer(
		getRedisConfig(config.RedisConfig),
		asynq.Config{
			Concurrency: config.Concurrency,
			Queues:      _queues,
		},
	)

	scheduler := asynq.NewScheduler(getRedisConfig(config.RedisConfig), nil)

	return &TaskQueue{
		client:    client,
		server:    server,
		scheduler: scheduler,
		mux:       asynq.NewServeMux(),
	}
}

func (tq *TaskQueue) EnqueueTask(taskType string, payload []byte, queueName Queue) error {
	task := asynq.NewTask(taskType, payload)
	_, err := tq.client.Enqueue(task, asynq.Queue(queueName.String()))

	return err
}

func (tq *TaskQueue) EnqueueTaskIn(taskType string, payload []byte, queueName Queue, delay time.Duration) error {
	task := asynq.NewTask(taskType, payload)
	_, err := tq.client.Enqueue(task, asynq.Queue(queueName.String()), asynq.ProcessIn(delay))

	return err
}

func (tq *TaskQueue) EnqueueTaskAt(taskType string, payload []byte, queueName Queue, runAt time.Time) error {
	task := asynq.NewTask(taskType, payload)
	_, err := tq.client.Enqueue(task, asynq.Queue(queueName.String()), asynq.ProcessAt(runAt))

	return err
}

func (tq *TaskQueue) RegisterHandler(taskType string, handler HandlerFunc) {
	tq.mux.HandleFunc(taskType, func(ctx context.Context, task *asynq.Task) error {
		return handler(ctx, Task{
			Type:    task.Type(),
			Payload: task.Payload(),
		})
	})
}

func (tq *TaskQueue) ScheduleTask(schedule string, taskType string, payload []byte) error {
	task := asynq.NewTask(taskType, payload)
	_, err := tq.scheduler.Register(schedule, task)

	return err
}

func (tq *TaskQueue) StartScheduler() error {
	if err := tq.scheduler.Start(); err != nil {
		return fmt.Errorf("could not start scheduler: %w", err)
	}

	return nil
}

func (tq *TaskQueue) StartServer() error {
	if err := tq.server.Run(tq.mux); err != nil {
		return fmt.Errorf("could not start worker server: %w", err)
	}

	return nil
}

func getRedisConfig(config RedisConfig) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Network:      config.Network,
		Addr:         net.JoinHostPort(config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		TLSConfig:    nil,
	}
}
