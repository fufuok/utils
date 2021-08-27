package utils

import (
	"net"
	"testing"
)

func TestIsPrivateIPString(t *testing.T) {
	for _, v := range []struct {
		ip   string
		flag bool
	}{
		{"::1", false},
		{"fe80::1", false},
		{"0.0.0.0", false},
		{"1.2.3.4", false},
		{"10.0.0.0", true},
		{"172.16.0.0", true},
		{"10.0.0.0", true},
		{"192.168.0.0", true},
		{"192.169.0.0", false},
	} {
		AssertEqual(t, v.flag, IsPrivateIPString(v.ip))
	}
}

func TestGetNotInternalIPv4(t *testing.T) {
	defIP4 := "7.7.7.7"

	for _, v := range []struct {
		in   string
		out  string
		def  string
		flag bool
	}{
		{"1.2.3.4", "1.2.3.4", defIP4, false},
		{"1.2.3.4", "1.2.3.4", defIP4, true},
		{"10.0.0.1", defIP4, defIP4, false},
		{"10.0.0.1", defIP4, defIP4, true},
		{"100.125.1.1", defIP4, defIP4, false},
		{"100.125.1.1", defIP4, defIP4, true},
		{"127.0.0.1", defIP4, defIP4, false},
		{"127.0.0.1", defIP4, defIP4, true},
		{"169.254.1.1", defIP4, defIP4, false},
		{"169.254.1.1", defIP4, defIP4, true},
		{"192.168.1.1", defIP4, defIP4, false},
		{"192.168.1.1", defIP4, defIP4, true},
		{"192.168.1.1", "192.168.1.1", "", true},
	} {
		AssertEqual(t, v.out, GetNotInternalIPv4String(v.in, v.def, v.flag))
		AssertEqual(t, net.ParseIP(v.out), GetNotInternalIPv4(net.ParseIP(v.in), net.ParseIP(v.def), v.flag))
	}
}

func TestIPv4AndLong(t *testing.T) {
	for _, v := range []struct {
		ipv4 string
		long int
	}{
		{"0.0.0.0", 0},
		{"0.0.0.1", 1},
		{"1.2.3.4", 16909060},
		{"10.0.0.0", 167772160},
		{"255.255.255.255", 4294967295},
		{"", -1},
	} {
		AssertEqual(t, v.long, IPv4String2Long(v.ipv4))
		AssertEqual(t, v.ipv4, Long2IPv4String(v.long))
	}

	// go1.17 net.ParseIP("009.001.01.1") == nil
	// Reject non-zero components with leading zeroes.
	// AssertEqual(t, 151060737, IPv4String2Long("009.001.01.1"))
	AssertEqual(t, -1, IPv4String2Long("ff"))
	AssertEqual(t, -1, IPv4String2Long("255.255.255.256"))
	AssertEqual(t, "", Long2IPv4String(4294967296))
}

func TestInIPNetString(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("1.1.1.1/24")
	ipNets := map[*net.IPNet]struct{}{ipNet: {}}
	AssertEqual(t, false, InIPNetString("abc", map[*net.IPNet]struct{}{}))
	AssertEqual(t, false, InIPNetString("::1", map[*net.IPNet]struct{}{}))
	AssertEqual(t, false, InIPNetString("0.0.0.0", map[*net.IPNet]struct{}{}))
	AssertEqual(t, true, InIPNetString("1.1.1.1", ipNets))
	AssertEqual(t, false, InIPNetString("1.1.2.1", ipNets))
	AssertEqual(t, true, InIPNetString("1.1.1.255", ipNets))

	_, ipNet, _ = net.ParseCIDR("0.0.0.0/0")
	ipNets = map[*net.IPNet]struct{}{ipNet: {}}
	AssertEqual(t, false, InIPNetString("abc", ipNets))
	AssertEqual(t, false, InIPNetString("::1", ipNets))
	AssertEqual(t, true, InIPNetString("0.0.0.0", ipNets))
	AssertEqual(t, true, InIPNetString("1.1.1.1", ipNets))
	AssertEqual(t, true, InIPNetString("1.1.1.1", ipNets))
	AssertEqual(t, true, InIPNetString("1.1.2.1", ipNets))
	AssertEqual(t, true, InIPNetString("1.1.1.255", ipNets))

	_, ipNet, _ = net.ParseCIDR("2001:db8::/32")
	ipNets = map[*net.IPNet]struct{}{ipNet: {}}
	AssertEqual(t, true, InIPNetString("2001:db8::1", ipNets))
}

func BenchmarkGetNotInternalIPv4String(b *testing.B) {
	cip := "100.125.1.1"
	fip := "1.1.1.1,2.2.2.2"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetNotInternalIPv4String(cip, fip)
	}
}

func BenchmarkGetNotInternalIPv4StringTrue(b *testing.B) {
	cip := "100.125.1.1"
	fip := ""
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetNotInternalIPv4String(cip, fip, true)
	}
}

// BenchmarkGetNotInternalIPv4String-8       	 7723636	       143.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4String-8       	 8165636	       135.1 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4String-8       	 8405780	       148.5 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 9055933	       145.5 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 7916724	       206.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 7862268	       151.3 ns/op	      16 B/op	       1 allocs/op

func BenchmarkInIPNet(b *testing.B) {
	ip := net.ParseIP("1.1.1.1")
	_, ipNet, _ := net.ParseCIDR("1.1.1.1/24")
	ipNets := map[*net.IPNet]struct{}{ipNet: {}}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = InIPNet(ip, ipNets)
	}
}

func BenchmarkInIPNetString(b *testing.B) {
	ip := "1.1.1.1"
	_, ipNet, _ := net.ParseCIDR("1.1.1.1/24")
	ipNets := map[*net.IPNet]struct{}{ipNet: {}}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = InIPNetString(ip, ipNets)
	}
}

// BenchmarkInIPNet-8         	17226256	        63.36 ns/op	       0 B/op	       0 allocs/op
// BenchmarkInIPNet-8         	19683216	        62.59 ns/op	       0 B/op	       0 allocs/op
// BenchmarkInIPNet-8         	19325880	        63.99 ns/op	       0 B/op	       0 allocs/op
// BenchmarkInIPNetString-8   	 8095532	       148.9 ns/op	      16 B/op	       1 allocs/op
// BenchmarkInIPNetString-8   	 7845603	       147.8 ns/op	      16 B/op	       1 allocs/op
// BenchmarkInIPNetString-8   	 7447910	       148.1 ns/op	      16 B/op	       1 allocs/op

func TestGetIPPort(t *testing.T) {
	var (
		IP       = "::1"
		PORT int = 111
	)
	addr := net.TCPAddr{
		IP:   net.ParseIP(IP),
		Port: PORT,
	}
	ip, port, err := GetIPPort(&addr)
	AssertEqual(t, IP, ip.String())
	AssertEqual(t, IP, ip.To16().String())
	AssertEqual(t, true, ip.To4() == nil)
	AssertEqual(t, PORT, port)
	AssertEqual(t, true, err == nil)

	udpAddr := net.UDPAddr{}
	ip, port, err = GetIPPort(&udpAddr)
	AssertEqual(t, 0, len(ip))
	AssertEqual(t, 0, port)
	AssertEqual(t, true, err == nil)
}
