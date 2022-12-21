package utils

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/fufuok/utils/assert"
)

var (
	testString = "  Fufu 中　文\u2728->?\n*\U0001F63A   "
	testBytes  = []byte(testString)
)

func TestGetString(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		def string
		out string
	}{
		{"Fufu\n 中　文", "", "Fufu\n 中　文"},
		{"", "", ""},
		{nil, "", ""},
		{nil, "NULL", "NULL"},
		{123, "", "123"},
		{123, "456", "123"},
		{123.45, "", "123.45"},
		{true, "", "true"},
		{false, "", "false"},
		{testBytes, "", testString},
		{[]int{0, 2, 1}, "", "[0 2 1]"},
	} {
		assert.Equal(t, v.out, GetString(v.in, v.def))
		assert.Equal(t, v.out, GetSafeString(MustString(v.in), v.def))
	}
	assert.Equal(t, "default", GetString(nil, "default"))
	assert.Equal(t, "default", GetSafeString("", "default"))
}

func TestGetSafeString(t *testing.T) {
	t.Parallel()
	b := []byte("Fufu")
	s := B2S(b)
	safeS1 := string(b)
	safeS2 := GetSafeB2S(b, "optional default")
	safeS3 := GetSafeString(s, "optional default")
	safeS4 := CopyString(s)
	assert.Equal(t, "Fufu", s)

	b[0] = 'X'

	assert.Equal(t, "Xufu", s)
	assert.Equal(t, "Fufu", safeS1)
	assert.Equal(t, "Fufu", safeS2)
	assert.Equal(t, "Fufu", safeS3)
	assert.Equal(t, "Fufu", safeS4)

	assert.Equal(t, "default", GetSafeB2S(nil, "default"))
}

func TestCopyString(t *testing.T) {
	t.Parallel()
	assert.Equal(t, testString, CopyString(testString))
}

func TestCopyB2S(t *testing.T) {
	t.Parallel()
	assert.Equal(t, testString, CopyB2S(testBytes))
}

func BenchmarkCopyB2S(b *testing.B) {
	bs := bytes.Repeat(testBytes, 10)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("CopyB2S", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = CopyB2S(bs)
		}
	})
	b.Run("Std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(bs)
		}
	})
}

// 仅补全函数, 实际直接使用 string(bs)
// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkCopyB2S
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkCopyB2S/CopyB2S-8              12065776                95.64 ns/op          288 B/op          1 allocs/op
// BenchmarkCopyB2S/CopyB2S-8              13230735                98.15 ns/op          288 B/op          1 allocs/op
// BenchmarkCopyB2S/CopyB2S-8              12620445                97.15 ns/op          288 B/op          1 allocs/op
// BenchmarkCopyB2S/Std-8                  11827356                99.76 ns/op          288 B/op          1 allocs/op
// BenchmarkCopyB2S/Std-8                  11875660                96.93 ns/op          288 B/op          1 allocs/op
// BenchmarkCopyB2S/Std-8                  12063805               100.8 ns/op           288 B/op          1 allocs/op

func TestJoinString(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	assert.Equal(t, strings.Join(val, ""), JoinString(val...))
	assert.Equal(t, "1,2/3", JoinString("1", ",", "2", "/", "3"))
}

func TestSearchString(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	assert.Equal(t, 0, SearchString(val, "a"))
	assert.Equal(t, 1, SearchString(val, "b"))
	assert.Equal(t, 2, SearchString(val, "c"))
	assert.Equal(t, -1, SearchString(val, "d"))
}

func TestInStrings(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	assert.Equal(t, true, InStrings(val, "a"))
	assert.Equal(t, false, InStrings(val, "d"))
}

func TestRemoveString(t *testing.T) {
	t.Parallel()
	ok := false
	val := []string{"a", "b", "c"}
	val, ok = RemoveString(val, "b")
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"a", "c"}, val)

	val, ok = RemoveString(val, "b")
	assert.Equal(t, false, ok)
	assert.Equal(t, []string{"a", "c"}, val)
}

func BenchmarkStringPlusLarge(b *testing.B) {
	a := "apiname"
	t := "2021-04-11T12:00:00+08:00"
	c := "192.168.1.100"
	d := RandString(300)
	e := "===0==="
	x := []string{a, e, t, e, c, e, d}

	b.Run("a+b", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a + e + t + e + c + e + d
		}
	})
	b.Run("JoinString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = JoinString(a, e, t, e, c, e, d)
		}
	})
	b.Run("Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf(`%s%s%s%s%s%s%s`, a, e, t, e, c, e, d)
		}
	})
	b.Run("Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Join(x, "")
		}
	})
}

func BenchmarkStringPlus(b *testing.B) {
	a := "ctime:"
	c := "2021-04-11T12:00:00+08:00"
	x := []string{a, c, a, c, a, c}

	b.Run("a+b", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a + c + a + c + a + c
		}
	})
	b.Run("JoinString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = JoinString(a, c, a, c, a, c)
		}
	})
	b.Run("Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s%s%s%s%s%s", a, c, a, c, a, c)
		}
	})
	b.Run("Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Join(x, "")
		}
	})
}

