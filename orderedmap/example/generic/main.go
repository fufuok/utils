//go:build go1.21
// +build go1.21

package main

import (
	"fmt"
	"sort"

	"github.com/fufuok/utils/orderedmap"
)

func main() {
	om := orderedmap.NewOf[int, int32]()
	om.Set(7, 1)
	om.Set(2, 2)
	om.Sort(func(a *orderedmap.PairOf[int, int32], b *orderedmap.PairOf[int, int32]) bool {
		return a.Value() > b.Value()
	})
	for _, k := range om.Keys() {
		v := om.MustGet(k)
		fmt.Println(k, v)
	}

	c := om.Clone()
	m := om.ToMap()
	_ = c
	_ = m

	// use NewOf[string, any]() instead of o := map[string]any{}
	o := orderedmap.NewOf[string, any]()

	// use Set instead of o["a"] = 1
	o.Set("a", 1)
	o.Set("c", 2)

	// add some value with special characters
	o.Set("b", "\\.<>[]{}_-")

	// use Get instead of i, ok := o["a"]
	val, ok := o.Get("a")
	fmt.Println("Get:", val, ok)

	// use Keys instead of for k, v := range o
	keys := o.Keys()
	for _, k := range keys {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	// use o.Delete instead of delete(o, key)
	o.Delete("b")

	fmt.Println("sort the keys:")
	o.SortKeys(sort.Strings)
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("sort by Pair:")
	o.Sort(func(a *orderedmap.PairOf[string, any], b *orderedmap.PairOf[string, any]) bool {
		return a.Value().(int) < b.Value().(int)
	})
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("sort by Pair(reverse):")
	o.Sort(func(a *orderedmap.PairOf[string, any], b *orderedmap.PairOf[string, any]) bool {
		return a.Value().(int) > b.Value().(int)
	})
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}
}

// Output:
// 2 2
// 7 1
// Get: 1 true
// a 1
// c 2
// b \.<>[]{}_-
// sort the keys:
// a 1
// c 2
// sort by Pair:
// a 1
// c 2
// sort by Pair(reverse):
// c 2
// a 1
