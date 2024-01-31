package help

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func Handler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx = appcontext.New(bgCtx)
	)

	impl{}.Process(ctx, b, update)
}

type impl struct{}

func (c impl) Process(ctx *appcontext.AppContext, b *bot.Bot, update *models.Update) {
	arguments := appcommand.ExtractParameters(update.Message.Text)

	ctx.Logger.Info("receive: /help", appcontext.Fields{
		"text":      update.Message.Text,
		"arguments": arguments,
	})

	var (
		text = content.Group.Help
		l    = len(arguments)
	)

	if l == 1 {
		switch arguments[0] {
		case config.Commands.Monitor.Base:
			text = content.Group.Monitor
		case config.Commands.Example.Base:
			text = content.Group.Example.Base
		}
	}

	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		Text:      text,
	}); err != nil {
		ctx.Logger.Error("send /help message", err, appcontext.Fields{})
	}
}
