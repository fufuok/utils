package utils

import (
	"bytes"
	"strings"
	"testing"
)

func TestGetBytes(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		def []byte
		out []byte
	}{
		{"Fufu\n 中　文", nil, []byte("Fufu\n 中　文")},
		{nil, nil, nil},
		{nil, []byte("NULL"), []byte("NULL")},
		{123, nil, []byte("123")},
		{123, []byte("456"), []byte("123")},
		{123.45, nil, []byte("123.45")},
		{true, nil, []byte("true")},
		{false, nil, []byte("false")},
		{testString, nil, testBytes},
		{[]int{0, 2, 1}, nil, []byte("[0 2 1]")},
	} {
		AssertEqual(t, v.out, GetBytes(v.in, v.def))
	}
	AssertEqual(t, 0, len(GetBytes(nil)))
	AssertEqual(t, []byte{}, GetSafeBytes(nil))
}

func TestGetSafeBytes(t *testing.T) {
	t.Parallel()
	b := []byte("Fufu")
	s := B2S(b)
	safeB1 := []byte(s)
	safeB2 := GetSafeS2B(s)
	safeB3 := GetSafeBytes(b)
	safeB4 := CopyBytes(b)
	AssertEqual(t, "Fufu", s)

	b[0] = 'X'

	AssertEqual(t, "Xufu", s)
	AssertEqual(t, []byte("Fufu"), safeB1)
	AssertEqual(t, []byte("Fufu"), safeB2)
	AssertEqual(t, []byte("Fufu"), safeB3)
	AssertEqual(t, []byte("Fufu"), safeB4)

	AssertEqual(t, []byte("default"), GetSafeS2B("", []byte("default")))
}

func TestCopyBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, testBytes, CopyBytes(testBytes))
}

func TestCopyS2B(t *testing.T) {
	t.Parallel()
	AssertEqual(t, testBytes, CopyS2B(testString))
}

func TestJoinBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, []byte("1,2,3"), JoinBytes([]byte("1"), []byte(","), []byte("2"), []byte(","), []byte("3")))
	AssertEqual(
		t,
		bytes.Join([][]byte{[]byte("1"), []byte("2"), []byte("3")}, []byte(",")),
		JoinBytes([]byte("1"), []byte(","), []byte("2"), []byte(","), []byte("3")),
	)
}

func BenchmarkCopyS2B(b *testing.B) {
	s := strings.Repeat(testString, 10)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("utils", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = CopyS2B(s)
		}
	})
	b.Run("default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = []byte(s)
		}
	})
	b.Run("append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = append(append([]byte{}, s...))
		}
	})
}

// 仅补全函数, 实际直接使用 []byte(string)
// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkCopyS2B
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkCopyS2B/utils-8                10197265               109.8 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/utils-8                12149991               107.5 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/utils-8                12256554               110.6 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/default-8              10587153               116.1 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/default-8              10294596               113.3 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/default-8               9938052               123.7 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/append-8                9818898               121.1 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/append-8               10450903               119.1 ns/op           320 B/op          1 allocs/op
// BenchmarkCopyS2B/append-8                9462102               117.1 ns/op           320 B/op          1 allocs/op

func TestToLowerBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, bytes.ToLower([]byte("A")), ToLowerBytes([]byte("A")))
	AssertEqual(t, bytes.ToLower([]byte("/TesT/")), ToLowerBytes([]byte("/TesT/")))
	AssertEqual(t, true, bytes.Equal(bytes.ToLower([]byte("/TesT/")), ToLowerBytes([]byte("/TesT/"))))
	AssertEqual(t, false, bytes.Equal(bytes.ToUpper([]byte("/TesT/")), ToLowerBytes([]byte("/TesT/"))))
}

func BenchmarkToLowerBytes(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = ToLowerBytes([]byte("/TesT/"))
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.ToLower([]byte("/TesT/"))
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkToLowerBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkToLowerBytes/utils-8           169237525                7.012 ns/op           0 B/op          0 allocs/op
// BenchmarkToLowerBytes/utils-8           164508339                7.176 ns/op           0 B/op          0 allocs/op
// BenchmarkToLowerBytes/utils-8           163584484                8.336 ns/op           0 B/op          0 allocs/op
// BenchmarkToLowerBytes/default-8         21546925                56.56 ns/op            8 B/op          1 allocs/op
// BenchmarkToLowerBytes/default-8         25344794                42.91 ns/op            8 B/op          1 allocs/op
// BenchmarkToLowerBytes/default-8         27910351                45.92 ns/op            8 B/op          1 allocs/op

func TestToUpperBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, bytes.ToUpper([]byte("a")), ToUpperBytes([]byte("a")))
	AssertEqual(t, bytes.ToUpper([]byte("/TesT/")), ToUpperBytes([]byte("/TesT/")))
	AssertEqual(t, true, bytes.Equal(bytes.ToUpper([]byte("/TesT/")), ToUpperBytes([]byte("/TesT/"))))
	AssertEqual(t, false, bytes.Equal(bytes.ToLower([]byte("/TesT/")), ToUpperBytes([]byte("/TesT/"))))
}

