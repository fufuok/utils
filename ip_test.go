package utils

import (
	"math/big"
	"net"
	"testing"

	"github.com/fufuok/utils/assert"
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
		assert.Equal(t, v.flag, IsPrivateIPString(v.ip))
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
		assert.Equal(t, v.out, GetNotInternalIPv4String(v.in, v.def, v.flag))
		assert.Equal(t, net.ParseIP(v.out), GetNotInternalIPv4(net.ParseIP(v.in), net.ParseIP(v.def), v.flag))
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
		{"255.255.255.255", IPv4Max},
		{"", -1},
	} {
		assert.Equal(t, v.long, IPv4String2Long(v.ipv4))
		assert.Equal(t, v.ipv4, Long2IPv4String(v.long))
	}

	// go1.17 net.ParseIP("009.001.01.1") == nil
	// Reject non-zero components with leading zeroes.
	// Equal(t, 151060737, IPv4String2Long("009.001.01.1"))
	assert.Equal(t, -1, IPv4String2Long("ff"))
	assert.Equal(t, -1, IPv4String2Long("255.255.255.256"))
	assert.Equal(t, "", Long2IPv4String(IPv4Max+1))
}

func TestIPv6AndInt(t *testing.T) {
	i := big.NewInt(0)
	ipStr := Int2IPv6String(i)
	ipInt := IPv6String2Int(ipStr)
	assert.Equal(t, "::", ipStr)
	assert.True(t, ipInt.Cmp(i) == 0, ipInt)

	i = big.NewInt(0)
	i.SetString("340282366920938463463374607431768211455", 10)
	ipStr = Int2IPv6String(i)
	ipInt = IPv6String2Int(ipStr)
	assert.Equal(t, "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", ipStr)
	assert.True(t, ipInt.Cmp(i) == 0, ipInt)

	i = big.NewInt(0)
	i.SetString("340282366920938463463374607431768211456", 10)
	ipStr = Int2IPv6String(i)
	ipInt = IPv6String2Int(ipStr)
	assert.Equal(t, "", ipStr)
	assert.True(t, ipInt.Cmp(big.NewInt(-1)) == 0, ipInt)

	i = big.NewInt(-2)
	ipStr = Int2IPv6String(i)
	ipInt = IPv6String2Int(ipStr)
	assert.Equal(t, "", ipStr)
	assert.True(t, ipInt.Cmp(big.NewInt(-1)) == 0, ipInt)
}

func TestInIPNetString(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("1.1.1.1/24")
	ipNets := map[*net.IPNet]struct{}{ipNet: {}}
	assert.Equal(t, false, InIPNetString("abc", map[*net.IPNet]struct{}{}))
	assert.Equal(t, false, InIPNetString("::1", map[*net.IPNet]struct{}{}))
	assert.Equal(t, false, InIPNetString("0.0.0.0", map[*net.IPNet]struct{}{}))
	assert.Equal(t, true, InIPNetString("1.1.1.1", ipNets))
	assert.Equal(t, false, InIPNetString("1.1.2.1", ipNets))
	assert.Equal(t, true, InIPNetString("1.1.1.255", ipNets))

	_, ipNet, _ = net.ParseCIDR("0.0.0.0/0")
	ipNets = map[*net.IPNet]struct{}{ipNet: {}}
	assert.Equal(t, false, InIPNetString("abc", ipNets))
	assert.Equal(t, false, InIPNetString("::1", ipNets))
	assert.Equal(t, true, InIPNetString("0.0.0.0", ipNets))
	assert.Equal(t, true, InIPNetString("1.1.1.1", ipNets))
	assert.Equal(t, true, InIPNetString("1.1.1.1", ipNets))
	assert.Equal(t, true, InIPNetString("1.1.2.1", ipNets))
	assert.Equal(t, true, InIPNetString("1.1.1.255", ipNets))

	_, ipNet, _ = net.ParseCIDR("2001:db8::/32")
	ipNets = map[*net.IPNet]struct{}{ipNet: {}}
	assert.Equal(t, true, InIPNetString("2001:db8::1", ipNets))
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
		IP   = "::1"
		PORT = 111
	)
	addr := net.TCPAddr{
		IP:   net.ParseIP(IP),
		Port: PORT,
	}
	ip, port, err := GetIPPort(&addr)
	assert.Equal(t, IP, ip.String())
	assert.Equal(t, IP, ip.To16().String())
	assert.Equal(t, true, ip.To4() == nil)
	assert.Equal(t, PORT, port)
	assert.Equal(t, true, err == nil)

	udpAddr := net.UDPAddr{}
	ip, port, err = GetIPPort(&udpAddr)
	assert.Equal(t, 0, len(ip))
	assert.Equal(t, 0, port)
	assert.Equal(t, true, err == nil)
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		ip string
		v4 bool
	}{
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"::1", false},
		{"::ffff:0.0.0.0", false},
		{"2001:4860:4860::8888", false},
	}
	for _, v := range tests {
		assert.Equal(t, v.v4, IsIPv4(v.ip), v.ip)
		assert.Equal(t, !v.v4, IsIPv6(v.ip), v.ip)
		assert.Equal(t, true, IsIP(v.ip), v.ip)

		ip, isIPv6 := ParseIP(v.ip)
		assert.Equal(t, true, ip != nil)
		assert.Equal(t, !v.v4, isIPv6)
	}
	assert.Equal(t, false, IsIPv4("123"))
	assert.Equal(t, false, IsIPv6("123"))
	ip, isIPv6 := ParseIP("123")
	assert.Equal(t, false, ip != nil)
	assert.Equal(t, false, isIPv6)
}

