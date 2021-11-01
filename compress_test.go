package utils

import (
	"bytes"
	"compress/gzip"
	"testing"
)

func TestCompress(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Zip(data)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestSpeed, len(dst))

	src, err := Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, gzip.BestCompression)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestCompression, len(dst))

	src, err = Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, 6)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.DefaultCompression, len(dst))

	src, err = Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")
}

func BenchmarkCompressZip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Zip(data)
	}
}

func BenchmarkCompressZipBestCompression(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ZipLevel(data, gzip.BestCompression)
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

func BenchmarkCompressZipUnZip(b *testing.B) {
	bs := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dec, _ := Zip(bs)
			src, err := Unzip(dec)
			if err != nil {
				b.Fatal(err)
			}
			if !EqualFoldBytes(src, bs) {
				b.Fatal("src != bs")
			}
		}
	})
}

// compress/gzip
// go test -run=^$ -benchtime=1s -benchmem -count=2 -bench=BenchmarkCompress
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkCompressZip-4                            150390              8039 ns/op               8 B/op          0 allocs/op
// BenchmarkCompressZip-4                            150478              8046 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              37694             31277 ns/op              21 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              37742             31568 ns/op              21 B/op          0 allocs/op
// BenchmarkCompressUnZip-4                          371936              3354 ns/op           14092 B/op          7 allocs/op
// BenchmarkCompressUnZip-4                          383005              3400 ns/op           14091 B/op          7 allocs/op
// BenchmarkCompressZipUnZip-4                       174706              6549 ns/op           16282 B/op          7 allocs/op
// BenchmarkCompressZipUnZip-4                       186423              6427 ns/op           16248 B/op          7 allocs/op

// klauspost/compress/gzip
// go test -run=^$ -benchtime=1s -benchmem -count=2 -bench=BenchmarkCompress
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/xy-data-router/internal/gzip
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkCompressZip-4                            192632              6230 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressZip-4                            189828              6209 ns/op               4 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              48205             24751 ns/op              20 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              48046             24737 ns/op              21 B/op          0 allocs/op
// BenchmarkCompressUnZip-4                          397161              3311 ns/op           14122 B/op          9 allocs/op
// BenchmarkCompressUnZip-4                          391064              3336 ns/op           14125 B/op          9 allocs/op
// BenchmarkCompressZipUnZip-4                       199250              5850 ns/op           16644 B/op          9 allocs/op
// BenchmarkCompressZipUnZip-4                       198687              5993 ns/op           16789 B/op          9 allocs/op
