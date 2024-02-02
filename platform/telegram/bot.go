package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/maid-bots/util/appcommand"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var isLinkPreviewDisable = true

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
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Help.WithSlash), helpHandler)

	// example
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Example.WithSlash), exampleHandler)

	// monitor
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Monitor.WithSlash), monitorHandler)

	// unrecognized
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, unrecognizedHandler)

	go b.Start(context.Background())

	fmt.Printf("⚡️ [telegram]: initialized \n")
}

var commands = []models.BotCommand{
	{
		Command:     appcommand.Root.Help.Name,
		Description: appcommand.Root.Help.Description,
	},
	{
		Command:     appcommand.Root.Monitor.Name,
		Description: appcommand.Root.Monitor.Description,
	},
	{
		Command:     appcommand.Root.Example.Name,
		Description: appcommand.Root.Example.Description,
	},
}
