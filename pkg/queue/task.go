package queue

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	taskTimeout time.Duration = 30 * time.Second
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
	options = append(options, asynq.MaxRetry(retryTimes))

	// task timeout
	options = append(options, asynq.Timeout(taskTimeout))

	// Enqueue task
	return i.Client.Enqueue(task, options...)
}

// ScheduledTask create new task and run at specific time
// cronSpec follow cron expression
// https://www.freeformatter.com/cron-expression-generator-quartz.html
func (i Instance) ScheduledTask(typename string, payload []byte, cronSpec string, retryTimes int) (string, error) {
	// create task and options
	task := asynq.NewTask(typename, payload)
	options := make([]asynq.Option, 0)

	// retry times
	if retryTimes < 0 {
		retryTimes = 0
	}
	options = append(options, asynq.MaxRetry(retryTimes))

	// task timeout
	options = append(options, asynq.Timeout(taskTimeout))

	return i.Scheduler.Register(cronSpec, task, options...)
}
