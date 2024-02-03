package netinspect

import (
	"net"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcontext"
	probing "github.com/prometheus-community/pro-bing"
)

type ICMP struct {
	PackageTransmitted int
	PackageReceived    int
	PackageLoss        float64
	IPAddress          string
	ResponseTimeInMs   int64
}

func IsValidICMP(ctx *appcontext.AppContext, input string) bool {
	span := sentryio.NewSpan(ctx.Context, "validate icmp", "")
	defer span.Finish()

	ip := net.ParseIP(input)
	if ip != nil {
		return true
	}

	// try resolving as a hostname
	_, err := net.ResolveIPAddr("ip", input)
	return err == nil
}

func CheckICMP(ctx *appcontext.AppContext, address string) (*ICMP, error) {
	span := sentryio.NewSpan(ctx.Context, "check icmp", "")
	defer span.Finish()

	// setup pinger
	pinger, err := probing.NewPinger(address)
	if err != nil {
		return nil, err
	}
	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	stats := pinger.Statistics()
	return &ICMP{
		PackageTransmitted: stats.PacketsSent,
		PackageReceived:    stats.PacketsRecv,
		PackageLoss:        stats.PacketLoss,
		IPAddress:          stats.IPAddr.String(),
		ResponseTimeInMs:   stats.AvgRtt.Milliseconds(),
	}, nil
}
