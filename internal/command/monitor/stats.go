package monitor

import (
	"errors"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"

	"github.com/namhq1989/maid-bots/internal/service"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Stats struct {
	Arguments map[string]string
	Platform  string
	ChatID    string
	User      modelcommand.User
}

func (c Stats) Process(ctx *appcontext.AppContext) (*modelresponse.Stats, error) {
	id := c.Arguments[appcommand.MonitorParameters.ID]
	if id == "" {
		return nil, errors.New("id is required")
	}

	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return nil, err
	}

	// find
	stats, err := monitorSvc.StatsByCode(ctx, user.ID, id)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
