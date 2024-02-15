package appcommand

type Command struct {
	Name        string
	WithSlash   string
	Description string
}

type ArgumentAction struct {
	Name        string
	Description string
}

type ArgumentTarget struct {
	Name        string
	Description string
}

type ArgumentType struct {
	Name        string
	Description string
}
