package dao

import (
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Monitor struct{}

func (Monitor) InsertOne(ctx *appcontext.AppContext, doc mongodb.Monitor) error {
	var (
		col = mongodb.MonitorCol()
	)

	_, err := col.InsertOne(ctx.Context, doc)
	if err != nil {
		ctx.Logger.Error("Monitor insert one", err, appcontext.Fields{"doc": doc})
	}
	return err
}

func (Monitor) CountByCondition(ctx *appcontext.AppContext, condition interface{}) (int64, error) {
	var (
		col = mongodb.MonitorCol()
	)

	count, err := col.CountDocuments(ctx.Context, condition)
	if err != nil {
		ctx.Logger.Error("Monitor count by condition", err, appcontext.Fields{"condition": condition})
	}

	return count, err
}
