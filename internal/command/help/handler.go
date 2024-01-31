package help

import (
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func ProcessMessage(ctx *appcontext.AppContext, message, platform string) string {
	c := command{
		message:  message,
		platform: platform,
	}
	return c.process(ctx)
}