// JoinString 最优, 一次性 + 号拼接性能也不错, 特别是大字符串拼接
// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkStringPlus
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkStringPlusLarge/a+b-8           6230061               202.0 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/a+b-8           6117298               190.1 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/a+b-8           6460974               196.4 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/JoinString-8     6311520               175.9 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/JoinString-8     7053520               177.1 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/JoinString-8     6110152               178.1 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/Sprintf-8       1660620              1126 ns/op             496 B/op          8 allocs/op
// BenchmarkStringPlusLarge/Sprintf-8       1703002               717.8 ns/op           496 B/op          8 allocs/op
// BenchmarkStringPlusLarge/Sprintf-8       1635656               698.6 ns/op           496 B/op          8 allocs/op
// BenchmarkStringPlusLarge/Join-8          5731154               214.4 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/Join-8          5426298               215.6 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlusLarge/Join-8          5483766               225.9 ns/op           384 B/op          1 allocs/op
// BenchmarkStringPlus/a+b-8               10656237               117.6 ns/op            96 B/op          1 allocs/op
// BenchmarkStringPlus/a+b-8               10076886               115.7 ns/op            96 B/op          1 allocs/op
// BenchmarkStringPlus/a+b-8               11085562               114.9 ns/op            96 B/op          1 allocs/op
// BenchmarkStringPlus/JoinString-8         13749231                91.51 ns/op           96 B/op          1 allocs/op
// BenchmarkStringPlus/JoinString-8         12951004                89.82 ns/op           96 B/op          1 allocs/op
// BenchmarkStringPlus/JoinString-8         13045179                91.75 ns/op           96 B/op          1 allocs/op
// BenchmarkStringPlus/Sprintf-8            2197731               539.0 ns/op           192 B/op          7 allocs/op
// BenchmarkStringPlus/Sprintf-8            2252364               539.5 ns/op           192 B/op          7 allocs/op
// BenchmarkStringPlus/Sprintf-8            2269317               542.7 ns/op           192 B/op          7 allocs/op
// BenchmarkStringPlus/Join-8               9639103               127.3 ns/op            96 B/op          1 allocs/op
// BenchmarkStringPlus/Join-8               9567973               124.3 ns/op            96 B/op          1 allocs/op
// BenchmarkStringPlus/Join-8               9280820               126.7 ns/op            96 B/op          1 allocs/op

func TestToLower(t *testing.T) {
	t.Parallel()
	assert.Equal(t, strings.ToLower("A"), ToLower("A"))
	assert.Equal(t, strings.ToLower(testString), ToLower(testString))
}

