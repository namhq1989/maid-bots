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

func ExtractArguments(input string) map[string]string {
	params := make(map[string]string)

	// remove the leading '/' character from the command
	input = strings.TrimPrefix(input, "/")

	// split the command into parts separated by spaces
	parts := strings.Split(input, " ")

	// iterate over the parts to extract parameters
	for _, part := range parts {
		// split each part into key-value pairs separated by '='
		pair := strings.Split(part, "=")
		if len(pair) == 2 {
			key := pair[0]
			value := pair[1]
			params[key] = value
		}
	}

	return params
}
