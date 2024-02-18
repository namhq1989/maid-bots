package service

import (
	"strings"
	"time"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"

	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HealthCheckRecord struct{}

func (HealthCheckRecord) NewRecord(ctx *appcontext.AppContext, doc mongodb.HealthCheckRecord) error {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] new record")
	defer span.Finish()

	var (
		d = dao.HealthCheckRecord{}
	)

	return d.InsertOne(ctx, doc)
}

func (HealthCheckRecord) GetRecentRecordOfMonitor(ctx *appcontext.AppContext, code string) (*mongodb.HealthCheckRecord, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] get recent record of monitor")
	defer span.Finish()

	var (
		d         = dao.HealthCheckRecord{}
		condition = bson.M{
			"code": code,
		}
	)

	return d.FindOneByCondition(ctx, condition, &options.FindOneOptions{
		Sort: bson.M{
			"createdAt": -1,
		},
	})
}

func (HealthCheckRecord) DeleteByMonitorCode(ctx *appcontext.AppContext, code string, ownerID primitive.ObjectID) error {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] delete by monitor code")
	defer span.Finish()

	var (
		d         = dao.HealthCheckRecord{}
		condition = bson.D{
			{Key: "owner", Value: ownerID},
			{Key: "code", Value: strings.ToLower(code)},
		}
	)

	return d.DeleteManyByCondition(ctx, condition)
}

func (HealthCheckRecord) GetResponseTimeMetricsInTimeRange(ctx *appcontext.AppContext, ownerID primitive.ObjectID, code string, startTime, endTime time.Time) (*modelresponse.HealthCheckRecordResponseTimeMetrics, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] get response metrics in time range")
	defer span.Finish()

	var (
		d = dao.HealthCheckRecord{}
	)

	metrics, err := d.GetMetricsInTimeRange(ctx, ownerID, code, startTime, endTime)
	if err != nil {
		return nil, err
	}

	return &modelresponse.HealthCheckRecordResponseTimeMetrics{
		Average:          metrics.AverageResponseTime,
		Max:              metrics.MaxResponseTime,
		Min:              metrics.MinResponseTime,
		UptimePercentage: metrics.UptimePercentage,
	}, nil
}

func (HealthCheckRecord) GetResponseTimeChartDataInTimeRange(ctx *appcontext.AppContext, ownerID primitive.ObjectID, code string, startTime, endTime time.Time) ([]modelresponse.HealthCheckRecordResponseTimeChartData, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] get response time chart data in time range")
	defer span.Finish()

	var (
		d = dao.HealthCheckRecord{}
	)

	chart, err := d.GetResponseTimeChartDataInTimeRange(ctx, ownerID, code, startTime, endTime)
	if err != nil {
		return nil, err
	}

	var result = make([]modelresponse.HealthCheckRecordResponseTimeChartData, 0, len(chart))
	for _, v := range chart {
		result = append(result, modelresponse.HealthCheckRecordResponseTimeChartData{
			Date: v.Date,
			Hour: v.Hour,
			Avg:  v.Avg,
		})
	}

	return result, nil
}
