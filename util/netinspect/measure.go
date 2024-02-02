package netinspect

import (
	"net/http"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Measure struct {
	ResponseTimeInMs int64
}

func MeasureResponseTime(ctx *appcontext.AppContext, url string) (*Measure, error) {
	span := sentryio.NewSpan(ctx.Context, "measure response time", "")
	defer span.Finish()

	startTime := time.Now()

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	return &Measure{
		ResponseTimeInMs: duration.Milliseconds(),
	}, nil
}
