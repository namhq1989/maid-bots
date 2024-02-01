package appcommand

var Root = struct {
	Help    Command
	Monitor Command
	Random  Command
	Example Command
}{
	Help: Command{
		Name:        "help",
		WithSlash:   "/help",
		Description: "Get assistance and information about any handler",
	},
	Monitor: Command{
		Name:        "monitor",
		WithSlash:   "/monitor",
		Description: "Register and manage monitoring targets, including domain, http, icmp, tcp",
	},
	Random: Command{
		Name:        "random",
		WithSlash:   "/random",
		Description: "Random things",
	},
	Example: Command{
		Name:        "example",
		WithSlash:   "/example",
		Description: "Examples",
	},
}
