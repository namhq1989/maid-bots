package telegram

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func SendMessage(ctx *appcontext.AppContext, chatID, message string) {
	span := sentryio.NewSpan(ctx.Context, "[platform][telegram] send message")
	defer span.Finish()

	if _, err := telegramBot.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    chatID,
		Text:      bot.EscapeMarkdown(message),
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: &isLinkPreviewDisable,
		},
	}); err != nil {
		ctx.Logger.Error("[telegram] send message", err, appcontext.Fields{
			"chatId":  chatID,
			"message": message,
		})
	}
}
