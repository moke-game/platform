package platformfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"google.golang.org/grpc"
)

// Dial opens a gRPC connection using the process TLS settings.
func Dial(host string, sec sfx.SecuritySettingsParams) (*grpc.ClientConn, error) {
	if sec.MTLSEnable {
		return tools.DialWithSecurity(
			host,
			sec.ClientCert,
			sec.ClientKey,
			sec.ServerName,
			sec.ServerCaCert,
		)
	}
	return tools.DialInsecure(host)
}

// NewClient dials host and wraps the connection with newFn.
func NewClient[T any](
	host string,
	sec sfx.SecuritySettingsParams,
	newFn func(grpc.ClientConnInterface) T,
) (T, error) {
	var zero T
	conn, err := Dial(host, sec)
	if err != nil {
		return zero, err
	}
	return newFn(conn), nil
}
