package monitor

import (
	"errors"
	"fmt"

	"github.com/namhq1989/maid-bots/content"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"github.com/namhq1989/maid-bots/util/netinspect"
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
	domainData, err := netinspect.ParseDomain(ctx, c.Value)
	if err != nil {
		ctx.Logger.Error("error parsing domain data", err, appcontext.Fields{})
		return nil, err
	}

	// get ssl data
	sslData, err := netinspect.CheckSSL(ctx, domainData.Name, domainData.Port)
	if err != nil {
		ctx.Logger.Error("error checking domain ssl", err, appcontext.Fields{})
		return nil, err
	}

	// set response
	result.Name = domainData.Name
	result.Scheme = domainData.Scheme
	result.SSL.Issuer = sslData.Issuer
	result.SSL.ExpireAt = modelresponse.NewTimeResponse(sslData.ExpireAt)

	// response time
	measure, err := netinspect.MeasureResponseTime(ctx, fmt.Sprintf("%s://%s", domainData.Scheme, domainData.Name))
	if err == nil {
		result.IsUp = true
	}
	result.ResponseTimeInMS = measure.ResponseTimeInMs

	// ip
	ipLookup, _ := netinspect.IPLookup(ctx, domainData.Name)
	result.IPResolves = ipLookup.List

	// set template
	result.Template = content.MonitorTemplateDomain

	return result, err
}

func (c Check) http(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	ctx.AddLogData(appcontext.Fields{"url": c.Value})

	var result = &modelresponse.Check{}

	// get url data
	urlData, err := netinspect.ParseURL(ctx, c.Value)

	// check ssl
	sslData, err := netinspect.CheckSSL(ctx, urlData.Host, urlData.Port)
	if err != nil {
		ctx.Logger.Error("error checking domain ssl", err, appcontext.Fields{})
		return nil, err
	}

	// set response
	result.Name = urlData.Value
	result.Scheme = urlData.Scheme
	result.SSL.Issuer = sslData.Issuer
	result.SSL.ExpireAt = modelresponse.NewTimeResponse(sslData.ExpireAt)

	// response time
	measure, err := netinspect.MeasureResponseTime(ctx, urlData.Value)
	if err == nil {
		result.IsUp = true
	}
	result.ResponseTimeInMS = measure.ResponseTimeInMs

	// ip look up
	if urlData.Type == netinspect.TypeDomain {
		ipLookup, err := netinspect.IPLookup(ctx, urlData.Host)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error looking up ip: %s", err.Error()))
		}
		result.IPResolves = ipLookup.List
	} else {
		result.IPResolves = []string{urlData.Value}
	}

	// set template
	result.Template = content.MonitorTemplateDomain

	return result, nil
}

func (c Check) tcp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}

func (c Check) icmp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}
