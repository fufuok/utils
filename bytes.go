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
