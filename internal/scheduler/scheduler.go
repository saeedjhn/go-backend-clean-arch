package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"log"
	"sync"
	"time"
)

type Config struct {
}

type Scheduler struct {
	sch    gocron.Scheduler
	config Config
}

func New() Scheduler {
	sch, _ := gocron.NewScheduler() // Check err
	return Scheduler{
		//config: config,
		sch: sch,
	}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	job, _ := s.sch.NewJob( // Check err
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(func(arg string) {
			// Do something
			log.Println(arg)
		}, "Hello"),
	)

	log.Println("ID for job", job.ID())

	// start the scheduler
	s.sch.Start()

	<-done
	// wait to finish job
	log.Println("stop scheduler..")

	// when you're done, shut it down
	err := s.sch.Shutdown()
	if err != nil {
		log.Println("shutdown err, ", err.Error())
	}
}

//func (s Scheduler) SetJob(time time.Duration, task func()) {
//	job, _ := s.sch.NewJob( // Check err
//		gocron.DurationJob(time),
//		gocron.NewTask(task),
//	)
//
//	log.Println("ID for job", job.ID())
//}

//func (s Scheduler) SetJob(
//	done <-chan bool,
//	wg *sync.WaitGroup,
//	time time.Duration,
//	function func(),
//) {
//	defer wg.Done()
//
//	job, _ := s.sch.NewJob( // Check err
//		gocron.DurationJob(time),
//		gocron.NewTask(func() {
//			function()
//		}),
//	)
//
//	log.Println("ID for job", job.ID())
//
//	// start the scheduler
//	s.sch.Start()
//
//	<-done
//	// wait to finish job
//	log.Println("stop scheduler..")
//
//	// when you're done, shut it down
//	err := s.sch.Shutdown()
//	if err != nil {
//		log.Println("shutdown err, ", err.Error())
//	}
//}
