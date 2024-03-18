//go:build go1.18
// +build go1.18

package xhash

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestGenHasher64(t *testing.T) {
	hasher0 := GenHasher64[string]()
	var h0 string
	h0 = "ff"
	assert.Equal(t, hasher0(h0), hasher0(h0))

	hasher1 := GenHasher64[int]()
	var h1 int
	h1 = 123
	assert.Equal(t, hasher1(h1), hasher1(h1))

	type L1 int
	type L2 L1
	hasher2 := GenHasher64[L2]()
	var h2 L2
	h2 = 123
	assert.Equal(t, hasher2(h2), hasher2(h2))
	assert.Equal(t, hasher1(h1), hasher2(h2))

	type foo struct {
		x int
		y int
	}
	hasher3 := GenHasher64[*foo]()
	h3 := new(foo)
	h31 := h3
	assert.Equal(t, hasher3(h3), hasher3(h31))

	hasher4 := GenHasher[float64]()
	assert.Equal(t, hasher4(3.1415926), hasher4(3.1415926))
	assert.NotEqual(t, hasher4(3.1415926), hasher4(3.1415927))

	hasher5 := GenHasher[complex128]()
	assert.Equal(t, hasher5(complex(3, 5)), hasher5(complex(3, 5)))
	assert.NotEqual(t, hasher5(complex(4, 5)), hasher5(complex(3, 5)))

	hasher6 := GenHasher[byte]()
	assert.Equal(t, hasher6('\n'), hasher6(10))
	assert.NotEqual(t, hasher6('\r'), hasher6('\n'))

	hasher7 := GenHasher[uintptr]()
	assert.Equal(t, hasher7(8), hasher7(8))
	assert.NotEqual(t, hasher7(7), hasher7(8))
}
