//go:build go1.18
// +build go1.18

package generic

import (
	"github.com/fufuok/utils"
	"github.com/fufuok/utils/generic/constraints"
)

// EqualsFn is a function that returns whether 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// HashFn is a function that returns the hash of 't'.
type HashFn[T any] func(t T) uint64

// Equals wraps the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	return a == b
}

// Less wraps the '<' operator for ordered types.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// Compare uses a less function to determine the ordering of 'a' and 'b'. It returns:
//
// * -1 if a < b
//
// * 1 if a > b
//
// * 0 if a == b
func Compare[T any](a, b T, less LessFn[T]) int {
	if less(a, b) {
		return -1
	} else if less(b, a) {
		return 1
	}
	return 0
}

// Max returns the max of a and b.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the min of a and b.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Clamp returns x constrained within [lo:hi] range.
// If x compares less than lo, returns lo; otherwise if hi compares less than x, returns hi; otherwise returns v.
func Clamp[T constraints.Ordered](x, lo, hi T) T {
	return Max(lo, Min(hi, x))
}

// MaxFunc returns the max of a and b using the less func.
func MaxFunc[T any](a, b T, less LessFn[T]) T {
	if less(b, a) {
		return a
	}
	return b
}

// MinFunc returns the min of a and b using the less func.
func MinFunc[T any](a, b T, less LessFn[T]) T {
	if less(a, b) {
		return a
	}
	return b
}

// ClampFunc returns x constrained within [lo:hi] range using the less func.
// If x compares less than lo, returns lo; otherwise if hi compares less than x, returns hi; otherwise returns v.
func ClampFunc[T any](x, lo, hi T, less LessFn[T]) T {
	return MaxFunc(lo, MinFunc(hi, x, less), less)
}

func HashInt(i int) uint64 {
	return utils.HashUint64(uint64(i))
}

func HashUint64(u uint64) uint64 {
	return utils.HashUint64(u)
}

func HashString(s string) uint64 {
	return utils.Sum64(s)
}

func HashBytes(b []byte) uint64 {
	return utils.SumBytes64(b)
}