func BenchmarkToLower(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = ToLower(testString)
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.ToLower(testString)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkToLower$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkToLower/utils-8                18853339                66.37 ns/op           32 B/op          1 allocs/op
// BenchmarkToLower/utils-8                18715989                62.75 ns/op           32 B/op          1 allocs/op
// BenchmarkToLower/utils-8                19896207                61.00 ns/op           32 B/op          1 allocs/op
// BenchmarkToLower/default-8               3519330               376.4 ns/op            32 B/op          1 allocs/op
// BenchmarkToLower/default-8               3166155               360.4 ns/op            32 B/op          1 allocs/op
// BenchmarkToLower/default-8               3497022               343.4 ns/op            32 B/op          1 allocs/op

func TestToUpper(t *testing.T) {
	t.Parallel()
	assert.Equal(t, strings.ToUpper("a"), ToUpper("a"))
	assert.Equal(t, strings.ToUpper(testString), ToUpper(testString))
}

func BenchmarkToUpper(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = ToUpper(testString)
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.ToUpper(testString)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkToUpper$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkToUpper/utils-8                18758617                58.78 ns/op           32 B/op          1 allocs/op
// BenchmarkToUpper/utils-8                20168100                59.67 ns/op           32 B/op          1 allocs/op
// BenchmarkToUpper/utils-8                20464125                59.51 ns/op           32 B/op          1 allocs/op
// BenchmarkToUpper/default-8               3468216               379.7 ns/op            32 B/op          1 allocs/op
// BenchmarkToUpper/default-8               3130531               375.2 ns/op            32 B/op          1 allocs/op
// BenchmarkToUpper/default-8               3007088               372.3 ns/op            32 B/op          1 allocs/op

func TestTrimLeft(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "test/////", TrimLeft("/////test/////", '/'))
	assert.Equal(t, "test/", TrimLeft("/test/", '/'))
	assert.Equal(t, "", TrimLeft(" ", ' '))
	assert.Equal(t, "", TrimLeft("  ", ' '))
	assert.Equal(t, "", TrimLeft("", ' '))
	assert.Equal(t, strings.TrimLeft(testString, " "), TrimLeft(testString, ' '))
}

func BenchmarkTrimLeft(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = TrimLeft(testString, ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.TrimLeft(testString, " ")
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrimLeft$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrimLeft/utils-8               315254535                3.406 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeft/utils-8               398772445                3.162 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeft/utils-8               384151815                3.144 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimLeft/default-8             17307051                62.65 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimLeft/default-8             19507880                62.69 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimLeft/default-8             19264666                61.05 ns/op           24 B/op          1 allocs/op

func TestTrimRight(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "/////test", TrimRight("/////test/////", '/'))
	assert.Equal(t, "/test", TrimRight("/test/", '/'))
	assert.Equal(t, "", TrimRight(" ", ' '))
	assert.Equal(t, "", TrimRight("  ", ' '))
	assert.Equal(t, "", TrimRight("", ' '))
	assert.Equal(t, strings.TrimRight(testString, " "), TrimRight(testString, ' '))
}

func BenchmarkTrimRight(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = TrimRight(testString, ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.TrimRight(testString, " ")
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrimRight$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrimRight/utils-8              321572618                3.683 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRight/utils-8              304095558                3.692 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRight/utils-8              329710909                3.726 ns/op           0 B/op          0 allocs/op
// BenchmarkTrimRight/default-8            13112961                93.11 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimRight/default-8            13165082                90.08 ns/op           24 B/op          1 allocs/op
// BenchmarkTrimRight/default-8            13624411                91.24 ns/op           24 B/op          1 allocs/op

func TestTrim(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "test", Trim("/////test/////", '/'))
	assert.Equal(t, "test", Trim("/test/", '/'))
	assert.Equal(t, "test", Trim("test", '/'))
	assert.Equal(t, "", Trim(" ", ' '))
	assert.Equal(t, "", Trim("  ", ' '))
	assert.Equal(t, "", Trim("", ' '))
	assert.Equal(t, strings.Trim(testString, " "), Trim(testString, ' '))
}

func BenchmarkTrim(b *testing.B) {
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = Trim(testString, ' ')
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.Trim(testString, " ")
		}
	})
	b.Run("trimspace", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strings.TrimSpace(testString)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkTrim$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkTrim/utils-8           138754935                8.755 ns/op           0 B/op          0 allocs/op
// BenchmarkTrim/utils-8           137675163                8.641 ns/op           0 B/op          0 allocs/op
// BenchmarkTrim/utils-8           136393192                8.635 ns/op           0 B/op          0 allocs/op
// BenchmarkTrim/default-8         10280344               115.6 ns/op            24 B/op          1 allocs/op
// BenchmarkTrim/default-8         10843497               118.6 ns/op            24 B/op          1 allocs/op
// BenchmarkTrim/default-8         10608081               115.2 ns/op            24 B/op          1 allocs/op
// BenchmarkTrim/trimspace-8       22668753                52.89 ns/op            0 B/op          0 allocs/op
// BenchmarkTrim/trimspace-8       21154394                49.38 ns/op            0 B/op          0 allocs/op
// BenchmarkTrim/trimspace-8       24796974                48.31 ns/op            0 B/op          0 allocs/op

func TestEqualFold(t *testing.T) {
	res := CopyString(testString)
	assert.Equal(t, true, EqualFold(res, ToUpper(testString)))
	assert.Equal(t, true, EqualFold(ToLower(res), ToUpper(testString)))
	assert.Equal(t, false, EqualFold(res, Trim(testString, ' ')))
	assert.Equal(t, false, EqualFold("\na", "*A"))
}

func BenchmarkEqualFold(b *testing.B) {
	s := ToUpper(testString)
	t := ToLower(testString)
	ok := false
	b.ResetTimer()
	b.Run("utils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ok = EqualFold(s, t)
			assert.Equal(b, true, ok)
		}
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ok = strings.EqualFold(s, t)
			assert.Equal(b, true, ok)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkEqualFold$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkEqualFold/utils-8              12863390                95.72 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFold/utils-8              12384744                95.51 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFold/utils-8              12885228                94.62 ns/op            0 B/op          0 allocs/op
// BenchmarkEqualFold/default-8             6135718               200.4 ns/op             0 B/op          0 allocs/op
// BenchmarkEqualFold/default-8             6059118               201.6 ns/op             0 B/op          0 allocs/op
// BenchmarkEqualFold/default-8             6149115               199.1 ns/op             0 B/op          0 allocs/op

var cutTests = []struct {
	s, sep        string
	before, after string
	found         bool
}{
	{"abc", "b", "a", "c", true},
	{"abc", "a", "", "bc", true},
	{"abc", "c", "ab", "", true},
	{"abc", "abc", "", "", true},
	{"abc", "", "", "abc", true},
	{"abc", "d", "abc", "", false},
	{"", "d", "", "", false},
	{"", "", "", "", true},
}

func TestCutString(t *testing.T) {
	for _, tt := range cutTests {
		if before, after, found := CutString(tt.s, tt.sep); before != tt.before || after != tt.after || found != tt.found {
			t.Errorf("Cut(%q, %q) = %q, %q, %v, want %q, %q, %v", tt.s, tt.sep, before, after, found, tt.before, tt.after, tt.found)
		}
	}
}
