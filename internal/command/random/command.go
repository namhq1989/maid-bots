package random

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
		arguments = appcommand.ExtractArguments(c.payload.Message)
	)

	ctx.Logger.Info("receive: /random", appcontext.Fields{
		"message":   c.payload.Message,
		"platform":  c.payload.Platform,
		"arguments": arguments,
	})

	var (
		l    = len(arguments)
		text = "invalid command"
	)

	if l == 0 {
		return content.Command.Help.Random
	}

	if arguments[appcommand.RandomNumberParameters.Type] == appcommand.RandomTypes.Number {
		h := Number{
			Arguments: arguments,
		}
		text = h.Process(ctx)
	} else if arguments[appcommand.RandomNumberParameters.Type] == appcommand.RandomTypes.String {
		h := String{
			Arguments: arguments,
		}
		text = h.Process(ctx)
	}

	return text
}
