package utils

import (
	"fmt"
	"strings"
	"testing"
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
		{[]byte("Fufu 中　文\u2728->?\n*\U0001F63A"), "", "Fufu 中　文\u2728->?\n*\U0001F63A"},
		{[]int{0, 2, 1}, "", "[0 2 1]"},
	} {
		AssertEqual(t, v.out, GetString(v.in, v.def))
		AssertEqual(t, v.out, GetSafeString(MustString(v.in), v.def))
	}
	AssertEqual(t, "default", GetString(nil, "default"))
	AssertEqual(t, "default", GetSafeString("", "default"))
}

func TestGetSafeString(t *testing.T) {
	t.Parallel()
	b := []byte("Fufu")
	s := B2S(b)
	safeS1 := string(b)
	safeS2 := GetSafeB2S(b, "optional default")
	safeS3 := GetSafeString(s, "optional default")
	safeS4 := CopyString(s)
	AssertEqual(t, "Fufu", s)

	b[0] = 'X'

	AssertEqual(t, "Xufu", s)
	AssertEqual(t, "Fufu", safeS1)
	AssertEqual(t, "Fufu", safeS2)
	AssertEqual(t, "Fufu", safeS3)
	AssertEqual(t, "Fufu", safeS4)

	AssertEqual(t, "default", GetSafeB2S(nil, "default"))
}

func TestCopyString(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "Fufu 中　文\u2728->?\n*\U0001F63A", CopyString("Fufu 中　文\u2728->?\n*\U0001F63A"))
}

func TestAddString(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	AssertEqual(t, strings.Join(val, ""), AddString(val...))
	AssertEqual(t, "1,2/3", AddString("1", ",", "2", "/", "3"))
}

func TestSearchString(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	AssertEqual(t, 0, SearchString(val, "a"))
	AssertEqual(t, 1, SearchString(val, "b"))
	AssertEqual(t, 2, SearchString(val, "c"))
	AssertEqual(t, -1, SearchString(val, "d"))
}

func TestInStrings(t *testing.T) {
	t.Parallel()
	val := []string{"a", "b", "c"}
	AssertEqual(t, true, InStrings(val, "a"))
	AssertEqual(t, false, InStrings(val, "d"))
}

func BenchmarkStringPlusBig(b *testing.B) {
	b.ReportAllocs()
	a := "apiname"
	t := "2021-04-11T12:00:00+08:00"
	c := "192.168.1.100"
	d := RandString(300)
	e := "===0==="
	for i := 0; i < b.N; i++ {
		_ = a + e + t + e + c + e + d
	}
}

func BenchmarkAddStringBig(b *testing.B) {
	a := "apiname"
	t := "2021-04-11T12:00:00+08:00"
	c := "192.168.1.100"
	d := RandString(300)
	e := "===0==="
	for i := 0; i < b.N; i++ {
		_ = AddString(a, e, t, e, c, e, d)
	}
}

func BenchmarkStringPlus(b *testing.B) {
	b.ReportAllocs()
	a := "2021-04-11T12:00:00+08:00"
	c := "2021-04-11T12:00:00+08:00"
	d := "192.168.1.100"
	for i := 0; i < b.N; i++ {
		_ = `{"_ctime":"` + a + `","_gtime":"` + c + `","_cip":"` + d + `"}`
	}
}

func BenchmarkAddString(b *testing.B) {
	a := "2021-04-11T12:00:00+08:00"
	c := "2021-04-11T12:00:00+08:00"
	d := "192.168.1.100"
	for i := 0; i < b.N; i++ {
		_ = AddString(`{"_ctime":"`, a, `","_gtime":"`, c, `","_cip":"`, d, `"}`)
	}
}

func BenchmarkSprintf(b *testing.B) {
	a := "2021-04-11T12:00:00+08:00"
	c := "2021-04-11T12:00:00+08:00"
	d := "192.168.1.100"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(`{"_ctime":"%s","_gtime":"%s","_cip":"%s"}`,
			a, c, d)
	}
}

