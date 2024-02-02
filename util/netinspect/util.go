package netinspect

import "time"

const (
	TypeDomain = "domain"
	TypeIP     = "ip"

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	PortHTTP  = 80
	PortHTTPS = 443

	tlsTimeout = 5 * time.Second
	tcpTimeout = 5 * time.Second
)

func getPortFromScheme(scheme string) int {
	if scheme == SchemeHTTPS {
		return PortHTTPS
	}
	return PortHTTP
}
