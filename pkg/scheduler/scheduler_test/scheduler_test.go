package scheduler_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/scheduler"
	"github.com/stretchr/testify/assert"
)

//go:generate go test -v -race -count=1 ./...

func Test_Configure_WhenCalled_ShouldInitializeScheduler(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err, "Configure should not return an error")
}

func Test_RepeatTaskEvery_WhenCalled_ShouldScheduleTask(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err)

	err = sch.RepeatTaskEvery(time.Second, func() {
		log.Println("Task executed every second")
	})
	require.NoError(t, err)

	sch.Start()
	time.Sleep(3 * time.Second)

	err = sch.Shutdown()
	require.NoError(t, err)
}

func Test_StartAt_WhenCalled_ShouldScheduleDelayedTask(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err)

	err = sch.StartAt(2*time.Second, func() {
		log.Println("Task executed after delay")
	})
	require.NoError(t, err)

	sch.Start()
	time.Sleep(5 * time.Second)

	err = sch.Shutdown()
	require.NoError(t, err)
}

func Test_Start_WhenCalled_ShouldStartScheduler(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		sch.Start()
	})
}

func Test_StopJobs_WhenCalled_ShouldStopAllJobs(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err)

	err = sch.RepeatTaskEvery(time.Second, func() {
		log.Println("Task executed every second")
	})
	require.NoError(t, err)

	sch.Start()
	time.Sleep(3 * time.Second)

	err = sch.StopJobs()
	require.NoError(t, err)
}

func Test_Shutdown_WhenCalled_ShouldShutdownScheduler(t *testing.T) {
	sch := scheduler.New()
	err := sch.Configure()
	require.NoError(t, err)

	sch.Start()

	err = sch.Shutdown()
	require.NoError(t, err)
}
