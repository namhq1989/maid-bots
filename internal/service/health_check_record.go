package service

import (
	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson"
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

func (HealthCheckRecord) GetRecentRecordOfTarget(ctx *appcontext.AppContext, targetCode string) (*mongodb.HealthCheckRecord, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] get recent record of target")
	defer span.Finish()

	var (
		d    = dao.HealthCheckRecord{}
		cond = bson.M{
			"code": targetCode,
		}
	)

	return d.FindOneByCondition(ctx, cond, &options.FindOneOptions{
		Sort: bson.M{
			"createdAt": -1,
		},
	})
}
