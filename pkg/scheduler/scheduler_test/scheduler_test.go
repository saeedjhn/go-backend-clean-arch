package scheduler_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/saeedjhn/go-domain-driven-design/pkg/scheduler"
	"github.com/stretchr/testify/assert"
)

//go:generate go test -v -race -count=1 ./...

func TestScheduler_Configure_InitializesScheduler(t *testing.T) {
	sch := scheduler.New()

	err := sch.Configure()

	assert.NoError(t, err, "scheduler should be initialized successfully")
}

func TestScheduler_RepeatTaskEvery_WithValidConfig_TaskRunsSuccessfully(t *testing.T) {
	sch := scheduler.New()
	_ = sch.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sch.RepeatTaskEvery(ctx, func() {}, time.Millisecond*10)

	require.NoError(t, err)

	_ = sch.Start()

	time.Sleep(50 * time.Millisecond)

	_ = sch.Shutdown()
}

func TestScheduler_RepeatTaskEvery_SchedulerNotInitialized_ReturnsError(t *testing.T) {
	sch := scheduler.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sch.RepeatTaskEvery(ctx, func() {}, time.Second)

	assert.Error(t, err)
}

func TestScheduler_StartAt_WithValidConfig_TaskRunsAtScheduledTime(t *testing.T) {
	sch := scheduler.New()
	_ = sch.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sch.StartAt(ctx, func() {}, time.Millisecond*50)

	require.NoError(t, err)

	_ = sch.Start()

	time.Sleep(70 * time.Millisecond)

	_ = sch.Shutdown()
}

func TestScheduler_StartAt_WithCancel_TaskCanceledAtScheduledTime(t *testing.T) {
	sch := scheduler.New()
	_ = sch.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := sch.StartAt(ctx, func() {}, time.Millisecond*50)

	require.NoError(t, err)

	_ = sch.Start()

	time.Sleep(70 * time.Millisecond)

	_ = sch.Shutdown()
}

func TestScheduler_StartAt_SchedulerNotInitialized_ReturnsError(t *testing.T) {
	sch := scheduler.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sch.StartAt(ctx, func() {}, time.Second)

	assert.Error(t, err)
}

func TestScheduler_Start_SchedulerNotInitialized_ReturnsError(t *testing.T) {
	sch := scheduler.New()

	err := sch.Start()

	assert.Error(t, err)
}

func TestScheduler_StopJobs_WithRunningTasks_TasksShouldStop(t *testing.T) {
	sch := scheduler.New()
	_ = sch.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = sch.RepeatTaskEvery(ctx, func() {}, time.Millisecond*10)

	_ = sch.Start()
	time.Sleep(30 * time.Millisecond)

	err := sch.StopJobs()
	assert.NoError(t, err)
}

func TestScheduler_StopJobs_SchedulerNotInitialized_ReturnsError(t *testing.T) {
	sch := scheduler.New()

	err := sch.StopJobs()

	assert.Error(t, err)
}

func TestScheduler_Shutdown_WithRunningTasks_SchedulerStopsSuccessfully(t *testing.T) {
	sch := scheduler.New()
	_ = sch.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = sch.RepeatTaskEvery(ctx, func() {}, time.Millisecond*10)

	_ = sch.Start()
	time.Sleep(30 * time.Millisecond)

	err := sch.Shutdown()
	assert.NoError(t, err)
}

func TestScheduler_Shutdown_SchedulerNotInitialized_ReturnsError(t *testing.T) {
	sch := scheduler.New()

	err := sch.Shutdown()

	assert.Error(t, err)
}
