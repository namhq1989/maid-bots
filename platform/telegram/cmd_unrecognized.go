package telegram

import (
	"context"

	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/internal/command/unrecognized"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func unrecognizedHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		// apm transaction
		t = sentryio.NewTransaction(bgCtx, "unrecognized", getAPMTransactionData(update))

		ctx    = appcontext.New(t.Context())
		result = unrecognized.ProcessMessage(ctx, update.Message.Text, config.Platform.Telegram, getUserID(update))
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
		ctx.Logger.Error("send unrecognized response", err, appcontext.Fields{})
	}
}
