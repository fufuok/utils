package utils

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Zip(data)
	AssertEqual(t, true, err == nil, "failed to zip")

	src, err := Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")
	t.Logf("origin len: %d, zipped len: %d", len(data), len(dst))
}

func BenchmarkCompressZip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Zip(data)
	}
}

func BenchmarkCompressUnZip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	dec, _ := Zip(data)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Unzip(dec)
		}
	})
}

// go test -run=^$ -benchtime=1s -benchmem -count=2 -bench=BenchmarkCompress
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkCompressZip-4            150130              8045 ns/op               8 B/op          0 allocs/op
// BenchmarkCompressZip-4            148700              7960 ns/op               8 B/op          0 allocs/op
// BenchmarkCompressUnZip-4          362565              3525 ns/op           14092 B/op          7 allocs/op
// BenchmarkCompressUnZip-4          380100              3504 ns/op           14093 B/op          7 allocs/op
