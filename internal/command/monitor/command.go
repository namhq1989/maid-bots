package monitor

import (
	"errors"
	"regexp"

	"github.com/namhq1989/maid-bots/content"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/go-telegram/bot"

	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type command struct {
	payload   modelcommand.Payload
	arguments map[string]string
}

func (c command) process(ctx *appcontext.AppContext) string {
	c.arguments = appcommand.ExtractArguments(c.payload.Message)

	ctx.Logger.Info("receive: /monitor", appcontext.Fields{
		"platform":  c.payload.Platform,
		"message":   c.payload.Message,
		"arguments": c.arguments,
	})

	var (
		l    = len(c.arguments)
		text = "invalid command"
	)

	if l == 0 {
		return content.Command.Help.Monitor
	}

	// validate data
	if err := c.validateData(ctx); err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	switch c.arguments[appcommand.MonitorParameters.Action] {
	case appcommand.MonitorActions.Check:
		return c.check(ctx)
	case appcommand.MonitorActions.Register:
		return c.register(ctx)
	case appcommand.MonitorActions.List:
		return c.list(ctx)
	case appcommand.MonitorActions.Remove:
		return c.remove(ctx)
	case appcommand.MonitorActions.Stats:
		return c.stats(ctx)
	}

	return text
}

func (c command) validateData(ctx *appcontext.AppContext) error {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] validate data")
	defer span.Finish()

	var (
		t = c.arguments[appcommand.MonitorParameters.Type]
		v = c.arguments[appcommand.MonitorParameters.Target]
	)

	// only check for actions: "check", "register"
	if t != appcommand.MonitorActions.Check && t != appcommand.MonitorActions.Register {
		return nil
	}

	// value
	switch t {
	case appcommand.MonitorTypes.Domain:
		re := `^([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.[a-zA-Z]{2,6}$`
		if !regexp.MustCompile(re).MatchString(v) {
			return errors.New("invalid domain format")
		}
	case appcommand.MonitorTypes.HTTP:
		re := `^https?://.*$`
		if !regexp.MustCompile(re).MatchString(v) {
			return errors.New("invalid http url format")
		}
	case appcommand.MonitorTypes.TCP:
		re := `^[\w.-]+:\d+$`
		if !regexp.MustCompile(re).MatchString(v) {
			return errors.New("invalid tcp format")
		}
	case appcommand.MonitorTypes.ICMP:
		re := `^([a-zA-Z0-9_-]+\.){1,}[a-zA-Z]{2,}$|^\d+\.\d+\.\d+\.\d+$`
		if !regexp.MustCompile(re).MatchString(v) {
			return errors.New("invalid icmp format")
		}
	default:
		return errors.New("invalid target type")
	}

	return nil
}

func (c command) check(ctx *appcontext.AppContext) string {
	h := Check{
		Arguments: c.arguments,
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
		Arguments: c.arguments,
		Platform:  c.payload.Platform,
		ChatID:    c.payload.ChatID,
		User:      c.payload.User,
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
		Arguments: c.arguments,
		Platform:  c.payload.Platform,
		ChatID:    c.payload.ChatID,
		User:      c.payload.User,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return result
}

func (c command) remove(ctx *appcontext.AppContext) string {
	h := Remove{
		Arguments: c.arguments,
		Platform:  c.payload.Platform,
		ChatID:    c.payload.ChatID,
		User:      c.payload.User,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return result
}

func (c command) stats(ctx *appcontext.AppContext) string {
	h := Stats{
		Arguments: c.arguments,
		Platform:  c.payload.Platform,
		ChatID:    c.payload.ChatID,
		User:      c.payload.User,
	}

	// process
	result, err := h.Process(ctx)
	if err != nil {
		return bot.EscapeMarkdown(err.Error())
	}

	return bot.EscapeMarkdown(result.ToMarkdown(ctx))
}
