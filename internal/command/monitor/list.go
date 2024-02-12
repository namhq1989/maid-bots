package monitor

import (
	"errors"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type List struct {
	Message    string
	Target     string
	Parameters *ListParameters
	User       modelcommand.User
}

type ListParameters struct {
	Keyword   string
	PageToken string
}

func (c List) Process(ctx *appcontext.AppContext) (string, error) {
	// check target
	if !appcommand.IsMonitorTargetValid(c.Target) && c.Target != appcommand.MonitorTargets.All.Name {
		return "", errors.New("invalid target, please check `/help monitor` for more information")
	}

	// TODO:
	// reference to /random number command
	// "list" command supports:
	// - query with target (domain, http, ...)
	// - filter with keyword. it means it needs an additional field for search in database
	// - pagination with page token. the result will return the token for next page

	return "listing ...", nil
}
