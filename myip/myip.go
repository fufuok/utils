package myip

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var externalIPAPI = map[string][]string{
	"ipv4": {
		"https://ipv4.ipw.cn/api/ip/myip",
		"https://api-ipv4.ip.sb/ip",
		"https://api.ipify.org",
		"http://ip-api.com/line/?fields=query",
		"http://ifconfig.me/ip",
		"http://ident.me",
		"http://myexternalip.com/raw",
		"http://ip.42.pl/short",
	},
	"ipv6": {
		"https://ipv6.ipw.cn/api/ip/myip",
		"https://api-ipv6.ip.sb/ip",
		"https://api64.ipify.org",
	},
}

// ExternalIPAny 获取外网地址
func ExternalIPAny(retries ...int) string {
	n := 1
	if len(retries) > 0 && retries[0] > 0 {
		n += retries[0]
	}

	ip := ""
	for i := 0; i < n; i++ {
		ip = ExternalIPv4()
		if ip == "" {
			ip = ExternalIPv6()
		}
		if ip != "" {
			break
		}
	}

	return ip

}

// ExternalIP 获取外网地址 (出口公网地址)
func ExternalIP(v ...string) string {
	if len(v) > 0 && v[0] != "ipv4" {
		return ExternalIPv6()
	}

	return ExternalIPv4()
}

// ExternalIPv4 获取外网地址 (IPv4)
func ExternalIPv4() string {
	return getExternalIP("ipv4")
}

// ExternalIPv6 获取外网地址 (IPv6)
func ExternalIPv6() string {
	if ip := getExternalIP("ipv6"); ip != "" && strings.Count(ip, ":") > 1 {
		return ip
	}

	return ""
}

// 逐项请求外网地址
func getExternalIP(v string) string {
	if v != "ipv4" {
		v = "ipv6"
	}

	for _, u := range externalIPAPI[v] {
		if ip, ok := getAPI(u); ok {
			return ip
		}
	}

	return ""
}

// 请求 API 获取公网 IP
func getAPI(u string) (string, bool) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(u)
	if err != nil {
		return "", false
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	ip := net.ParseIP(strings.TrimSpace(string(b)))
	if ip != nil {
		return ip.String(), true
	}

	return "", resp.StatusCode == http.StatusOK
}

// InternalIPv4 获取内网地址 (IPv4)
func InternalIPv4() string {
	return InternalIP("", "udp4")
}

// InternalIPv6 获取内网地址 (临时 IPv6 地址)
func InternalIPv6() string {
	return InternalIP("[2001:4860:4860::8888]:53", "udp6")
}

// InternalIP 获取内网地址 (出口本地地址)
func InternalIP(dstAddr, network string) string {
	if dstAddr == "" {
		dstAddr = "8.8.8.8:53"
	}
	if network == "" {
		network = "udp"
	}

	conn, err := net.DialTimeout(network, dstAddr, time.Second)
	if err != nil {
		return ""
	}

	defer func() {
		_ = conn.Close()
	}()

	addr := conn.LocalAddr().String()
	ip := net.ParseIP(addr).String()
	if ip == "<nil>" {
		ip, _, _ = net.SplitHostPort(addr)
	}

	return ip
}

// LocalIP 获取本地地址 (第一个)
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok &&
				!ipnet.IP.IsLinkLocalUnicast() && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// LocalIPv4s 获取所有本地地址 IPv4
func LocalIPv4s() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok &&
				!ipnet.IP.IsLinkLocalUnicast() && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}
