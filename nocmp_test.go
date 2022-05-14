// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package utils

import (
	"reflect"
	"testing"
)

func TestNoCmpComparability(t *testing.T) {
	tests := []struct {
		desc       string
		give       interface{}
		comparable bool
	}{
		{
			desc: "NoCmp struct",
			give: NoCmp{},
		},
		{
			desc: "struct with NoCmp embedded",
			give: struct{ NoCmp }{},
		},
		{
			desc:       "pointer to struct with NoCmp embedded",
			give:       &struct{ NoCmp }{},
			comparable: true,
		},

		// All exported types must be uncomparable.
		{desc: "Bool", give: Bool{}},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			typ := reflect.TypeOf(tt.give)
			AssertEqualf(t, tt.comparable, typ.Comparable(),
				"type %v comparablity mismatch", typ)
		})
	}
}

// NoCmp must not add to the size of a struct in-memory.
func TestNoCmpSize(t *testing.T) {
	type x struct{ _ int }

	before := reflect.TypeOf(x{}).Size()

	type y struct {
		_ NoCmp
		_ x
	}

	after := reflect.TypeOf(y{}).Size()

	AssertEqual(t, before, after,
		"expected NoCmp to have no effect on struct size")
}

// This test will fail to compile if we disallow copying of NoCmp.
//
// We need to allow this so that users can do,
//
//   var x atomic.Int32
//   x = atomic.NewInt32(1)
func TestNoCmpCopy(t *testing.T) {
	type foo struct{ _ NoCmp }

	t.Run("struct copy", func(t *testing.T) {
		a := foo{}
		b := a
		_ = b // unused
	})

	t.Run("pointer copy", func(t *testing.T) {
		a := &foo{}
		b := *a
		_ = b // unused
	})
}
