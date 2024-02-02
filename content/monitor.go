package content

const (
	MonitorTemplateDefault = "default"
	MonitorTemplateDomain  = "domain"
)

type MonitorCheckResult struct {
	Default string
	Domain  string
}

type MonitorResult struct {
	Check MonitorCheckResult
}
