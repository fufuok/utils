package main

import (
	"encoding/json"
	"fmt"

	"github.com/fufuok/utils/orderedmap"
)

func main() {
	// use New() instead of o := map[string]interface{}{}
	o := orderedmap.New()

	// use SetEscapeHTML() to whether escape problematic HTML characters or not, defaults is true
	o.SetEscapeHTML(false)

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
	o.Delete("a")

	// serialize to a json string using encoding/json
	bytes, _ := json.Marshal(o)
	fmt.Println("JSON:", string(bytes))
	prettyBytes, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println("JSON pretty:", string(prettyBytes))

	// deserialize a json string using encoding/json
	// all maps (including nested maps) will be parsed as orderedmaps
	s := `{"a": 1, "c": 1.23, "d": 1.23, "b": -1, "e": 7}`
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sort the keys:")
	o.SortKeys()
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("sort by Pair:")
	o.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(float64) < b.Value().(float64)
	})
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("sort by Pair(reverse):")
	o.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(float64) > b.Value().(float64)
	})
	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("size:", o.Size())
}

// Output:
// Get: 1 true
// a 1
// c 2
// b \.<>[]{}_-
// JSON: {"c":2,"b":"\\.\u003c\u003e[]{}_-"}
// JSON pretty: {
//  "c": 2,
//  "b": "\\.\u003c\u003e[]{}_-"
// }
// sort the keys:
// a 1
// b -1
// c 1.23
// d 1.23
// e 7
// sort by Pair:
// b -1
// a 1
// c 1.23
// d 1.23
// e 7
// sort by Pair(reverse):
// e 7
// c 1.23
// d 1.23
// a 1
// b -1
// size: 5
