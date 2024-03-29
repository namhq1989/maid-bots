package service

import (
	"errors"
	"strings"
	"time"

	"github.com/namhq1989/maid-bots/pkg/chart"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

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

func (s Monitor) randomCode(ctx *appcontext.AppContext, ownerID primitive.ObjectID) string {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] random code")
	defer span.Finish()

	var code string
	for {
		code = random.StringWithLength(ctx, codeLength)
		exists := s.isCodeExisted(ctx, code, ownerID)
		if !exists {
			break
		}
	}

	return strings.ToLower(code)
}

func (Monitor) isCodeExisted(ctx *appcontext.AppContext, code string, ownerID primitive.ObjectID) bool {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] is code existed")
	defer span.Finish()

	var (
		d         = dao.Monitor{}
		condition = bson.D{
			{Key: "owner", Value: ownerID},
			{Key: "code", Value: strings.ToLower(code)},
		}
	)

	total, _ := d.CountByCondition(ctx, condition)
	return total > 0
}

func (Monitor) IsTargetExisted(ctx *appcontext.AppContext, monitorType mongodb.MonitorType, target string, userID primitive.ObjectID) bool {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] is target existed")
	defer span.Finish()

	var (
		d         = dao.Monitor{}
		condition = bson.D{
			{Key: "owner", Value: userID},
			{Key: "type", Value: monitorType},
			{Key: "target", Value: target},
		}
	)

	total, _ := d.CountByCondition(ctx, condition)
	return total > 0
}

func (s Monitor) CreateDomain(ctx *appcontext.AppContext, domain, scheme string, ownerID primitive.ObjectID) (*mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] create domain")
	defer span.Finish()

	var (
		d    = dao.Monitor{}
		code = s.randomCode(ctx, ownerID)
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
		Interval:  mongodb.MonitorInterval60Seconds,
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

func (s Monitor) CreateHTTP(ctx *appcontext.AppContext, http string, ownerID primitive.ObjectID) (*mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] create http")
	defer span.Finish()

	var (
		d    = dao.Monitor{}
		code = s.randomCode(ctx, ownerID)
	)

	var doc = mongodb.Monitor{
		ID:        mongodb.NewObjectID(),
		Owner:     ownerID,
		Code:      code,
		Target:    http,
		Type:      mongodb.MonitorTypeHTTP,
		Interval:  mongodb.MonitorInterval30Seconds,
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

func (s Monitor) CreateTCP(ctx *appcontext.AppContext, tcp string, ownerID primitive.ObjectID) (*mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] create tcp")
	defer span.Finish()

	var (
		d    = dao.Monitor{}
		code = s.randomCode(ctx, ownerID)
	)

	var doc = mongodb.Monitor{
		ID:        mongodb.NewObjectID(),
		Owner:     ownerID,
		Code:      code,
		Target:    tcp,
		Type:      mongodb.MonitorTypeTCP,
		Interval:  mongodb.MonitorInterval60Seconds,
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

func (s Monitor) CreateICMP(ctx *appcontext.AppContext, icmp string, ownerID primitive.ObjectID) (*mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] create icmp")
	defer span.Finish()

	var (
		d    = dao.Monitor{}
		code = s.randomCode(ctx, ownerID)
	)

	var doc = mongodb.Monitor{
		ID:        mongodb.NewObjectID(),
		Owner:     ownerID,
		Code:      code,
		Target:    icmp,
		Type:      mongodb.MonitorTypeICMP,
		Interval:  mongodb.MonitorInterval60Seconds,
		CreatedAt: time.Now(),
	}

	if err := d.InsertOne(ctx, doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

type MonitorFindByUserIDFilter struct {
	Type    string
	Keyword string
	Page    int64
}

func (Monitor) FindByOwnerID(ctx *appcontext.AppContext, ownerID primitive.ObjectID, filter MonitorFindByUserIDFilter) ([]mongodb.Monitor, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] find by owner id")
	defer span.Finish()

	var (
		d               = dao.Monitor{}
		limit     int64 = 10
		skip            = limit * filter.Page
		condition       = bson.D{
			{Key: "owner", Value: ownerID},
		}
	)

	// set filter
	if filter.Type != "" {
		condition = append(condition, bson.E{Key: "type", Value: filter.Type})
	}

	// find options
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.M{"createdAt": -1},
	}

	return d.FindByCondition(ctx, condition, opts)
}

func (Monitor) DeleteByID(ctx *appcontext.AppContext, monitorID string, ownerID primitive.ObjectID) error {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] delete by id")
	defer span.Finish()

	if err := mongodb.RunTransaction(ctx, func(sessionCtx mongo.SessionContext) error {
		// assign context
		ctx.Context = sessionCtx

		// remove target by code
		{
			var (
				d         = dao.Monitor{}
				condition = bson.D{
					{Key: "owner", Value: ownerID},
					{Key: "code", Value: strings.ToLower(monitorID)},
				}
			)

			// delete
			isDeleted, err := d.DeleteOneByCondition(ctx, condition)
			if err != nil {
				return err
			}

			if !isDeleted {
				return errors.New("monitor not found")
			}
		}

		// remove health check record by monitor code
		{
			var (
				hcrSvc = HealthCheckRecord{}
			)

			if err := hcrSvc.DeleteByMonitorCode(ctx, monitorID, ownerID); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (Monitor) StatsByCode(ctx *appcontext.AppContext, ownerID primitive.ObjectID, code string) (*modelresponse.Stats, error) {
	span := sentryio.NewSpan(ctx.Context, "[service][monitor] stats")
	defer span.Finish()

	var (
		d         = dao.Monitor{}
		condition = bson.D{
			{Key: "owner", Value: ownerID},
			{Key: "code", Value: strings.ToLower(code)},
		}
	)

	// find monitor first
	monitor, err := d.FindOneByCondition(ctx, condition)
	if err != nil {
		return nil, err
	}
	if monitor == nil {
		return nil, errors.New("monitor not found")
	}

	// statistics
	var (
		hrcSvc    = HealthCheckRecord{}
		now       = time.Now()
		oneDayAgo = now.Add(-24 * time.Hour)
		result    = &modelresponse.Stats{}
	)

	// response time
	if result.ResponseTime, err = hrcSvc.GetResponseTimeMetricsInTimeRange(ctx, ownerID, code, oneDayAgo, now); err != nil {
		return nil, err
	}

	// assign data
	result.Monitor = modelresponse.Monitor{
		Code:      monitor.Code,
		Type:      string(monitor.Type),
		Target:    monitor.Target,
		Interval:  monitor.Interval,
		CreatedAt: modelresponse.NewTimeResponse(monitor.CreatedAt),
	}

	// chart
	result.Chart, _ = hrcSvc.GetResponseTimeChartDataInTimeRange(ctx, ownerID, code, oneDayAgo, now)

	c := chart.MonitorResponseTimeLine{Data: result.Chart}
	result.ChartImageName, _ = c.ToImage(ctx)

	return result, nil
}
