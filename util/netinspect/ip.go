package netinspect

import (
	"net"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type IP struct {
	List []string
}

func IPLookup(ctx *appcontext.AppContext, host string) (*IP, error) {
	span := sentryio.NewSpan(ctx.Context, "ip lookup", "")
	defer span.Finish()

	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	// convert each IP address from net.IP to string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	return &IP{
		List: ipStrings,
	}, nil
}
