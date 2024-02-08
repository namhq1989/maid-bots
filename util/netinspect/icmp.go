package netinspect

import (
	"errors"
	"net"
	"time"

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
	span := sentryio.NewSpan(ctx.Context, "[util][netinspect] is valid icmp")
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
	span := sentryio.NewSpan(ctx.Context, "[util][netinspect] check icmp")
	defer span.Finish()

	// setup pinger
	pinger, err := probing.NewPinger(address)
	if err != nil {
		return nil, err
	}
	pinger.Timeout = 10 * time.Second
	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	stats := pinger.Statistics()
	if stats.PacketLoss == 100 {
		return nil, errors.New("target unreachable")
	}
	return &ICMP{
		PackageTransmitted: stats.PacketsSent,
		PackageReceived:    stats.PacketsRecv,
		PackageLoss:        stats.PacketLoss,
		IPAddress:          stats.IPAddr.String(),
		ResponseTimeInMs:   stats.AvgRtt.Milliseconds(),
	}, nil
}
