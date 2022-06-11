//go:build go1.18
// +build go1.18

package btree_test

import (
	"fmt"
	"math/rand"
	"testing"

	g "github.com/fufuok/utils/generic"
	"github.com/fufuok/utils/generic/btree"
)

func checkeq[K any, V comparable](cm *btree.Tree[K, V], get func(k K) (V, bool), t *testing.T) {
	cm.Each(func(key K, val V) {
		if ov, ok := get(key); !ok {
			t.Fatalf("key %v should exist", key)
		} else if val != ov {
			t.Fatalf("value mismatch: %v != %v", val, ov)
		}
	})
}

func TestCrossCheck(t *testing.T) {
	stdm := make(map[int]int)
	tree := btree.New[int, int](g.Less[int])

	const nops = 1000

	for i := 0; i < nops; i++ {
		key := rand.Intn(100)
		val := rand.Int()
		op := rand.Intn(2)

		switch op {
		case 0:
			stdm[key] = val
			tree.Put(key, val)
		case 1:
			var del int
			for k := range stdm {
				del = k
				break
			}
			delete(stdm, del)
			tree.Remove(del)
		}

		checkeq(tree, func(k int) (int, bool) {
			v, ok := stdm[int(k)]
			return v, ok
		}, t)
	}
}

func Example() {
	tree := btree.New[int, string](g.Less[int])

	tree.Put(42, "foo")
	tree.Put(-10, "bar")
	tree.Put(0, "baz")

	tree.Each(func(key int, val string) {
		fmt.Println(key, val)
	})

	// Output:
	// -10 bar
	// 0 baz
	// 42 foo
}
