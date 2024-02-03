package unrecognized

import (
	"github.com/namhq1989/maid-bots/internal/command/example"
	"github.com/namhq1989/maid-bots/internal/command/help"
	"github.com/namhq1989/maid-bots/internal/command/monitor"
	"github.com/namhq1989/maid-bots/internal/command/random"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	message  string
	platform string
	userID   string
}

func (c command) process(ctx *appcontext.AppContext) string {
	var (
		arguments = appcommand.ExtractParameters(c.message)
	)

	// sometimes users click on the handler from "/help" content
	// and for some reason this library can't process the handler
	// so this function will re-check again to make sure the message is a handler
	cmd := appcommand.ExtractCommand(c.message)
	if cmd != "" {
		switch cmd {
		case appcommand.Root.Help.WithSlash:
			return help.ProcessMessage(ctx, c.message, c.platform)
		case appcommand.Root.Example.WithSlash:
			return example.ProcessMessage(ctx, c.message, c.platform)
		case appcommand.Root.Monitor.WithSlash:
			return monitor.ProcessMessage(ctx, c.message, c.platform, c.userID)
		case appcommand.Root.Random.WithSlash:
			return random.ProcessMessage(ctx, c.message, c.platform, c.userID)
		}
	}

	ctx.Logger.Info("receive: unrecognized message", appcontext.Fields{
		"message":   c.message,
		"platform":  c.platform,
		"arguments": arguments,
	})

	return "invalid command"
}
