package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestCompressZip(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Zip(data)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.BestSpeed, len(dst))

	src, err := Unzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, zlib.BestCompression)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.BestCompression, len(dst))

	src, err = Unzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")

	dst, err = ZipLevel(data, 6)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), zlib.DefaultCompression, len(dst))

	src, err = Unzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")
}

func TestCompressGzip(t *testing.T) {
	data := bytes.Repeat(testBytes, 100)

	dst, err := Gzip(data)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestSpeed, len(dst))

	src, err := Ungzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")

	dst, err = GzipLevel(data, gzip.BestCompression)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.BestCompression, len(dst))

	src, err = Ungzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")

	dst, err = GzipLevel(data, 6)
	assert.Equal(t, true, err == nil, "failed to zip")
	t.Logf("origin len: %d, zipped(level: %d) len: %d", len(data), gzip.DefaultCompression, len(dst))

	src, err = Ungzip(dst)
	assert.Equal(t, true, err == nil, "failed to unzip")

	assert.Equal(t, true, bytes.Equal(data, src), "data != src")
}

func BenchmarkCompress_Zip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Zip(data)
	}
}

func BenchmarkCompress_Zip_BestCompression(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ZipLevel(data, gzip.BestCompression)
	}
}

func BenchmarkCompress_Unzip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	dec, _ := Zip(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Unzip(dec)
	}
}

func BenchmarkCompress_Zip_Parallel(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Zip(data)
		}
	})
}

func BenchmarkCompress_Unzip_Parallel(b *testing.B) {
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

func BenchmarkCompress_ZipUnzip_Parallel(b *testing.B) {
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

func BenchmarkCompress_Gzip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Gzip(data)
	}
}

func BenchmarkCompress_Gzip_BestCompression(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GzipLevel(data, gzip.BestCompression)
	}
}

func BenchmarkCompress_Ungzip(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	dec, _ := Gzip(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Ungzip(dec)
	}
}

func BenchmarkCompress_Gzip_Parallel(b *testing.B) {
	data := bytes.Repeat(testBytes, 100)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Gzip(data)
		}
	})
}

func BenchmarkCompress_Ungzip_Parallel(b *testing.B) {
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

func BenchmarkCompress_GzipUngzip(b *testing.B) {
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
// BenchmarkCompress_Zip-4                           137169              8728 ns/op              96 B/op          1 allocs/op
// BenchmarkCompress_Zip-4                           136028              8758 ns/op             104 B/op          1 allocs/op
// BenchmarkCompress_Zip_BestCompression-4            31932             35699 ns/op             105 B/op          1 allocs/op
// BenchmarkCompress_Zip_BestCompression-4            31845             35171 ns/op              80 B/op          1 allocs/op
// BenchmarkCompress_Unzip-4                          84660             14576 ns/op           52604 B/op         12 allocs/op
// BenchmarkCompress_Unzip-4                          80158             15248 ns/op           52604 B/op         12 allocs/op
// BenchmarkCompress_Zip_Parallel-4                  504540              2354 ns/op             100 B/op          1 allocs/op
// BenchmarkCompress_Zip_Parallel-4                  490543              2429 ns/op              96 B/op          1 allocs/op
// BenchmarkCompress_Unzip_Parallel-4                102718             11924 ns/op           52606 B/op         12 allocs/op
// BenchmarkCompress_Unzip_Parallel-4                 82278             14111 ns/op           52604 B/op         12 allocs/op
// BenchmarkCompress_ZipUnzip_Parallel-4              85200             15105 ns/op           61635 B/op         13 allocs/op
// BenchmarkCompress_ZipUnzip_Parallel-4              85476             14631 ns/op           61918 B/op         13 allocs/op
// BenchmarkCompress_Gzip-4                          155563              7617 ns/op             119 B/op          1 allocs/op
// BenchmarkCompress_Gzip-4                          155354              7601 ns/op             112 B/op          1 allocs/op
// BenchmarkCompress_Gzip_BestCompression-4           32972             34728 ns/op             120 B/op          1 allocs/op
// BenchmarkCompress_Gzip_BestCompression-4           33518             33674 ns/op              96 B/op          1 allocs/op
// BenchmarkCompress_Ungzip-4                        173500              7082 ns/op           12037 B/op          6 allocs/op
// BenchmarkCompress_Ungzip-4                        162517              7543 ns/op           12038 B/op          6 allocs/op
// BenchmarkCompress_Gzip_Parallel-4                 584652              2074 ns/op             114 B/op          1 allocs/op
// BenchmarkCompress_Gzip_Parallel-4                 507086              2096 ns/op             114 B/op          1 allocs/op
// BenchmarkCompress_Ungzip_Parallel-4               318189              3876 ns/op           12042 B/op          6 allocs/op
// BenchmarkCompress_Ungzip_Parallel-4               313982              4064 ns/op           12043 B/op          6 allocs/op
// BenchmarkCompress_GzipUngzip-4                    177799              6656 ns/op           14329 B/op          7 allocs/op
// BenchmarkCompress_GzipUngzip-4                    192142              6538 ns/op           14509 B/op          7 allocs/op
