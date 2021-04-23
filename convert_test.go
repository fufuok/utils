package utils

import (
	"strings"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := StringToBytes(s)
		AssertEqual(t, expected, actual)
		AssertEqual(t, len(expected), len(actual))
	}

	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := StringToBytes(expected)

	AssertEqual(t, []byte(expected), actual)
}

func TestString2Bytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := String2Bytes(s)
		AssertEqual(t, expected, actual)
		AssertEqual(t, len(expected), len(actual))
	}

	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := String2Bytes(expected)

	AssertEqual(t, []byte(expected), actual)
}

func TestStr2Bytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := Str2Bytes(s)
		AssertEqual(t, expected, actual)
		AssertEqual(t, len(expected), len(actual))
	}

	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := Str2Bytes(expected)

	AssertEqual(t, []byte(expected), actual)
}

func TestS2B(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := S2B(s)
		AssertEqual(t, expected, actual)
		AssertEqual(t, len(expected), len(actual))
	}

	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := S2B(expected)

	AssertEqual(t, []byte(expected), actual)
}

func TestB2S(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		b := RandBytes(64)
		AssertEqual(t, string(b), B2S(b))
	}

	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := B2S([]byte(expected))

	AssertEqual(t, expected, actual)
}

func TestMustJSONString(t *testing.T) {
	t.Parallel()
	actual := MustJSONString(map[string]interface{}{
		"_c": "中 文",
		"a":  true,
		"b":  1.23,
	})

	AssertEqual(t, true, strings.Contains(actual, `"a":true`))
	AssertEqual(t, true, strings.Contains(actual, `"b":1.23`))
	AssertEqual(t, true, strings.Contains(actual, `"_c":"中 文"`))
}

func TestMustString(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		out string
	}{
		{"Is string?", "Is string?"},
		{0, "0"},
		{0.005, "0.005"},
		{nil, ""},
		{true, "true"},
		{false, "false"},
		{[]byte("Fufu 中　文\u2728->?\n*\U0001F63A"), "Fufu 中　文\u2728->?\n*\U0001F63A"},
		{[]int{0, 2, 1}, "[0 2 1]"},
		{map[string]interface{}{"a": 0, "b": true, "C": []byte("c")}, "map[C:[99] a:0 b:true]"},
	} {
		AssertEqual(t, v.out, MustString(v.in))
	}
}

func TestMustInt(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		out int
	}{
		{"2", 2},
		{0b0010, 2},
		{10, 10},
		{0o77, 63},
		{0xff, 255},
		{-1, -1},
		{true, 1},
		{"0x", 0},
		{false, 0},
		{uint(11), 0},
		{1.005, 0},
		{nil, 0},
	} {
		AssertEqual(t, v.out, MustInt(v.in))
	}
}

func TestMustBool(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		out bool
	}{
		{"1", true},
		{"t", true},
		{"T", true},
		{"TRUE", true},
		{"true", true},
		{"True", true},
		{true, true},
		{1, true},
		{0x01, true},
		{false, false},
		{0, false},
		{"2", false},
		{nil, false},
		{"TrUe", false},
	} {
		AssertEqual(t, v.out, MustBool(v.in))
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	s := strings.Repeat("Fufu 中　文\u2728->?\n*\U0001F63A", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(s)
	}
}

func BenchmarkString2Bytes(b *testing.B) {
	s := strings.Repeat("Fufu 中　文\u2728->?\n*\U0001F63A", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = String2Bytes(s)
	}
}

func BenchmarkStr2Bytes(b *testing.B) {
	s := strings.Repeat("Fufu 中　文\u2728->?\n*\U0001F63A", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes(s)
	}
}

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat("Fufu 中　文\u2728->?\n*\U0001F63A", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = S2B(s)
	}
}

func BenchmarkStdStringToBytes(b *testing.B) {
	s := strings.Repeat("Fufu 中　文\u2728->?\n*\U0001F63A", 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

// BenchmarkStringToBytes-8                	1000000000	         0.379 ns/op	       0 B/op	       0 allocs/op
// BenchmarkString2Bytes-8                 	1000000000	         0.375 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStr2Bytes-8                    	1000000000	         0.301 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2B-8                          	1000000000	         0.345 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStdStringToBytes-8             	   28250	     41335 ns/op	  262144 B/op	       1 allocs/op
