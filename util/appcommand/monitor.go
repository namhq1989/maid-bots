package appcommand

import "slices"

var MonitorActions = struct {
	Check    ArgumentAction
	Register ArgumentAction
	List     ArgumentAction
	Remove   ArgumentAction
	Stats    ArgumentAction
}{
	Check: ArgumentAction{
		Name:        "check",
		Description: "Perform a one-time check on the target",
	},
	Register: ArgumentAction{
		Name:        "register",
		Description: "Register a new target for continuous monitoring according to a predefined schedule",
	},
	List: ArgumentAction{
		Name:        "list",
		Description: "Display a list of all registered monitoring targets",
	},
	Remove: ArgumentAction{
		Name:        "remove",
		Description: "Remove a specific target by its id",
	},
	Stats: ArgumentAction{
		Name:        "stats",
		Description: "Get statistic of a specific target by its id",
	},
}

var MonitorTargets = struct {
	Domain ArgumentTarget
	HTTP   ArgumentTarget
	ICMP   ArgumentTarget
	TCP    ArgumentTarget
}{
	Domain: ArgumentTarget{
		Name:        "domain",
		Description: "Domain name",
	},
	HTTP: ArgumentTarget{
		Name:        "http",
		Description: "http/https url",
	},
	ICMP: ArgumentTarget{
		Name:        "icmp",
		Description: "Domain/IP without port",
	},
	TCP: ArgumentTarget{
		Name:        "tcp",
		Description: "Domain/IP with port",
	},
}

var MonitorActionsArray = []string{
	MonitorActions.Check.Name,
	MonitorActions.Register.Name,
	MonitorActions.List.Name,
	MonitorActions.Remove.Name,
	MonitorActions.Stats.Name,
}

var MonitorTargetsArray = []string{
	MonitorTargets.Domain.Name,
	MonitorTargets.HTTP.Name,
	MonitorTargets.ICMP.Name,
	MonitorTargets.TCP.Name,
}

func IsMonitorActionValid(a string) bool {
	return slices.Contains(MonitorActionsArray, a)
}

func IsMonitorTargetValid(t string) bool {
	return slices.Contains(MonitorTargetsArray, t)
}
