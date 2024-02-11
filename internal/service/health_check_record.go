package service

import (
	"time"

	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	last5Minutes = time.Minute * 5
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

func (HealthCheckRecord) IsTargetDownRecently(ctx *appcontext.AppContext, targetCode string) (bool, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][health check record] is target down recently")
	defer span.Finish()

	var (
		d    = dao.HealthCheckRecord{}
		cond = bson.M{
			"code":   targetCode,
			"status": mongodb.HealthCheckRecordStatusDown,
			"createdAt": bson.M{
				"$gte": time.Now().Add(-last5Minutes),
			},
		}
	)

	total, err := d.CountByCondition(ctx, cond)
	return total > 0, err
}
