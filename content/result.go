package content

type PlatformResult struct {
	Telegram string
	Discord  string
	Slack    string
}

var Result = struct {
	MonitorCheck PlatformResult
}{
	MonitorCheck: PlatformResult{},
}

func result() {
	// monitor
	// check
	Result.MonitorCheck.Telegram = readFile("content/result/monitor/check/telegram.md")
}
