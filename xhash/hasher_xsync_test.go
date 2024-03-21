//go:build go1.18
// +build go1.18

package xhash

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMakeHashFunc(t *testing.T) {
	type User struct {
		Name string
		City string
	}

	sHasher := MakeHasher[string]()
	iHasher := MakeHasher[User]()

	// Not that much to test TBH.
	// check that hash is not always the same
	for i := 0; ; i++ {
		if sHasher("foo") != sHasher("bar") {
			break
		}
		if i >= 100 {
			t.Error("sHasher is always the same")
			break
		}
	}

	if sHasher("foo") != sHasher("foo") {
		t.Error("sHasher is not deterministic")
	}

	if iHasher(User{Name: "Ivan", City: "Sofia"}) != iHasher(User{Name: "Ivan", City: "Sofia"}) {
		t.Error("iHasher is not deterministic")
	}
}

func BenchmarkMakeHashFunc(b *testing.B) {
	type Point struct {
		X, Y, Z int
	}

	type User struct {
		ID        int
		FirstName string
		LastName  string
		IsActive  bool
		City      string
	}

	type PadInside struct {
		A int
		B byte
		C int
	}

	type PadTrailing struct {
		A int
		B byte
	}

	doBenchmarkMakeHashFunc(b, int64(116))
	doBenchmarkMakeHashFunc(b, int32(116))
	doBenchmarkMakeHashFunc(b, 3.14)
	doBenchmarkMakeHashFunc(b, "test key test key test key test key test key test key test key test key test key ")
	doBenchmarkMakeHashFunc(b, Point{1, 2, 3})
	doBenchmarkMakeHashFunc(b, User{ID: 1, FirstName: "Ivan", LastName: "Ivanov", IsActive: true, City: "Sofia"})
	doBenchmarkMakeHashFunc(b, PadInside{})
	doBenchmarkMakeHashFunc(b, PadTrailing{})
	doBenchmarkMakeHashFunc(b, [1024]byte{})
	doBenchmarkMakeHashFunc(b, [128]Point{})
	doBenchmarkMakeHashFunc(b, [128]User{})
	doBenchmarkMakeHashFunc(b, [128]PadInside{})
	doBenchmarkMakeHashFunc(b, [128]PadTrailing{})
}

func doBenchmarkMakeHashFunc[T comparable](b *testing.B, val T) {
	hash := MakeHasher[T]()
	b.Run(fmt.Sprintf("%T", val), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = hash(val)
		}
	})
}

func TestCollision_MakeHasher(t *testing.T) {
	n := 20_000_000
	sHasher := MakeHasher[string]()
	iHasher := MakeHasher[int]()
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

func BenchmarkHasher_MakeHasher(b *testing.B) {
	sHasher := MakeHasher[string]()
	iHasher := MakeHasher[int]()
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

func BenchmarkHasher_Parallel_MakeHasher(b *testing.B) {
	sHasher := MakeHasher[string]()
	s := strings.Repeat(testString, 10)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = sHasher(s)
		}
	})
}

// go test -run=^$ -benchmem -count=2 -bench=BenchmarkHasher
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils/xhash
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// BenchmarkHasher_MakeHasher/string-16            48853290                24.60 ns/op            7 B/op          0 allocs/op
// BenchmarkHasher_MakeHasher/string-16            50446735                26.41 ns/op            7 B/op          0 allocs/op
// BenchmarkHasher_MakeHasher/int-16               314223793                3.722 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_MakeHasher/int-16               327831469                3.681 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_Parallel_MakeHasher-16          759237190                1.611 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_Parallel_MakeHasher-16          800960254                1.598 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_GenHasher64/string-16           43828345                26.33 ns/op            7 B/op          0 allocs/op
// BenchmarkHasher_GenHasher64/string-16           43891936                26.14 ns/op            7 B/op          0 allocs/op
// BenchmarkHasher_GenHasher64/int-16              470061992                2.594 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_GenHasher64/int-16              459792141                2.596 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_Parallel_GenHasher64-16         333951942                3.606 ns/op           0 B/op          0 allocs/op
// BenchmarkHasher_Parallel_GenHasher64-16         321419834                3.627 ns/op           0 B/op          0 allocs/op
