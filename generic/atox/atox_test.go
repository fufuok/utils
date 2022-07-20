//go:build go1.18
// +build go1.18

package atox_test

import (
	"testing"

	"github.com/fufuok/utils/generic/atox"
)

type Uint8 uint8

func TestN(t *testing.T) {
	eq := func(x, y any) {
		if x != y {
			t.Errorf("%v (%T) != %v (%T)", x, x, y, y)
		}
	}
	eq(atox.Must[float32]("1.0"), float32(1.0))
	eq(atox.Must[float64]("1.0"), float64(1.0))
	eq(atox.Must[uint8]("1"), uint8(1))
	eq(atox.Must[Uint8]("1"), Uint8(1))
}
