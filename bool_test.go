package utils

import (
	"encoding/json"
	"testing"
)

func TestBool_NewBool(t *testing.T) {
	v := new(Bool)
	AssertEqual(t, false, v.Load())
	v.Toggle()
	AssertEqual(t, true, v.Load())

	v = NewBool(false)
	AssertEqual(t, false, v.Load())
	v = NewBool(true)
	AssertEqual(t, true, v.Load())

	v = NewTrue()
	AssertEqual(t, true, v.Load())
	v = NewFalse()
	AssertEqual(t, false, v.Load())
}

func TestBool_Store(t *testing.T) {
	var v Bool
	v.Store(true)
	AssertEqual(t, true, v.Load())
	v.StoreFalse()
	AssertEqual(t, false, v.Load())
	v.StoreTrue()
	AssertEqual(t, true, v.Load())
}

func TestBool_Swap(t *testing.T) {
	var v Bool
	AssertEqual(t, v.Swap(false), false)
	AssertEqual(t, false, v.Load())
	AssertEqual(t, v.Swap(true), false)
	AssertEqual(t, true, v.Load())
	AssertEqual(t, v.Swap(false), true)
	AssertEqual(t, false, v.Load())
}

func TestBool_CAS(t *testing.T) {
	tests := []struct {
		init   bool
		oldV   bool
		newV   bool
		result bool
		store  bool
	}{
		{false, true, false, false, false},
		{false, false, true, true, true},
		{false, true, true, false, false},
		{false, false, false, true, false},
		{true, true, false, true, false},
		{true, false, true, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
	}
	for _, d := range tests {
		v := NewBool(d.init)
		AssertEqual(t, d.result, v.CAS(d.oldV, d.newV), "CAS")
		AssertEqual(t, d.store, v.Load(), "Store")
	}
}

func TestBool_Toggle(t *testing.T) {
	var v Bool
	AssertEqual(t, false, v.Toggle())
	AssertEqual(t, true, v.Toggle())
	AssertEqual(t, false, v.Toggle())
	AssertEqual(t, true, v.Load())
}

func TestBool_Marshal(t *testing.T) {
	bs := []byte("true")
	var v Bool
	err := json.Unmarshal(bs, &v)
	AssertEqual(t, nil, err)
	AssertEqual(t, true, v.Load())

	bs, err = json.Marshal(v.Load())
	AssertEqual(t, nil, err)
	AssertEqual(t, []byte("true"), bs)

	AssertEqual(t, "true", v.String())
}

func BenchmarkBool(b *testing.B) {
	var v Bool
	b.Run("Store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v.StoreTrue()
		}
	})
	b.Run("Load", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v.Load()
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkBool
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkBool/Store-4           208978531                5.768 ns/op           0 B/op          0 allocs/op
// BenchmarkBool/Store-4           210560200                5.690 ns/op           0 B/op          0 allocs/op
// BenchmarkBool/Store-4           209799655                5.733 ns/op           0 B/op          0 allocs/op
// BenchmarkBool/Load-4            1000000000               0.3181 ns/op          0 B/op          0 allocs/op
// BenchmarkBool/Load-4            1000000000               0.3209 ns/op          0 B/op          0 allocs/op
// BenchmarkBool/Load-4            1000000000               0.3225 ns/op          0 B/op          0 allocs/op
