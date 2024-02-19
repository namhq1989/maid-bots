package help

import (
	"github.com/namhq1989/maid-bots/content"
	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	payload modelcommand.Payload
}

func (c command) process(ctx *appcontext.AppContext) modelcommand.Result {
	var (
		arguments = appcommand.ExtractArguments(c.payload.Message)
	)

	ctx.Logger.Info("receive: /help", appcontext.Fields{
		"message":   c.payload.Message,
		"platform":  c.payload.Platform,
		"arguments": arguments,
	})

	return modelcommand.Result{
		Text: content.Command.Help.Base,
	}
}
