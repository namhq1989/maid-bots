package telegram

import (
	"context"
	"fmt"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
		Commands:     commands,
		Scope:        models.BotCommandScope(&models.BotCommandScopeDefault{}),
		LanguageCode: "en",
	})

	// help
	b.RegisterHandler(bot.HandlerTypeMessageText, appcommand.Root.Help.WithSlash, bot.MatchTypePrefix, helpHandler)

	// example
	b.RegisterHandler(bot.HandlerTypeMessageText, appcommand.Root.Example.WithSlash, bot.MatchTypePrefix, exampleHandler)

	// monitor
	b.RegisterHandler(bot.HandlerTypeMessageText, appcommand.Root.Monitor.WithSlash, bot.MatchTypePrefix, monitorHandler)

	// text
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, defaultHandler)

	go b.Start(context.Background())

	fmt.Printf("⚡️ [telegram]: initialized \n")
}

var commands = []models.BotCommand{
	{
		Command:     appcommand.Root.Help.Base,
		Description: appcommand.Root.Help.Description,
	},
	{
		Command:     appcommand.Root.Monitor.Base,
		Description: appcommand.Root.Monitor.Description,
	},
	{
		Command:     appcommand.Root.Example.Base,
		Description: appcommand.Root.Example.Description,
	},
}
