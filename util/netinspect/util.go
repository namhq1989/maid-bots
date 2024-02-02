package netinspect

const (
	TypeDomain = "domain"
	TypeIP     = "ip"

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	PortHTTP  = 80
	PortHTTPS = 443
)

func getPortFromScheme(scheme string) int {
	if scheme == SchemeHTTPS {
		return PortHTTPS
	}
	return PortHTTP
}
