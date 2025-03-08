package taskqueue_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/saeedjhn/go-domain-driven-design/pkg/taskqueue"
)

//go:generate go test -v -race -count=1 ./...

func TestTaskQueue_EnqueueTask_ValidTask_ReturnsNoError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})
	err := tq.EnqueueTask("example_task", []byte(`{"key": "value"}`))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTaskQueue_EnqueueTaskIn_ValidTaskWithDelay_ReturnsNoError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.EnqueueTaskIn("example_task", []byte(`{"key": "value"}`), taskqueue.WithDelay(10*time.Second))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTaskQueue_EnqueueTaskAt_ValidTaskWithTime_ReturnsNoError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.EnqueueTaskAt("example_task", []byte(`{"key": "value"}`), taskqueue.WithAt(time.Now().Add(10*time.Minute)))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTaskQueue_RegisterHandler_ValidHandler_RegistersSuccessfully(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	handler := func(_ context.Context, _ taskqueue.Task) error {
		return nil
	}

	tq.RegisterHandler("example_task", handler)
}

func TestTaskQueue_ScheduleTask_InvalidSchedule_ReturnsError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.ScheduleTask("invalid_schedule", "example_task", []byte(`{"key": "value"}`))
	if err == nil {
		t.Error("Expected an error due to invalid schedule, but got none")
	}
}

func TestTaskQueue_ScheduleTask_ValidSchedule_ReturnsNoError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.ScheduleTask("@every 1m", "example_task", []byte(`{"key": "value"}`))
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestTaskQueue_EnqueueTask_InvalidRedisConfig_ReturnsError(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: "server_invalid",
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.EnqueueTask("example_task", []byte(`{"key": "value"}`))
	if err == nil {
		t.Error("Expected an error due to invalid Redis config, but got none")
	}
}

func TestTaskQueue_ShutdownScheduler_ValidScheduler_ShutsDownSuccessfully(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	err := tq.StartScheduler()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	tq.ShutdownScheduler()
}

func TestTaskQueue_ShutdownServer_ValidServer_ShutsDownSuccessfully(t *testing.T) {
	t.Parallel()

	server := miniredis.RunT(t)
	tq := taskqueue.New(taskqueue.Config{
		RedisConfig: taskqueue.RedisConfig{
			Host: server.Host(),
			Port: server.Port(),
		},
		Concurrency: 10,
	})

	go func() {
		err := tq.StartServer()
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	}()
	time.Sleep(2 * time.Second)

	// process, err := os.FindProcess(os.Getpid())
	// if err != nil {
	// 	t.Fatalf("Failed to find process: %v", err)
	// }

	// err = process.Signal(syscall.SIGTSTP) // Ctrl + z
	// if err != nil {
	// 	t.Fatalf("Failed to send SIGTSTP signal: %v", err)
	// }
	// err = process.Signal(syscall.SIGTERM) // Ctrl + c
	// if err != nil {
	// 	t.Fatalf("Failed to send SIGTERM signal: %v", err)
	// }
	// err = process.Signal(syscall.SIGINT) // Ctrl + c
	// if err != nil {
	// 	t.Fatalf("Failed to send SIGINT signal: %v", err)

	tq.ShutdownServer()
}
