package dao

import (
	"errors"

	"github.com/namhq1989/maid-bots/util/appcontext"

	modelmongodb "github.com/namhq1989/maid-bots/internal/models/mongodb"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{}

func (User) FindOneByCondition(ctx *appcontext.AppContext, condition interface{}, opts ...*options.FindOneOptions) (doc *modelmongodb.User, err error) {
	var (
		col = mongodb.UserCol()
	)

	err = col.FindOne(ctx.Context, condition, opts...).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.Logger.Error("User find one by condition", err, appcontext.Fields{"condition": condition})
	}
	return doc, err
}

func (User) InsertOne(ctx *appcontext.AppContext, doc modelmongodb.User) error {
	var (
		col = mongodb.UserCol()
	)

	_, err := col.InsertOne(ctx.Context, doc)
	if err != nil {
		ctx.Logger.Error("User insert one", err, appcontext.Fields{"doc": doc})
	}
	return err
}
