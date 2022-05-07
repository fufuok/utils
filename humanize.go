package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Ref: dustin/go-humanize
// IEC Sizes.
// kibis of bits
const (
	Byte = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
	TiByte
	PiByte
	EiByte
)

// SI Sizes.
const (
	IByte = 1
	KByte = IByte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

var bytesSizeTable = map[string]uint64{
	"b":   Byte,
	"kib": KiByte,
	"kb":  KByte,
	"mib": MiByte,
	"mb":  MByte,
	"gib": GiByte,
	"gb":  GByte,
	"tib": TiByte,
	"tb":  TByte,
	"pib": PiByte,
	"pb":  PByte,
	"eib": EiByte,
	"eb":  EByte,
	// Without suffix
	"":   Byte,
	"ki": KiByte,
	"k":  KByte,
	"mi": MiByte,
	"m":  MByte,
	"gi": GiByte,
	"g":  GByte,
	"ti": TiByte,
	"t":  TByte,
	"pi": PiByte,
	"p":  PByte,
	"ei": EiByte,
	"e":  EByte,
	// bps
	"bps":  IByte,
	"kbps": KByte,
	"mbps": MByte,
	"gbps": GByte,
	"tbps": TByte,
	"pbps": PByte,
	"ebps": EByte,
}

func Logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

// HumanBaseBytes 数字的数量级表示
func HumanBaseBytes(v uint64, base float64, sizes []string) string {
	if v < 10 {
		return fmt.Sprintf("%d %s", v, sizes[0])
	}
	e := math.Floor(Logn(float64(v), base))
	n := float64(len(sizes) - 1)
	if e > n {
		e = n
	}
	suffix := sizes[int(e)]
	val := math.Floor(float64(v)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %v"
	if val < 10 {
		f = "%.1f %v"
	}

	return fmt.Sprintf(f, val, suffix)
}

// HumanGBMB 转为 ** GB ** MB
// 1 GB = 1024 MB
func HumanGBMB(v uint64) string {
	sizes := []string{"B", "KB", "MB"}
	g := v / GiByte
	if g == 0 {
		return HumanBaseBytes(v, 1024, sizes)
	}

	m := v - g*GiByte
	return fmt.Sprintf("%d GB %s", g, HumanBaseBytes(m, 1024, sizes))
}

// HumanIntBytes 1 KB = 1000 B
func HumanIntBytes(v int) string {
	return HumanBytes(uint64(v))
}

// HumanIntIBytes 1 KiB = 1024 B
func HumanIntIBytes(v int) string {
	return HumanIBytes(uint64(v))
}

// HumanIntKbps 1 Kbps = 1000 bit
func HumanIntKbps(v int) string {
	return HumanKbps(uint64(v))
}

// HumanBytes 1 KB = 1000 B
// e.g. HumanBytes(82854982) -> 83 MB
func HumanBytes(v uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return HumanBaseBytes(v, 1000, sizes)
}

// HumanIBytes 1 KiB = 1024 B
// e.g. HumanIBytes(82854982) -> 79 MiB
func HumanIBytes(v uint64) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return HumanBaseBytes(v, 1024, sizes)
}

// HumanKbps 1 Kbps = 1000 bit, 传输速率(bit per second, 位每秒)
// e.g. HumanKbps(82854982) -> 83 Mbps
func HumanKbps(v uint64) string {
	sizes := []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps", "Ebps"}
	return HumanBaseBytes(v, 1000, sizes)
}

// ParseHumanBytes 解析数字的数量级表示
// e.g. ParseHumanBytes("42 MB") -> 42000000, nil
// e.g. ParseHumanBytes("42 mib") -> 44040192, nil
func ParseHumanBytes(s string) (uint64, error) {
	lastDigit := 0
	hasComma := false
	for _, r := range s {
		if !(unicode.IsDigit(r) || r == '.' || r == ',') {
			break
		}
		if r == ',' {
			hasComma = true
		}
		lastDigit++
	}

	num := s[:lastDigit]
	if hasComma {
		num = strings.Replace(num, ",", "", -1)
	}

	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}

	extra := ToLower(strings.TrimSpace(s[lastDigit:]))
	if m, ok := bytesSizeTable[extra]; ok {
		f *= float64(m)
		if f >= math.MaxUint64 {
			return 0, fmt.Errorf("too large: %v", s)
		}
		return uint64(f), nil
	}

	return 0, fmt.Errorf("unhandled size name: %v", extra)
}

// MustParseHumanBytes 解析数字的数量级表示
// e.g. MustParseHumanBytes("42 MB") -> 42000000
// e.g. MustParseHumanBytes("-42 mib", 123) -> 123
func MustParseHumanBytes(s string, defaultVal ...uint64) uint64 {
	num, err := ParseHumanBytes(s)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return num
}
