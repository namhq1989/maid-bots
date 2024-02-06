package telegram

import (
	"fmt"
	"regexp"
	"strconv"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
)

func generateRegexp(cmd string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s`, cmd))
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
