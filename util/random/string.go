package random

import (
	"math/rand"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

var (
	letterRunes       = []rune("abcdefghijkmnpqrstuvwxyz23456789")
	letterRunesLength = len(letterRunes)
)

func StringWithLength(ctx *appcontext.AppContext, l int) string {
	span := sentryio.NewSpan(ctx.Context, "[util] random string with length", "")
	defer span.Finish()

	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, l)
	for i := range b {
		b[i] = letterRunes[randSeed.Intn(letterRunesLength)]
	}
	return string(b)
}
