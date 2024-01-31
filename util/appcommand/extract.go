package appcommand

import (
	"regexp"
	"strings"
)

func ExtractCommand(input string) string {
	re := regexp.MustCompile(`^/(\w+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return "/" + matches[1]
	}

	// No match found
	return ""
}

func ExtractParameters(input string) []string {
	re := regexp.MustCompile(`^/(\w+)(?:\s+(.*))?`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return strings.Fields(matches[2])
	}
	return []string{}
}