// 高频率使用场景, bytes copy 最优, 一次性 + 号拼接性能也不错, 特别是大字符串拼接 (go1.15.6)
// BenchmarkStringPlusBig-8   	43791952	       272 ns/op	     704 B/op	       1 allocs/op
// BenchmarkStringPlusBig-8   	44648138	       272 ns/op	     704 B/op	       1 allocs/op
// BenchmarkStringPlusBig-8   	43970554	       271 ns/op	     704 B/op	       1 allocs/op
// BenchmarkAddStringBig-8    	41560038	       259 ns/op	     704 B/op	       1 allocs/op
// BenchmarkAddStringBig-8    	45365892	       261 ns/op	     704 B/op	       1 allocs/op
// BenchmarkAddStringBig-8    	43232275	       263 ns/op	     704 B/op	       1 allocs/op
// BenchmarkStringPlus-8      	95079327	       127 ns/op	     112 B/op	       1 allocs/op
// BenchmarkStringPlus-8      	90380433	       128 ns/op	     112 B/op	       1 allocs/op
// BenchmarkStringPlus-8      	94727542	       129 ns/op	     112 B/op	       1 allocs/op
// BenchmarkAddString-8       	100000000	       107 ns/op	     112 B/op	       1 allocs/op
// BenchmarkAddString-8       	100000000	       107 ns/op	     112 B/op	       1 allocs/op
// BenchmarkAddString-8       	100000000	       107 ns/op	     112 B/op	       1 allocs/op
// BenchmarkSprintf-8         	31582645	       389 ns/op	     160 B/op	       4 allocs/op
// BenchmarkSprintf-8         	30813024	       393 ns/op	     160 B/op	       4 allocs/op
// BenchmarkSprintf-8         	31000237	       388 ns/op	     160 B/op	       4 allocs/op

func BenchmarkSprintfAny_X(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%v%d",
			[]byte("码~ 顶替&#"),
			"2021-04-11T12:00:00+08:00",
			true,
			123,
		)
	}
}

func BenchmarkAddStringAny_X(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AddString(
			MustString([]byte("码~ 顶替&#")),
			"2021-04-11T12:00:00+08:00",
			MustString(true),
			MustString(123),
		)
	}
}

func BenchmarkSprintfString_X(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s",
			"码~ 顶替&#",
			"2021-04-11T12:00:00+08:00",
			"true",
			"123",
		)
	}
}

func BenchmarkAddString_X(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AddString(
			"码~ 顶替&#",
			"2021-04-11T12:00:00+08:00",
			"true",
			"123",
		)
	}
}

func BenchmarkStringsJoin_X(b *testing.B) {
	s := []string{
		"码~ 顶替&#",
		"2021-04-11T12:00:00+08:00",
		"true",
		"123",
	}
	for i := 0; i < b.N; i++ {
		_ = strings.Join(s, "")
	}
}

// BenchmarkSprintfAny_X-8      	 3408980	       319 ns/op	      96 B/op	       3 allocs/op
// BenchmarkSprintfAny_X-8      	 3749611	       334 ns/op	      96 B/op	       3 allocs/op
// BenchmarkSprintfAny_X-8      	 3161620	       331 ns/op	      96 B/op	       3 allocs/op
// BenchmarkAddStringAny_X-8    	 3043008	       397 ns/op	     104 B/op	       5 allocs/op
// BenchmarkAddStringAny_X-8    	 3076884	       384 ns/op	     104 B/op	       5 allocs/op
// BenchmarkAddStringAny_X-8    	 3154803	       380 ns/op	     104 B/op	       5 allocs/op
// BenchmarkSprintfString_X-8   	 5598256	       238 ns/op	      48 B/op	       1 allocs/op
// BenchmarkSprintfString_X-8   	 5505559	       223 ns/op	      48 B/op	       1 allocs/op
// BenchmarkSprintfString_X-8   	 5557806	       214 ns/op	      48 B/op	       1 allocs/op
// BenchmarkAddString_X-8       	18094471	        67.9 ns/op	      48 B/op	       1 allocs/op
// BenchmarkAddString_X-8       	18498307	        72.5 ns/op	      48 B/op	       1 allocs/op
// BenchmarkAddString_X-8       	17632718	        71.9 ns/op	      48 B/op	       1 allocs/op
// BenchmarkStringsJoin_X-8     	12587984	       107 ns/op	      48 B/op	       1 allocs/op
// BenchmarkStringsJoin_X-8     	10302427	       149 ns/op	      48 B/op	       1 allocs/op
// BenchmarkStringsJoin_X-8     	12493851	       102 ns/op	      48 B/op	       1 allocs/op
