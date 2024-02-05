//go:build go1.21
// +build go1.21

package orderedmap

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestOrderedMapOf_Clone(t *testing.T) {
	o := NewOf[int]()
	o.Set("b", 2)
	o.Set("a", 1)
	o.Set("c", 3)

	c := o.Clone()

	o.Set("a", 42)
	c.Set("a", 43)

	assert.Equal(t, 42, o.MustGet("a"))
	assert.Equal(t, 43, c.MustGet("a"))
}

func TestOrderedMapOf_ToMap(t *testing.T) {
	o := NewOf[int]()
	o.Set("b", 2)
	o.Set("a", 1)
	o.Set("c", 3)

	m := o.ToMap()

	o.Set("a", 42)
	m["a"] = 43

	assert.Equal(t, 42, o.MustGet("a"))
	assert.Equal(t, 43, m["a"])
}
