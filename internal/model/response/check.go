package modelresponse

import (
	"strings"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"

	"github.com/namhq1989/maid-bots/content"
)

type Check struct {
	Template         string
	Name             string    `json:"name"`
	IsUp             bool      `json:"isUp"`
	ResponseTimeInMS int64     `json:"responseTime"`
	Scheme           string    `json:"scheme"`
	IPResolves       []string  `json:"ipResolves"`
	SSL              CheckSSL  `json:"ssl"`
	ICMP             CheckICMP `json:"icmp"`
}

type CheckSSL struct {
	ExpireAt *TimeResponse `json:"expireAt"`
	Issuer   string        `json:"issuer"`
}

type CheckICMP struct {
	PackageTransmitted int     `json:"packageTransmitted"`
	PackageReceived    int     `json:"packageReceived"`
	PackageLoss        float64 `json:"packageLoss"`
}

func (m Check) ToMarkdown(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "[model][check] convert to markdown")
	defer span.Finish()

	var status = "Up ✅"
	if !m.IsUp {
		status = "Down ❌"
	}

	var https = "https ✅"
	var isHttps = m.Scheme == "https"
	if !isHttps {
		https = "http ❌"
	}

	result := ""

	switch m.Template {
	case content.MonitorTemplateDomain:
		result = content.Response.Monitor.Check.Domain
		result = strings.ReplaceAll(result, "$name", m.Name)
		result = strings.ReplaceAll(result, "$status", status)
		result = strings.ReplaceAll(result, "$response_time", formatReadableInt(m.ResponseTimeInMS))
		result = strings.ReplaceAll(result, "$scheme", https)
		result = strings.ReplaceAll(result, "$ip_resolves", strings.Join(m.IPResolves, ", "))
		if isHttps {
			result = strings.ReplaceAll(result, "$ssl_issuer", m.SSL.Issuer)
			result = strings.ReplaceAll(result, "$ssl_expires", m.SSL.ExpireAt.FormatYYYYMMDD())
		}
		return result
	case content.MonitorTemplateHTTP:
		result = content.Response.Monitor.Check.HTTP
		result = strings.ReplaceAll(result, "$name", m.Name)
		result = strings.ReplaceAll(result, "$status", status)
		result = strings.ReplaceAll(result, "$response_time", formatReadableInt(m.ResponseTimeInMS))
		result = strings.ReplaceAll(result, "$scheme", https)
		return result
	case content.MonitorTemplateTCP:
		result = content.Response.Monitor.Check.TCP
		result = strings.ReplaceAll(result, "$name", m.Name)
		result = strings.ReplaceAll(result, "$status", status)
		result = strings.ReplaceAll(result, "$response_time", formatReadableInt(m.ResponseTimeInMS))
		return result
	case content.MonitorTemplateICMP:
		result = content.Response.Monitor.Check.ICMP
		result = strings.ReplaceAll(result, "$name", m.Name)
		result = strings.ReplaceAll(result, "$status", status)
		result = strings.ReplaceAll(result, "$ip_resolves", strings.Join(m.IPResolves, ", "))
		result = strings.ReplaceAll(result, "$response_time", formatReadableInt(m.ResponseTimeInMS))
		result = strings.ReplaceAll(result, "$pk_transmitted", formatReadableInt(int64(m.ICMP.PackageTransmitted)))
		result = strings.ReplaceAll(result, "$pk_received", formatReadableInt(int64(m.ICMP.PackageReceived)))
		result = strings.ReplaceAll(result, "$pk_loss", formatReadableInt(int64(m.ICMP.PackageLoss)))
		return result
	}

	return result
}
