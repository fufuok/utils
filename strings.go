package utils

import (
	"encoding/base64"
)

// 获取字符串结果, 可选指定默认值
func GetString(v interface{}, defaultVal ...string) string {
	s := MustString(v)
	if len(s) == 0 && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// Immutable, string to string
func CopyString(s string) string {
	tmp := make([]byte, len(s))
	copy(tmp, s)
	return B2S(tmp)
}

// 拼接字符串, 返回 bytes from bytes.Join()
func AddStringBytes(s ...string) []byte {
	switch len(s) {
	case 0:
		return []byte{}
	case 1:
		return []byte(s[0])
	}

	n := 0
	for _, v := range s {
		n += len(v)
	}

	b := make([]byte, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], v)
	}

	return b
}

// 拼接字符串
func AddString(s ...string) string {
	return B2S(AddStringBytes(s...))
}

// 搜索字符串位置(左, 第一个)
func SearchString(ss []string, s string) int {
	for i, v := range ss {
		if s == v {
			return i
		}
	}
	return -1
}

// 检查字符串是否存在于 slice
func InStrings(ss []string, s string) bool {
	return SearchString(ss, s) != -1
}

// Base64 编码
func B64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64 解码
func B64Decode(s string) []byte {
	if b, err := base64.StdEncoding.DecodeString(s); err == nil {
		return b
	}

	return nil
}

// Base64 解码, 安全 URL, 替换: "+/" 为 "-_"
func B64UrlEncode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

// Base64 解码
func B64UrlDecode(s string) []byte {
	if b, err := base64.URLEncoding.DecodeString(s); err == nil {
		return b
	}

	return nil
}
