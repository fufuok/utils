//go:build go1.18
// +build go1.18

package xslices

import (
	"strings"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/generic"
)

// ToString 将切片拼接成字符串
func ToString[T any](xs []T, sep string) string {
	switch len(xs) {
	case 0:
		return ""
	case 1:
		return utils.MustString(xs[0])
	default:
	}

	ss := make([]string, len(xs))
	for i, v := range xs {
		ss[i] = utils.MustString(v)
	}
	return strings.Join(ss, sep)
}

// Average 求数字切片的平均值
func Average[T generic.Numeric](xs []T) float64 {
	switch len(xs) {
	case 0:
		return 0
	case 1:
		return float64(xs[0])
	default:
	}

	var m T
	for _, x := range xs {
		m += x
	}
	return float64(m) / float64(len(xs))
}
