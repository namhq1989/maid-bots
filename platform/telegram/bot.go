package telegram

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/namhq1989/maid-bots/pkg/redis"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/namhq1989/maid-bots/platform/telegram/command"

	"github.com/go-telegram/bot/models"

	"github.com/go-telegram/bot"
)

func Init(enabled bool, token string) {
	if !enabled {
		fmt.Printf("⚡️ [telegram]: disabled \n")
		return
	}

	opts := []bot.Option{
		bot.WithCheckInitTimeout(10 * time.Second),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	commands := make([]models.BotCommand, 0)
	for _, v := range command.Commands {
		commands = append(commands, v.BotCommand)
	}
	_, _ = b.SetMyCommands(context.Background(), &bot.SetMyCommandsParams{
		Commands:     commands,
		Scope:        models.BotCommandScope(&models.BotCommandScopeDefault{}),
		LanguageCode: "en",
	})

	// help
	b.RegisterHandler(bot.HandlerTypeMessageText, command.HelpCommand.WithSlash, bot.MatchTypePrefix, command.HelpHandler)

	// domain
	b.RegisterHandler(bot.HandlerTypeMessageText, command.DomainCheckCommand.WithSlash, bot.MatchTypePrefix, command.DomainCheckHandler)

	// not a command
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, defaultHandler)

	go b.Start(context.Background())

	fmt.Printf("⚡️ [telegram]: initialized \n")
}

func defaultHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx = appcontext.New(bgCtx)
	)

	ctx.Logger.Info("receive message that is not a command", appcontext.Fields{
		"chat": update.Message.Chat.ID,
	})

	value := redis.GetValueByKey(fmt.Sprintf("%d", update.Message.Chat.ID))
	// trim value since it has `"` in value
	value = strings.Trim(value, `"`)

	switch value {
	// domain
	case command.DomainCheckCommand.Command:
		command.DomainCheck{Domain: update.Message.Text}.Process(ctx, b, update)
	default:
		ctx.Logger.Text("no meaning message")
		return
	}
}
