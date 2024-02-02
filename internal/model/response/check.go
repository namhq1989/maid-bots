package modelresponse

import (
	"fmt"
	"strings"

	"github.com/namhq1989/maid-bots/content"
)

type Check struct {
	Template         string
	Name             string
	IsUp             bool     `json:"isUp"`
	ResponseTimeInMS int64    `json:"responseTime"`
	Scheme           string   `json:"scheme"`
	IPResolves       []string `json:"ipResolves"`
	SSL              CheckSSL `json:"ssl"`
}

type CheckSSL struct {
	ExpireAt *TimeResponse `json:"expireAt"`
	Issuer   string        `json:"issuer"`
}

func (m Check) ToMarkdown() string {
	var status = "Up ✅"
	if !m.IsUp {
		status = "Down ❌"
	}

	var https = "https ✅"
	var isHttps = m.Scheme == "https"
	if !isHttps {
		https = "http ❌"
	}

	result := content.Response.Monitor.Check.Default
	isTemplateDomain := m.Template == content.MonitorTemplateDomain
	if isTemplateDomain {
		result = content.Response.Monitor.Check.Domain
	}

	result = strings.ReplaceAll(result, "$name", m.Name)
	result = strings.ReplaceAll(result, "$status", status)
	result = strings.ReplaceAll(result, "$response_time", fmt.Sprintf("%d", m.ResponseTimeInMS))

	if isTemplateDomain {
		result = strings.ReplaceAll(result, "$scheme", https)
		result = strings.ReplaceAll(result, "$ip_resolves", strings.Join(m.IPResolves, ", "))
		if isHttps {
			result = strings.ReplaceAll(result, "$ssl_issuer", m.SSL.Issuer)
			result = strings.ReplaceAll(result, "$ssl_expires", m.SSL.ExpireAt.FormatYYYYMMDD())
		}
	}

	return result
}
