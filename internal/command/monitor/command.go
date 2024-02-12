package monitor

import (
	"errors"
	"fmt"
	"regexp"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/go-telegram/bot"

	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	payload   modelcommand.Payload
	argAction string
	argTarget string
	argValue  string
}

func (c command) process(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] process")
	defer span.Finish()

	var (
		arguments = appcommand.ExtractParameters(c.payload.Message)
	)

	ctx.Logger.Info("receive: /monitor", appcontext.Fields{
		"platform":  c.payload.Platform,
		"message":   c.payload.Message,
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
			ctx.Logger.Error("validate arguments failed", err, appcontext.Fields{})
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
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] validate arguments")
	defer span.Finish()

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
		re := `^([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.[a-zA-Z]{2,6}$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid domain format")
		}
	case appcommand.MonitorTargets.HTTP.Name:
		re := `^https?://.*$`
		if !regexp.MustCompile(re).MatchString(c.argValue) {
			return errors.New("invalid http url format")
		}
	case appcommand.MonitorTargets.TCP.Name:
		re := `^[\w.-]+:\d+$`
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
	h := Check{
		Target: c.argTarget,
		Value:  c.argValue,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return bot.EscapeMarkdown(result.ToMarkdown(ctx))
}

func (c command) register(ctx *appcontext.AppContext) string {
	h := Register{
		Target:   c.argTarget,
		Value:    c.argValue,
		Platform: c.payload.Platform,
		ChatID:   c.payload.ChatID,
		User:     c.payload.User,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return result
}

func (c command) list(ctx *appcontext.AppContext) string {
	h := List{
		Message: c.payload.Message,
		Target:  c.argTarget,
		User:    c.payload.User,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return bot.EscapeMarkdown(result)
}

func (c command) remove(_ *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("removing %s %s with user %s ...", c.argTarget, c.argValue, c.payload.User.ID))
}

func (c command) stats(_ *appcontext.AppContext) string {
	return bot.EscapeMarkdown(fmt.Sprintf("getting stats of %s %s with user %s ...", c.argTarget, c.argValue, c.payload.User.ID))
}
