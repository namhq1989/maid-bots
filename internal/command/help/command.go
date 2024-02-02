package help

import (
	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	message  string
	platform string
}

func (c command) process(ctx *appcontext.AppContext) string {
	var (
		arguments = appcommand.ExtractParameters(c.message)
	)

	ctx.Logger.Info("receive: /help", appcontext.Fields{
		"message":   c.message,
		"platform":  c.platform,
		"arguments": arguments,
	})

	var (
		text = content.Command.Help.Base
		l    = len(arguments)
	)

	if l == 1 {
		switch arguments[0] {
		case appcommand.Root.Monitor.Name:
			text = content.Command.Help.Monitor
		case appcommand.Root.Example.Name:
			text = content.Command.Example.Base
		}
	}

	return text
}
