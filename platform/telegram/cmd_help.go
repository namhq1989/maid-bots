package telegram

import (
	"context"

	"github.com/namhq1989/maid-bots/internal/command/help"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func helpHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	ctx := appcontext.New(bgCtx)

	// apm transaction
	t := sentryio.NewTransaction(bgCtx, appcommand.Root.Help.WithSlash, getAPMTransactionData(ctx, update))
	defer t.Finish()

	// re-assign context
	ctx.Context = t.Context()

	// process
	result := help.ProcessMessage(ctx, getPayload(update))

	// respond
	respond(ctx, b, update, appcommand.Root.Help.Name, result.Text)
}
