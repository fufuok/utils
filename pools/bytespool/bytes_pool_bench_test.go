package bytespool

import (
	"testing"
)

func BenchmarkCapacityPools(b *testing.B) {
	b.Run("New", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bs := New(1024)
			Release(bs)
		}
	})
	b.Run("Make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bs := Make(1024)
			Release(bs)
		}
	})
	b.Run("New.Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				bs := New(1024)
				Release(bs)
			}
		})
	})
	b.Run("Make.Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				bs := Make(1024)
				Release(bs)
			}
		})
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils/pools/bytespool
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkCapacityPools/New-4            60133353                19.34 ns/op            0 B/op          0 allocs/op
// BenchmarkCapacityPools/New-4            61206177                19.49 ns/op            0 B/op          0 allocs/op
// BenchmarkCapacityPools/Make-4           61580971                19.58 ns/op            0 B/op          0 allocs/op
// BenchmarkCapacityPools/Make-4           61389439                19.71 ns/op            0 B/op          0 allocs/op
// BenchmarkCapacityPools/New.Parallel-4           240337632                5.041 ns/op           0 B/op          0 allocs/op
// BenchmarkCapacityPools/New.Parallel-4           235125742                5.133 ns/op           0 B/op          0 allocs/op
// BenchmarkCapacityPools/Make.Parallel-4          229302106                5.073 ns/op           0 B/op          0 allocs/op
// BenchmarkCapacityPools/Make.Parallel-4          238298523                5.308 ns/op           0 B/op          0 allocs/op
