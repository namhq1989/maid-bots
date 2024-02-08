package service

import (
	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
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
