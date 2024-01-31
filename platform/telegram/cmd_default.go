package telegram

import (
	"context"
	"fmt"
	"strings"

	"github.com/namhq1989/maid-bots/util/appcommand"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/pkg/redis"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

// TODO
// move default handler to internal/command

func defaultHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx  = appcontext.New(bgCtx)
		text = update.Message.Text
	)

	ctx.Logger.Info("receive message that is not a handler", appcontext.Fields{
		"chat": update.Message.Chat.ID,
		"text": text,
	})

	// sometimes users click on the handler from "/help" content
	// and for some reason this library can't process the handler
	// so this function will re-check again to make sure the message is a handler
	cmd, hasCmd := appcommand.ExtractCommand(text)
	if hasCmd {
		switch cmd {
		case appcommand.Root.Help.WithSlash:
			helpHandler(bgCtx, b, update)
		case appcommand.Root.Monitor.WithSlash:
			monitorHandler(bgCtx, b, update)
		case appcommand.Root.Example.WithSlash:
			exampleHandler(bgCtx, b, update)
		}
	} else {
		// if it is not a handler
		// then we need to check if it is the next step of the previous handler
		value := redis.GetValueByKey(fmt.Sprintf("%d", update.Message.Chat.ID))
		// trim value since it has `"` in value
		value = strings.Trim(value, `"`)

		fmt.Println("value: ", value)
	}

	// switch value {
	// // domain
	// case handler.DomainCheckCommand.Command:
	// 	handler.DomainCheck{Domain: update.Message.Text}.Process(ctx, b, update)
	// default:
	// 	ctx.Logger.Text("no meaning message")
	// 	return
	// }
}
