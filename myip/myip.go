package myip

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/fufuok/utils"
)

var externalIPAPI = map[string][]string{
	"ipv4": {
		"https://4.ipw.cn",
		"http://ipinfo.io/ip",
		"http://ifconfig.me/ip",
		"http://ident.me",
		"http://myexternalip.com/raw",

		"https://api64.ipify.org",
		"https://api.ipify.org",
	},
	"ipv6": {
		"https://6.ipw.cn",
		"http://ifconfig.me/ip",
		"http://ident.me",
		"http://myexternalip.com/raw",
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
		ip = getExternalIP("ipv4", false)
		if ip == "" {
			ip = getExternalIP("ipv6", false)
		}
		if ip != "" {
			break
		}
	}
	return ip
}

// ExternalIP 获取外网地址 (出口公网地址)
func ExternalIP(v ...string) string {
	if len(v) > 0 && v[0] == "ipv6" {
		return ExternalIPv6()
	}
	return ExternalIPv4()
}

// ExternalIPv4 获取外网地址 (IPv4)
func ExternalIPv4() string {
	return getExternalIP("ipv4", true)
}

// ExternalIPv6 获取外网地址 (IPv6)
func ExternalIPv6() string {
	return getExternalIP("ipv6", true)
}

// 逐项请求外网地址
func getExternalIP(v string, strict bool) string {
	if v != "ipv6" {
		v = "ipv4"
	}

	for _, u := range externalIPAPI[v] {
		if ip, ok := getAPI(u); ok {
			if strict {
				switch v {
				case "ipv6":
					if strings.Count(ip, ":") <= 1 {
						continue
					}
				default:
					if strings.Count(ip, ":") > 0 || strings.Count(ip, ".") != 3 {
						continue
					}
				}
			}
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
	return "", false
}

// InternalIPAny 获取内网地址
func InternalIPAny() string {
	ip := InternalIPv4()
	if ip == "" {
		ip = InternalIPv6()
	}
	return ip
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

// LocalIP 获取本地地址 (第一个), 可指定要排除的接口, 比如: "lo", "vpp"
func LocalIP(exclude ...string) string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	var ip net.IP
	for _, i := range ifaces {
		if utils.InStrings(exclude, i.Name) {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				ip = net.IPv4zero
			}
			if !ip.IsLinkLocalUnicast() && !ip.IsLoopback() && ip.To4() != nil {
				return ip.String()
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

// InterfaceAddrs 获取所有带 IP 的接口和对应的所有 IP
// 排除本地链路地址和环回地址
func InterfaceAddrs(v ...string) (map[string][]net.IP, error) {
	ifAddrs := make(map[string][]net.IP)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ifAddrs, err
	}

	var (
		ip net.IP
		t  string
	)
	if len(v) > 0 {
		t = strings.ToLower(v[0])
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return ifAddrs, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				ip = net.IPv4zero
			}
			if ip.IsLinkLocalUnicast() || ip.IsLoopback() {
				continue
			}
			switch t {
			case "ipv6":
				if ip.To4() != nil {
					continue
				}
			case "ipv4":
				if ip.To4() == nil {
					continue
				}
			}
			ifAddrs[i.Name] = append(ifAddrs[i.Name], ip)
		}
	}
	return ifAddrs, nil
}
