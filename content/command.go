package content

type ListCommands struct {
	Base    string
	Monitor string
	Random  string
}

var Command = struct {
	Help    ListCommands
	Example ListCommands
}{
	Help:    ListCommands{},
	Example: ListCommands{},
}

func command() {
	// help
	Command.Help.Base = readFile("content/command/help/base.md")
	Command.Help.Monitor = readFile("content/command/help/monitor.md")
	Command.Help.Random = readFile("content/command/help/random.md")
}
