package domain

import (
	"net"
)

func IPLookup(url string) ([]string, error) {
	ips, err := net.LookupIP(url)
	if err != nil {
		return []string{}, err
	}

	// convert each IP address from net.IP to string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	return ipStrings, nil
}
