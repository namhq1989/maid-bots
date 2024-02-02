package netinspect

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type TCP struct {
	ResponseTimeInMs int64
}

func CheckTCP(ctx *appcontext.AppContext, address string) (*TCP, error) {
	span := sentryio.NewSpan(ctx.Context, "check tcp", "")
	defer span.Finish()

	if !isValidTCP(address) {
		return nil, errors.New("invalid tcp format")
	}

	var result = &TCP{}

	// dial
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("error dialing tcp: %s", err.Error())
	}
	defer func() { _ = conn.Close() }()

	result.ResponseTimeInMs = time.Since(startTime).Milliseconds()
	return result, nil
}

func isValidTCP(input string) bool {
	_, _, err := net.SplitHostPort(input)
	return err == nil
}
