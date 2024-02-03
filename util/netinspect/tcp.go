package netinspect

import (
	"errors"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"net"
)

func CheckTCP(ctx *appcontext.AppContext, address string) error {
	span := sentryio.NewSpan(ctx.Context, "check tcp", "")
	defer span.Finish()

	if !isValidTCP(address) {
		return errors.New("invalid tcp format")
	}

	return nil
}

func isValidTCP(input string) bool {
	_, _, err := net.SplitHostPort(input)
	return err == nil
}
