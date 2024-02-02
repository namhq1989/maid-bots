package content

var Response = struct {
	Monitor MonitorResult
}{
	Monitor: MonitorResult{},
}

func response() {
	// monitor
	Response.Monitor.Check.Default = readFile("content/response/monitor/check/default.md")
	Response.Monitor.Check.Domain = readFile("content/response/monitor/check/domain.md")
}
