//go:build go1.18
// +build go1.18

package utils

import (
	"testing"
)

func TestGenHasher64(t *testing.T) {
	hasher0 := GenHasher64[string]()
	var h0 string
	h0 = "ff"
	AssertEqual(t, hasher0(h0), hasher0(h0))

	hasher1 := GenHasher64[int]()
	var h1 int
	h1 = 123
	AssertEqual(t, hasher1(h1), hasher1(h1))

	type L1 int
	type L2 L1
	hasher2 := GenHasher64[L2]()
	var h2 L2
	h2 = 123
	AssertEqual(t, hasher2(h2), hasher2(h2))
	AssertEqual(t, hasher1(h1), hasher2(h2))

	type foo struct {
		x int
		y int
	}
	hasher3 := GenHasher64[*foo]()
	var h3 = new(foo)
	h31 := h3
	AssertEqual(t, hasher3(h3), hasher3(h31))

	hasher4 := GenHasher[float64]()
	AssertEqual(t, hasher4(3.1415926), hasher4(3.1415926))
	AssertNotEqual(t, hasher4(3.1415926), hasher4(3.1415927))

	hasher5 := GenHasher[complex128]()
	AssertEqual(t, hasher5(complex(3, 5)), hasher5(complex(3, 5)))
	AssertNotEqual(t, hasher5(complex(4, 5)), hasher5(complex(3, 5)))

	hasher6 := GenHasher[byte]()
	AssertEqual(t, hasher6('\n'), hasher6(10))
	AssertNotEqual(t, hasher6('\r'), hasher6('\n'))

	hasher7 := GenHasher[uintptr]()
	AssertEqual(t, hasher7(8), hasher7(8))
	AssertNotEqual(t, hasher7(7), hasher7(8))
}
