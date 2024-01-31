package appcommand

var MonitorSubCommands = struct {
	Check    SubCommand
	Register SubCommand
	List     SubCommand
	Remove   SubCommand
	Stats    SubCommand
}{
	Check: SubCommand{
		Base:        "check",
		Description: "Perform a one-time check on the target",
	},
	Register: SubCommand{
		Base:        "register",
		Description: "Register a new target for continuous monitoring according to a predefined schedule",
	},
	List: SubCommand{
		Base:        "list",
		Description: "Display a list of all registered monitoring targets",
	},
	Remove: SubCommand{
		Base:        "remove",
		Description: "Remove a specific target by its id",
	},
	Stats: SubCommand{
		Base:        "stats",
		Description: "Get statistic of a specific target by its id",
	},
}
