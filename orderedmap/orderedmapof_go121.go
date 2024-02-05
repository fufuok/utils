//go:build go1.21
// +build go1.21

package orderedmap

import (
	"maps"
	"slices"
)

func (o *OrderedMapOf[V]) Clone() *OrderedMapOf[V] {
	return &OrderedMapOf[V]{
		keys:   slices.Clone(o.keys),
		values: maps.Clone(o.values),
	}
}

func (o *OrderedMapOf[V]) ToMap() map[string]V {
	return maps.Clone(o.values)
}
