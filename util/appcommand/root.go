package appcommand

import "slices"

var Root = struct {
	Help    Command
	Monitor Command
	Random  Command
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
}

var RootCommandsArray = []string{
	Root.Help.WithSlash,
	Root.Monitor.WithSlash,
	Root.Random.WithSlash,
}

func IsRootCommandValid(v string) bool {
	return slices.Contains(RootCommandsArray, v)
}
