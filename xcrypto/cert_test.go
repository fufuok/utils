package xcrypto

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/fufuok/utils/assert"
)

func TestGetCertificate(t *testing.T) {
	network := "tcp"
	addr := "www.microsoft.com"
	timeout := 5 * time.Second
	tlsConfig := &tls.Config{InsecureSkipVerify: false}
	cert, err := GetCertificate(network, addr, timeout, tlsConfig)
	assert.Nil(t, err)
	assert.NotNil(t, cert)

	addr = "https://www.microsoft.com"
	cert, err = GetCertificate(network, addr, 0, nil)
	assert.Nil(t, err)
	assert.NotNil(t, cert)

	t.Logf("DNSNames: %v\nNotBefore: %s\nNotAfter: %s\n", cert.DNSNames, cert.NotBefore, cert.NotAfter)
}
