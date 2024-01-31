package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/maid-bots/platform/telegram/command/example"

	"github.com/namhq1989/maid-bots/platform/telegram/command"

	"github.com/namhq1989/maid-bots/config"

	"github.com/namhq1989/maid-bots/platform/telegram/command/help"
	"github.com/namhq1989/maid-bots/platform/telegram/command/monitor"

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

	// set commands
	_, _ = b.SetMyCommands(context.Background(), &bot.SetMyCommandsParams{
		Commands:     command.Commands,
		Scope:        models.BotCommandScope(&models.BotCommandScopeDefault{}),
		LanguageCode: "en",
	})

	// help
	b.RegisterHandler(bot.HandlerTypeMessageText, config.Commands.Help.WithSlash, bot.MatchTypePrefix, help.Handler)

	// example
	b.RegisterHandler(bot.HandlerTypeMessageText, config.Commands.Example.WithSlash, bot.MatchTypePrefix, example.Handler)

	// monitor
	b.RegisterHandler(bot.HandlerTypeMessageText, config.Commands.Monitor.WithSlash, bot.MatchTypePrefix, monitor.Handler)

	// text
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, defaultHandler)

	go b.Start(context.Background())

	fmt.Printf("⚡️ [telegram]: initialized \n")
}
