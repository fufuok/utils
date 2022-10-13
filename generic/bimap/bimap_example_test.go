//go:build go1.18
// +build go1.18

package bimap_test

import (
	"fmt"

	"github.com/fufuok/utils/generic/bimap"
)

func Example() {
	var bimap bimap.Bimap[int, string]

	bimap.Add(1, "foo")
	bimap.Add(2, "bar")
	bimap.Add(3, "moo")
	bimap.Add(4, "doo")

	fmt.Println(bimap.GetForward(4))
	fmt.Println(bimap.GetReverse("moo"))
	fmt.Println(bimap.GetReverse("unknown"))
	// Output:
	// doo true
	// 3 true
	// 0 false
}
