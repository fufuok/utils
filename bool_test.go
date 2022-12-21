package utils

import (
	"encoding/json"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestBool_NewBool(t *testing.T) {
	v := new(Bool)
	assert.Equal(t, false, v.Load())
	v.Toggle()
	assert.Equal(t, true, v.Load())

	v = NewBool(false)
	assert.Equal(t, false, v.Load())
	v = NewBool(true)
	assert.Equal(t, true, v.Load())

	v = NewTrue()
	assert.Equal(t, true, v.Load())
	v = NewFalse()
	assert.Equal(t, false, v.Load())
}

func TestBool_Store(t *testing.T) {
	var v Bool
	v.Store(true)
	assert.Equal(t, true, v.Load())
	v.StoreFalse()
	assert.Equal(t, false, v.Load())
	v.StoreTrue()
	assert.Equal(t, true, v.Load())
}

func TestBool_Swap(t *testing.T) {
	var v Bool
	assert.Equal(t, v.Swap(false), false)
	assert.Equal(t, false, v.Load())
	assert.Equal(t, v.Swap(true), false)
	assert.Equal(t, true, v.Load())
	assert.Equal(t, v.Swap(false), true)
	assert.Equal(t, false, v.Load())
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
		assert.Equal(t, d.result, v.CAS(d.oldV, d.newV), "CAS")
		assert.Equal(t, d.store, v.Load(), "Store")
	}
}

func TestBool_Toggle(t *testing.T) {
	var v Bool
	assert.Equal(t, false, v.Toggle())
	assert.Equal(t, true, v.Toggle())
	assert.Equal(t, false, v.Toggle())
	assert.Equal(t, true, v.Load())
}

func TestBool_Marshal(t *testing.T) {
	bs := []byte("true")
	var v Bool
	err := json.Unmarshal(bs, &v)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, v.Load())

	bs, err = json.Marshal(v.Load())
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("true"), bs)

	assert.Equal(t, "true", v.String())
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
