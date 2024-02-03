package netinspect

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Measure struct {
	ResponseTimeInMs int64
}

func MeasureHTTPResponseTime(ctx *appcontext.AppContext, url string) (*Measure, error) {
	span := sentryio.NewSpan(ctx.Context, "measure response time", "")
	defer span.Finish()

	startTime := time.Now()

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()

	return &Measure{
		ResponseTimeInMs: time.Since(startTime).Milliseconds(),
	}, nil
}

func MeasureTCPResponseTime(ctx *appcontext.AppContext, address string) (*Measure, error) {
	span := sentryio.NewSpan(ctx.Context, "measure tcp response time", "")
	defer span.Finish()

	// dial
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", address, tcpTimeout)
	if err != nil {
		return nil, fmt.Errorf("error dialing tcp: %s", err.Error())
	}
	defer func() { _ = conn.Close() }()

	return &Measure{
		ResponseTimeInMs: time.Since(startTime).Milliseconds(),
	}, nil
}
