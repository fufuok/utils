package utils

import (
	"strings"
	"unicode/utf8"
)

const (
	toLowerTable = "\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18" +
		"\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuvwxyz[\\]^_`" +
		"abcdefghijklmnopqrstuvwxyz{|}~\u007f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e" +
		"\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6" +
		"\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe" +
		"\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6" +
		"\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee" +
		"\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"
	toUpperTable = "\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18" +
		"\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ{|}~\u007f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e" +
		"\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6" +
		"\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe" +
		"\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6" +
		"\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee" +
		"\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"
)

// GetString 获取字符串结果, 可选指定默认值
func GetString(v interface{}, defaultVal ...string) string {
	s := MustString(v)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// GetSafeString Immutable, 可选指定默认值
func GetSafeString(s string, defaultVal ...string) string {
	s = CopyString(s)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// GetSafeB2S Immutable, 可选指定默认值
func GetSafeB2S(b []byte, defaultVal ...string) string {
	if len(b) == 0 {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return ""
	}
	return string(b)
}

// CopyString Immutable, string to string
// e.g. fiberParam := utils.CopyString(c.Params("test"))
// e.g. utils.CopyString(s[500:1000]) // 可以让 s 被 GC 回收
// strings.Clone(s) // go1.18
func CopyString(s string) string {
	if s == "" {
		return ""
	}
	tmp := make([]byte, len(s))
	copy(tmp, s)
	return B2S(tmp)
}

// CopyB2S Immutable, []byte to string
// string(b)
func CopyB2S(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return B2S(CopyBytes(b))
}

// JoinStringBytes 拼接字符串, 返回 bytes from bytes.Join()
func JoinStringBytes(s ...string) []byte {
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

// JoinString 拼接字符串
func JoinString(s ...string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	case 2:
		return s[0] + s[1]
	case 3:
		return s[0] + s[1] + s[2]
	default:
		return B2S(JoinStringBytes(s...))
	}
}

// SearchString 搜索字符串位置(左, 第一个)
// 用 slices.Index 替代
func SearchString(ss []string, s string) int {
	for i := range ss {
		if s == ss[i] {
			return i
		}
	}
	return -1
}

// InStrings 检查字符串是否存在于 slice
// 用 slices.Contains 替代
func InStrings(ss []string, s string) bool {
	return SearchString(ss, s) != -1
}

// RemoveString 删除字符串元素
func RemoveString(ss []string, s string) ([]string, bool) {
	for i := range ss {
		if s == ss[i] {
			return append(ss[:i], ss[i+1:]...), true
		}
	}
	return ss, false
}

// TrimSlice 清除 slice 中各元素的空白, 并删除空白项
// 注意: 原切片将被修改
func TrimSlice(ss []string) []string {
	if len(ss) == 0 {
		return ss
	}
	idx := 0
	for _, v := range ss {
		v := strings.TrimSpace(v)
		if v != "" {
			ss[idx] = v
			idx++
		}
	}
	return ss[:idx]
}

// ToLower converts ascii string to lower-case
// Ref: fiber
func ToLower(b string) string {
	res := make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toLowerTable[res[i]]
	}

	return B2S(res)
}

// ToUpper converts ascii string to upper-case
// Ref: fiber
func ToUpper(b string) string {
	res := make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toUpperTable[res[i]]
	}

	return B2S(res)
}

// TrimLeft is the equivalent of strings.TrimLeft
// Ref: fiber
func TrimLeft(s string, cutset byte) string {
	lenStr, start := len(s), 0
	for start < lenStr && s[start] == cutset {
		start++
	}
	return s[start:]
}

// TrimRight is the equivalent of strings.TrimRight
// Ref: fiber
func TrimRight(s string, cutset byte) string {
	lenStr := len(s)
	for lenStr > 0 && s[lenStr-1] == cutset {
		lenStr--
	}
	return s[:lenStr]
}

// Trim is the equivalent of strings.Trim
// Ref: fiber
func Trim(s string, cutset byte) string {
	i, j := 0, len(s)-1
	for ; i <= j; i++ {
		if s[i] != cutset {
			break
		}
	}
	for ; i < j; j-- {
		if s[j] != cutset {
			break
		}
	}

	return s[i : j+1]
}

// EqualFold tests ascii strings for equality case-insensitively
// Ref: fiber
func EqualFold(b, s string) bool {
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

// CutString xslices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
// Ref: go1.18
func CutString(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// RuneSubString 多字节字符串截取
// 长度支持 0(空字符串) / >0(正常截取) / <0(从结尾处往前截取)
// 长度等于或超过字符串长度, 返回原字符串值, 此时忽略 suffix 参数
func RuneSubString(s string, length int, suffix string) string {
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
		s = RuneReverse(s)
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
		return suffix + RuneReverse(r)
	}
	return r + suffix
}

// RuneReverse 反转多字节字符串
func RuneReverse(s string) string {
	var i, n, m int
	b := make([]byte, len(s))
	for m < len(s) {
		_, n = utf8.DecodeRuneInString(s[i:])
		m = i + n
		copy(b[len(b)-m:], s[i:m])
		i = m
	}
	return B2S(b)
}

// Reverse 反转字符串
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return B2S(b)
}

// IsNumeric 检查字符串是否全是 ASCII 数字: 0-9
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// IsLetter 检查字符串是否全是 ASCII 字母
func IsLetter(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
			return false
		}
	}
	return true
}

// IsLetterOrNumeric 检查字符串是否全是 ASCII 字母或数字
func IsLetterOrNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

// TruncStr 截断字符串
func TruncStr(s string, maxLen int, suffix string) string {
	if maxLen <= 0 || s == "" {
		return ""
	}
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	var i, n int
	for i = range s {
		if n == maxLen {
			break
		}
		n++
	}
	return s[:i] + suffix
}

// LastString 获取文本内容最后一个分隔符(单字节: sep[0])后的内容
func LastString(s, sep string, defaultStr ...string) string {
	idx := strings.LastIndex(s, sep)
	if idx != -1 && idx+1 < len(s) {
		return s[idx+1:]
	}
	if len(defaultStr) > 0 {
		return defaultStr[0]
	}
	return ""
}

// FirstString 获取文本内容第一个分隔符(单字节: sep[0])前的内容
func FirstString(s, sep string, defaultStr ...string) string {
	idx := strings.Index(s, sep)
	if idx != -1 {
		return s[:idx]
	}
	if len(defaultStr) > 0 {
		return defaultStr[0]
	}
	return ""
}
