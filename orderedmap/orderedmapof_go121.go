//go:build go1.21
// +build go1.21

package orderedmap

import (
	"maps"
	"slices"
)

func (o *OrderedMapOf[K, V]) Clone() *OrderedMapOf[K, V] {
	return &OrderedMapOf[K, V]{
		keys:   slices.Clone(o.keys),
		values: maps.Clone(o.values),
	}
}

func (o *OrderedMapOf[K, V]) ToMap() map[K]V {
	return maps.Clone(o.values)
}
