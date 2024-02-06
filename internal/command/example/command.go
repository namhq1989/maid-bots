package example

import (
	"github.com/namhq1989/maid-bots/content"
	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	payload modelcommand.Payload
}

func (c command) process(ctx *appcontext.AppContext) string {
	var (
		arguments = appcommand.ExtractParameters(c.payload.Message)
	)

	ctx.Logger.Info("receive: /example", appcontext.Fields{
		"message":   c.payload.Message,
		"platform":  c.payload.Platform,
		"arguments": arguments,
	})

	var (
		text = content.Command.Example.Base
		l    = len(arguments)
	)

	if l == 1 {
		switch arguments[0] {
		case appcommand.Root.Monitor.Name:
			text = content.Command.Example.Monitor
		case appcommand.Root.Random.Name:
			text = content.Command.Example.Random
		}
	}

	return text
}
