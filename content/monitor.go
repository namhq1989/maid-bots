package content

type MonitorCheckResult struct {
	Default string
	Domain  string
}

type MonitorResult struct {
	Check MonitorCheckResult
}
