//go:build go1.18
// +build go1.18

// Package avl provides an implementation of an AVL tree. An AVL tree is a
// self-balancing binary search tree. It stores key-value pairs that are sorted
// based on the key, and maintains that the tree is always balanced, ensuring
// logarithmic-time for all operations.
package avl

import (
	g "github.com/fufuok/utils/generic"
)

// Tree implements an AVL tree.
type Tree[K, V any] struct {
	root *node[K, V]
	less g.LessFn[K]
}

// New returns an empty AVL tree.
func New[K, V any](less g.LessFn[K]) *Tree[K, V] {
	return &Tree[K, V]{
		less: less,
	}
}

// Put associates 'key' with 'value'.
func (t *Tree[K, V]) Put(key K, value V) {
	t.root = t.root.add(key, value, t.less)
}

// Remove removes the value associated with 'key'.
func (t *Tree[K, V]) Remove(key K) {
	t.root = t.root.remove(key, t.less)
}

// Get returns the value associated with 'key'.
func (t *Tree[K, V]) Get(key K) (V, bool) {
	n := t.root.search(key, t.less)
	if n == nil {
		var v V
		return v, false
	}
	return n.value, true
}

// Each calls 'fn' on every node in the tree in order
func (t *Tree[K, V]) Each(fn func(key K, val V)) {
	t.root.each(fn)
}

// Height returns the height of the tree.
func (t *Tree[K, V]) Height() int {
	return t.root.getHeight()
}

// Size returns the number of elements in the tree.
func (t *Tree[K, V]) Size() int {
	return t.root.size()
}

type node[K, V any] struct {
	key   K
	value V

	height int
	left   *node[K, V]
	right  *node[K, V]
}

func (n *node[K, V]) add(key K, value V, less g.LessFn[K]) *node[K, V] {
	if n == nil {
		return &node[K, V]{
			key:    key,
			value:  value,
			height: 1,
			left:   nil,
			right:  nil,
		}
	}

	if g.Compare(key, n.key, less) < 0 {
		n.left = n.left.add(key, value, less)
	} else if g.Compare(key, n.key, less) > 0 {
		n.right = n.right.add(key, value, less)
	} else {
		n.value = value
	}
	return n.rebalanceTree()
}

func (n *node[K, V]) remove(key K, less g.LessFn[K]) *node[K, V] {
	if n == nil {
		return nil
	}
	if g.Compare(key, n.key, less) < 0 {
		n.left = n.left.remove(key, less)
	} else if g.Compare(key, n.key, less) > 0 {
		n.right = n.right.remove(key, less)
	} else {
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.value = rightMinNode.value
			n.right = n.right.remove(rightMinNode.key, less)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

func (n *node[K, V]) search(key K, less g.LessFn[K]) *node[K, V] {
	if n == nil {
		return nil
	}
	if g.Compare(key, n.key, less) < 0 {
		return n.left.search(key, less)
	} else if g.Compare(key, n.key, less) > 0 {
		return n.right.search(key, less)
	} else {
		return n
	}
}

func (n *node[K, V]) each(fn func(key K, val V)) {
	if n == nil {
		return
	}
	n.left.each(fn)
	fn(n.key, n.value)
	n.right.each(fn)
}

func (n *node[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node[K, V]) recalculateHeight() {
	n.height = 1 + g.Max(n.left.getHeight(), n.right.getHeight())
}

func (n *node[K, V]) rebalanceTree() *node[K, V] {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor <= -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor >= 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *node[K, V]) rotateLeft() *node[K, V] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node[K, V]) rotateRight() *node[K, V] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node[K, V]) findSmallest() *node[K, V] {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

func (n *node[K, V]) size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.size() + n.right.size()
}
