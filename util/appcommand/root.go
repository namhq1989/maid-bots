package appcommand

var Root = struct {
	Help    Command
	Monitor Command
	Random  Command
	Example Command
}{
	Help: Command{
		Base:        "help",
		WithSlash:   "/help",
		Description: "Get assistance and information about any handler",
	},
	Monitor: Command{
		Base:        "monitor",
		WithSlash:   "/monitor",
		Description: "Register and manage monitoring targets, including domain, http, icmp, tcp",
	},
	Random: Command{
		Base:        "random",
		WithSlash:   "/random",
		Description: "Random things",
	},
	Example: Command{
		Base:        "example",
		WithSlash:   "/example",
		Description: "Examples",
	},
}
