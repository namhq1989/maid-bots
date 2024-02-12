package dao

import (
	"errors"

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
