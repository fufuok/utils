package utils

import (
	"strings"
	"testing"
	"time"

	"github.com/fufuok/utils/assert"
)

func TestStringToBytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := StringToBytes(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := StringToBytes(expected)
	assert.Equal(t, []byte(expected), actual)
}

func TestString2Bytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := String2Bytes(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := String2Bytes(expected)
	assert.Equal(t, []byte(expected), actual)
}

func TestStr2Bytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := Str2Bytes(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := Str2Bytes(expected)
	assert.Equal(t, []byte(expected), actual)
}

func TestStrToBytes(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := StrToBytes(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := StrToBytes(expected)
	assert.Equal(t, []byte(expected), actual)
}

func TestS2B(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := S2B(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := S2B(expected)
	assert.Equal(t, []byte(expected), actual)

	assert.Equal(t, true, S2B("") == nil)
	assert.Equal(t, testBytes, S2B(testString))
}

func TestB2S(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		b := RandBytes(64)
		assert.Equal(t, string(b), B2S(b))
	}

	expected := testString
	actual := B2S([]byte(expected))
	assert.Equal(t, expected, actual)

	assert.Equal(t, true, B2S(nil) == "")
	assert.Equal(t, testString, B2S(testBytes))
}

func TestMustJSONString(t *testing.T) {
	t.Parallel()
	js := map[string]interface{}{
		"_c": "中 文",
		"a":  true,
		"b":  1.23,
	}
	actual := MustJSONString(&js)

	assert.Equal(t, true, strings.Contains(actual, `"a":true`))
	assert.Equal(t, true, strings.Contains(actual, `"b":1.23`))
	assert.Equal(t, true, strings.Contains(actual, `"_c":"中 文"`))

	actualIndent := MustJSONIndentString(&js)
	assert.Equal(t, true, strings.Contains(actualIndent, "  "))
}

func TestMustString(t *testing.T) {
	now := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)
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
		{[]byte(testString), testString},
		{[]int{0, 2, 1}, "[0 2 1]"},
		{map[string]interface{}{"a": 0, "b": true, "C": []byte("c")}, "map[C:[99] a:0 b:true]"},
		{now, "2022-01-02 03:04:05"},
		{&Bool{}, "false"},
	} {
		assert.Equal(t, v.out, MustString(v.in))
	}
	assert.Equal(t, "2022-01-02T03:04:05Z", MustString(now, time.RFC3339))
}

func TestMustInt(t *testing.T) {
	for _, v := range []struct {
		in  interface{}
		out int
	}{
		{"2", 2},
		{"  2 \n ", 2},
		{0b0010, 2},
		{10, 10},
		{0o77, 63},
		{0xff, 255},
		{-1, -1},
		{true, 1},
		{"0x", 0},
		{false, 0},
		{uint(11), 11},
		{uint64(11), 11},
		{int64(11), 11},
		{float32(11.0), 11},
		{1.005, 1},
		{nil, 0},
	} {
		assert.Equal(t, v.out, MustInt(v.in))
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
		{2, true},
		{2.1, true},
		{0x01, true},
		{false, false},
		{0.1, false},
		{0, false},
		{"2", false},
		{nil, false},
		{"TrUe", false},
	} {
		assert.Equal(t, v.out, MustBool(v.in))
	}
}

func TestB64Encode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "6Kej56CBL+e8lueggX4g6aG25pu/JiM=", B64Encode(S2B("解码/编码~ 顶替&#")))
}

func TestB64UrlEncode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "6Kej56CBL-e8lueggX4g6aG25pu_JiM=", B64UrlEncode(S2B("解码/编码~ 顶替&#")))
}

func TestB64Decode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []byte("解码/编码~ 顶替&#"), B64Decode("6Kej56CBL+e8lueggX4g6aG25pu/JiM="))
}

func TestB64UrlDecode(t *testing.T) {
	for _, v := range []struct {
		in  string
		out []byte
	}{
		{"6Kej56CBL-e8lueggX4g6aG25pu_JiM=", []byte("解码/编码~ 顶替&#")},
		{"123", nil},
	} {
		assert.Equal(t, v.out, B64UrlDecode(v.in))
	}
}

func BenchmarkS2BStringToBytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(s)
	}
}

func BenchmarkS2BString2Bytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = String2Bytes(s)
	}
}

func BenchmarkS2BStr2Bytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes(s)
	}
}

func BenchmarkS2BStrToBytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StrToBytes(s)
	}
}

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = S2B(s)
	}
}

func BenchmarkS2BStdStringToBytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

// BenchmarkS2BStringToBytes-8      	1000000000	         0.378 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2BString2Bytes-8       	1000000000	         0.325 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2BStr2Bytes-8          	1000000000	         0.334 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2BStrToBytes-8         	1000000000	         0.421 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2B-8                   	1000000000	         0.325 ns/op	       0 B/op	       0 allocs/op
// BenchmarkS2BStdStringToBytes-8   	   27127	     43473 ns/op	  262144 B/op	       1 allocs/op
