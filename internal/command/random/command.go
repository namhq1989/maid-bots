package random

import (
	"github.com/namhq1989/maid-bots/content"
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

	ctx.Logger.Info("receive: /random", appcontext.Fields{
		"message":   c.message,
		"platform":  c.platform,
		"arguments": arguments,
	})

	var (
		l    = len(arguments)
		text = "invalid command"
	)

	if l == 0 {
		// /random

		// just skip it and respond the content of `/help random` command
		return content.Command.Help.Random
	} else if l == 1 {
		// 	/monitor $arg1

		if arguments[0] == appcommand.RandomTargets.Number.Name {
			return content.Command.Help.RandomNumber
		} else if arguments[0] == appcommand.RandomTargets.String.Name {
			return content.Command.Help.RandomString
		} else {
			return content.Command.Help.Random
		}
	} else {
		// 	/monitor $arg1 $arg2

		switch arguments[0] {
		case appcommand.RandomTargets.Number.Name:
			h := Number{
				Message: c.message,
			}
			text = h.Process(ctx)
		case appcommand.RandomTargets.String.Name:
			h := String{
				Message: c.message,
				Target:  arguments[1],
			}
			text = h.Process(ctx)
		}

		return text
	}
}