func TestIPv42LongLittle(t *testing.T) {
	ipv4 := "1.2.3.4"
	longLittle := IPv4String2LongLittle(ipv4)
	assert.Equal(t, "4.3.2.1", Long2IPv4String(longLittle))
	assert.Equal(t, ipv4, LongLittle2IPv4String(longLittle))
}

func TestParseHostPort(t *testing.T) {
	tests := []struct {
		s    string
		ip   string
		port uint16
		v6   bool
	}{
		{"0.0.0.0:80", "0.0.0.0", 80, false},
		{"255.255.255.255:0", "255.255.255.255", 0, false},
		{"[::1]:22", "::1", 22, true},
		{"[::ffff:0.0.0.0]", "0.0.0.0", 0, true},
		{"[2001:4860:4860::8888]:777", "2001:4860:4860::8888", 777, true},
	}
	for _, v := range tests {
		host, port, isv6, err := ParseHostPort(v.s)
		assert.Equal(t, true, host != nil)
		assert.Equal(t, v.ip, host.String())
		assert.Equal(t, v.port, port)
		assert.Equal(t, v.v6, isv6)
		assert.Equal(t, nil, err)
	}
	host, _, _, err := ParseHostPort("0:1")
	assert.Equal(t, false, host != nil)
	assert.Equal(t, false, err == nil)
}

func TestParseIPx(t *testing.T) {
	tests := []struct {
		s      string
		isNil  bool
		isIPv6 bool
	}{
		{"", true, false},
		{" 0.0.0.0", true, false},
		{"[2001::]", true, true},
		{"-1", true, false},
		{"abc", true, false},

		{"0.0.0.0", false, false},
		{"255.255.255.255", false, false},
		{"::1", false, true},
		{"::ffff:0.0.0.0", false, true},
		{"2001:4860:4860::8888", false, true},
		{"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", false, true},
	}
	for _, v := range tests {
		ip, isIPv6 := ParseIPx(v.s)
		assert.Equal(t, v.isNil, ip == nil)
		assert.Equal(t, v.isIPv6, isIPv6)

		ip, isIPv6 = ParseIPxWithNumeric(v.s)
		assert.Equal(t, v.isNil, ip == nil)
		assert.Equal(t, v.isIPv6, isIPv6)
	}

	ip, isIPv6 := ParseIPx("0")
	assert.Nil(t, ip)
	assert.False(t, isIPv6)

	tests = []struct {
		s      string
		isNil  bool
		isIPv6 bool
	}{
		{"-1", true, false},
		{"0", false, false},
		{"4294967295", false, false},
		{"4294967296", false, true},
		{"340282366920938463463374607431768211455", false, true},
		{"340282366920938463463374607431768211456", true, true},
	}
	for _, v := range tests {
		ip, isIPv6 = ParseIPxWithNumeric(v.s)
		assert.Equal(t, v.isNil, ip == nil)
		assert.Equal(t, v.isIPv6, isIPv6)
	}
}

