//go:build go1.21
// +build go1.21

package xslices

import (
	"cmp"
)

func Max[T cmp.Ordered](xs ...T) (y T) {
	switch len(xs) {
	case 0:
		return
	case 1:
		return xs[0]
	case 2:
		return max(xs[0], xs[1])
	default:
	}

	y = xs[0]
	for _, v := range xs[1:] {
		y = max(y, v)
	}
	return
}

func Min[T cmp.Ordered](xs ...T) (y T) {
	switch len(xs) {
	case 0:
		return
	case 1:
		return xs[0]
	case 2:
		return min(xs[0], xs[1])
	default:
	}

	y = xs[0]
	for _, v := range xs[1:] {
		y = min(y, v)
	}
	return
}
