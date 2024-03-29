package modelresponse

import (
	"strings"

	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type HealthCheckRecordResponseTimeMetrics struct {
	Average          float64 `json:"average"`
	Max              float64 `json:"max"`
	Min              float64 `json:"min"`
	UptimePercentage float64 `json:"uptimePercentage"`
}

type HealthCheckRecordResponseTimeChartData struct {
	Date string  `json:"date"`
	Hour int     `json:"hour"`
	Avg  float64 `json:"avg"`
}

type Stats struct {
	Monitor        Monitor                                  `json:"monitor"`
	ResponseTime   *HealthCheckRecordResponseTimeMetrics    `json:"responseTime"`
	Chart          []HealthCheckRecordResponseTimeChartData `json:"chart"`
	ChartImageName string                                   `json:"-"`
}

func (m Stats) ToMarkdown(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "[model][stats] convert to markdown")
	defer span.Finish()

	result := content.Response.Monitor.Stats

	result = strings.ReplaceAll(result, "$id", m.Monitor.Code)
	result = strings.ReplaceAll(result, "$type", m.Monitor.Type)
	result = strings.ReplaceAll(result, "$target", m.Monitor.Target)
	result = strings.ReplaceAll(result, "$interval", formatReadableInt(int64(m.Monitor.Interval)))
	result = strings.ReplaceAll(result, "$created_at", m.Monitor.CreatedAt.FormatYYYYMMDD())

	result = strings.ReplaceAll(result, "$average_response_time", formatReadableFloat64(m.ResponseTime.Average, 0))
	result = strings.ReplaceAll(result, "$max_response_time", formatReadableFloat64(m.ResponseTime.Max, 0))
	result = strings.ReplaceAll(result, "$min_response_time", formatReadableFloat64(m.ResponseTime.Min, 0))
	result = strings.ReplaceAll(result, "$uptime_percentage", formatReadableFloat64(m.ResponseTime.UptimePercentage, 2))

	return result
}
