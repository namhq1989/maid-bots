package unrecognized

import (
	"fmt"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/internal/command/help"
	"github.com/namhq1989/maid-bots/internal/command/monitor"
	"github.com/namhq1989/maid-bots/internal/command/random"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
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

	// sometimes users click on the handler from "/help" content
	// and for some reason this library can't process the handler
	// so this function will re-check again to make sure the message can handle correctly
	cmd := appcommand.ExtractCommand(c.payload.Message)
	if !appcommand.IsRootCommandValid(cmd) {
		ctx.Logger.Info("receive: unrecognized message", appcontext.Fields{
			"message":   c.payload.Message,
			"platform":  c.payload.Platform,
			"arguments": arguments,
		})

		return modelcommand.Result{Text: "invalid command"}
	}

	// apm transaction
	t := sentryio.NewTransaction(ctx.Context, cmd, map[string]string{
		"platform":  c.payload.Platform,
		"message":   c.payload.Message,
		"userId":    c.payload.User.ID,
		"username":  c.payload.User.Username,
		"requestId": ctx.RequestID,
		"traceId":   ctx.TraceID,
	})
	defer t.Finish()

	// re-assign context
	ctx.Context = t.Context()

	switch cmd {
	case appcommand.Root.Help.WithSlash:
		return help.ProcessMessage(ctx, c.payload)
	case appcommand.Root.Monitor.WithSlash:
		return monitor.ProcessMessage(ctx, c.payload)
	case appcommand.Root.Random.WithSlash:
		return random.ProcessMessage(ctx, c.payload)
	}

	ctx.Logger.Info(fmt.Sprintf("receive: %s message", cmd), appcontext.Fields{
		"message":   c.payload.Message,
		"platform":  c.payload.Platform,
		"arguments": arguments,
	})

	return modelcommand.Result{
		Text: "invalid command",
	}
}
