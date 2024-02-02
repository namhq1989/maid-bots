package telegram

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
)

func generateRegexp(cmd string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s`, cmd))
}

func getUserID(update *models.Update) string {
	return strconv.FormatInt(update.Message.From.ID, 10)
}

func getAPMTransactionData(ctx *appcontext.AppContext, update *models.Update) map[string]string {
	return map[string]string{
		"platform":  config.Platform.Telegram,
		"message":   update.Message.Text,
		"username":  update.Message.From.Username,
		"userId":    getUserID(update),
		"requestId": ctx.RequestID,
		"traceId":   ctx.TraceID,
	}
}
