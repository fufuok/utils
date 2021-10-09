package utils

import (
	"strings"
)

// ValidOptionalPort reports whether port is either an empty string
// or matches /^:\d*$/
func ValidOptionalPort(port string) bool {
	if port == "" {
		return true
	}
	if port[0] != ':' {
		return false
	}
	for _, b := range port[1:] {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}

// SplitHostPort separates host and port. If the port is not valid, it returns
// the entire input as host, and it doesn't check the validity of the host.
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
func SplitHostPort(hostPort string) (host, port string) {
	host = hostPort

	colon := strings.LastIndexByte(host, ':')
	if colon != -1 && ValidOptionalPort(host[colon:]) {
		host, port = host[:colon], host[colon+1:]
	}

	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}

	return
}

// ReplaceHost 返回 b 的主机名 + a 的端口
// e.g. ReplaceHost("a.cn:77", "b.cn:88") == "b.cn:77"
func ReplaceHost(a, b string) string {
	_, port := SplitHostPort(a)
	host, _ := SplitHostPort(b)

	// [fe80::1]:80
	if strings.Contains(host, ":") {
		host = "[" + host + "]"
	}

	// b.cn [fe80::1]
	if port == "" {
		return host
	}

	// b.cn:77
	return host + ":" + port
}
