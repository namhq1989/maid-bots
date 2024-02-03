package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
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
	result := random.ProcessMessage(ctx, update.Message.Text, config.Platform.Telegram, getUserID(update))

	// respond
	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: &isLinkPreviewDisable,
		},
		Text: result,
	}); err != nil {
		ctx.Logger.Error("send /random response", err, appcontext.Fields{})
	}
}
