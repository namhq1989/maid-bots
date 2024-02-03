package content

var Response = struct {
	Monitor MonitorResult
}{
	Monitor: MonitorResult{},
}

func response() {
	// monitor
	Response.Monitor.Check.Domain = readFile("content/response/monitor/check/domain.md")
	Response.Monitor.Check.HTTP = readFile("content/response/monitor/check/http.md")
	Response.Monitor.Check.TCP = readFile("content/response/monitor/check/tcp.md")
	Response.Monitor.Check.ICMP = readFile("content/response/monitor/check/icmp.md")
}
