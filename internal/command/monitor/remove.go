package monitor

import (
	"errors"
	"fmt"

	"github.com/namhq1989/maid-bots/internal/service"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Remove struct {
	Arguments map[string]string
	Platform  string
	ChatID    string
	User      modelcommand.User
}

func (c Remove) Process(ctx *appcontext.AppContext) (string, error) {
	id := c.Arguments[appcommand.MonitorParameters.ID]
	if id == "" {
		return "", errors.New("id is required")
	}

	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// remove
	if err = monitorSvc.DeleteByID(ctx, id, user.ID); err != nil {
		return "", err
	}

	return fmt.Sprintf("target `%s` has been successfully removed", id), nil
}
