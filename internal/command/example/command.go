package example

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

	ctx.Logger.Info("receive: /example", appcontext.Fields{
		"message":   c.message,
		"platform":  c.platform,
		"arguments": arguments,
	})

	var (
		text = content.Group.Example.Base
		l    = len(arguments)
	)

	if l == 1 {
		switch arguments[0] {
		case appcommand.Root.Monitor.Base:
			text = content.Group.Example.Monitor
		}
	}

	return text
}