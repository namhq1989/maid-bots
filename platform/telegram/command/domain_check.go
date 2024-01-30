package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	modelresponse "github.com/namhq1989/maid-bots/internal/models/response"
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func DomainCheckHandler(bgCtx context.Context, b *bot.Bot, update *models.Update) {
	var (
		ctx = appcontext.New(bgCtx)
	)

	if update.Message.Text == DomainCheckCommand.WithSlash {
		// add redis
		AddRedisKey(update.Message.Chat.ID, DomainCheckCommand.Command)

		_, _ = b.SendMessage(ctx.Context, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Please input your domain",
		})
	} else {
		parts := strings.Split(update.Message.Text, " ")
		DomainCheck{Domain: parts[1]}.Process(ctx, b, update)
	}
}

type DomainCheck struct {
	Domain string
}

func (c DomainCheck) Process(ctx *appcontext.AppContext, b *bot.Bot, update *models.Update) {
	ctx.Logger.Info("receive: /domain_check", appcontext.Fields{
		"text": update.Message.Text,
	})

	// del redis key
	defer DelRedisKey(update.Message.Chat.ID)

	response, err := service.Domain{Name: c.Domain}.Check(ctx)
	if err != nil {
		_, _ = b.SendMessage(ctx.Context, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Got error: %s", err.Error()),
		})
	} else {
		_, _ = b.SendMessage(ctx.Context, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      c.formatDomainCheckResponseHTMLTable(response),
			ParseMode: models.ParseModeHTML,
		})
	}
}

func (c DomainCheck) formatDomainCheckResponseHTMLTable(response *modelresponse.PublicCheckDomain) string {
	var builder strings.Builder

	builder.WriteString("<b>Result:</b>\n")
	builder.WriteString("<b>-------</b>\n")
	builder.WriteString(fmt.Sprintf("<i>Domain Name</i>: %s\n", response.DomainName))
	builder.WriteString(fmt.Sprintf("<i>Status</i>: %s\n", c.getStatusEmojiHTML(response.IsUp)))
	builder.WriteString(fmt.Sprintf("<i>HTTPS</i>: %s\n", c.getBoolEmojiHTML(response.IsHTTPS)))
	builder.WriteString(fmt.Sprintf("<i>Response Time</i>: %d ms\n", response.ResponseTimeInMS))
	builder.WriteString(fmt.Sprintf("<i>Scheme</i>: %s\n", response.Scheme))
	builder.WriteString(fmt.Sprintf("<i>Expires At</i>: %s\n", response.ExpireAt.FormatYYYYMMDD()))
	builder.WriteString(fmt.Sprintf("<i>Issuer</i>: %s\n", response.Issuer))
	builder.WriteString(fmt.Sprintf("<i>IP Resolves</i>: %s\n", strings.Join(response.IPResolves, ", ")))

	return builder.String()
}

func (c DomainCheck) getStatusEmojiHTML(status bool) string {
	if status {
		return "<i>Up</i>  ✅"
	}
	return "<i>Down</i> ❌"
}

func (c DomainCheck) getBoolEmojiHTML(value bool) string {
	if value {
		return "<i>Yes</i>  ✅"
	}
	return "<i>No</i> ❌"
}
