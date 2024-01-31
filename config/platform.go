package config

var Platform = struct {
	Telegram string
	Slack    string
	Discord  string
	Web      string
}{
	Telegram: "telegram",
	Slack:    "slack",
	Discord:  "discord",
	Web:      "web",
}
