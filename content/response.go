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

	Response.Monitor.Stats.Domain = readFile("content/response/monitor/stats/domain.md")
	Response.Monitor.Stats.HTTP = readFile("content/response/monitor/stats/http.md")
	Response.Monitor.Stats.TCP = readFile("content/response/monitor/stats/tcp.md")
	Response.Monitor.Stats.ICMP = readFile("content/response/monitor/stats/icmp.md")
}
