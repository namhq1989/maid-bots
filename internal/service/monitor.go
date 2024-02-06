package service

import (
	"time"

	"github.com/namhq1989/maid-bots/util/random"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/internal/dao"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	codeLength = 4
)

type Monitor struct{}

func (s Monitor) randomCode(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "[service] random code", "")
	defer span.Finish()

	var code = ""

	for {
		code = random.StringWithLength(ctx, codeLength)
		exists := s.isCodeExisted(ctx, code)
		if !exists {
			break
		}
	}

	return code
}

func (Monitor) isCodeExisted(ctx *appcontext.AppContext, code string) bool {
	var (
		d         = dao.Monitor{}
		condition = bson.M{
			"code": code,
		}
	)

	total, _ := d.CountByCondition(ctx, condition)
	return total > 0
}

func (Monitor) IsTargetExisted(ctx *appcontext.AppContext, monitorType mongodb.MonitorType, target string, userID primitive.ObjectID) bool {
	var (
		d         = dao.Monitor{}
		condition = bson.D{
			{"owner", userID},
			{"type", monitorType},
			{"target", target},
		}
	)

	total, _ := d.CountByCondition(ctx, condition)
	return total > 0
}

func (s Monitor) CreateDomain(ctx *appcontext.AppContext, domain, scheme string, ownerID primitive.ObjectID) (*mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service] create domain", "")
	defer span.Finish()

	var (
		d    = dao.Monitor{}
		code = s.randomCode(ctx)
	)

	var doc = mongodb.Monitor{
		ID:     mongodb.NewObjectID(),
		Owner:  ownerID,
		Code:   code,
		Target: domain,
		Type:   mongodb.MonitorTypeDomain,
		Data: mongodb.MonitorMetadata{
			Scheme: scheme,
		},
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
