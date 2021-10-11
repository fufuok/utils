package utils

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	AssertEqual(t, true, RandInt(1, 2) == 1)
	AssertEqual(t, true, RandInt(-1, 0) == -1)
	AssertEqual(t, true, RandInt(0, 5) >= 0)
	AssertEqual(t, true, RandInt(0, 5) < 5)
	AssertEqual(t, 0, RandInt(2, 2))
	AssertEqual(t, 0, RandInt(3, 2))
}

func TestRandString(t *testing.T) {
	a, b := RandString(77), RandString(77)
	AssertEqual(t, 77, len(a))
	AssertEqual(t, false, a == b)
}

func TestRandBytes(t *testing.T) {
	a, b := RandBytes(77), RandBytes(77)
	AssertEqual(t, 77, len(a))
	AssertEqual(t, false, MemHashb(a) == MemHashb(b))
}

func TestRandHex(t *testing.T) {
	AssertEqual(t, true, len(RandHex(16)) == 32)
}

func BenchmarkRandBytes(b *testing.B) {
	b.ReportAllocs()
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

// go test -run=^$ -bench=BenchmarkRandBytes -benchtime=1s -count=3
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkRandBytes/RandBytes-8            343431              3733 ns/op
// BenchmarkRandBytes/RandBytes-8            374635              3312 ns/op
// BenchmarkRandBytes/RandBytes-8            348879              3745 ns/op
// BenchmarkRandBytes/FastRandBytes-8      13103540               118.8 ns/op
// BenchmarkRandBytes/FastRandBytes-8      13080572                98.85 ns/op
// BenchmarkRandBytes/FastRandBytes-8      11399919               100.9 ns/op
