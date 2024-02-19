package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/internal/command/random"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func randomHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	ctx := appcontext.New(bgCtx)

	// apm transaction
	t := sentryio.NewTransaction(bgCtx, appcommand.Root.Random.WithSlash, getAPMTransactionData(ctx, update))
	defer t.Finish()

	// re-assign context
	ctx.Context = t.Context()

	// process
	result := random.ProcessMessage(ctx, getPayload(update))

	// respond
	respond(ctx, b, update, appcommand.Root.Random.Name, result.Text)
}
