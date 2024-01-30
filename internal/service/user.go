package service

import (
	"time"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/namhq1989/maid-bots/internal/dao"
	modelmongodb "github.com/namhq1989/maid-bots/internal/models/mongodb"
	"github.com/namhq1989/maid-bots/pkg/sso"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (User) FindByGitHubID(ctx *appcontext.AppContext, githubID string) (*modelmongodb.User, error) {
	var (
		d         = dao.User{}
		condition = bson.M{
			"github.id": githubID,
		}
	)

	return d.FindOneByCondition(ctx, condition)
}

func (User) FindByGoogleID(ctx *appcontext.AppContext, googleID string) (*modelmongodb.User, error) {
	var (
		d         = dao.User{}
		condition = bson.M{
			"google.id": googleID,
		}
	)

	return d.FindOneByCondition(ctx, condition)
}

func (User) CreateWithGitHubData(ctx *appcontext.AppContext, githubData sso.GitHubUserData) (*modelmongodb.User, error) {
	var (
		d = dao.User{}
	)

	user := modelmongodb.User{
		ID:     primitive.NewObjectID(),
		Name:   githubData.Name,
		Avatar: githubData.Avatar,
		Google: nil,
		GitHub: &modelmongodb.UserSocialProviderInformation{
			ID:     githubData.ID,
			Name:   githubData.Name,
			Email:  githubData.Email,
			Avatar: githubData.Avatar,
		},
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (User) CreateWithGoogleData(ctx *appcontext.AppContext, googleData sso.GoogleUserData) (*modelmongodb.User, error) {
	var (
		d = dao.User{}
	)

	user := modelmongodb.User{
		ID:     primitive.NewObjectID(),
		Name:   googleData.Name,
		Avatar: googleData.Avatar,
		GitHub: nil,
		Google: &modelmongodb.UserSocialProviderInformation{
			ID:     googleData.ID,
			Name:   googleData.Name,
			Email:  googleData.Email,
			Avatar: googleData.Avatar,
		},
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return &user, nil
}
