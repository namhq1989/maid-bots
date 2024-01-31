package config

type Command struct {
	Base        string
	WithSlash   string
	Description string
}

var Commands = struct {
	Help    Command
	Monitor Command
	Example Command
}{
	Help: Command{
		Base:        "help",
		WithSlash:   "/help",
		Description: "Get assistance and information about any command",
	},
	Monitor: Command{
		Base:        "monitor",
		WithSlash:   "/monitor",
		Description: "Register and manage monitoring targets, including domain, http, icmp, tcp",
	},
	Example: Command{
		Base:        "example",
		WithSlash:   "/example",
		Description: "Examples",
	},
}
