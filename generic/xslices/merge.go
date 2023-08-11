//go:build go1.18
// +build go1.18

package xslices

// Merge 浅拷贝合并多个切片, 不影响原切片
func Merge[E any](s []E, ss ...[]E) []E {
	if len(ss) == 0 || len(s) == 0 {
		return s
	}
	n := len(s)
	for _, v := range ss {
		n += len(v)
	}
	d := make([]E, 0, n)
	d = append(d, s...)
	for _, v := range ss {
		d = append(d, v...)
	}
	return d
}
