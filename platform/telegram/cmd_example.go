package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/internal/command/example"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func exampleHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		// apm transaction
		t = sentryio.NewTransaction(bgCtx, appcommand.Root.Example.WithSlash, getAPMTransactionData(update))

		ctx    = appcontext.New(t.Context())
		result = example.ProcessMessage(ctx, update.Message.Text, config.Platform.Telegram)
	)
	defer t.Finish()

	// respond
	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: &isLinkPreviewDisable,
		},
		Text: result,
	}); err != nil {
		ctx.Logger.Error("send /example response", err, appcontext.Fields{})
	}
}
