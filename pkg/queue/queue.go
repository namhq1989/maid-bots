package queue

import (
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

// Instance ...
type Instance struct {
	Client    *asynq.Client
	Server    *asynq.ServeMux
	Scheduler *asynq.Scheduler
	Inspector *asynq.Inspector
}

var instance Instance

func Init(redisURL string, concurrency int) {
	// redis connection
	redisConn := getRedisConnFromURL(redisURL)

	// init instance
	instance.Server = initServer(redisConn, concurrency)
	instance.Scheduler = initScheduler(redisConn)
	instance.Client = initClient(redisConn)
	instance.Inspector = initInspector(redisConn)

	fmt.Printf("⚡️ [queue]: initialized \n")
}

func initServer(redisConn asynq.RedisClientOpt, concurrency int) *asynq.ServeMux {
	// set unrecognized for concurrency
	if concurrency == 0 {
		concurrency = 100
	}

	// unrecognized retry delay in 10s
	retryDelayFunc := func(n int, e error, t *asynq.Task) time.Duration {
		return 10 * time.Second
	}

	// init server
	server := asynq.NewServer(redisConn, asynq.Config{
		Concurrency:    concurrency,
		RetryDelayFunc: retryDelayFunc,
		Queues: map[string]int{
			queueCronjob: 10,
			queueDefault: 3,
		},
	})

	// init mux server
	mux := asynq.NewServeMux()

	// run server
	go func() {
		if err := server.Run(mux); err != nil {
			msg := fmt.Sprintf("error when initializing queue SERVER: %s", err.Error())
			panic(msg)
		}
	}()

	return mux
}

func initScheduler(redisConn asynq.RedisClientOpt) *asynq.Scheduler {
	// always run at HCM timezone
	l, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	// init scheduler
	scheduler := asynq.NewScheduler(redisConn, &asynq.SchedulerOpts{
		Location: l,
	})

	// run scheduler
	go func() {
		if err := scheduler.Run(); err != nil {
			msg := fmt.Sprintf("error when initializing queue SCHEDULER: %s", err.Error())
			panic(msg)
		}
	}()

	return scheduler
}

func initClient(redisConn asynq.RedisClientOpt) *asynq.Client {
	client := asynq.NewClient(redisConn)
	if client == nil {
		panic("error when initializing queue CLIENT")
	}
	return client
}

func initInspector(redisConn asynq.RedisClientOpt) *asynq.Inspector {
	return asynq.NewInspector(redisConn)
}

// GetInstance ...
func GetInstance() Instance {
	return instance
}
