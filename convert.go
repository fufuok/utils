package utils

import (
	"fmt"
	"strconv"
	"unsafe"

	json "github.com/json-iterator/go"
)

// 更安全的 reflect.StringHeader
type stringHeader struct {
	data unsafe.Pointer
	len  int
}

// 更安全的 reflect.SliceHeader
type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

// string 转为 []byte
func S2B(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&sliceHeader{
		data: (*stringHeader)(unsafe.Pointer(&s)).data,
		len:  len(s),
		cap:  len(s),
	}))
}

// []byte 转为 string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 转 JSON 返回 []byte
func MustJSON(v interface{}) []byte {
	js, _ := json.Marshal(&v)
	return js
}

// 转 JSON 返回 string
func MustJSONString(v interface{}) string {
	return B2S(MustJSON(v))
}

// 强制转为字符串
func MustString(v interface{}) string {
	switch s := v.(type) {
	default:
		return fmt.Sprintf("%v", v)
	case string:
		return s
	case []byte:
		return B2S(s)
	case nil:
		return ""
	}
}

// 强制转为整数 (int)
func MustInt(v interface{}) int {
	switch i := v.(type) {
	default:
		d, ok := i.(int)
		if ok {
			return d
		}
		return 0
	case string:
		v, err := strconv.ParseInt(i, 0, 0)
		if err == nil {
			return int(v)
		}
		return 0
	case bool:
		if i {
			return 1
		}
		return 0
	case nil:
		return 0
	}
}

// 强制转为 bool
func MustBool(v interface{}) bool {
	switch t := v.(type) {
	default:
		if MustInt(v) == 1 {
			return true
		}
	case string:
		switch t {
		case "1", "t", "T", "true", "TRUE", "True":
			return true
		}
	}

	return false
}
