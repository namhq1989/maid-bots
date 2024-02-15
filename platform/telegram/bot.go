package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/util/appcommand"
)

var (
	telegramBot          *bot.Bot
	isLinkPreviewDisable = true
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
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Help.WithSlash), helpHandler)

	// monitor
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Monitor.WithSlash), monitorHandler)

	// random
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Random.WithSlash), randomHandler)

	// unrecognized
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, unrecognizedHandler)

	// assign
	telegramBot = b

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
		Command:     appcommand.Root.Random.Name,
		Description: appcommand.Root.Random.Description,
	},
}
