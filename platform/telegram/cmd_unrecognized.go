package telegram

import (
	"context"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/internal/command/unrecognized"
	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func unrecognizedHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	ctx := appcontext.New(bgCtx)

	// process
	payload := modelcommand.Payload{
		Platform: config.Platform.Telegram,
		Message:  update.Message.Text,
		User:     getUser(update),
	}
	result := unrecognized.ProcessMessage(ctx, payload)

	// respond
	respond(ctx, b, update, "unrecognized", result)
}
