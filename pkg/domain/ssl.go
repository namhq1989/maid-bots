package domain

import (
	"crypto/tls"
	"net"
	"strconv"
	"time"
)

type SSLData struct {
	IsHTTPS  bool
	ExpireAt time.Time
	Issuer   string
}

func CheckSSL(domain string, port int) (result *SSLData, err error) {
	result = &SSLData{}

	// ssl data
	portStr := strconv.Itoa(port)
	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout: 10 * time.Second,
	}, "tcp", net.JoinHostPort(domain, portStr), nil)
	result.IsHTTPS = err == nil
	if err != nil {
		result.ExpireAt = time.Time{}
	} else {
		defer func() { _ = conn.Close() }()

		// get the state of the connection to access the certificates
		state := conn.ConnectionState()

		// check if there are any peer certificates
		if len(state.PeerCertificates) > 0 {
			// the first certificate in the chain is the leaf certificate (your domain's certificate)
			leafCert := state.PeerCertificates[0]

			result.Issuer = leafCert.Issuer.CommonName
			result.ExpireAt = leafCert.NotAfter
		}
	}

	return
}
