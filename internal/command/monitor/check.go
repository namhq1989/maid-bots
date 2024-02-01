package monitor

import (
	"errors"
	"fmt"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"github.com/namhq1989/maid-bots/util/domain"
)

type Check struct {
	Target string
	Value  string
}

func (c Check) Process(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	switch c.Target {
	case appcommand.MonitorTargets.Domain.Name:
		return c.domain(ctx)
	case appcommand.MonitorTargets.HTTP.Name:
		return c.http(ctx)
	case appcommand.MonitorTargets.TCP.Name:
		return c.tcp(ctx)
	case appcommand.MonitorTargets.ICMP.Name:
		return c.icmp(ctx)
	}

	return nil, errors.New(fmt.Sprintf("target %s is not supported", c.Target))
}

func (c Check) domain(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	ctx.AddLogData(appcontext.Fields{"domain": c.Value})

	var result = &modelresponse.Check{}

	// get domain data
	domainData, err := domain.Parse(c.Value)
	if err != nil {
		ctx.Logger.Error("error parsing domain data", err, appcontext.Fields{})
		return nil, err
	}

	// get ssl data
	sslData, err := domain.CheckSSL(domainData.Name, domainData.Port)
	if err != nil {
		ctx.Logger.Error("error checking domain ssl", err, appcontext.Fields{})
		return nil, err
	}

	// set result
	result.Scheme = domainData.Scheme
	result.SSL.Issuer = sslData.Issuer
	result.SSL.ExpireAt = modelresponse.NewTimeResponse(sslData.ExpireAt)

	// response time
	result.ResponseTimeInMS, err = domain.MeasureResponseTime(fmt.Sprintf("%s://%s", domainData.Scheme, domainData.Name))
	if err == nil {
		result.IsUp = true
	}

	// ip
	result.IPResolves, _ = domain.IPLookup(domainData.Name)

	return result, err
}

func (c Check) http(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}

func (c Check) tcp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}

func (c Check) icmp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}
