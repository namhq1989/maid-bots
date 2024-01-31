package command

import (
	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/config"
)

var Commands = []models.BotCommand{
	{
		Command:     config.Commands.Help.Base,
		Description: config.Commands.Help.Description,
	},
	{
		Command:     config.Commands.Monitor.Base,
		Description: config.Commands.Monitor.Description,
	},
	{
		Command:     config.Commands.Example.Base,
		Description: config.Commands.Example.Description,
	},
}
