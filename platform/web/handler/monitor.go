package handler

import (
	"github.com/labstack/echo/v4"
	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/internal/service"
	routepayload "github.com/namhq1989/maid-bots/platform/web/route/payload"
	"github.com/namhq1989/maid-bots/util/echocontext"
	"github.com/namhq1989/maid-bots/util/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct{}

// List godoc
// @tags Monitor
// @summary List
// @id monitor-list
// @accept json
// @produce json
// @router /monitors/list [get]
func (Monitor) List(c echo.Context) error {
	var (
		ctx        = echocontext.GetContext(c)
		ownerID, _ = primitive.ObjectIDFromHex("65c4c6f7c4c0c54bb33eedf0")
		query      = echocontext.GetQuery(c).(routepayload.MonitorList)
		s          = service.Monitor{}
	)

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
