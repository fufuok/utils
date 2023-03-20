//go:build go1.18
// +build go1.18

// Copyright 2022 Changkun Ou <changkun.de>. All rights reserved.
// Use of this source code is governed by a GPLv3 license that
// can be found in the LICENSE file.

package deepcopy_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fufuok/utils/generic/deepcopy"
	"github.com/fufuok/utils/generic/maps"
	"github.com/fufuok/utils/generic/slices"
)

func ExampleValue() {
	tests := []any{
		`"Now cut that out!"`,
		39,
		true,
		false,
		2.14,
		[]string{
			"Phil Harris",
			"Rochester van Jones",
			"Mary Livingstone",
			"Dennis Day",
		},
		[2]string{
			"Jell-O",
			"Grape-Nuts",
		},
	}

	for _, expected := range tests {
		actual := deepcopy.Value(expected)
		fmt.Println(actual)
	}
	// Output:
	// "Now cut that out!"
	// 39
	// true
	// false
	// 2.14
	// [Phil Harris Rochester van Jones Mary Livingstone Dennis Day]
	// [Jell-O Grape-Nuts]
}

type Foo struct {
	Foo *Foo
	Bar int
}

func ExampleMap() {
	x := map[string]*Foo{
		"foo": {Bar: 1},
		"bar": {Bar: 2},
	}
	y := deepcopy.Value(x)
	for _, k := range []string{"foo", "bar"} { // to ensure consistent order
		fmt.Printf("x[\"%v\"] = y[\"%v\"]: %v\n", k, k, x[k] == y[k])
		fmt.Printf("x[\"%v\"].Foo = y[\"%v\"].Foo: %v\n", k, k, x[k].Foo == y[k].Foo)
		fmt.Printf("x[\"%v\"].Bar = y[\"%v\"].Bar: %v\n", k, k, x[k].Bar == y[k].Bar)
	}
	// Output:
	// x["foo"] = y["foo"]: false
	// x["foo"].Foo = y["foo"].Foo: false
	// x["foo"].Bar = y["foo"].Bar: true
	// x["bar"] = y["bar"]: false
	// x["bar"].Foo = y["bar"].Foo: false
	// x["bar"].Bar = y["bar"].Bar: true
}

func TestInterface(t *testing.T) {
	x := []any{nil}
	y := deepcopy.Value(x)
	if !reflect.DeepEqual(x, y) || len(y) != 1 {
		t.Errorf("expect %v == %v; y had length %v (expected 1)", x, y, len(y))
	}
	var a any
	b := deepcopy.Value(a)
	if a != b {
		t.Errorf("expected %v == %v", a, b)
	}
}

func TestAvoidInfiniteLoops(t *testing.T) {
	x := &Foo{
		Bar: 4,
	}
	x.Foo = x
	y := deepcopy.Value(x)
	if x == y {
		t.Fatalf("expect x==y returns false")
	}
	if x != x.Foo {
		t.Fatalf("expect x==x.Foo returns true")
	}
	if y != y.Foo {
		t.Fatalf("expect y==y.Foo returns true")
	}
}

func TestSpecialKind(t *testing.T) {
	x := func() int { return 42 }
	y := deepcopy.Value(x)
	if y() != 42 || reflect.DeepEqual(x, y) {
		t.Fatalf("copied value is equal")
	}

	a := map[bool]any{true: x}
	b := deepcopy.Value(a)
	if reflect.DeepEqual(a, b) {
		t.Fatalf("copied value is equal")
	}

	c := []any{x}
	d := deepcopy.Value(c)
	if reflect.DeepEqual(c, d) {
		t.Fatalf("copied value is equal")
	}
}

type Test struct {
	t time.Time
}

// TODO: fix
// func Test_Time(t *testing.T) {
// 	t1 := Test{time.Now()}
// 	t2 := deepcopy.Value(t1)
// 	if t1.t.String() != t2.t.String() {
// 		t.Logf("t1: %s, t2: %s", t1.t.Format(time.RFC3339), t2.t.Format(time.RFC3339))
// 		t.Logf("t1: %v, t2: %v", t1, t2)
// 		t.Fatalf("t1: %v, t2: %v", t1.t.Location(), t2.t.Location())
// 	}
// }

func Example() {
	type tFN struct {
		fn func() int
	}
	type tMap map[bool][]tFN
	m1 := tMap{true: []tFN{{func() int { return 1 }}}}

	// 深拷贝
	m2 := deepcopy.Value(m1)
	// equal: false 1
	fmt.Printf("equal: %v %+v\n", reflect.DeepEqual(m1, m2), m2[true][0].fn())

	// 浅拷贝
	m3 := maps.Clone[tMap](m1)
	// equal: true 1
	fmt.Printf("equal: %v %+v\n", reflect.DeepEqual(m1, m3), m3[true][0].fn())

	// 改变 m1 后
	m1[true][0] = tFN{func() int { return 0 }}
	// m1: 0, m2: 1, m3: 0
	fmt.Printf("m1: %d, m2: %d, m3: %d\n", m1[true][0].fn(), m2[true][0].fn(), m3[true][0].fn())

	s1 := []tMap{m1}
	s2 := deepcopy.Value(s1)
	s3 := slices.Clone(s1)
	m1[true][0] = tFN{func() int { return 2 }}
	// s1: 2, s2: 0, s3: 2
	fmt.Printf("s1: %d, s2: %d, s3: %d\n", s1[0][true][0].fn(), s2[0][true][0].fn(), s3[0][true][0].fn())

	// Output:
	// equal: false 1
	// equal: true 1
	// m1: 0, m2: 1, m3: 0
	// s1: 2, s2: 0, s3: 2
}
