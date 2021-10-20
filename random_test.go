package utils

import (
	"bytes"
	"math"
	"math/rand"
	"sync"
	"testing"

)

func TestRandInt(t *testing.T) {
	t.Parallel()
	AssertEqual(t, true, RandInt(1, 2) == 1)
	AssertEqual(t, true, RandInt(-1, 0) == -1)
	AssertEqual(t, true, RandInt(0, 5) >= 0)
	AssertEqual(t, true, RandInt(0, 5) < 5)
	AssertEqual(t, 2, RandInt(2, 2))
	AssertEqual(t, 2, RandInt(3, 2))
}

func TestRandUint32(t *testing.T) {
	t.Parallel()
	AssertEqual(t, true, RandUint32(1, 2) == 1)
	AssertEqual(t, true, RandUint32(0, 5) < 5)
	AssertEqual(t, uint32(2), RandUint32(2, 2))
	AssertEqual(t, uint32(2), RandUint32(3, 2))
}

func TestFastIntn(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		AssertEqual(t, true, FastRandn(uint32(i)) < uint32(i))
		AssertEqual(t, true, FastIntn(i) < i)
	}
	AssertEqual(t, 0, FastIntn(-2))
	AssertEqual(t, 0, FastIntn(0))
	AssertEqual(t, true, FastIntn(math.MaxUint32) < math.MaxUint32)
	AssertEqual(t, true, FastIntn(math.MaxInt64) < math.MaxInt64)
}

func BenchmarkRandInt(b *testing.B) {
	b.Run("RandInt", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = RandInt(0, i)
		}
	})
	b.Run("RandUint32", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = RandUint32(0, uint32(i))
		}
	})
	b.Run("FastIntn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = FastIntn(i)
		}
	})
	b.Run("Rand.Intn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = Rand.Intn(i)
		}
	})
	b.Run("std.rand.Intn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = rand.Intn(i)
		}
	})
}

func BenchmarkRandIntParallel(b *testing.B) {
	b.Run("RandInt", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandInt(0, math.MaxInt32)
			}
		})
	})
	b.Run("RandUint32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandUint32(0, math.MaxInt32)
			}
		})
	})
	b.Run("FastIntn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = FastIntn(math.MaxInt32)
			}
		})
	})
	b.Run("Rand.Intn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Rand.Intn(math.MaxInt32)
			}
		})
	})
	var mu sync.Mutex
	b.Run("std.rand.Intn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				_ = rand.Intn(math.MaxInt32)
				mu.Unlock()
			}
		})
	})
}

func TestRandString(t *testing.T) {
	t.Parallel()
	a, b := RandString(777), RandString(777)
	AssertEqual(t, 777, len(a))
	AssertEqual(t, false, a == b)
	AssertEqual(t, "", RandString(-1))
}

func TestRandBytes(t *testing.T) {
	t.Parallel()
	a, b := RandBytes(777), RandBytes(777)
	AssertEqual(t, 777, len(a))
	AssertEqual(t, 777, len(b))
	AssertEqual(t, false, bytes.Equal(a, b))
}

func TestFastRandBytes(t *testing.T) {
	t.Parallel()
	a, b := FastRandBytes(777), FastRandBytes(777)
	AssertEqual(t, 777, len(a))
	AssertEqual(t, 777, len(b))
	AssertEqual(t, false, bytes.Equal(a, b))
}

func TestRandHex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, 32, len(RandHex(16)))
	AssertEqual(t, 14, len(RandHex(7)))
}

func BenchmarkRandBytes(b *testing.B) {
	b.Run("RandString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = RandString(20)
		}
	})
	b.Run("RandBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = RandBytes(20)
		}
	})
	b.Run("FastRandBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FastRandBytes(20)
		}
	})
}

func BenchmarkRandBytesParallel(b *testing.B) {
	b.Run("RandString", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandString(20)
			}
		})
	})
	b.Run("RandBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandBytes(20)
			}
		})
	})
	b.Run("FastRandBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = FastRandBytes(20)
			}
		})
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/random
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkRandInt/RandInt-4                     300554511                3.982 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/RandInt-4                     299879216                4.007 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/RandUint32-4                  266118160                4.466 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/RandUint32-4                  266337582                4.507 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/FastIntn-4                    313761758                3.834 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/FastIntn-4                    312968804                3.811 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/Rand.Intn-4                    35001715                34.30 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/Rand.Intn-4                    34904052                34.54 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/std.rand.Intn-4                56418733                21.57 ns/op            0 B/op          0 allocs/op
// BenchmarkRandInt/std.rand.Intn-4                56331698                21.41 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/RandInt-4            1000000000                1.060 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/RandInt-4            1000000000                1.045 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/RandUint32-4          990860647                1.197 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/RandUint32-4         1000000000                1.182 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/FastIntn-4           1000000000                1.060 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/FastIntn-4           1000000000                1.055 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/Rand.Intn-4           130758892                9.132 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/Rand.Intn-4           130173494                9.065 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/std.rand.Intn-4        13878208                88.73 ns/op            0 B/op          0 allocs/op
// BenchmarkRandIntParallel/std.rand.Intn-4        13828624                89.97 ns/op            0 B/op          0 allocs/op

// BenchmarkRandBytes/RandString-4                 18142927                64.27 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytes/RandString-4                 18944168                64.43 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytes/RandBytes-4                   1730853                694.3 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytes/RandBytes-4                   1719566                687.4 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytes/FastRandBytes-4              18185881                64.99 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytes/FastRandBytes-4              18052567                65.73 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandString-4         63093751                19.00 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandString-4         67928510                19.29 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandBytes-4           1309642                916.2 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandBytes-4           1315711                916.6 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/FastRandBytes-4      64837544                20.05 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/FastRandBytes-4      65973478                19.52 ns/op           24 B/op          1 allocs/op
