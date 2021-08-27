package utils

import (
	"errors"
	"net"
)

// IsPrivate reports whether ip is a private address, according to
// RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
// Ref: go1.17+ func (ip IP) IsPrivate() bool
func IsPrivateIP(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		// Following RFC 1918, Section 3. Private Address Space which says:
		//   The Internet Assigned Numbers Authority (IANA) has reserved the
		//   following three blocks of the IP address space for private internets:
		//     10.0.0.0        -   10.255.255.255  (10/8 prefix)
		//     172.16.0.0      -   172.31.255.255  (172.16/12 prefix)
		//     192.168.0.0     -   192.168.255.255 (192.168/16 prefix)
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1]&0xf0 == 16) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	// Following RFC 4193, Section 8. IANA Considerations which says:
	//   The IANA has assigned the FC00::/7 prefix to "Unique Local Unicast".
	return len(ip) == net.IPv6len && ip[0]&0xfe == 0xfc
}

// IsPrivateIPString 是否为私有 IP
func IsPrivateIPString(ip string) bool {
	return IsPrivateIP(net.ParseIP(ip))
}

// IsInternalIPv4 是否为内网 IPv4, 包含 NAT 专用网段 RFC6598, 比如华为云 ELB 的 100.125.0.0/16
func IsInternalIPv4(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 ||
		ip4[0] == 100 && ip4[1] >= 64 ||
		ip4[0] == 169 && ip4[1] == 254 ||
		ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 ||
		ip4[0] == 192 && ip4[1] == 168
}

// GetNotInternalIPv4 如果是内网 IPv4 则使用默认值, flag 为真是必定返回一个 IP
func GetNotInternalIPv4(ip, defaultIP net.IP, flag ...bool) net.IP {
	if IsInternalIPv4(ip) {
		if defaultIP == nil && len(flag) > 0 && flag[0] {
			return ip
		}

		return defaultIP
	}

	return ip
}

// IsInternalIPv4String 是否为内网 IPv4
func IsInternalIPv4String(ip string) bool {
	return IsInternalIPv4(net.ParseIP(ip))
}

// GetNotInternalIPv4String 如果是内网 IPv4 则使用默认值
func GetNotInternalIPv4String(ip, defaultIP string, flag ...bool) string {
	if IsInternalIPv4String(ip) {
		if defaultIP == "" && len(flag) > 0 && flag[0] {
			return ip
		}

		return defaultIP
	}

	return ip
}

// IPv42Long IPv4 转数值
func IPv42Long(ip net.IP) int {
	ip4 := ip.To4()
	if ip4 == nil {
		return -1
	}

	return int(ip4[0])<<24 | int(ip4[1])<<16 | int(ip4[2])<<8 | int(ip4[3])
}

// Long2IPv4 数值转 IPv4
func Long2IPv4(n int) net.IP {
	if n > 4294967295 || n < 0 {
		return nil
	}

	ip4 := make(net.IP, net.IPv4len)
	ip4[0] = byte(n >> 24)
	ip4[1] = byte(n >> 16)
	ip4[2] = byte(n >> 8)
	ip4[3] = byte(n)

	return ip4
}

// IPv4String2Long IPv4 字符串转数值
func IPv4String2Long(ip string) int {
	return IPv42Long(net.ParseIP(ip))
}

// Long2IPv4String 数值转 IPv4 字符串
func Long2IPv4String(n int) string {
	ip4 := Long2IPv4(n)
	if ip4 == nil {
		return ""
	}

	return Long2IPv4(n).String()
}

// InIPNetString 是否包含在指定 IPNet 列表中
func InIPNetString(ip string, ipNets map[*net.IPNet]struct{}) bool {
	return InIPNet(net.ParseIP(ip), ipNets)
}

// InIPNet 是否包含在指定 IPNet 列表中
func InIPNet(ip net.IP, ipNets map[*net.IPNet]struct{}) bool {
	for ipNet := range ipNets {
		if ipNet.Contains(ip) {
			return true
		}
	}

	return false
}

// GetIPPort 返回 IP 和 端口
func GetIPPort(addr net.Addr) (ip net.IP, port int, err error) {
	switch v := addr.(type) {
	case *net.UDPAddr:
		ip = v.IP
		port = v.Port
	case *net.TCPAddr:
		ip = v.IP
		port = v.Port
	default:
		err = errors.New("not TCPAddr or UDPAddr")
	}

	return
}
