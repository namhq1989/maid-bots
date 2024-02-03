package content

const (
	MonitorTemplateDomain = "domain"
	MonitorTemplateHTTP   = "http"
	MonitorTemplateTCP    = "tcp"
	MonitorTemplateICMP   = "icmp"
)

type MonitorCheckResult struct {
	Domain string
	HTTP   string
	TCP    string
	ICMP   string
}

type MonitorResult struct {
	Check MonitorCheckResult
}
