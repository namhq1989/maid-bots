package telegram

import (
	"context"

	"github.com/namhq1989/maid-bots/internal/command/monitor"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func monitorHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx    = appcontext.New(bgCtx)
		result = monitor.ProcessMessage(ctx, update.Message.Text, config.Platform.Telegram)
	)

	// respond
	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		Text:      result,
	}); err != nil {
		ctx.Logger.Error("send /monitor message", err, appcontext.Fields{})
	}
}
