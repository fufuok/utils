package utils

import (
	"testing"
)

func TestS2B(t *testing.T) {
	t.Parallel()
	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := S2B(expected)

	AssertEqual(t, []byte(expected), actual)
}

func TestB2S(t *testing.T) {
	t.Parallel()
	expected := "Fufu 中　文\u2728->?\n*\U0001F63A"
	actual := B2S([]byte(expected))

	AssertEqual(t, expected, actual)
}

func TestMustJSONString(t *testing.T) {
	t.Parallel()
	expected := `{"_c":"中 文","a":true,"b":1.23}`
	actual := MustJSONString(map[string]interface{}{
		"_c": "中 文",
		"a":  true,
		"b":  1.23,
	})

	AssertEqual(t, expected, actual)
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
