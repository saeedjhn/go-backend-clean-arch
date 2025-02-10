package main

import (
	"context"
	"log"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/taskqueue"
)

const _concurrency = 10

func main() {
	taskType := "send_welcome_email_at"

	taskQueue := taskqueue.New(taskqueue.Config{
		Concurrency: _concurrency,
		RedisConfig: taskqueue.RedisConfig{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "password123",
		},
	})

	taskQueue.RegisterHandler(taskType, func(_ context.Context, task taskqueue.Task) error {
		t := task.Type
		email := string(task.Payload)

		log.Printf("Sending welcome email to - %s:  %s...", t, email)

		return nil
	})

	email := "user@example.com"
	if err := taskQueue.EnqueueTaskAt(
		taskType,
		[]byte(email),
		taskqueue.WithAt(time.Now().Add(30*time.Second)),
	); err != nil {
		log.Fatalf("Failed to enqueue task for email %s: %v", email, err)
	}

	go func() {
		err := taskQueue.StartServer()
		if err != nil {
			log.Fatalf("Failed to start task queue server: %v", err)
		}
	}()

	select {}
}

// func main() {
// 	taskType := "send_welcome_email_in"
//
// 	taskQueue := taskqueue.New(taskqueue.Config{
// 		Concurrency: _concurrency,
// 		RedisConfig: taskqueue.RedisConfig{
// 			Host:     "127.0.0.1",
// 			Port:     "6379",
// 			Password: "password123",
// 		},
// 	})
//
// 	taskQueue.RegisterHandler(taskType, func(_ context.Context, task taskqueue.Task) error {
// 		t := task.Type
// 		email := string(task.Payload)
//
// 		log.Printf("Sending welcome email to - %s:  %s...", t, email)
//
// 		return nil
// 	})
//
// 	email := "user@example.com"
// 	if err := taskQueue.EnqueueTaskIn(taskType, []byte(email), taskqueue.WithDelay(20*time.Second)); err != nil {
// 		log.Fatalf("Failed to enqueue task for email %s: %v", email, err)
// 	}
//
// 	go func() {
// 		err := taskQueue.StartServer()
// 		if err != nil {
// 			log.Fatalf("Failed to start task queue server: %v", err)
// 		}
// 	}()
//
// 	select {}
// }

// func main() {
// 	taskType := "send_welcome_email_schedule"
//
// 	taskQueue := taskqueue.New(taskqueue.Config{
// 		Concurrency: _concurrency,
// 		RedisConfig: taskqueue.RedisConfig{
// 			Host:     "127.0.0.1",
// 			Port:     "6379",
// 			Password: "password123",
// 		},
// 	})
//
// 	taskQueue.RegisterHandler(taskType, func(_ context.Context, task taskqueue.Task) error {
// 		t := task.Type
// 		email := string(task.Payload)
//
// 		log.Printf("Sending welcome email to - %s:  %s...", t, email)
//
// 		return nil
// 	})
//
// 	email := "user@example.com"
// 	if err := taskQueue.EnqueueTask(taskType, []byte(email)); err != nil {
// 		log.Fatalf("Failed to enqueue task for email %s: %v", email, err)
// 	}
//
// 	schedule := "* * * * *"
// 	if err := taskQueue.ScheduleTask(schedule, taskType, []byte(email)); err != nil {
// 		log.Fatalf("Failed to schedule task for email %s at schedule %s: %v", email, schedule, err)
// 	} else {
// 		log.Println("Task scheduled successfully!")
// 	}
//
// 	go func() {
// 		err := taskQueue.StartScheduler()
// 		if err != nil {
// 			log.Fatalf("Failed to start task scheduler: %v", err)
// 		}
// 	}()
//
// 	go func() {
// 		err := taskQueue.StartServer()
// 		if err != nil {
// 			log.Fatalf("Failed to start task queue server: %v", err)
// 		}
// 	}()
//
// 	select {}
// }
