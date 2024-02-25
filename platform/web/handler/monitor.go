package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/response"
)

type Monitor struct{}

// List godoc
// @tags Monitor
// @summary List
// @id monitor-list
// @security ApiKeyAuth
// @accept json
// @produce json
// @router /monitors/list [get]
func (Monitor) List(c echo.Context) error {
	var (
		ctx    = echocontext.GetContext(c)
		userID = echocontext.GetUserID(c)
		query  = echocontext.GetQuery(c).(routepayload.MonitorList)
		s      = service.Monitor{}
	)

	// convert owner id
	ownerID, err := mongodb.ObjectIDFromString(userID)
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	docs, err := s.FindByOwnerID(ctx, ownerID, service.MonitorFindByUserIDFilter{
		Page: query.Page,
	})
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	var result = make([]modelresponse.Monitor, 0)
	for _, doc := range docs {
		result = append(result, modelresponse.Monitor{
			Code:      doc.Code,
			Type:      string(doc.Type),
			Target:    doc.Target,
			Interval:  doc.Interval,
			CreatedAt: modelresponse.NewTimeResponse(doc.CreatedAt),
		})
	}

	return response.R200(c, "", echo.Map{
		"monitors": result,
	})
}

// Stats godoc
// @tags Monitor
// @summary Stats
// @id monitor-stats
// @security ApiKeyAuth
// @accept json
// @produce json
// @router /monitors/stats [get]
func (Monitor) Stats(c echo.Context) error {
	var (
		ctx    = echocontext.GetContext(c)
		userID = echocontext.GetUserID(c)
		code   = c.Param("code")
		s      = service.Monitor{}
	)

	// convert owner id
	ownerID, err := mongodb.ObjectIDFromString(userID)
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	result, err := s.StatsByCode(ctx, ownerID, code)
	if err != nil {
		return response.R400(c, err.Error(), echo.Map{})
	}

	defer func() { _ = os.Remove(result.ChartImageName) }()

	return response.R200(c, "", echo.Map{
		"stats": result,
	})
}
