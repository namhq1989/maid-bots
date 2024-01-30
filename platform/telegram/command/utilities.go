package command

import (
	"fmt"
	"time"

	"github.com/go-telegram/bot/models"
	"github.com/namhq1989/maid-bots/pkg/redis"
)

type Command struct {
	models.BotCommand
	Sample    string
	WithSlash string
}

var HelpCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "help",
		Description: "Help about any command",
	},
	Sample:    "/help",
	WithSlash: "/help",
}

//
// DOMAIN
//

var DomainCheckCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "domain_check",
		Description: "Check domain information",
	},
	Sample:    "/domain_check domain_name",
	WithSlash: "/domain_check",
}

var DomainWatchCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "domain_watch",
		Description: "Watch domain's availability",
	},
	Sample:    "/domain_watch domain_name",
	WithSlash: "/domain_watch",
}

var DomainListCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "domain_list",
		Description: "List all watching domains",
	},
	Sample:    "/domain_list",
	WithSlash: "/domain_list",
}

var DomainEditCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "domain_edit",
		Description: "Edit a domain",
	},
	Sample:    "/domain_edit domain_name OR domain_id",
	WithSlash: "/domain_edit",
}

var DomainDeleteCommand = Command{
	BotCommand: models.BotCommand{
		Command:     "domain_delete",
		Description: "Delete a domain",
	},
	Sample:    "/domain_delete domain_name OR domain_id",
	WithSlash: "/domain_delete",
}

//
// COMMANDS
//

var Commands = []Command{
	HelpCommand,

	// domain
	DomainCheckCommand,
	DomainWatchCommand,
	DomainListCommand,
	DomainEditCommand,
	DomainDeleteCommand,
}

//
// FUNCTIONS
//

func AddRedisKey(chatID int64, cmd string) {
	redis.SetKeyValue(fmt.Sprintf("%d", chatID), cmd, 30*time.Minute)
}

func DelRedisKey(chatID int64) {
	redis.DelKey(fmt.Sprintf("%d", chatID))
}
