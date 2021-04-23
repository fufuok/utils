package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/fufuok/utils/json"
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

// Ref: fasthttp
func String2Bytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// Ref: csdn.weixin_43705457
func Str2Bytes(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s                                                  // 把s的地址付给b
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s) // 修改容量为长度
	return
}

// Ref: csdn.u010853261
func StringToBytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&sliceHeader{
		data: (*stringHeader)(unsafe.Pointer(&s)).data,
		len:  len(s),
		cap:  len(s),
	}))
}

// StringToBytes converts string to byte slice without a memory allocation.
// Ref: gin
func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString
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
