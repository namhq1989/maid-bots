package monitor

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/namhq1989/maid-bots/config"

	"github.com/go-telegram/bot"

	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	message   string
	platform  string
	userID    string
	argAction string
	argTarget string
	argValue  string
}

func (c command) process(ctx *appcontext.AppContext) string {
	var (
		arguments = appcommand.ExtractParameters(c.message)
	)

	ctx.Logger.Info("receive: /monitor", appcontext.Fields{
		"message":   c.message,
		"platform":  c.platform,
		"arguments": arguments,
	})

	var (
		l    = len(arguments)
		text = "invalid command"
	)

	if l == 0 || l == 1 {
		// /monitor

		// this command requires at least 2 arguments
		// just skip it and respond the content of `/help monitor` command
		return content.Command.Help.Monitor
	} else if l == 2 {
		// 	/monitor $arg1 $arg2

		if arguments[0] == appcommand.MonitorActions.List.Name {
			// if `$arg1` is equal to `list`
			c.argAction = arguments[0]
			c.argTarget = arguments[1]
			return c.list(ctx)
		} else {
			// for other scenarios with insufficient arguments, respond the content of `/example monitor` command
			return content.Command.Example.Monitor
		}
	} else {
		// 	/monitor $arg1 $arg2 $arg3
		c.argAction = arguments[0]
		c.argTarget = arguments[1]
		c.argValue = arguments[2]

		if err := c.validateArguments(ctx); err != nil {
			ctx.Logger.Error("failed to validate arguments", err, appcontext.Fields{})
			return err.Error()
		}

		switch c.argAction {
		case appcommand.MonitorActions.Check.Name:
			return c.check(ctx)
		case appcommand.MonitorActions.Register.Name:
			return c.register(ctx)
		case appcommand.MonitorActions.List.Name:
			return c.list(ctx)
		case appcommand.MonitorActions.Remove.Name:
			return c.remove(ctx)
		case appcommand.MonitorActions.Stats.Name:
			return c.stats(ctx)
		}
		return text
	}
}

func (c command) validateArguments(ctx *appcontext.AppContext) error {
	// exception for action "list" and target "all"
	if c.argAction == appcommand.MonitorActions.List.Name && c.argTarget == "all" {
		return nil
	}

	// action
	if !appcommand.IsMonitorActionValid(c.argAction) {
		return errors.New("invalid action, please check `/help monitor` for more information")
	}

	// target
	if !appcommand.IsMonitorTargetValid(c.argTarget) {
		return errors.New("invalid target, please check `/help monitor` for more information")
	}

	// value
	switch c.argTarget {
	case appcommand.MonitorTargets.Domain.Name:
		re := `^(https?://)?([a-zA-Z0-9_-]+\.){1,}[a-zA-Z]{2,}(/.*)?$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid domain format")
		}
	case appcommand.MonitorTargets.HTTP.Name:
		re := `^https?://.*$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid http url format")
		}
	case appcommand.MonitorTargets.TCP.Name:
		re := `^([a-zA-Z0-9_-]+\.){1,}[a-zA-Z]{2,}:\d+$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid tcp format")
		}
	case appcommand.MonitorTargets.ICMP.Name:
		re := `^([a-zA-Z0-9_-]+\.){1,}[a-zA-Z]{2,}$|^\d+\.\d+\.\d+\.\d+$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid icmp format")
		}
	}

	return nil
}

func (c command) check(ctx *appcontext.AppContext) string {
	check := Check{
		Target: c.argTarget,
		Value:  c.argValue,
	}

	// process
	result, err := check.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	// assign data
	result.Name = c.argValue
	result.Target = c.argTarget

	// convert data for platforms
	var text = ""
	switch c.platform {
	case config.Platform.Telegram:
		text = bot.EscapeMarkdown(result.ToTelegram())
	case config.Platform.Discord:
		text = "Discord result"
		// text = bot.EscapeMarkdown(result.ToDiscord())
	case config.Platform.Slack:
		text = "Slack result"
		// text = bot.EscapeMarkdown(result.ToSlack())
	}

	return text
}

func (c command) register(ctx *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("registering %s %s with user %s ...", c.argTarget, c.argValue, c.userID))
}

func (c command) list(ctx *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("listing monitoring domains of Telegram user id %s ...", c.userID))
}

func (c command) remove(ctx *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("removing %s %s with user %s ...", c.argTarget, c.argValue, c.userID))
}

func (c command) stats(ctx *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("getting stats of %s %s with user %s ...", c.argTarget, c.argValue, c.userID))
}