// Ref: gofiber/utils
func Test_IsIPv4(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, IsIPv4("255.255.255.255"))
	assert.Equal(t, true, IsIPv4("174.23.33.100"))
	assert.Equal(t, true, IsIPv4("127.0.0.1"))
	assert.Equal(t, true, IsIPv4("0.0.0.0"))

	assert.Equal(t, false, IsIPv4("255.255.255.256"))
	assert.Equal(t, false, IsIPv4(".0.0.0"))
	assert.Equal(t, false, IsIPv4("0.0.0."))
	assert.Equal(t, false, IsIPv4("0.0.0"))
	assert.Equal(t, false, IsIPv4(".0.0.0."))
	assert.Equal(t, false, IsIPv4("0.0.0.0.0"))
	assert.Equal(t, false, IsIPv4("0"))
	assert.Equal(t, false, IsIPv4(""))
	assert.Equal(t, false, IsIPv4("2345:0425:2CA1::0567:5673:23b5"))
	assert.Equal(t, false, IsIPv4("invalid"))
	assert.Equal(t, false, IsIPv4("189.12.34.260"))
	assert.Equal(t, false, IsIPv4("189.12.260.260"))
	assert.Equal(t, false, IsIPv4("189.260.260.260"))
	assert.Equal(t, false, IsIPv4("999.999.999.999"))
	assert.Equal(t, false, IsIPv4("9999.9999.9999.9999"))
}

// Ref: gofiber/utils
func Benchmark_IsIPv4(b *testing.B) {
	ip := "174.23.33.100"
	var res bool

	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = IsIPv4(ip)
		}
		assert.Equal(b, true, res)
	})

	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = net.ParseIP(ip) != nil
		}
		assert.Equal(b, true, res)
	})
}

// Ref: gofiber/utils
func Test_IsIPv6(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, IsIPv6("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"))
	assert.Equal(t, true, IsIPv6("9396:9549:b4f7:8ed0:4791:1330:8c06:e62d"))
	assert.Equal(t, true, IsIPv6("2345:0425:2CA1::0567:5673:23b5"))
	assert.Equal(t, true, IsIPv6("2001:1:2:3:4:5:6:7"))
	assert.Equal(t, true, IsIPv6("2001::"))
	assert.Equal(t, true, IsIPv6("::"))

	assert.Equal(t, false, IsIPv6("2001::1::1"))
	assert.Equal(t, false, IsIPv6("2001::fffg"))
	assert.Equal(t, false, IsIPv6("1.1.1.1"))
	assert.Equal(t, false, IsIPv6("2001:1:2:3:4:5:6:"))
	assert.Equal(t, false, IsIPv6(":1:2:3:4:5:6:"))
	assert.Equal(t, false, IsIPv6("1:2:3:4:5:6:"))
	assert.Equal(t, false, IsIPv6(""))
	assert.Equal(t, false, IsIPv6("invalid"))
}

// Ref: gofiber/utils
func Benchmark_IsIPv6(b *testing.B) {
	ip := "9396:9549:b4f7:8ed0:4791:1330:8c06:e62d"
	var res bool

	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = IsIPv6(ip)
		}
		assert.Equal(b, true, res)
	})

	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = net.ParseIP(ip) != nil
		}
		assert.Equal(b, true, res)
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=Benchmark_IsIP
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_IsIPv4/fiber-16               88282953                13.25 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv4/fiber-16               90268933                13.23 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv4/default-16             45390993                25.53 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv4/default-16             46329868                25.49 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv6/fiber-16               27761413                44.99 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv6/fiber-16               27327100                43.97 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv6/default-16             21168326                57.92 ns/op            0 B/op          0 allocs/op
// Benchmark_IsIPv6/default-16             20407336                57.98 ns/op            0 B/op          0 allocs/op
