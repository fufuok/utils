package xcrypto

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"strings"
	"time"
)

var (
	ErrInvalidParam = errors.New("invalid parameter")
	ErrInvalidCert  = errors.New("invalid certificate")
)

// GetCertificate 获取域名证书信息
func GetCertificate(network, addr string, timeout time.Duration, tlsConf *tls.Config) (*x509.Certificate, error) {
	addr = strings.TrimSpace(addr)
	if addr == "" {
		return nil, ErrInvalidParam
	}
	if !strings.HasSuffix(addr, ":443") {
		addr += ":443"
	}
	if strings.HasPrefix(addr, "https://") {
		addr = addr[8:]
	}

	dialer := new(net.Dialer)
	if timeout > 0 {
		dialer.Timeout = timeout
	}
	if tlsConf == nil {
		tlsConf = new(tls.Config)
	}
	conn, err := tls.DialWithDialer(dialer, network, addr, tlsConf)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		return certs[0], nil
	}
	return nil, ErrInvalidCert
}
