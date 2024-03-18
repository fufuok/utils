//go:build go1.18
// +build go1.18

package orderedmap

import (
	"sort"
)

type PairOf[K comparable, V any] struct {
	key   K
	value V
}

func (kv *PairOf[K, V]) Key() K {
	return kv.key
}

func (kv *PairOf[K, V]) Value() V {
	return kv.value
}

type ByPairOf[K comparable, V any] struct {
	Pairs    []*PairOf[K, V]
	LessFunc func(a *PairOf[K, V], j *PairOf[K, V]) bool
}

func (a ByPairOf[K, V]) Len() int           { return len(a.Pairs) }
func (a ByPairOf[K, V]) Swap(i, j int)      { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }
func (a ByPairOf[K, V]) Less(i, j int) bool { return a.LessFunc(a.Pairs[i], a.Pairs[j]) }

type OrderedMapOf[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

func NewOf[K comparable, V any]() *OrderedMapOf[K, V] {
	o := OrderedMapOf[K, V]{}
	o.keys = []K{}
	o.values = map[K]V{}
	return &o
}

func (o *OrderedMapOf[K, V]) Get(key K) (V, bool) {
	val, exists := o.values[key]
	return val, exists
}

func (o *OrderedMapOf[K, V]) MustGet(key K) V {
	val, _ := o.values[key]
	return val
}

func (o *OrderedMapOf[K, V]) Set(key K, value V) {
	_, exists := o.values[key]
	if !exists {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *OrderedMapOf[K, V]) Delete(key K) {
	// check key is in use
	_, ok := o.values[key]
	if !ok {
		return
	}
	// remove from keys
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			break
		}
	}
	// remove from values
	delete(o.values, key)
}

func (o *OrderedMapOf[K, V]) Keys() []K {
	return o.keys
}

func (o *OrderedMapOf[K, V]) Values() map[K]V {
	return o.values
}

// SortKeys Sort the map keys using your sort func
func (o *OrderedMapOf[K, V]) SortKeys(sortFunc func(keys []K)) {
	sortFunc(o.keys)
}

// Sort the map using your sort func
func (o *OrderedMapOf[K, V]) Sort(lessFunc func(a *PairOf[K, V], b *PairOf[K, V]) bool) {
	pairs := make([]*PairOf[K, V], len(o.keys))
	for i, key := range o.keys {
		pairs[i] = &PairOf[K, V]{key, o.values[key]}
	}

	sort.Sort(ByPairOf[K, V]{pairs, lessFunc})

	for i, pair := range pairs {
		o.keys[i] = pair.key
	}
}
