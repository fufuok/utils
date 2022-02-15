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

// ToLowerBytes is the equivalent of bytes.ToLower
// Ref: fiber
func ToLowerBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toLowerTable[b[i]]
	}
	return b
}

// ToUpperBytes is the equivalent of bytes.ToUpper
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

// EqualFoldBytes the equivalent of bytes.EqualFold
// Ref: fiber
func EqualFoldBytes(b, s []byte) (equals bool) {
	n := len(b)
	equals = n == len(s)
	if equals {
		for i := 0; i < n; i++ {
			if equals = b[i]|0x20 == s[i]|0x20; !equals {
				break
			}
		}
	}
	return
}
