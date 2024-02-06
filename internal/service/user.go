package service

import (
	"fmt"
	"time"

	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/pkg/mongodb"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/sso"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct{}

func (User) FindByGitHubID(ctx *appcontext.AppContext, githubID string) (*mongodb.User, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][user] find by GitHub id")
	defer span.Finish()

	var (
		d         = dao.User{}
		condition = bson.M{
			"github.id": githubID,
		}
	)

	return d.FindOneByCondition(ctx, condition)
}

func (User) FindByGoogleID(ctx *appcontext.AppContext, googleID string) (*mongodb.User, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][user] find by Google id")
	defer span.Finish()

	var (
		d         = dao.User{}
		condition = bson.M{
			"google.id": googleID,
		}
	)

	return d.FindOneByCondition(ctx, condition)
}

func (User) CreateWithGitHubData(ctx *appcontext.AppContext, githubData sso.GitHubUserData) (*mongodb.User, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][user] create with GitHub data")
	defer span.Finish()

	var (
		d = dao.User{}
	)

	user := mongodb.User{
		ID:     mongodb.NewObjectID(),
		Name:   githubData.Name,
		Avatar: githubData.Avatar,
		Google: nil,
		GitHub: &mongodb.UserSocialProviderInformation{
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

func (User) CreateWithGoogleData(ctx *appcontext.AppContext, googleData sso.GoogleUserData) (*mongodb.User, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][user] create with Google data")
	defer span.Finish()

	var (
		d = dao.User{}
	)

	user := mongodb.User{
		ID:     mongodb.NewObjectID(),
		Name:   googleData.Name,
		Avatar: googleData.Avatar,
		GitHub: nil,
		Google: &mongodb.UserSocialProviderInformation{
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

func (User) FindOrCreateWithPlatformID(ctx *appcontext.AppContext, platform string, id string) (*mongodb.User, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][user] find or create with platform id")
	defer span.Finish()

	var (
		d         = dao.User{}
		condition = bson.M{
			fmt.Sprintf("platform.%s", platform): id,
		}
	)

	// find
	user, err := d.FindOneByCondition(ctx, condition)
	if err != nil {
		return nil, err
	}

	// return if user existed
	if user != nil {
		return user, nil
	}

	// create new user
	user = &mongodb.User{
		ID:        mongodb.NewObjectID(),
		Name:      "Anonymous",
		Avatar:    "",
		Platform:  mongodb.UserPlatform{},
		Google:    nil,
		GitHub:    nil,
		CreatedAt: time.Now(),
	}

	switch platform {
	case config.Platform.Telegram:
		user.Platform.Telegram = id
	case config.Platform.Slack:
		user.Platform.Slack = id
	case config.Platform.Discord:
		user.Platform.Discord = id
	}

	if err = d.InsertOne(ctx, *user); err != nil {
		return nil, err
	}
	return user, nil
}
