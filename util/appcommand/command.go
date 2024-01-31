package appcommand

type Command struct {
	Base        string
	WithSlash   string
	Description string
}

type SubCommand struct {
	Base        string
	Description string
}
