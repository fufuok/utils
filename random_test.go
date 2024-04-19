package utils

import (
	"math"
	"math/rand"
	"strings"
	"sync"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestRandInt(t *testing.T) {
	t.Parallel()
	assert.Equal(t, true, RandInt(1, 2) == 1)
	assert.Equal(t, true, RandInt(-1, 0) == -1)
	assert.Equal(t, true, RandInt(0, 5) >= 0)
	assert.Equal(t, true, RandInt(0, 5) < 5)
	assert.Equal(t, 2, RandInt(2, 2))
	assert.Equal(t, 2, RandInt(3, 2))
}

func TestRandUint32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, true, RandUint32(1, 2) == 1)
	assert.Equal(t, true, RandUint32(0, 5) < 5)
	assert.Equal(t, uint32(2), RandUint32(2, 2))
	assert.Equal(t, uint32(2), RandUint32(3, 2))
}

func TestFastIntn(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		assert.Equal(t, true, FastRandn(uint32(i)) < uint32(i))
		assert.Equal(t, true, FastIntn(i) < i)
	}
	assert.Equal(t, 0, FastIntn(-2))
	assert.Equal(t, 0, FastIntn(0))
	assert.Equal(t, true, FastIntn(math.MaxUint32) < math.MaxUint32)
	assert.Equal(t, true, FastIntn(math.MaxInt64) < math.MaxInt64)
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
	fns := []func(n int) string{RandString, RandAlphaString, RandHexString, RandDecString}
	ss := []string{letterBytes, alphaBytes, hexBytes, decBytes}
	for i, fn := range fns {
		a, b := fn(777), fn(777)
		assert.Equal(t, 777, len(a))
		assert.NotEqual(t, a, b)
		assert.Equal(t, "", fn(-1))
		for _, s := range ss[i] {
			assert.True(t, strings.ContainsRune(a, s))
		}
	}
}

func TestRandBytesLetters(t *testing.T) {
	t.Parallel()
	letters := ""
	assert.Nil(t, RandBytesLetters(10, letters))
	letters = "a"
	assert.Nil(t, RandBytesLetters(10, letters))
	letters = "ab"
	s := B2S(RandBytesLetters(10, letters))
	assert.Equal(t, 10, len(s))
	assert.True(t, strings.Contains(s, "a"))
	assert.True(t, strings.Contains(s, "b"))
	letters = "xxxxxxxxxxxx"
	s = B2S(RandBytesLetters(100, letters))
	assert.Equal(t, 100, len(s))
	assert.Equal(t, strings.Repeat("x", 100), s)
}

func BenchmarkRandBytesParallel(b *testing.B) {
	b.Run("RandBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandBytes(20)
			}
		})
	})
	b.Run("RandAlphaBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandAlphaBytes(20)
			}
		})
	})
	b.Run("RandHexBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandHexBytes(20)
			}
		})
	})
	b.Run("RandDecBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandDecBytes(20)
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

// go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=BenchmarkRandBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// BenchmarkRandBytesParallel/RandBytes-16                 95409290                12.65 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandBytes-16                 90086031                12.75 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandAlphaBytes-16            79601335                14.97 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandAlphaBytes-16            76708616                14.81 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandHexBytes-16              39585378                28.88 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandHexBytes-16              43593310                29.04 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandDecBytes-16              32723065                36.39 ns/op           24 B/op          1 allocs/op
// BenchmarkRandBytesParallel/RandDecBytes-16              33422029                36.33 ns/op           24 B/op          1 allocs/op
