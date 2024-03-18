//go:build go1.18
// +build go1.18

package orderedmap

import (
	"reflect"
	"sort"
	"testing"
)

func TestOrderedMapOf(t *testing.T) {
	o := NewOf[string, any]()
	// number
	o.Set("number", 3)
	v, _ := o.Get("number")
	if v.(int) != 3 {
		t.Error("Set number")
	}
	// string
	o.Set("string", "x")
	v, _ = o.Get("string")
	if v.(string) != "x" {
		t.Error("Set string")
	}
	// string slice
	o.Set("strings", []string{
		"t",
		"u",
	})
	v, _ = o.Get("strings")
	if v.([]string)[0] != "t" {
		t.Error("Set strings first index")
	}
	if v.([]string)[1] != "u" {
		t.Error("Set strings second index")
	}
	// mixed slice
	o.Set("mixed", []interface{}{
		1,
		"1",
	})
	v, _ = o.Get("mixed")
	if v.([]interface{})[0].(int) != 1 {
		t.Error("Set mixed int")
	}
	if v.([]interface{})[1].(string) != "1" {
		t.Error("Set mixed string")
	}
	// overriding existing key
	o.Set("number", 4)
	v, _ = o.Get("number")
	if v.(int) != 4 {
		t.Error("Override existing key")
	}
	// Keys method
	keys := o.Keys()
	expectedKeys := []string{
		"number",
		"string",
		"strings",
		"mixed",
	}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Error("Keys method", key, "!=", expectedKeys[i])
		}
	}
	// Values method
	values := o.Values()
	expectedValues := map[string]interface{}{
		"number":  4,
		"string":  "x",
		"strings": []string{"t", "u"},
		"mixed":   []interface{}{1, "1"},
	}
	if !reflect.DeepEqual(values, expectedValues) {
		t.Error("Values method returned unexpected map")
	}
	// delete
	o.Delete("strings")
	o.Delete("not a key being used")
	if len(o.Keys()) != 3 {
		t.Error("Delete method")
	}
	_, ok := o.Get("strings")
	if ok {
		t.Error("Delete did not remove 'strings' key")
	}
}

func TestOrderedMapOf_SortKeys(t *testing.T) {
	o := NewOf[string, int]()
	o.Set("b", 2)
	o.Set("a", 1)
	o.Set("c", 3)

	o.SortKeys(sort.Strings)

	// Check the root keys
	expectedKeys := []string{
		"a",
		"b",
		"c",
	}
	k := o.Keys()
	for i := range k {
		if k[i] != expectedKeys[i] {
			t.Error("SortKeys root key order", i, k[i], "!=", expectedKeys[i])
		}
	}
}

func TestOrderedMapOf_Sort(t *testing.T) {
	o := NewOf[string, float64]()
	o.Set("b", 2.1)
	o.Set("a", 1.1)
	o.Set("c", 3)

	o.Sort(func(a *PairOf[string, float64], b *PairOf[string, float64]) bool {
		return a.value > b.value
	})

	// Check the root keys
	expectedKeys := []string{
		"c",
		"b",
		"a",
	}
	k := o.Keys()
	for i := range k {
		if k[i] != expectedKeys[i] {
			t.Error("Sort root key order", i, k[i], "!=", expectedKeys[i])
		}
	}
}
