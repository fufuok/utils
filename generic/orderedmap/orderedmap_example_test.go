//go:build go1.18
// +build go1.18

package orderedmap_test

import (
	"fmt"

	"github.com/fufuok/utils/generic/orderedmap"
)

func ExampleNewOrderedMap() {
	m := orderedmap.NewOrderedMap[string, any]()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set("123", true)

	fmt.Println(m.Len())
	m.Delete("qux")
	fmt.Println(m.Len())

	// Iterate through all elements from oldest to newest:
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}

	// You can also use Back and Prev to iterate in reverse:
	for el := m.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Key, el.Value)
	}

	// Output:
	// 3
	// 2
	// foo bar
	// 123 true
	// 123 true
	// foo bar
}
