package monitor

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

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

	// set response
	result.Name = domainData.Name
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
	// parse the URL
	parsedURL, err := url.Parse(c.Value)
	if err != nil {
		ctx.Logger.Error("error parsing url data", err, appcontext.Fields{})
		return nil, err
	}

	// Send an HTTP HEAD request to the URL
	resp, err := http.Get(c.Value)
	if err != nil {
		ctx.Logger.Error("error making GET request", err, appcontext.Fields{})
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	ctx.Logger.Print("resp", resp)
	fmt.Println("url", resp.Request.URL.String())

	if resp.Request.URL.String() != c.Value {
		parsedURL, _ = url.Parse(resp.Request.URL.String())
		c.Value = parsedURL.String()
	}

	// resolve IP address
	ipAddr, err := net.ResolveIPAddr("ip", parsedURL.Hostname())
	if err != nil {
		ctx.Logger.Error("error resolving ip address", err, appcontext.Fields{})
		return nil, err
	}

	// Build the full URL
	fullURL := parsedURL.String()

	// Measure response time
	startTime := time.Now()
	resp, err = http.Get(fullURL)
	if err != nil {
		ctx.Logger.Error("error measuring response time", err, appcontext.Fields{})
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// Calculate response time
	responseTime := time.Since(startTime)

	return &modelresponse.Check{
		Name:             c.Value,
		IsUp:             true,
		ResponseTimeInMS: responseTime.Milliseconds(),
		Scheme:           parsedURL.Scheme,
		IPResolves:       []string{ipAddr.String()},
		SSL:              modelresponse.CheckSSL{},
	}, nil
}

func (c Check) tcp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}

func (c Check) icmp(ctx *appcontext.AppContext) (*modelresponse.Check, error) {
	return nil, nil
}
