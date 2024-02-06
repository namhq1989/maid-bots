package help

import (
	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

func ProcessMessage(ctx *appcontext.AppContext, payload modelcommand.Payload) string {
	c := command{
		payload: payload,
	}
	return c.process(ctx)
}
