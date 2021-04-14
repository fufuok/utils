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
	AssertEqual(t, *new([]byte), GetBytes(nil))
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
