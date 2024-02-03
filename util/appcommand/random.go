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

var RandomParameters = struct {
	Type   string
	Min    string
	Max    string
	Count  string
	Unique string
}{
	Type:   "type",
	Min:    "min",
	Max:    "max",
	Count:  "count",
	Unique: "unique",
}

var RandomTargetsArray = []string{
	RandomTargets.Number.Name,
	RandomTargets.String.Name,
}

func IsRandomTargetValid(t string) bool {
	return slices.Contains(RandomTargetsArray, t)
}
