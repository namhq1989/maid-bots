package service

import (
	"fmt"

	modelresponse "github.com/namhq1989/maid-bots/internal/models/response"
	"github.com/namhq1989/maid-bots/pkg/domain"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Domain struct {
	Name string
}

func (s Domain) Check(ctx *appcontext.AppContext) (*modelresponse.PublicCheckDomain, error) {
	ctx.AddLogData(appcontext.Fields{"domain": s.Name})

	var result = &modelresponse.PublicCheckDomain{}

	// get domain data
	domainData, err := domain.Parse(s.Name)
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
	result.DomainName = domainData.Name
	result.Scheme = domainData.Scheme
	result.IsHTTPS = sslData.IsHTTPS
	result.Issuer = sslData.Issuer
	result.ExpireAt = modelresponse.NewTimeResponse(sslData.ExpireAt)

	// response time
	result.ResponseTimeInMS, err = domain.MeasureResponseTime(fmt.Sprintf("%s://%s", domainData.Scheme, domainData.Name))
	if err == nil {
		result.IsUp = true
	}

	// ip
	result.IPResolves, _ = domain.IPLookup(domainData.Name)

	return result, err
}
