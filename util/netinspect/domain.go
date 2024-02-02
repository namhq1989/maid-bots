package netinspect

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Domain struct {
	IsHTTPS bool
	Name    string
	Scheme  string
	Port    int
}

func ParseDomain(ctx *appcontext.AppContext, input string) (*Domain, error) {
	span := sentryio.NewSpan(ctx.Context, "parse domain", "")
	defer span.Finish()

	if !isValidDomain(input) {
		return nil, errors.New("invalid domain name")
	}

	// add an unrecognized scheme if none is present to help with parsing
	if !strings.Contains(input, "://") {
		input = "https://" + input
	}

	// parse the input
	parsedURL, err := url.Parse(input)
	if err != nil {
		return nil, errors.New("invalid URL format")
	}

	// validate the hostname (domain)
	domain := parsedURL.Hostname()
	if domain == "" {
		return nil, errors.New("invalid host name")
	}

	// validate the scheme
	scheme := parsedURL.Scheme
	if scheme != SchemeHTTP && scheme != SchemeHTTPS {
		return nil, errors.New("invalid scheme: only http and https are supported")
	}

	// return the extracted and validated scheme and domain
	return &Domain{
		Name:    domain,
		Scheme:  scheme,
		Port:    getPortFromScheme(scheme),
		IsHTTPS: scheme == SchemeHTTPS,
	}, nil
}

func isValidDomain(domain string) bool {
	regex := `^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}(/([a-zA-Z0-9/-]+)?)?$`
	re := regexp.MustCompile(regex)
	return re.MatchString(domain)
}
