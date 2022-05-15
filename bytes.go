package utils

import (
	"bytes"
)

// GetBytes 先转为字符串再转为 []byte, 可选指定默认值
func GetBytes(v interface{}, defaultVal ...[]byte) []byte {
	switch b := v.(type) {
	default:
		bs := S2B(MustString(v))
		if len(bs) == 0 && len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return bs
	case []byte:
		return b
	}
}

// GetSafeBytes Immutable, 可选指定默认值
func GetSafeBytes(b []byte, defaultVal ...[]byte) []byte {
	b = CopyBytes(b)
	if len(b) == 0 && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return b
}

// GetSafeS2B Immutable, 可选指定默认值
func GetSafeS2B(s string, defaultVal ...[]byte) []byte {
	b := []byte(s)
	if len(b) == 0 && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return b
}

// CopyBytes Immutable, []byte to []byte
func CopyBytes(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

// CopyS2B Immutable, string to []byte
// []byte(s)
func CopyS2B(s string) []byte {
	tmp := make([]byte, len(s))
	copy(tmp, s)
	return tmp
}

// JoinBytes 拼接 []byte
func JoinBytes(b ...[]byte) []byte {
	return bytes.Join(b, nil)
}

// ToLowerBytes converts ascii slice to lower-case
// Ref: fiber
func ToLowerBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toLowerTable[b[i]]
	}
	return b
}

// ToUpperBytes converts ascii slice to upper-case
// Ref: fiber
func ToUpperBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toUpperTable[b[i]]
	}
	return b
}

// TrimLeftBytes is the equivalent of bytes.TrimLeft
// Ref: fiber
func TrimLeftBytes(b []byte, cutset byte) []byte {
	lenStr, start := len(b), 0
	for start < lenStr && b[start] == cutset {
		start++
	}
	return b[start:]
}

// TrimRightBytes is the equivalent of bytes.TrimRight
// Ref: fiber
func TrimRightBytes(b []byte, cutset byte) []byte {
	lenStr := len(b)
	for lenStr > 0 && b[lenStr-1] == cutset {
		lenStr--
	}
	return b[:lenStr]
}

// TrimBytes is the equivalent of bytes.Trim
// Ref: fiber
func TrimBytes(b []byte, cutset byte) []byte {
	i, j := 0, len(b)-1
	for ; i <= j; i++ {
		if b[i] != cutset {
			break
		}
	}
	for ; i < j; j-- {
		if b[j] != cutset {
			break
		}
	}

	return b[i : j+1]
}

// EqualFoldBytes tests ascii slices for equality case-insensitively
// Ref: fiber
func EqualFoldBytes(b, s []byte) bool {
	if len(b) != len(s) {
		return false
	}
	for i := len(b) - 1; i >= 0; i-- {
		if toUpperTable[b[i]] != toUpperTable[s[i]] {
			return false
		}
	}
	return true
}

// CutBytes slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, nil, false.
//
// Cut returns slices of the original slice s, not copies.
// Ref: go1.18
func CutBytes(s, sep []byte) (before, after []byte, found bool) {
	if i := bytes.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}
