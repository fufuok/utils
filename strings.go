package utils

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
	s := string(b)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// CopyString Immutable, string to string
// e.g. fiberParam := utils.CopyString(c.Params("test"))
// e.g. utils.CopyString(s[500:1000]) // 可以让 s 被 GC 回收
func CopyString(s string) string {
	return B2S([]byte(s))
}

// CopyB2S Immutable, []byte to string
func CopyB2S(b []byte) string {
	// string(b)
	return B2S(CopyBytes(b))
}

// AddStringBytes 拼接字符串, 返回 bytes from bytes.Join()
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

// AddString 拼接字符串
func AddString(s ...string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	default:
		return B2S(AddStringBytes(s...))
	}
}

// SearchString 搜索字符串位置(左, 第一个)
func SearchString(ss []string, s string) int {
	for i := range ss {
		if s == ss[i] {
			return i
		}
	}
	return -1
}

// InStrings 检查字符串是否存在于 slice
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
