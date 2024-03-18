package orderedmap

import (
	"strconv"
	"testing"
)

var count = 1_000_000

func BenchmarkOrderedMap_Set(b *testing.B) {
	o := New()
	for i := 0; i < b.N; i++ {
		o.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkOrderedMap_Get(b *testing.B) {
	o := New()
	for i := 0; i < count; i++ {
		o.Set(strconv.Itoa(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o.Get(strconv.Itoa(i))
	}
}

func BenchmarkOrderedMap_Iterate(b *testing.B) {
	o := New()
	for i := 0; i < count; i++ {
		o.Set(strconv.Itoa(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range o.Keys() {
			o.Get(k)
		}
	}
}

// go test -run=^$ -benchmem -benchtime=1s -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils/orderedmap
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// BenchmarkOrderedMap_Set-16               3375973               414.8 ns/op           283 B/op          2 allocs/op
// BenchmarkOrderedMap_Get-16               9953631               124.6 ns/op             7 B/op          0 allocs/op
// BenchmarkOrderedMap_Iterate-16                21          56389484 ns/op               0 B/op          0 allocs/op
