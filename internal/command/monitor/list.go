package monitor

import (
	"fmt"
	"strconv"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type List struct {
	Arguments  map[string]string
	Platform   string
	ChatID     string
	Parameters *ListParameters
	User       modelcommand.User
}

type ListParameters struct {
	Type    string
	Keyword string
	Page    int64
}

func (c *List) mapParametersToStruct(ctx *appcontext.AppContext) error {
	span := sentryio.NewSpan(ctx.Context, "map parameters to struct")
	defer span.Finish()

	var (
		parameters         = ListParameters{}
		totalMatchedParams = 0
	)

	for key, value := range c.Arguments {
		switch key {
		case appcommand.MonitorParameters.Type:
			parameters.Type = value
			totalMatchedParams++
		case appcommand.MonitorParameters.Keyword:
			parameters.Keyword = value
			totalMatchedParams++
		case appcommand.MonitorParameters.Page:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Page = int64(v)
			totalMatchedParams++
		default:
			continue
		}
	}

	c.Parameters = &parameters
	return nil
}

func (c *List) find(ctx *appcontext.AppContext) (string, error) {
	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user by platform id
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// find
	monitors, err := monitorSvc.FindByUserID(ctx, user.ID, service.MonitorFindByUserIDFilter{
		Type:    c.Parameters.Type,
		Keyword: c.Parameters.Keyword,
		Page:    c.Parameters.Page,
	})
	if err != nil {
		return "", err
	}

	if len(monitors) == 0 {
		return "No monitors found", nil
	}

	// setup content
	var content string
	content += "```\n"
	content += fmt.Sprintf(" %-5s | %-8s | %-50s \n", "Id", "Type", "Target")
	content += "-------|----------|-----------------\n"
	for _, monitor := range monitors {
		content += fmt.Sprintf(" %-5s | %-8s | %-50s \n", monitor.Code, monitor.Type, monitor.Target)
	}
	content += "```\n"

	return content, nil
}

func (c *List) Process(ctx *appcontext.AppContext) (string, error) {
	// map to struct
	if err := c.mapParametersToStruct(ctx); err != nil {
		return "", err
	}

	// find and return
	return c.find(ctx)
}
