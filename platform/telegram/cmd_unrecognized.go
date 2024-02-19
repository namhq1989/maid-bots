package telegram

import (
	"context"

	"github.com/namhq1989/maid-bots/internal/command/unrecognized"
	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func unrecognizedHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	ctx := appcontext.New(bgCtx)

	// process
	result := unrecognized.ProcessMessage(ctx, getPayload(update))

	// respond
	respond(ctx, b, update, "unrecognized", result.Text)
}