func BenchmarkToUpperBytes(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = ToUpperBytes([]byte("/TesT/"))
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.ToUpper([]byte("/TesT/"))
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkToUpperBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkToUpperBytes/utils-8           170378092                7.061 ns/op           0 B/op          0 allocs/op
// BenchmarkToUpperBytes/utils-8           170443166                7.380 ns/op           0 B/op          0 allocs/op
// BenchmarkToUpperBytes/utils-8           172179099                7.373 ns/op           0 B/op          0 allocs/op
// BenchmarkToUpperBytes/default-8         24008786                46.90 ns/op            8 B/op          1 allocs/op
// BenchmarkToUpperBytes/default-8         29689668                44.83 ns/op            8 B/op          1 allocs/op
// BenchmarkToUpperBytes/default-8         29171811                42.58 ns/op            8 B/op          1 allocs/op

func TestTrimLeftBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, []byte("test/////"), TrimLeftBytes([]byte("/////test/////"), '/'))
	AssertEqual(t, []byte("test/"), TrimLeftBytes([]byte("/test/"), '/'))
	AssertEqual(t, 0, len(TrimLeftBytes([]byte(" "), ' ')))
	AssertEqual(t, 0, len(TrimLeftBytes([]byte("  "), ' ')))
	AssertEqual(t, 0, len(TrimLeftBytes([]byte(""), ' ')))
	AssertEqual(t, bytes.TrimLeft([]byte("  TesT  "), " "), TrimLeftBytes([]byte("  TesT  "), ' '))
}

func BenchmarkTrimLeftBytes(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = TrimLeftBytes([]byte("  TesT  "), ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.TrimLeft([]byte("  TesT  "), " ")
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrimLeftBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrimLeftBytes/utils-8          347581670                3.051 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeftBytes/utils-8          362493957                3.227 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeftBytes/utils-8          385871320                3.401 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeftBytes/default-8        15500971                71.58 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimLeftBytes/default-8        15734938                73.00 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimLeftBytes/default-8        19290650                65.52 ns/op           24 B/op          1 allocs/op

func TestTrimRightBytes(t *testing.T) {
	t.Parallel()
	AssertEqual(t, []byte("/////test"), TrimRightBytes([]byte("/////test/////"), '/'))
	AssertEqual(t, []byte("/test"), TrimRightBytes([]byte("/test/"), '/'))
	AssertEqual(t, 0, len(TrimRightBytes([]byte(" "), ' ')))
	AssertEqual(t, 0, len(TrimRightBytes([]byte("  "), ' ')))
	AssertEqual(t, 0, len(TrimRightBytes([]byte(""), ' ')))
	AssertEqual(t, bytes.TrimRight([]byte("  TesT  "), " "), TrimRightBytes([]byte("  TesT  "), ' '))
}

func BenchmarkTrimRightBytes(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = TrimRightBytes([]byte("  TesT  "), ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.TrimRight([]byte("  TesT  "), " ")
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrimRightBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrimRightBytes/utils-8                 381420740                3.423 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRightBytes/utils-8                 394509609                3.076 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRightBytes/utils-8                 396171007                3.118 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRightBytes/default-8               17798528                66.47 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimRightBytes/default-8               18517689                64.92 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimRightBytes/default-8               18744405                65.73 ns/op           24 B/op          1 allocs/op

func TestTrimBytes(t *testing.T) {
	AssertEqual(t, []byte("test"), TrimBytes([]byte("/////test/////"), '/'))
	AssertEqual(t, []byte("test"), TrimBytes([]byte("/test/"), '/'))
	AssertEqual(t, []byte("test"), TrimBytes([]byte("test"), '/'))
	AssertEqual(t, 0, len(TrimBytes([]byte(" "), ' ')))
	AssertEqual(t, 0, len(TrimBytes([]byte("  "), ' ')))
	AssertEqual(t, 0, len(TrimBytes([]byte(""), ' ')))
	AssertEqual(t, bytes.Trim([]byte("  TesT  "), " "), TrimBytes([]byte("  TesT  "), ' '))
}

func BenchmarkTrimBytes(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = TrimBytes([]byte("  TesT  "), ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.Trim([]byte("  TesT  "), " ")
		}
	})
	b.Run("trimspace", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = bytes.TrimSpace([]byte("  TesT  "))
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrimBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrimBytes/utils-8              142327920                8.657 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimBytes/utils-8              150808749                8.059 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimBytes/utils-8              144292578                8.129 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimBytes/default-8            14513946                82.10 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimBytes/default-8            14985987                79.10 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimBytes/default-8            15338756                78.82 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimBytes/trimspace-8          100000000               11.56 ns/op            0 B/op          0 allocs/op
// BenchmarkTrimBytes/trimspace-8          100000000               11.92 ns/op            0 B/op          0 allocs/op
// BenchmarkTrimBytes/trimspace-8          100000000               10.94 ns/op            0 B/op          0 allocs/op

func TestEqualFoldBytes(t *testing.T) {
	res := []byte("  tESt  ")
	AssertEqual(t, true, EqualFoldBytes(res, ToUpperBytes([]byte("  TesT  "))))
	AssertEqual(t, true, EqualFoldBytes(ToLowerBytes(res), ToUpperBytes([]byte("  TesT  "))))
	AssertEqual(t, false, EqualFoldBytes(res, TrimBytes([]byte("  TesT  "), ' ')))
}

func BenchmarkEqualFoldBytes(b *testing.B) {
	s := ToUpperBytes(CopyBytes(testBytes))
	t := ToLowerBytes(CopyBytes(testBytes))
	ok := false
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ok = EqualFoldBytes(s, t)
			AssertEqual(b, true, ok)
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ok = bytes.EqualFold(s, t)
			AssertEqual(b, true, ok)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkEqualFoldBytes
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkEqualFoldBytes/utils-8                 11882070                95.38 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFoldBytes/utils-8                 12616200                97.33 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFoldBytes/utils-8                 12393940                99.75 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFoldBytes/default-8                5822217               209.0 ns/op             0 B/op          0 allocs/op
// BenchmarkEqualFoldBytes/default-8                6151662               197.9 ns/op             0 B/op          0 allocs/op
// BenchmarkEqualFoldBytes/default-8                5932452               199.1 ns/op             0 B/op          0 allocs/op
