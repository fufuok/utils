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
	}
	AssertEqual(t, "", GetString(nil))
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

func TestB64Encode(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "6Kej56CBL+e8lueggX4g6aG25pu/JiM=", B64Encode("解码/编码~ 顶替&#"))
}

func TestB64UrlEncode(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "6Kej56CBL-e8lueggX4g6aG25pu_JiM=", B64UrlEncode("解码/编码~ 顶替&#"))
}

func TestB64Decode(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "解码/编码~ 顶替&#", B64Decode("6Kej56CBL+e8lueggX4g6aG25pu/JiM="))
}

func TestB64UrlDecode(t *testing.T) {
	for _, v := range []struct {
		in, out string
	}{
		{"6Kej56CBL-e8lueggX4g6aG25pu_JiM=", "解码/编码~ 顶替&#"},
		{"123", ""},
	} {
		AssertEqual(t, v.out, B64UrlDecode(v.in))
	}
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
