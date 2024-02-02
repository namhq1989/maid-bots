package content

type ListCommands struct {
	Base    string
	Monitor string
}

var Command = struct {
	Help    ListCommands
	Example ListCommands
}{
	Help: ListCommands{
		Base:    "",
		Monitor: "",
	},
	Example: ListCommands{
		Base:    "",
		Monitor: "",
	},
}

func command() {
	// help
	Command.Help.Base = readFile("content/command/help.md")
	Command.Help.Monitor = readFile("content/command/help_monitor.md")

	// example
	Command.Example.Base = readFile("content/command/example.md")
	Command.Example.Monitor = readFile("content/command/example_monitor.md")
}
