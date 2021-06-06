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
}

func Logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

// HumanBaseBytes 数字的数量级表示
func HumanBaseBytes(v uint64, base float64, sizes []string) string {
	if v < 10 {
		return fmt.Sprintf("%d B", v)
	}
	e := math.Floor(Logn(float64(v), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(v)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %v"
	if val < 10 {
		f = "%.1f %v"
	}

	return fmt.Sprintf(f, val, suffix)
}

// HumanIntBytes 数字的数量级表示
func HumanIntBytes(v int) string {
	return HumanBytes(uint64(v))
}

// HumanIntIBytes 数字的数量级表示
func HumanIntIBytes(v int) string {
	return HumanIBytes(uint64(v))
}

// HumanBytes 数字的数量级表示
// e.g. HumanBytes(82854982) -> 83 MB
func HumanBytes(v uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return HumanBaseBytes(v, 1000, sizes)
}

// HumanIBytes 数字的数量级表示
// e.g. HumanIBytes(82854982) -> 79 MiB
func HumanIBytes(v uint64) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return HumanBaseBytes(v, 1024, sizes)
}

// ParseHumanBytes 解析数字的数量级表示
// e.g. ParseBytes("42 MB") -> 42000000, nil
// e.g. ParseBytes("42 mib") -> 44040192, nil
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

	extra := strings.ToLower(strings.TrimSpace(s[lastDigit:]))
	if m, ok := bytesSizeTable[extra]; ok {
		f *= float64(m)
		if f >= math.MaxUint64 {
			return 0, fmt.Errorf("too large: %v", s)
		}
		return uint64(f), nil
	}

	return 0, fmt.Errorf("unhandled size name: %v", extra)
}
