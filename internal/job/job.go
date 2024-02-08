package job

import "github.com/namhq1989/maid-bots/pkg/queue"

func Init() {
	var (
		q = queue.GetInstance()
	)

	monitor := Monitor{}
	monitor.setup(q, monitorCheck.interval60Seconds)
	q.Server.HandleFunc(monitorCheck.interval60Seconds.Task, monitor.check)
}
