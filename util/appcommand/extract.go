package appcommand

import (
	"regexp"
	"strings"
)

func ExtractCommand(input string) (string, bool) {
	re := regexp.MustCompile(`^/(\w+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return "/" + matches[1], true
	}

	// No match found
	return "", false
}

func ExtractParameters(input string) []string {
	re := regexp.MustCompile(`^/(\w+)(?:\s+(.*))?`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return strings.Fields(matches[2])
	}
	return []string{}
}
