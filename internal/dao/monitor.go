package dao

import (
	"errors"

	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Monitor struct{}

func (Monitor) InsertOne(ctx *appcontext.AppContext, doc mongodb.Monitor) error {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] insert one")
	defer span.Finish()

	var (
		col = mongodb.MonitorCol()
	)

	_, err := col.InsertOne(ctx.Context, doc)
	if err != nil {
		ctx.Logger.Error("Monitor insert one", err, appcontext.Fields{"doc": doc})
	}
	return err
}

func (Monitor) FindOneByCondition(ctx *appcontext.AppContext, condition interface{}, opts ...*options.FindOneOptions) (doc *mongodb.Monitor, err error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] find one by condition")
	defer span.Finish()

	var (
		col = mongodb.MonitorCol()
	)

	err = col.FindOne(ctx.Context, condition, opts...).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.Logger.Error("Monitor find one by condition", err, appcontext.Fields{"condition": condition})
		return nil, err
	}
	return doc, nil
}

func (Monitor) FindByCondition(ctx *appcontext.AppContext, condition interface{}, opts ...*options.FindOptions) ([]mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] find by condition")
	defer span.Finish()

	var (
		col  = mongodb.MonitorCol()
		docs = make([]mongodb.Monitor, 0)
	)

	// find
	cursor, err := col.Find(ctx.Context, condition, opts...)
	if err != nil {
		ctx.Logger.Error("Monitor find by condition", err, appcontext.Fields{"condition": condition})
		return docs, err
	}

	// parse
	defer func() { _ = cursor.Close(ctx.Context) }()
	err = cursor.All(ctx.Context, &docs)
	if err != nil {
		ctx.Logger.Error("Monitor parse documents", err, appcontext.Fields{"condition": condition})
		return docs, err
	}

	return docs, nil
}

func (Monitor) CountByCondition(ctx *appcontext.AppContext, condition interface{}) (int64, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] count by condition")
	defer span.Finish()

	var (
		col = mongodb.MonitorCol()
	)

	count, err := col.CountDocuments(ctx.Context, condition)
	if err != nil {
		ctx.Logger.Error("Monitor count by condition", err, appcontext.Fields{"condition": condition})
	}

	return count, err
}

func (Monitor) UpdateOneByCondition(ctx *appcontext.AppContext, filter interface{}, data interface{}) error {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] update one by condition")
	defer span.Finish()

	var (
		col = mongodb.MonitorCol()
	)

	_, err := col.UpdateOne(ctx.Context, filter, data)
	if err != nil {
		ctx.Logger.Error("Monitor count by condition", err, appcontext.Fields{"filter": filter, "data": data})
	}

	return err
}

func (Monitor) DeleteOneByCondition(ctx *appcontext.AppContext, condition interface{}) (bool, error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][monitor] delete one by condition")
	defer span.Finish()

	var (
		col = mongodb.MonitorCol()
	)

	result, err := col.DeleteOne(ctx.Context, condition)
	if err != nil {
		ctx.Logger.Error("Monitor delete one by condition", err, appcontext.Fields{"condition": condition})
		return false, err
	}

	return result.DeletedCount > 0, nil
}
