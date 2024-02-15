package appcommand

var MonitorActions = struct {
	Check    string
	Register string
	List     string
	Remove   string
	Stats    string
}{
	Check:    "check",
	Register: "register",
	List:     "list",
	Remove:   "remove",
	Stats:    "stats",
}

var MonitorTypes = struct {
	Domain string
	HTTP   string
	ICMP   string
	TCP    string
}{
	Domain: "domain",
	HTTP:   "http",
	ICMP:   "icmp",
	TCP:    "tcp",
}

var MonitorParameters = struct {
	Action  string
	Type    string
	Target  string
	Keyword string
	Page    string
}{
	Action:  "action",
	Type:    "type",
	Target:  "target",
	Keyword: "keyword",
	Page:    "page",
}
