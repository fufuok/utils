//go:build go1.18
// +build go1.18

package xhash

import (
	"testing"

	"github.com/fufuok/utils"
)

func TestGenHasher64(t *testing.T) {
	hasher0 := GenHasher64[string]()
	var h0 string
	h0 = "ff"
	utils.AssertEqual(t, hasher0(h0), hasher0(h0))

	hasher1 := GenHasher64[int]()
	var h1 int
	h1 = 123
	utils.AssertEqual(t, hasher1(h1), hasher1(h1))

	type L1 int
	type L2 L1
	hasher2 := GenHasher64[L2]()
	var h2 L2
	h2 = 123
	utils.AssertEqual(t, hasher2(h2), hasher2(h2))
	utils.AssertEqual(t, hasher1(h1), hasher2(h2))

	type foo struct {
		x int
		y int
	}
	hasher3 := GenHasher64[*foo]()
	var h3 = new(foo)
	h31 := h3
	utils.AssertEqual(t, hasher3(h3), hasher3(h31))

	hasher4 := GenHasher[float64]()
	utils.AssertEqual(t, hasher4(3.1415926), hasher4(3.1415926))
	utils.AssertNotEqual(t, hasher4(3.1415926), hasher4(3.1415927))

	hasher5 := GenHasher[complex128]()
	utils.AssertEqual(t, hasher5(complex(3, 5)), hasher5(complex(3, 5)))
	utils.AssertNotEqual(t, hasher5(complex(4, 5)), hasher5(complex(3, 5)))

	hasher6 := GenHasher[byte]()
	utils.AssertEqual(t, hasher6('\n'), hasher6(10))
	utils.AssertNotEqual(t, hasher6('\r'), hasher6('\n'))

	hasher7 := GenHasher[uintptr]()
	utils.AssertEqual(t, hasher7(8), hasher7(8))
	utils.AssertNotEqual(t, hasher7(7), hasher7(8))
}
