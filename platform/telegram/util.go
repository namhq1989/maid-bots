package telegram

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/namhq1989/maid-bots/pkg/sentryio"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
)

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

func generateRegexp(cmd string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s`, cmd))
}

func getPayload(update *models.Update) modelcommand.Payload {
	return modelcommand.Payload{
		Platform: config.Platform.Telegram,
		ChatID:   fmt.Sprintf("%d", update.Message.Chat.ID),
		Message:  update.Message.Text,
		User:     getUser(update),
	}
}

func getUser(update *models.Update) modelcommand.User {
	return modelcommand.User{
		ID:       strconv.FormatInt(update.Message.From.ID, 10),
		Name:     fmt.Sprintf("%s %s", update.Message.From.LastName, update.Message.From.FirstName),
		Username: update.Message.From.Username,
	}
}

func getAPMTransactionData(ctx *appcontext.AppContext, update *models.Update) map[string]string {
	user := getUser(update)

	return map[string]string{
		"platform":  config.Platform.Telegram,
		"message":   update.Message.Text,
		"username":  user.Username,
		"userId":    user.ID,
		"requestId": ctx.RequestID,
		"traceId":   ctx.TraceID,
	}
}
