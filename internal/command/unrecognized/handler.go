package unrecognized

import (
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func ProcessMessage(ctx *appcontext.AppContext, message, platform, userID string) string {
	c := command{
		message:  message,
		platform: platform,
		userID:   userID,
	}
	return c.process(ctx)
}
