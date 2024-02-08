package netinspect

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type SSL struct {
	IsHTTPS  bool
	ExpireAt time.Time
	Issuer   string
}

func CheckSSL(ctx *appcontext.AppContext, host string, port int) (result *SSL, err error) {
	span := sentryio.NewSpan(ctx.Context, "[util][netinspect] check ssl")
	defer span.Finish()

	result = &SSL{}

	// get data
	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout: tlsTimeout,
	}, "tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), nil)
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
