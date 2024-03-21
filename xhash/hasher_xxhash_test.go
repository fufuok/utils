//go:build go1.18
// +build go1.18

package xhash

import (
	"strconv"
	"strings"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestGenHasher64(t *testing.T) {
	hasher0 := GenHasher64[string]()
	var h0 string
	h0 = "ff"
	assert.Equal(t, hasher0(h0), hasher0(h0))

	hasher1 := GenHasher64[int]()
	var h1 int
	h1 = 123
	assert.Equal(t, hasher1(h1), hasher1(h1))

	type L1 int
	type L2 L1
	hasher2 := GenHasher64[L2]()
	var h2 L2
	h2 = 123
	assert.Equal(t, hasher2(h2), hasher2(h2))
	assert.Equal(t, hasher1(h1), hasher2(h2))

	type foo struct {
		x int
		y int
	}
	hasher3 := GenHasher64[*foo]()
	h3 := new(foo)
	h31 := h3
	assert.Equal(t, hasher3(h3), hasher3(h31))

	hasher4 := GenHasher[float64]()
	assert.Equal(t, hasher4(3.1415926), hasher4(3.1415926))
	assert.NotEqual(t, hasher4(3.1415926), hasher4(3.1415927))

	hasher5 := GenHasher[complex128]()
	assert.Equal(t, hasher5(complex(3, 5)), hasher5(complex(3, 5)))
	assert.NotEqual(t, hasher5(complex(4, 5)), hasher5(complex(3, 5)))

	hasher6 := GenHasher[byte]()
	assert.Equal(t, hasher6('\n'), hasher6(10))
	assert.NotEqual(t, hasher6('\r'), hasher6('\n'))

	hasher7 := GenHasher[uintptr]()
	assert.Equal(t, hasher7(8), hasher7(8))
	assert.NotEqual(t, hasher7(7), hasher7(8))
}

func TestCollision_GenHasher64(t *testing.T) {
	n := 20_000_000
	sHasher := GenHasher64[string]()
	iHasher := GenHasher64[int]()
	ms := make(map[uint64]string)
	mi := make(map[uint64]int)
	for i := 0; i < n; i++ {
		s := strconv.Itoa(i)
		hs := sHasher(s)
		hi := iHasher(i)
		if v, ok := ms[hs]; ok {
			t.Fatalf("Collision: %s(%d) == %s(%d)", v, sHasher(v), s, sHasher(s))
		}
		if v, ok := mi[hi]; ok {
			t.Fatalf("Collision: %d(%d) == %d(%d)", v, iHasher(v), i, iHasher(i))
		}
		ms[hs] = s
		mi[hi] = i
	}

	hi := iHasher(7)
	if _, ok := mi[hi]; !ok {
		t.Fatalf("The number 7 should exist")
	}
	if len(ms) != len(mi) || len(ms) != n {
		t.Fatalf("Hash count: %d, %d, %d", len(ms), len(mi), n)
	}
}

func BenchmarkHasher_GenHasher64(b *testing.B) {
	sHasher := GenHasher64[string]()
	iHasher := GenHasher64[int]()
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = sHasher(strconv.Itoa(i))
		}
	})
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = iHasher(i)
		}
	})
}

func BenchmarkHasher_Parallel_GenHasher64(b *testing.B) {
	sHasher := GenHasher64[string]()
	s := strings.Repeat(testString, 10)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = sHasher(s)
		}
	})
}
