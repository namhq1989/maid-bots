package appcommand

import "slices"

var RandomTargets = struct {
	Number ArgumentTarget
	String ArgumentTarget
}{
	Number: ArgumentTarget{
		Name:        "number",
		Description: "Random a number",
	},
	String: ArgumentTarget{
		Name:        "string",
		Description: "Random a string",
	},
}

var RandomTargetsArray = []string{
	RandomTargets.Number.Name,
	RandomTargets.String.Name,
}

func IsRandomTargetValid(t string) bool {
	return slices.Contains(RandomTargetsArray, t)
}
