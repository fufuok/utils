package assert

import (
	"unicode/utf8"
	"unsafe"
)

func runeSubString(s string, length int, suffix string) string {
	if s == "" || length == 0 {
		return ""
	}

	n := 0
	m := length
	if length < 0 {
		m = -m
	}
	for range s {
		n++
	}
	if m >= n {
		return s
	}

	if length < 0 {
		s = runeReverse(s)
	}

	n = 0
	i := 0
	for i = range s {
		if n == m {
			break
		}
		n++
	}
	r := s[:i]
	if length < 0 {
		return suffix + runeReverse(r)
	}
	return r + suffix
}

func runeReverse(s string) string {
	var i, n, m int
	b := make([]byte, len(s))
	for m < len(s) {
		_, n = utf8.DecodeRuneInString(s[i:])
		m = i + n
		copy(b[len(b)-m:], s[i:m])
		i = m
	}
	return *(*string)(unsafe.Pointer(&b))
}
