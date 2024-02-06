package dao

import (
	"errors"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{}

func (User) FindOneByCondition(ctx *appcontext.AppContext, condition interface{}, opts ...*options.FindOneOptions) (doc *mongodb.User, err error) {
	span := sentryio.NewSpan(ctx.Context, "[dao][user] find one by condition")
	defer span.Finish()

	var (
		col = mongodb.UserCol()
	)

	err = col.FindOne(ctx.Context, condition, opts...).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.Logger.Error("User find one by condition", err, appcontext.Fields{"condition": condition})
		return nil, err
	}
	return doc, nil
}

func (User) InsertOne(ctx *appcontext.AppContext, doc mongodb.User) error {
	span := sentryio.NewSpan(ctx.Context, "[dao][user] insert one")
	defer span.Finish()

	var (
		col = mongodb.UserCol()
	)

	_, err := col.InsertOne(ctx.Context, doc)
	if err != nil {
		ctx.Logger.Error("User insert one", err, appcontext.Fields{"doc": doc})
	}
	return err
}
