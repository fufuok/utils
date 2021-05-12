package utils

import (
	"bytes"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Round 四舍五入, ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision (十进制小数点后数字的数目) 进行四舍五入的结果
// precision 也可以是负数或零
// Ref: thinkeridea/go-extend
func Round(val float64, precision int) float64 {
	if precision == 0 {
		return math.Round(val)
	}

	p := math.Pow10(precision)
	if precision < 0 {
		return math.Floor(val*p+0.5) * math.Pow10(-precision)
	}

	return math.Floor(val*p+0.5) / p
}

// 浮点数转千分位分隔字符串
// Ref: dustin/go-humanize
// e.g. Commaf(834142.32) -> 834,142.32
func Commaf(v float64) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}

	return buf.String()
}

// big.Float 千分位分隔字符串
// Ref: dustin/go-humanize
func BigCommaf(v *big.Float) string {
	buf := &bytes.Buffer{}
	if v.Sign() < 0 {
		buf.Write([]byte{'-'})
		v.Abs(v)
	}

	comma := []byte{','}

	parts := strings.Split(v.Text('f', -1), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}

	return buf.String()
}
