package monitor

import (
	"fmt"
	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Register struct {
	Target   string
	Value    string
	Platform string
	UserID   string
}

func (c Register) Process(ctx *appcontext.AppContext) (string, error) {
	// check first
	check := Check{
		Target: c.Target,
		Value:  c.Value,
	}

	result, err := check.Process(ctx)
	if err != nil {
		return "", err
	}

	switch c.Target {
	case appcommand.MonitorTargets.Domain.Name:
		return c.domain(ctx, result)
	case appcommand.MonitorTargets.HTTP.Name:
		return c.http(ctx, result)
	case appcommand.MonitorTargets.TCP.Name:
		return c.tcp(ctx, result)
	case appcommand.MonitorTargets.ICMP.Name:
		return c.icmp(ctx, result)
	}

	return "", fmt.Errorf("target %s is not supported", c.Target)
}

func (c Register) domain(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	return "", nil
}

func (c Register) http(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	return "", nil
}

func (c Register) tcp(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	return "", nil
}

func (c Register) icmp(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	return "", nil
}
