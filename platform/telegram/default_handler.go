package telegram

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/pkg/redis"
	"github.com/namhq1989/maid-bots/platform/telegram/command/help"
	"github.com/namhq1989/maid-bots/platform/telegram/command/monitor"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func defaultHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx  = appcontext.New(bgCtx)
		text = update.Message.Text
	)

	ctx.Logger.Info("receive message that is not a command", appcontext.Fields{
		"chat": update.Message.Chat.ID,
		"text": text,
	})

	// sometimes users click on the command from "/help" content
	// and for some reason this library can't process the command
	// so this function will re-check again to make sure the message is a command
	cmd, hasCmd := extractCommand(text)
	if hasCmd {
		switch cmd {
		case config.Commands.Help.WithSlash:
			help.Handler(bgCtx, b, update)
		case config.Commands.Monitor.WithSlash:
			monitor.Handler(bgCtx, b, update)
		}
	} else {
		// if it is not a command
		// then we need to check if it is the next step of the previous command
		value := redis.GetValueByKey(fmt.Sprintf("%d", update.Message.Chat.ID))
		// trim value since it has `"` in value
		value = strings.Trim(value, `"`)

		fmt.Println("value: ", value)
	}

	// switch value {
	// // domain
	// case command.DomainCheckCommand.Command:
	// 	command.DomainCheck{Domain: update.Message.Text}.Process(ctx, b, update)
	// default:
	// 	ctx.Logger.Text("no meaning message")
	// 	return
	// }
}

func extractCommand(input string) (string, bool) {
	re := regexp.MustCompile(`^/(\w+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return "/" + matches[1], true
	}

	// No match found
	return "", false
}
