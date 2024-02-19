package chart

import (
	"errors"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/snapshot-chromedp/render"
	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type MonitorResponseTimeLine struct {
	Data []modelresponse.HealthCheckRecordResponseTimeChartData
}

func (c MonitorResponseTimeLine) ToImage(ctx *appcontext.AppContext) (string, error) {
	l := len(c.Data)

	if l == 0 {
		return "", errors.New("no data")
	}

	var (
		line      = charts.NewLine()
		xAxisData = make([]string, 0, l)
		yAxisData = make([]opts.LineData, 0, l)
		fileName  = randomImageName()
	)

	line.SetGlobalOptions(
		charts.WithAnimation(false),
	)

	for _, item := range c.Data {
		xAxisData = append(xAxisData, item.Date)
		yAxisData = append(yAxisData, opts.LineData{Value: item.Avg})
	}

	line.SetXAxis(xAxisData).
		AddSeries("Average Response Time (Milliseconds) / UTC", yAxisData)

	render.MakeChartSnapshot(line.RenderContent(), fileName)

	ctx.Logger.Info("line chart generated", appcontext.Fields{
		"filename": fileName,
	})

	return fileName, nil
}
