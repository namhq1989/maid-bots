package netinspect

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type URL struct {
	Type   string
	Value  string
	Host   string
	Scheme string
	Port   int
}

func ParseURL(ctx *appcontext.AppContext, input string) (*URL, error) {
	span := sentryio.NewSpan(ctx.Context, "parse url", "")
	defer span.Finish()

	// parse the URL
	parsedURL, err := url.Parse(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing url: %s", err.Error()))
	}

	var result = &URL{}

	// check if the host is an ip address
	parts := strings.Split(parsedURL.Host, ":")
	ip := net.ParseIP(parts[0])
	if ip != nil {
		result.Type = TypeIP
		result.Value = input
		result.Host = ip.String()

		// if there is a port specified in the URL, extract and display it
		if len(parts) > 1 {
			port, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, errors.New("invalid url port")
			}
			result.Port = port
		} else {
			result.Port = getPortFromScheme(parsedURL.Scheme)
		}
	} else {
		if parsedURL.Scheme == SchemeHTTP {
			// get final url
			finalURL, err := getFinalURL(input)
			if err == nil {
				parsedURL, _ = url.Parse(finalURL)
			}
		}

		result.Type = TypeDomain
		result.Value = parsedURL.String()
		result.Host = parsedURL.Host
		result.Scheme = parsedURL.Scheme
		result.Port = getPortFromScheme(parsedURL.Scheme)
	}

	return result, nil
}

func getFinalURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error making GET request: %s", err.Error()))
	}
	defer func() { _ = resp.Body.Close() }()

	return resp.Request.URL.String(), nil
}
