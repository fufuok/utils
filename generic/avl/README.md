<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# avl

```go
import "github.com/fufuok/utils/generic/avl"
```

Package avl provides an implementation of an AVL tree\. An AVL tree is a self\-balancing binary search tree\. It stores key\-value pairs that are sorted based on the key\, and maintains that the tree is always balanced\, ensuring logarithmic\-time for all operations\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	g "github.com/fufuok/utils/generic"
	"github.com/fufuok/utils/generic/avl"
)

func main() {
	tree := avl.New[int, string](g.Less[int])

	tree.Put(42, "foo")
	tree.Put(-10, "bar")
	tree.Put(0, "baz")
	tree.Put(10, "quux")
	tree.Remove(10)

	tree.Each(func(key int, val string) {
		fmt.Println(key, val)
	})

}
```

#### Output

```
-10 bar
0 baz
42 foo
```

</p>
</details>

## Index

- [type Tree](<#type-tree>)
  - [func New[K, V any](less g.LessFn[K]) *Tree[K, V]](<#func-new>)
  - [func (t *Tree[K, V]) Each(fn func(key K, val V))](<#func-treek-v-each>)
  - [func (t *Tree[K, V]) Get(key K) (V, bool)](<#func-treek-v-get>)
  - [func (t *Tree[K, V]) Height() int](<#func-treek-v-height>)
  - [func (t *Tree[K, V]) Put(key K, value V)](<#func-treek-v-put>)
  - [func (t *Tree[K, V]) Remove(key K)](<#func-treek-v-remove>)
  - [func (t *Tree[K, V]) Size() int](<#func-treek-v-size>)


## type [Tree](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L12-L15>)

Tree implements an AVL tree\.

```go
type Tree[K, V any] struct {
    // contains filtered or unexported fields
}
```

### func [New](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L18>)

```go
func New[K, V any](less g.LessFn[K]) *Tree[K, V]
```

New returns an empty AVL tree\.

### func \(\*Tree\[K\, V\]\) [Each](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L45>)

```go
func (t *Tree[K, V]) Each(fn func(key K, val V))
```

Each calls 'fn' on every node in the tree in order

### func \(\*Tree\[K\, V\]\) [Get](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L35>)

```go
func (t *Tree[K, V]) Get(key K) (V, bool)
```

Get returns the value associated with 'key'\.

### func \(\*Tree\[K\, V\]\) [Height](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L50>)

```go
func (t *Tree[K, V]) Height() int
```

Height returns the height of the tree\.

### func \(\*Tree\[K\, V\]\) [Put](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L25>)

```go
func (t *Tree[K, V]) Put(key K, value V)
```

Put associates 'key' with 'value'\.

### func \(\*Tree\[K\, V\]\) [Remove](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L30>)

```go
func (t *Tree[K, V]) Remove(key K)
```

Remove removes the value associated with 'key'\.

### func \(\*Tree\[K\, V\]\) [Size](<https://gitee.com/fufuok/utils/blob/master/generic/avl/avl.go#L55>)

```go
func (t *Tree[K, V]) Size() int
```

Size returns the number of elements in the tree\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)