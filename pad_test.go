package utils

import (
	"fmt"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestPad(t *testing.T) {
	src := "test"
	tests := []struct {
		pad, want string
		n         int
	}{
		{"0", src, 3},
		{"0", "0" + src, 5},
		{"0", "0" + src + "0", 6},
		{"0", "00" + src + "0", 7},
		{"0", "00" + src + "00", 8},

		{"123", "1" + src, 5},
		{"123", "1" + src + "1", 6},
		{"123", "12" + src + "1", 7},
		{"123", "12" + src + "12", 8},
		{"123", "123" + src + "12", 9},
		{"123", "1231" + src + "123", 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, Pad(src, v.pad, v.n))
	}
}

func TestLeftPad(t *testing.T) {
	src := "test"
	tests := []struct {
		pad, want string
		n         int
	}{
		{"0", src, 3},
		{"0", "0" + src, 5},
		{"0", "00" + src, 6},
		{"0", "000" + src, 7},

		{"123", "1" + src, 5},
		{"123", "12" + src, 6},
		{"123", "123" + src, 7},
		{"123", "1231" + src, 8},
		{"123", "1231231" + src, 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, LeftPad(src, v.pad, v.n))
	}
	assert.Equal(t, fmt.Sprintf("%32s", src), LeftPad(src, " ", 32))
	assert.Equal(t, fmt.Sprintf("%032s", "111"), LeftPad("111", "0", 32))
}

func TestRightPad(t *testing.T) {
	src := "test"
	tests := []struct {
		pad, want string
		n         int
	}{
		{"0", src, 3},
		{"0", src + "0", 5},
		{"0", src + "00", 6},
		{"0", src + "000", 7},

		{"123", src + "1", 5},
		{"123", src + "12", 6},
		{"123", src + "123", 7},
		{"123", src + "1231", 8},
		{"123", src + "1231231", 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, RightPad(src, v.pad, v.n))
	}
	assert.Equal(t, fmt.Sprintf("%-32s", src), RightPad(src, " ", 32))
	assert.Equal(t, fmt.Sprintf("%s%032s", "111", "")[:32], RightPad("111", "0", 32))
}

func TestPadBytes(t *testing.T) {
	src := []byte("test")
	tests := []struct {
		pad, want []byte
		n         int
	}{
		{[]byte("0"), src, 3},
		{[]byte("0"), JoinBytes([]byte("0"), src), 5},
		{[]byte("0"), JoinBytes([]byte("0"), src, []byte("0")), 6},
		{[]byte("0"), JoinBytes([]byte("00"), src, []byte("0")), 7},
		{[]byte("0"), JoinBytes([]byte("00"), src, []byte("00")), 8},

		{[]byte("123"), JoinBytes([]byte("1"), src), 5},
		{[]byte("123"), JoinBytes([]byte("1"), src, []byte("1")), 6},
		{[]byte("123"), JoinBytes([]byte("12"), src, []byte("1")), 7},
		{[]byte("123"), JoinBytes([]byte("12"), src, []byte("12")), 8},
		{[]byte("123"), JoinBytes([]byte("123"), src, []byte("12")), 9},
		{[]byte("123"), JoinBytes([]byte("1231"), src, []byte("123")), 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, PadBytes(src, v.pad, v.n))
	}
}

func TestLeftPadBytes(t *testing.T) {
	src := []byte("test")
	tests := []struct {
		pad, want []byte
		n         int
	}{
		{[]byte("0"), src, 3},
		{[]byte("0"), JoinBytes([]byte("0"), src), 5},
		{[]byte("0"), JoinBytes([]byte("00"), src), 6},
		{[]byte("0"), JoinBytes([]byte("000"), src), 7},

		{[]byte("123"), JoinBytes([]byte("1"), src), 5},
		{[]byte("123"), JoinBytes([]byte("12"), src), 6},
		{[]byte("123"), JoinBytes([]byte("123"), src), 7},
		{[]byte("123"), JoinBytes([]byte("1231"), src), 8},
		{[]byte("123"), JoinBytes([]byte("1231231"), src), 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, LeftPadBytes(src, v.pad, v.n))
	}
}

func TestRightPadBytes(t *testing.T) {
	src := []byte("test")
	tests := []struct {
		pad, want []byte
		n         int
	}{
		{[]byte("0"), src, 3},
		{[]byte("0"), JoinBytes(src, []byte("0")), 5},
		{[]byte("0"), JoinBytes(src, []byte("00")), 6},
		{[]byte("0"), JoinBytes(src, []byte("000")), 7},

		{[]byte("123"), JoinBytes(src, []byte("1")), 5},
		{[]byte("123"), JoinBytes(src, []byte("12")), 6},
		{[]byte("123"), JoinBytes(src, []byte("123")), 7},
		{[]byte("123"), JoinBytes(src, []byte("1231")), 8},
		{[]byte("123"), JoinBytes(src, []byte("1231231")), 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.want, RightPadBytes(src, v.pad, v.n))
	}
}

func BenchmarkPad(b *testing.B) {
	s := "111"
	b.Run("LeftPad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = LeftPad(s, "0", 32)
		}
	})
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%032s", s)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=^BenchmarkPad$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkPad/LeftPad-8          15095967                82.43 ns/op           32 B/op          1 allocs/op
// BenchmarkPad/LeftPad-8          15392232                81.87 ns/op           32 B/op          1 allocs/op
// BenchmarkPad/LeftPad-8          14027343                85.44 ns/op           32 B/op          1 allocs/op
// BenchmarkPad/fmt.Sprintf-8       4820535                289.7 ns/op           48 B/op          2 allocs/op
// BenchmarkPad/fmt.Sprintf-8       4064318                327.1 ns/op           48 B/op          2 allocs/op
// BenchmarkPad/fmt.Sprintf-8       4150326                279.9 ns/op           48 B/op          2 allocs/op
