package utils

import (
	"bytes"
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
		{[]byte("Fufu 中　文\u2728->?\n*\U0001F63A"), nil, []byte("Fufu 中　文\u2728->?\n*\U0001F63A")},
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
	AssertEqual(t, []byte("Fufu 中　文\u2728->?\n*\U0001F63A"), CopyBytes([]byte("Fufu 中　文\u2728->?\n*\U0001F63A")))
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
