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

// Average 求数字切片的平均值, 可选指定保留小数位数
func Average[T generic.Numeric](xs []T, precision ...int) float64 {
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
	v := float64(m) / float64(len(xs))
	if len(precision) > 0 && precision[0] > 0 {
		return utils.Round(v, precision[0])
	}
	return v
}

// Chunk 按指定 size 对元素切片分组
func Chunk[T any, S ~[]T](ss S, size int) []S {
	if size <= 0 {
		if ss == nil {
			return []S{}
		}
		return []S{ss}
	}

	length := len(ss)
	num := length / size
	if length%size != 0 {
		num += 1
	}

	res := make([]S, 0, num)
	for i := 0; i < num; i++ {
		last := (i + 1) * size
		if last > length {
			last = length
		}
		res = append(res, ss[i*size:last:last])
	}
	return res
}
