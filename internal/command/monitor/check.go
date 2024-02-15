package monitor

import (
	"fmt"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/content"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"github.com/namhq1989/maid-bots/util/netinspect"
)

type Check struct {
	Arguments map[string]string
}

func (c Check) Process(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	var (
		t = c.Arguments[appcommand.MonitorParameters.Type]
	)

	switch t {
	case appcommand.MonitorTypes.Domain:
		return c.domain(ctx)
	case appcommand.MonitorTypes.HTTP:
		return c.http(ctx)
	case appcommand.MonitorTypes.TCP:
		return c.tcp(ctx)
	case appcommand.MonitorTypes.ICMP:
		return c.icmp(ctx)
	}

	return nil, fmt.Errorf("type \"%s\" is not supported", t)
}

func (c Check) domain(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] domain")
	defer span.Finish()

	var (
		target = c.Arguments[appcommand.MonitorParameters.Target]
	)

	ctx.AddLogData(appcontext.Fields{"domain": target})

	var result = &modelresponse.Check{}

	// get domain data
	domainData, err := netinspect.ParseDomain(ctx, target)
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
	measure, err := netinspect.MeasureHTTPResponseTime(ctx, fmt.Sprintf("%s://%s", domainData.Scheme, domainData.Name))
	if err != nil {
		ctx.Logger.Error("error measuring http", err, appcontext.Fields{})
		return nil, err
	}
	result.IsUp = measure.IsUp
	result.ResponseTimeInMS = measure.ResponseTimeInMs

	// ip
	ipLookup, _ := netinspect.IPLookup(ctx, domainData.Name)
	result.IPResolves = ipLookup.List

	// set template
	result.Template = content.MonitorTemplateDomain

	return result, err
}

func (c Check) http(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] http")
	defer span.Finish()

	var (
		target = c.Arguments[appcommand.MonitorParameters.Target]
	)

	ctx.AddLogData(appcontext.Fields{"url": target})

	var result = &modelresponse.Check{}

	// get url data
	urlData, err := netinspect.ParseURL(ctx, target)
	if err != nil {
		ctx.Logger.Error("error parsing url data", err, appcontext.Fields{})
		return nil, err
	}

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
	measure, err := netinspect.MeasureHTTPResponseTime(ctx, urlData.Value)
	if err != nil {
		ctx.Logger.Error("error measuring http", err, appcontext.Fields{})
		return nil, err
	}
	result.IsUp = measure.IsUp
	result.ResponseTimeInMS = measure.ResponseTimeInMs

	// ip look up
	if urlData.Type == netinspect.TypeDomain {
		ipLookup, err := netinspect.IPLookup(ctx, urlData.Host)
		if err != nil {
			return nil, fmt.Errorf("error looking up ip: %s", err.Error())
		}
		result.IPResolves = ipLookup.List
	} else {
		result.IPResolves = []string{urlData.Value}
	}

	// set template
	result.Template = content.MonitorTemplateHTTP

	return result, nil
}

func (c Check) tcp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] tcp")
	defer span.Finish()

	var (
		target = c.Arguments[appcommand.MonitorParameters.Target]
	)

	ctx.AddLogData(appcontext.Fields{"tcp": target})

	var result = &modelresponse.Check{}

	// check
	err := netinspect.CheckTCP(ctx, target)
	if err != nil {
		ctx.Logger.Error("error checking tcp data", err, appcontext.Fields{})
		return nil, err
	}

	// measure
	measure, err := netinspect.MeasureTCPResponseTime(ctx, target)
	if err != nil {
		ctx.Logger.Error("error measuring tcp data", err, appcontext.Fields{})
		return nil, err
	}

	// set data
	result.Template = content.MonitorTemplateTCP
	result.ResponseTimeInMS = measure.ResponseTimeInMs
	result.IsUp = measure.IsUp
	result.Name = target

	return result, nil
}

func (c Check) icmp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	span := sentryio.NewSpan(ctx.Context, "[monitor][check] icmp")
	defer span.Finish()

	var (
		target = c.Arguments[appcommand.MonitorParameters.Target]
	)

	ctx.AddLogData(appcontext.Fields{"tcp": target})

	var result = &modelresponse.Check{}

	// validate
	if !netinspect.IsValidICMP(ctx, target) {
		return nil, fmt.Errorf("invalid icmp address %s", target)
	}

	// check
	icmpData, err := netinspect.CheckICMP(ctx, target)
	if err != nil {
		ctx.Logger.Error("error checking icmp data", err, appcontext.Fields{})
		return nil, err
	}

	// set data
	result.Template = content.MonitorTemplateICMP
	result.ResponseTimeInMS = icmpData.ResponseTimeInMs
	result.IsUp = true
	result.Name = target
	result.IPResolves = []string{icmpData.IPAddress}
	result.ICMP = modelresponse.CheckICMP{
		PackageTransmitted: icmpData.PackageTransmitted,
		PackageReceived:    icmpData.PackageReceived,
		PackageLoss:        icmpData.PackageLoss,
	}

	return result, nil
}
