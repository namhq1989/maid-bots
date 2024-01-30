package command

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func HelpHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx = appcontext.New(bgCtx)
	)

	Help{}.Process(ctx, b, update)
}

type Help struct{}

func (c Help) Process(ctx *appcontext.AppContext, b *bot.Bot, update *models.Update) {
	ctx.Logger.Text("receive: /help")

	var text = `
		<b>Available Commands:</b>
	`

	text += c.generateBaseOnCommands()

	text += `
Feel free to use these commands to interact with the bot and manage your watching domains. If you need further assistance, use the /help command or reach out to the bot owner.
	`

	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeHTML,
		Text:      text,
	}); err != nil {
		ctx.Logger.Error("send /help message", err, appcontext.Fields{})
	}
}

func (Help) generateBaseOnCommands() string {
	s := ""
	for _, c := range Commands {
		s += fmt.Sprintf(`
			%s
			%s
		`, c.BotCommand.Description, c.Sample)
	}

	return s
}
