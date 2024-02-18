package dao

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HealthCheckRecord struct{}

func (HealthCheckRecord) InsertOne(ctx *appcontext.AppContext, doc mongodb.HealthCheckRecord) error {
	span := sentryio.NewSpan(ctx.Context, "[dao][health check record] insert one")
	defer span.Finish()

	var (
		col = mongodb.HealthCheckRecordCol()
	)

	_, err := col.InsertOne(ctx.Context, doc)
	if err != nil {
		ctx.Logger.Error("HealthCheckRecord insert one", err, appcontext.Fields{"doc": doc})
	}
	return err
}

func (HealthCheckRecord) CountByCondition(ctx *appcontext.AppContext, condition interface{}) (int64, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][health check record] count by condition")
	defer span.Finish()

	var (
		col = mongodb.HealthCheckRecordCol()
	)

	count, err := col.CountDocuments(ctx.Context, condition)
	if err != nil {
		ctx.Logger.Error("HealthCheckRecord count by condition", err, appcontext.Fields{"condition": condition})
	}

	return count, err
}

func (HealthCheckRecord) FindOneByCondition(ctx *appcontext.AppContext, condition interface{}, opts ...*options.FindOneOptions) (doc *mongodb.HealthCheckRecord, err error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][health check record] find one by condition")
	defer span.Finish()

	var (
		col = mongodb.HealthCheckRecordCol()
	)

	err = col.FindOne(ctx.Context, condition, opts...).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.Logger.Error("HealthCheckRecord find one by condition", err, appcontext.Fields{"condition": condition})
		return nil, err
	}
	return doc, nil
}

func (HealthCheckRecord) DeleteManyByCondition(ctx *appcontext.AppContext, condition interface{}) error {
	span := sentryio.NewSpan(ctx.Context, "[dao][health check record] delete many by condition")
	defer span.Finish()

	var (
		col = mongodb.HealthCheckRecordCol()
	)

	_, err := col.DeleteMany(ctx.Context, condition)
	if err != nil {
		ctx.Logger.Error("HealthCheckRecord delete many by condition", err, appcontext.Fields{"condition": condition})
	}

	return err
}

type MetricsInTimeRange struct {
	AverageResponseTime float64 `bson:"averageResponseTime"`
	MaxResponseTime     float64 `bson:"maxResponseTime"`
	MinResponseTime     float64 `bson:"minResponseTime"`
	UptimePercentage    float64 `bson:"uptimePercentage"`
}

func (HealthCheckRecord) GetMetricsInTimeRange(ctx *appcontext.AppContext, ownerID primitive.ObjectID, code string, startTime, endTime time.Time) (*MetricsInTimeRange, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] get metrics in time range")
	defer span.Finish()

	var (
		col      = mongodb.HealthCheckRecordCol()
		pipeline = bson.A{
			bson.M{
				"$match": bson.M{
					"owner": ownerID,
					"code":  code,
					"createdAt": bson.M{
						"$gte": startTime,
						"$lte": endTime,
					},
				},
			},
			bson.M{
				"$group": bson.M{
					"_id":                 nil,
					"averageResponseTime": bson.M{"$avg": "$responseTimeInMs"},
					"maxResponseTime":     bson.M{"$max": "$responseTimeInMs"},
					"minResponseTime":     bson.M{"$min": "$responseTimeInMs"},
					"totalDuration":       bson.M{"$sum": 1},
					"uptimeDuration": bson.M{
						"$sum": bson.M{
							"$cond": bson.A{
								bson.M{"$eq": []interface{}{"$status", "up"}},
								1, // Add 1 if condition is true
								0, // Add 0 if condition is false
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"averageResponseTime": 1,
					"maxResponseTime":     1,
					"minResponseTime":     1,
					"uptimePercentage": bson.M{
						"$multiply": bson.A{
							bson.M{"$divide": []interface{}{"$uptimeDuration", "$totalDuration"}},
							100,
						},
					},
				},
			},
		}
	)

	// execute aggregation query
	cursor, err := col.Aggregate(ctx.Context, pipeline)
	if err != nil {
		ctx.Logger.Error("HealthCheckRecord get metrics in time range aggregation", err, appcontext.Fields{
			"ownerID": ownerID.Hex(),
			"code":    code,
			"start":   startTime.String(),
			"end":     endTime.String(),
		})
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx.Context) }()

	// process aggregation results
	var result MetricsInTimeRange
	if cursor.Next(ctx.Context) {
		if err = cursor.Decode(&result); err != nil {
			ctx.Logger.Error("HealthCheckRecord get metrics in time range decode result", err, appcontext.Fields{
				"ownerID": ownerID.Hex(),
				"code":    code,
				"start":   startTime.String(),
				"end":     endTime.String(),
			})
			return nil, err
		}
	}

	return &result, nil
}

type ResponseTimeChartData struct {
	Date string  `bson:"date"`
	Hour int     `bson:"hour"`
	Avg  float64 `bson:"avg"`
}

func (d HealthCheckRecord) GetResponseTimeChartDataInTimeRange(ctx *appcontext.AppContext, ownerID primitive.ObjectID, code string, startTime, endTime time.Time) ([]ResponseTimeChartData, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] get response time chart data in time range")
	defer span.Finish()

	var (
		col      = mongodb.HealthCheckRecordCol()
		pipeline = bson.A{
			bson.M{
				"$match": bson.M{
					"owner": ownerID,
					"code":  code,
					"createdAt": bson.M{
						"$gte": startTime,
						"$lte": endTime,
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"hour":             bson.M{"$hour": "$createdAt"},
					"date":             bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$createdAt"}},
					"responseTimeInMs": 1,
				},
			},
			bson.M{
				"$group": bson.M{
					"_id":             bson.M{"hour": "$hour", "date": "$date"},
					"avgResponseTime": bson.M{"$avg": "$responseTimeInMs"},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id":  0,
					"date": "$_id.date",
					"hour": "$_id.hour",
					"avg":  "$avgResponseTime",
				},
			},
			bson.M{
				"$sort": bson.M{"date": 1, "hour": 1},
			},
		}
	)

	// execute aggregation query
	cursor, err := col.Aggregate(ctx.Context, pipeline)
	if err != nil {
		ctx.Logger.Error("HealthCheckRecord get response time chart data in time range aggregation", err, appcontext.Fields{
			"ownerID": ownerID.Hex(),
			"code":    code,
			"start":   startTime.String(),
			"end":     endTime.String(),
		})
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx.Context) }()

	// Map the MongoDB result to ResponseData structs
	var result []ResponseTimeChartData
	for cursor.Next(ctx.Context) {
		var record bson.M
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}
		mappedData := d.mapChartData(record)
		if mappedData != nil {
			result = append(result, *mappedData)
		}
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d HealthCheckRecord) mapChartData(result bson.M) *ResponseTimeChartData {
	var data ResponseTimeChartData

	// Extract date, hour, and avg from the result
	date, ok := result["date"].(string)
	if !ok {
		return nil
	}
	hour, ok := result["hour"].(int32)
	if !ok {
		return nil
	}
	avg, ok := result["avg"].(float64)
	if !ok {
		return nil
	}

	// Assign values to ResponseData struct
	data.Date = date
	data.Hour = int(hour)
	data.Avg = avg

	return &data
}
