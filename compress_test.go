package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"testing"
)

func TestCompressZip(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Zip(data)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.BestSpeed, len(dst))

	src, err := Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, zlib.BestCompression)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.BestCompression, len(dst))

	src, err = Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, 6)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.DefaultCompression, len(dst))

	src, err = Unzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")
}

func TestCompressGzip(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Gzip(data)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestSpeed, len(dst))

	src, err := Ungzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = GzipLevel(data, gzip.BestCompression)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestCompression, len(dst))

	src, err = Ungzip(dst)
	AssertEqual(t, true, err == nil, "failed to unzip")

	AssertEqual(t, true, bytes.Equal(data, src), "data != src")

	dst, err = GzipLevel(data, 6)
	AssertEqual(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.DefaultCompression, len(dst))

	src, err = Ungzip(dst)
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

func BenchmarkCompressUnzip(b *testing.B) {
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

func BenchmarkCompressZipUnzip(b *testing.B) {
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

func BenchmarkCompressGzip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Gzip(data)
	}
}

func BenchmarkCompressGzipBestCompression(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GzipLevel(data, gzip.BestCompression)
	}
}

func BenchmarkCompressUngzip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	dec, _ := Gzip(data)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Ungzip(dec)
		}
	})
}

func BenchmarkCompressGzipUngzip(b *testing.B) {
	bs := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dec, _ := Gzip(bs)
			src, err := Ungzip(dec)
			if err != nil {
				b.Fatal(err)
			}
			if !EqualFoldBytes(src, bs) {
				b.Fatal("src != bs")
			}
		}
	})
}

// go test -run=^$ -benchtime=1s -benchmem -count=2 -bench=BenchmarkCompress
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkCompressZip-4                            133740              8911 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressZip-4                            134272              8975 ns/op               8 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              35164             33991 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressZipBestCompression-4              35014             33796 ns/op              23 B/op          0 allocs/op
// BenchmarkCompressUnzip-4                          148243              8101 ns/op           52604 B/op         12 allocs/op
// BenchmarkCompressUnzip-4                          151365              8247 ns/op           52604 B/op         12 allocs/op
// BenchmarkCompressZipUnzip-4                        88929             13559 ns/op           61415 B/op         12 allocs/op
// BenchmarkCompressZipUnzip-4                        84874             13709 ns/op           61295 B/op         12 allocs/op
// BenchmarkCompressGzip-4                           151142              7880 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressGzip-4                           151334              7877 ns/op               7 B/op          0 allocs/op
// BenchmarkCompressGzipBestCompression-4             36963             32587 ns/op               0 B/op          0 allocs/op
// BenchmarkCompressGzipBestCompression-4             37431             32252 ns/op              21 B/op          0 allocs/op
// BenchmarkCompressUngzip-4                         365432              3492 ns/op           12043 B/op          6 allocs/op
// BenchmarkCompressUngzip-4                         363258              3426 ns/op           12043 B/op          6 allocs/op
// BenchmarkCompressGzipUngzip-4                     186872              6677 ns/op           14367 B/op          6 allocs/op
// BenchmarkCompressGzipUngzip-4                     183661              6602 ns/op           13901 B/op          6 allocs/op
