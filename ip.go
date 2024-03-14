package utils

import (
	"errors"
	"math"
	"math/big"
	"net"
	"strconv"
	"strings"
)

var (
	ErrInvalidHostPort = errors.New("invalid Host or Port")
)

// IsPrivateIP reports whether ip is a private address, according to
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

// IPv62Int IPv6 转数值
func IPv62Int(ip net.IP) *big.Int {
	ip6 := ip.To16()
	if ip6 == nil {
		return big.NewInt(-1)
	}
	ipInt := big.NewInt(0)
	ipInt.SetBytes(ip6)
	return ipInt
}

// IPv42Long IPv4 转数值
func IPv42Long(ip net.IP) int {
	ip4 := ip.To4()
	if ip4 == nil {
		return -1
	}
	return int(ip4[0])<<24 | int(ip4[1])<<16 | int(ip4[2])<<8 | int(ip4[3])
}

// IPv42LongLittle IPv4 转小端数值
func IPv42LongLittle(ip net.IP) int {
	ip4 := ip.To4()
	if ip4 == nil {
		return -1
	}
	return int(ip4[3])<<24 | int(ip4[2])<<16 | int(ip4[1])<<8 | int(ip4[0])
}

// Int2IPv6 数值转 IPv4
func Int2IPv6(ipInt *big.Int) net.IP {
	if ipInt.Sign() == -1 {
		return nil
	}

	ipBytes := ipInt.Bytes()
	n := len(ipBytes)
	if n > 16 {
		return nil
	}

	// 前面补零, 补齐 16 位
	if n < 16 {
		padding := make([]byte, 16-len(ipBytes))
		ipBytes = append(padding, ipBytes...)
	}

	ip := net.IP(ipBytes)
	return ip
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

// LongLittle2IPv4 小端数值转 IPv4
func LongLittle2IPv4(n int) net.IP {
	if n > 4294967295 || n < 0 {
		return nil
	}

	ip4 := make(net.IP, net.IPv4len)
	ip4[3] = byte(n >> 24)
	ip4[2] = byte(n >> 16)
	ip4[1] = byte(n >> 8)
	ip4[0] = byte(n)
	return ip4
}

// IPv6String2Int IPv6 字符串转数值
func IPv6String2Int(ip string) *big.Int {
	return IPv62Int(net.ParseIP(ip))
}

// IPv4String2Long IPv4 字符串转数值
func IPv4String2Long(ip string) int {
	return IPv42Long(net.ParseIP(ip))
}

// IPv4String2LongLittle IPv4 字符串转数值(小端)
func IPv4String2LongLittle(ip string) int {
	return IPv42LongLittle(net.ParseIP(ip))
}

// Int2IPv6String 数值转 IPv6 字符串
func Int2IPv6String(n *big.Int) string {
	ip6 := Int2IPv6(n)
	if ip6 == nil {
		return ""
	}
	return ip6.String()
}

// Long2IPv4String 数值转 IPv4 字符串
func Long2IPv4String(n int) string {
	ip4 := Long2IPv4(n)
	if ip4 == nil {
		return ""
	}
	return ip4.String()
}

// LongLittle2IPv4String 数值(小端)转 IPv4 字符串
func LongLittle2IPv4String(n int) string {
	ip4 := LongLittle2IPv4(n)
	if ip4 == nil {
		return ""
	}
	return ip4.String()
}

// InIPNetString 是否包含在指定 IPNet 列表中
func InIPNetString(ip string, ipNets map[*net.IPNet]struct{}) bool {
	return InIPNet(net.ParseIP(ip), ipNets)
}

// InIPNet 是否包含在指定 IPNet 列表中
func InIPNet(ip net.IP, ipNets map[*net.IPNet]struct{}) bool {
	if len(ip) == 0 {
		return false
	}

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

// IsIP 判断是否为合法 IPv4 / IPv6
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsIPv4 判断是否为合法 IPv4
func IsIPv4(ip string) bool {
	return ParseIPv4(ip) != nil
}

// IsIPv6 判断是否为合法 IPv6
func IsIPv6(ip string) bool {
	return ParseIPv6(ip) != nil
}

// ParseIPv4 判断是否为合法 IPv4 并解析
func ParseIPv4(ip string) net.IP {
	if strings.Contains(ip, ".") && !strings.Contains(ip, ":") {
		return net.ParseIP(ip)
	}
	return nil
}

// ParseIPv6 判断是否为合法 IPv6 并解析
func ParseIPv6(ip string) net.IP {
	if strings.Contains(ip, ":") {
		return net.ParseIP(ip)
	}
	return nil
}

// ParseIP 解析 IP 并返回是否为 IPv6
func ParseIP(s string) (net.IP, bool) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, false
	}
	return ip, strings.Contains(s, ":")
}

// ParseIPx 解析 IP, 包括数字形态
func ParseIPx(s string) net.IP {
	if s == "" {
		return nil
	}

	allNumeric := true
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '.' || char == ':' {
			return net.ParseIP(s)
		}
		if char < '0' || char > '9' {
			allNumeric = false
		}
	}

	// 数字转 IPv4
	if allNumeric {
		if n, err := strconv.Atoi(s); err == nil {
			return Long2IPv4(n)
		}
	}
	return nil
}

// ParseHostPort 解析 IP 和端口
func ParseHostPort(s string) (net.IP, uint16, bool, error) {
	h, p := SplitHostPort(s)
	ip, isIPv6 := ParseIP(h)
	port := MustInt(p)
	if ip == nil || port > math.MaxUint16 {
		return nil, 0, false, ErrInvalidHostPort
	}
	return ip, uint16(port), isIPv6, nil
}
