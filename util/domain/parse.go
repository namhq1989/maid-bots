package domain

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type Information struct {
	IsHTTPS bool
	Name    string
	Scheme  string
	Port    int
}

const (
	schemeHTTP = "http"
	portHTTP   = 80

	schemeHTTPS = "https"
	portHTTPS   = 443
)

func Parse(input string) (*Information, error) {
	if !isValidDomain(input) {
		return nil, errors.New("invalid domain name")
	}

	// add a unrecognized scheme if none is present to help with parsing
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
	if scheme != schemeHTTP && scheme != schemeHTTPS {
		return nil, errors.New("invalid scheme: only http and https are supported")
	}

	// return the extracted and validated scheme and domain
	return &Information{
		Name:    domain,
		Scheme:  scheme,
		Port:    getPortFromScheme(scheme),
		IsHTTPS: scheme == schemeHTTPS,
	}, nil
}

func isValidDomain(domain string) bool {
	regex := `^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}(/([a-zA-Z0-9/-]+)?)?$`
	re := regexp.MustCompile(regex)
	return re.MatchString(domain)
}

func getPortFromScheme(scheme string) int {
	if scheme == schemeHTTPS {
		return portHTTPS
	}

	return portHTTP
}
