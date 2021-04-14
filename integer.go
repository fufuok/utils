package utils

// 整数取大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 整数取小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 获取 int 结果, 可选指定默认值(若给定了默认值,则返回正整数或 0)
func GetInt(v interface{}, defaultInt ...int) int {
	i := MustInt(v)
	if i <= 0 && len(defaultInt) > 0 {
		return defaultInt[0]
	}
	return i
}

// 搜索整数位置(左, 第一个)
func SearchInt(slice []int, n int) int {
	for i, v := range slice {
		if n == v {
			return i
		}
	}

	return -1
}

// 检查整数是否存在于 slice
func InInts(slice []int, n int) bool {
	return SearchInt(slice, n) != -1
}
