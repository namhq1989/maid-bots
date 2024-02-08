package job

import (
	"context"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"github.com/namhq1989/maid-bots/internal/command/monitor"
	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/queue"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type monitorCheckData struct {
	Task       string `json:"task"`
	CronSpec   string `json:"cronSpec"`
	Interval   int    `json:"interval"`
	RetryTimes int    `json:"retryTimes"`
}

var monitorCheck = struct {
	interval60Seconds monitorCheckData
}{
	interval60Seconds: monitorCheckData{
		Task:       "monitor:check:60s",
		CronSpec:   "*/1 * * * *",
		Interval:   60,
		RetryTimes: 3,
	},
}

type Monitor struct{}

func (Monitor) setup(q queue.Instance, data monitorCheckData) {
	var (
		ctx = appcontext.New(context.Background())
	)

	ctx.AddLogData(appcontext.Fields{
		"task":       data.Task,
		"cronSpec":   data.CronSpec,
		"retryTimes": data.RetryTimes,
	})

	b, _ := json.Marshal(data)
	entryID, err := q.ScheduleTask(data.Task, b, data.CronSpec, data.RetryTimes)
	if err != nil {
		ctx.Logger.Error("error when initializing cronjob", err, appcontext.Fields{})
		panic(err)
	}

	ctx.Logger.Info(fmt.Sprintf("[queue] task '%s' scheduled successfully", data.Task), appcontext.Fields{
		"entryId": entryID,
	})
}

func (j Monitor) check(bgCtx context.Context, task *asynq.Task) error {
	var (
		ctx  = appcontext.New(bgCtx)
		data monitorCheckData
	)

	// parse data
	if err := json.Unmarshal(task.Payload(), &data); err != nil {
		ctx.Logger.Error("cannot parse task payload", err, appcontext.Fields{
			"payload": string(task.Payload()),
		})
		return err
	}

	// process
	total, success, err := j.process(ctx, data.Interval)

	// write result
	result := map[string]interface{}{
		"total":   total,
		"success": success,
		"error":   "",
	}
	if err != nil {
		result["error"] = err.Error()
	}

	b, _ := json.Marshal(result)
	if _, err := task.ResultWriter().Write(b); err != nil {
		ctx.Logger.Error("cannot write task result", err, appcontext.Fields{})
	}

	return err
}

func (Monitor) process(ctx *appcontext.AppContext, intervalSeconds int) (total int64, success int64, err error) {
	// count total first
	var (
		d         = dao.Monitor{}
		condition = bson.M{
			"interval": intervalSeconds,
		}
	)

	total, err = d.CountByCondition(ctx, condition)
	if err != nil {
		return 0, 0, err
	}

	// loop
	var (
		hcrService       = service.HealthCheckRecord{}
		page       int64 = 0
		limit      int64 = 100
		failure    int64 = 0
	)

	for {
		offset := page * limit

		records, err := d.FindByCondition(ctx, condition, &options.FindOptions{
			Limit: &limit,
			Skip:  &offset,
			Sort:  bson.M{"createdAt": 1},
		})
		if err != nil {
			break
		}

		for _, record := range records {
			var (
				doc = mongodb.HealthCheckRecord{
					ID:        mongodb.NewObjectID(),
					Owner:     record.Owner,
					Type:      record.Type,
					Code:      record.Code,
					Target:    record.Target,
					CreatedAt: time.Now(),
				}
				c = monitor.Check{
					Target: string(record.Type),
					Value:  record.Target,
				}
			)

			result, err := c.Process(ctx)
			if err != nil {
				failure++
				doc.Status = mongodb.HealthCheckRecordStatusDown
				doc.Description = err.Error()
			} else {
				doc.Status = mongodb.HealthCheckRecordStatusUp
				doc.ResponseTimeInMs = result.ResponseTimeInMS
			}

			// save history
			_ = hcrService.NewRecord(ctx, doc)
		}

		// break if end of data
		if len(records) < int(limit) {
			break
		}

		// increase page
		page++
	}

	return
}
