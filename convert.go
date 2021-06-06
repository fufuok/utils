package utils

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

// String2Bytes Ref: fasthttp
func String2Bytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// StringToBytes Ref: csdn.u010853261
func StringToBytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&sliceHeader{
		data: (*stringHeader)(unsafe.Pointer(&s)).data,
		len:  len(s),
		cap:  len(s),
	}))
}

// Str2Bytes Ref: csdn.weixin_43705457
func Str2Bytes(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

// StrToBytes Ref: Allenxuxu / toolkit
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// S2B StringToBytes converts string to byte slice without a memory allocation.
// Ref: gin
func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// B2S BytesToString
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// MustJSONIndent 转 Jsoniter 返回 []byte
func MustJSONIndent(v interface{}) []byte {
	js, _ := json.MarshalIndent(&v, "", "  ")
	return js
}

// MustJSONIndentString 转 Jsoniter Indent 返回 string
func MustJSONIndentString(v interface{}) string {
	return B2S(MustJSONIndent(v))
}

// MustJSON 转 Jsoniter 返回 []byte
func MustJSON(v interface{}) []byte {
	js, _ := json.Marshal(&v)
	return js
}

// MustJSONString 转 Jsoniter 返回 string
func MustJSONString(v interface{}) string {
	return B2S(MustJSON(v))
}

// MustString 强制转为字符串
func MustString(v interface{}) string {
	switch s := v.(type) {
	default:
		return fmt.Sprint(v)
	case string:
		return s
	case []byte:
		return B2S(s)
	case error:
		return s.Error()
	case nil:
		return ""
	case bool:
		return strconv.FormatBool(s)
	case int:
		return strconv.Itoa(s)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int32:
		return strconv.Itoa(int(s))
	case int64:
		return strconv.FormatInt(s, 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(s, 10)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	}
}

// MustInt 强制转为整数 (int)
func MustInt(v interface{}) int {
	switch i := v.(type) {
	default:
		d, ok := i.(int)
		if ok {
			return d
		}
		return 0
	case string:
		v, err := strconv.Atoi(strings.TrimSpace(i))
		if err == nil {
			return v
		}
		return 0
	case bool:
		if i {
			return 1
		}
		return 0
	case nil:
		return 0
	case int:
		return i
	case int8:
		return int(i)
	case int16:
		return int(i)
	case int32:
		return int(i)
	case int64:
		return int(i)
	case uint:
		return int(i)
	case uint8:
		return int(i)
	case uint16:
		return int(i)
	case uint32:
		return int(i)
	case uint64:
		return int(i)
	case float32:
		return int(i)
	case float64:
		return int(i)
	}
}

// MustBool 强制转为 bool
func MustBool(v interface{}) bool {
	switch t := v.(type) {
	default:
		if MustInt(v) != 0 {
			return true
		}
	case bool:
		return t
	case string:
		switch t {
		case "1", "t", "T", "true", "TRUE", "True":
			return true
		}
	}

	return false
}

// B64Encode Base64 编码
func B64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// B64Decode Base64 解码
func B64Decode(s string) []byte {
	if b, err := base64.StdEncoding.DecodeString(s); err == nil {
		return b
	}

	return nil
}

// B64UrlEncode Base64 解码, 安全 URL, 替换: "+/" 为 "-_"
func B64UrlEncode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

// B64UrlDecode Base64 解码
func B64UrlDecode(s string) []byte {
	if b, err := base64.URLEncoding.DecodeString(s); err == nil {
		return b
	}

	return nil
}
