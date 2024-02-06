package telegram

import (
	"context"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"

	"github.com/namhq1989/maid-bots/internal/command/monitor"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func monitorHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	ctx := appcontext.New(bgCtx)

	// apm transaction
	t := sentryio.NewTransaction(bgCtx, appcommand.Root.Monitor.WithSlash, getAPMTransactionData(ctx, update))
	defer t.Finish()

	// re-assign context
	ctx.Context = t.Context()

	// process
	payload := modelcommand.Payload{
		Platform: config.Platform.Telegram,
		Message:  update.Message.Text,
		User:     getUser(update),
	}
	result := monitor.ProcessMessage(ctx, payload)

	// respond
	respond(ctx, b, update, appcommand.Root.Monitor.Name, result)
}
