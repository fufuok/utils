//go:build go1.18
// +build go1.18

package orderedmap

import (
	"sort"
)

type PairOf[V any] struct {
	key   string
	value V
}

func (kv *PairOf[V]) Key() string {
	return kv.key
}

func (kv *PairOf[V]) Value() V {
	return kv.value
}

type ByPairOf[V any] struct {
	Pairs    []*PairOf[V]
	LessFunc func(a *PairOf[V], j *PairOf[V]) bool
}

func (a ByPairOf[V]) Len() int           { return len(a.Pairs) }
func (a ByPairOf[V]) Swap(i, j int)      { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }
func (a ByPairOf[V]) Less(i, j int) bool { return a.LessFunc(a.Pairs[i], a.Pairs[j]) }

type OrderedMapOf[V any] struct {
	keys   []string
	values map[string]V
}

func NewOf[V any]() *OrderedMapOf[V] {
	o := OrderedMapOf[V]{}
	o.keys = []string{}
	o.values = map[string]V{}
	return &o
}

func (o *OrderedMapOf[V]) Get(key string) (V, bool) {
	val, exists := o.values[key]
	return val, exists
}

func (o *OrderedMapOf[V]) MustGet(key string) V {
	val, _ := o.values[key]
	return val
}

func (o *OrderedMapOf[V]) Set(key string, value V) {
	_, exists := o.values[key]
	if !exists {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *OrderedMapOf[V]) Delete(key string) {
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

func (o *OrderedMapOf[V]) Keys() []string {
	return o.keys
}

func (o *OrderedMapOf[V]) Values() map[string]V {
	return o.values
}

// SortKeys Sort the map keys using your sort func
func (o *OrderedMapOf[V]) SortKeys(sortFunc ...func(keys []string)) {
	if len(sortFunc) > 0 {
		sortFunc[0](o.keys)
		return
	}
	sort.Strings(o.keys)
}

// Sort the map using your sort func
func (o *OrderedMapOf[V]) Sort(lessFunc func(a *PairOf[V], b *PairOf[V]) bool) {
	pairs := make([]*PairOf[V], len(o.keys))
	for i, key := range o.keys {
		pairs[i] = &PairOf[V]{key, o.values[key]}
	}

	sort.Sort(ByPairOf[V]{pairs, lessFunc})

	for i, pair := range pairs {
		o.keys[i] = pair.key
	}
}
