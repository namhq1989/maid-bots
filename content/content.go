package content

import (
	"fmt"
	"os"
	"path/filepath"
)

func Load() {
	command()
	result()
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
