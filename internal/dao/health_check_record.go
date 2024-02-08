package dao

import (
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
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
