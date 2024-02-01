package telegram

import (
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
	"strconv"
)

func getAPMTransactionData(update *models.Update) map[string]string {
	return map[string]string{
		"platform": config.Platform.Telegram,
		"message":  update.Message.Text,
		"username": update.Message.From.Username,
		"userId":   strconv.FormatInt(update.Message.From.ID, 10),
	}
}
