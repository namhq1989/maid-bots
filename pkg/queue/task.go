package queue

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	queueDefault = "default"
	queueCronjob = "cronjob"

	taskTimeout   time.Duration = 30 * time.Second
	taskRetention               = 24 * 7 * time.Hour
)

// RunTask ...
func (i Instance) RunTask(queue string, payload []byte, retryTimes int) (*asynq.TaskInfo, error) {
	// create task and options
	task := asynq.NewTask(queue, payload)
	options := make([]asynq.Option, 0)

	// retry times
	if retryTimes < 0 {
		retryTimes = 0
	}

	// append options
	options = append(options,
		asynq.MaxRetry(retryTimes),
		asynq.Timeout(taskTimeout),
		asynq.Retention(taskRetention),
	)

	// Enqueue task
	return i.Client.Enqueue(task, options...)
}

// ScheduleTask create new task and run at specific time
// cronSpec follow cron expression
// https://www.freeformatter.com/cron-expression-generator-quartz.html
func (i Instance) ScheduleTask(typename string, payload []byte, cronSpec string, retryTimes int) (string, error) {
	// create task and options
	task := asynq.NewTask(typename, payload)
	options := make([]asynq.Option, 0)

	// retry times
	if retryTimes < 0 {
		retryTimes = 0
	}

	// append options
	options = append(options,
		asynq.Queue(queueCronjob),
		asynq.MaxRetry(retryTimes),
		asynq.Timeout(taskTimeout),
		asynq.Retention(taskRetention),
	)

	return i.Scheduler.Register(cronSpec, task, options...)
}

func (i Instance) RemoveScheduler(id string) error {
	return i.Scheduler.Unregister(id)
}
