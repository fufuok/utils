package utils

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// MaxInt 整数取大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt 整数取小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SumInt 整数和
func SumInt(v ...int) int {
	x := 0
	for _, n := range v {
		x += n
	}
	return x
}

// GetInt 获取 int 结果, 可选指定默认值(若给定了默认值,则返回正整数或 0)
func GetInt(v interface{}, defaultInt ...int) int {
	i := MustInt(v)
	if i <= 0 && len(defaultInt) > 0 {
		return defaultInt[0]
	}
	return i
}

// SearchInt 搜索整数位置(左, 第一个)
func SearchInt(slice []int, n int) int {
	for i, v := range slice {
		if n == v {
			return i
		}
	}

	return -1
}

// InInts 检查整数是否存在于 slice
func InInts(slice []int, n int) bool {
	return SearchInt(slice, n) != -1
}

// Commai 整数转千分位分隔字符串
func Commai(v int) string {
	return Comma(int64(v))
}

// Comma 整数转千分位分隔字符串
// Ref: dustin/go-humanize
// e.g. Comma(834142) -> 834,142
func Comma(v int64) string {
	sign := ""

	// Min int64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))

	return sign + strings.Join(parts[j:], ",")
}

// Commau 整数转千分位分隔字符串
// Ref: dustin/go-humanize
func Commau(v uint64) string {
	sign := ""
	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatUint(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))

	return sign + strings.Join(parts[j:], ",")
}

// BigComma big.Int 千分位分隔字符串
// Ref: dustin/go-humanize
func BigComma(b *big.Int) string {
	sign := ""
	if b.Sign() < 0 {
		sign = "-"
		b.Abs(b)
	}

	athousand := big.NewInt(1000)
	c := (&big.Int{}).Set(b)
	_, m := Bigoom(c, athousand)
	parts := make([]string, m+1)
	j := len(parts) - 1

	mod := &big.Int{}
	for b.Cmp(athousand) >= 0 {
		b.DivMod(b, athousand, mod)
		parts[j] = strconv.FormatInt(mod.Int64(), 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		j--
	}
	parts[j] = strconv.Itoa(int(b.Int64()))
	return sign + strings.Join(parts[j:], ",")
}

// Bigoom big.Int 总数量级
// Ref: dustin/go-humanize
func Bigoom(n, b *big.Int) (float64, int) {
	mag := 0
	m := &big.Int{}
	for n.Cmp(b) >= 0 {
		n.DivMod(n, b, m)
		mag++
	}
	return float64(n.Int64()) + (float64(m.Int64()) / float64(b.Int64())), mag
}
