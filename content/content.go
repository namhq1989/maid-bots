package content

import (
	"fmt"
	"os"
	"path/filepath"
)

type Example struct {
	Base    string
	Monitor string
}

var Group = struct {
	Help    string
	Monitor string
	Example Example
}{
	Help:    "",
	Monitor: "",
	Example: Example{
		Base:    "",
		Monitor: "",
	},
}

func Load() {
	// help
	Group.Help = readFile("content/help.md")

	// monitor
	Group.Monitor = readFile("content/monitor.md")

	// example
	Group.Example.Base = readFile("content/example.md")
	Group.Example.Monitor = readFile("content/example_monitor.md")
}

func readFile(fPath string) string {
	path, _ := os.Getwd()

	// read the content of the file
	c, err := os.ReadFile(filepath.Join(path, fPath))
	if err != nil {
		panic(fmt.Sprintf("error when reading file %s: %s", fPath, err.Error()))
	}
	return string(c)
}
