package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
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

	// random
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, generateRegexp(appcommand.Root.Random.WithSlash), randomHandler)

	// unrecognized
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, unrecognizedHandler)

	go b.Start(context.Background())

	fmt.Printf("⚡️ [telegram]: initialized \n")
}

func respond(ctx *appcontext.AppContext, b *bot.Bot, update *models.Update, command, result string) {
	span := sentryio.NewSpan(ctx.Context, "[platform][telegram] respond")
	defer span.Finish()

	// respond
	if _, err := b.SendMessage(ctx.Context, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: &isLinkPreviewDisable,
		},
		Text: result,
	}); err != nil {
		ctx.Logger.Error(fmt.Sprintf("send /%s response", command), err, appcontext.Fields{})
	}
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
	{
		Command:     appcommand.Root.Example.Name,
		Description: appcommand.Root.Example.Description,
	},
}
