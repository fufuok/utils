// Package jsongen forked from darjun/json-gen
package jsongen

import (
	"strconv"

	"github.com/fufuok/utils/pools/bytespool"
)

type Value interface {
	// Serialize 将值序列化为字符串, 追加到 buf 并返回
	Serialize(buf []byte) []byte
	// Size 返回值最终字节数
	Size() int
}

// V 表示标准的 JSON 值，例如: 123，1.23，true 等
// 字符串值是以双引号包裹的字符串, 如: "abc"
type V string

func (v V) Serialize(buf []byte) []byte {
	return append(buf, v...)
}

func (v V) Size() int {
	return len(v)
}

// RawString 附加原生 JSON 数据字符串形式, 如直接附加: `[1,{"A":1}]`
type RawString string

func (s RawString) Serialize(buf []byte) []byte {
	return append(buf, s...)
}

func (s RawString) Size() int {
	return len(s)
}

// RawBytes 附加原生 JSON 数据
type RawBytes []byte

func (b RawBytes) Serialize(buf []byte) []byte {
	return append(buf, b...)
}

func (b RawBytes) Size() int {
	return len(b)
}

// Array 数组类的 JSON 数据
type Array struct {
	values []Value
}

// NewArray 创建 JSON 数组, 无添加值时, 结果中至少附加一个空数组: []
func NewArray() *Array {
	return &Array{}
}

func (a *Array) Serialize(buf []byte) []byte {
	if cap(buf) == 0 {
		buf = bytespool.Make(a.Size())
	}

	buf = append(buf, '[')
	count := len(a.values)
	for i, e := range a.values {
		buf = e.Serialize(buf)
		if i != count-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, ']')
}

func (a *Array) Size() int {
	size := 0
	for _, e := range a.values {
		size += e.Size()
	}

	// for []
	size += 2

	if len(a.values) > 1 {
		// for ,
		size += len(a.values) - 1
	}
	return size
}

// AppendUint 追加单个或多个 uint64 到数组: [1,2] => [1,2,3,4]
func (a *Array) AppendUint(vv ...uint64) {
	for _, v := range vv {
		a.values = append(a.values, V(strconv.FormatUint(v, 10)))
	}
}

// AppendInt 追加单个或多个 int64 到数组: [1,2] => [1,2,3,4]
func (a *Array) AppendInt(vv ...int64) {
	for _, v := range vv {
		a.values = append(a.values, V(strconv.FormatInt(v, 10)))
	}
}

// AppendFloat 追加单个或多个 float64 到数组: [1,2] => [1,2,3.1,4]
func (a *Array) AppendFloat(vv ...float64) {
	for _, v := range vv {
		a.values = append(a.values, V(strconv.FormatFloat(v, 'f', -1, 64)))
	}
}

// AppendBool 追加单个或多个 bool 到数组: [1,2] => [1,2,true,false]
func (a *Array) AppendBool(vv ...bool) {
	for _, v := range vv {
		a.values = append(a.values, V(strconv.FormatBool(v)))
	}
}

// AppendString 追加单个或多个 string 到数组: [1,2] => [1,2,"A","b"]
// a.AppendString("A", "b")
func (a *Array) AppendString(vv ...string) {
	for _, v := range vv {
		a.values = append(a.values, V(EscapeString(v)))
	}
}

// AppendMap 追加单个或多个 map 到数组: [1,2] => [1,2,{"A":1},{"b":true}]
func (a *Array) AppendMap(vv ...*Map) {
	for _, v := range vv {
		if v != nil {
			a.values = append(a.values, v)
		}
	}
}

// AppendArray 追加单个或多个 array 到数组: [1,2] => [1,2,[{"A":1}],[true]]
func (a *Array) AppendArray(vv ...*Array) {
	for _, v := range vv {
		if v != nil {
			a.values = append(a.values, v)
		}
	}
}

// AppendRawString 追加单个或多个原生 JSON 字符串, 如: [1,2] => [1,2,[2,{"A":1}]]
// a.AppendRawString(`[2,{"A":1}]`)
func (a *Array) AppendRawString(ss ...string) {
	if len(ss) == 0 {
		return
	}
	vv := make([]Value, 0, len(ss))
	for _, v := range ss {
		if v != "" {
			vv = append(vv, RawString(v))
		}
	}
	a.values = append(a.values, vv...)
}

// AppendRawBytes 追加单个或多个原生 JSON 数据
func (a *Array) AppendRawBytes(bb ...[]byte) {
	if len(bb) == 0 {
		return
	}
	vv := make([]Value, 0, len(bb))
	for _, v := range bb {
		if v != nil {
			vv = append(vv, RawBytes(v))
		}
	}
	a.values = append(a.values, vv...)
}

// AppendUintArray 追加 uint64 数组: [1,2] => [1,2,[3,4,5]]
func (a *Array) AppendUintArray(vv []uint64) {
	sub := NewArray()
	sub.AppendUint(vv...)
	a.values = append(a.values, sub)
}

// AppendIntArray 追加 int64 数组: [1,2] => [1,2,[3,4,5]]
func (a *Array) AppendIntArray(vv []int64) {
	sub := NewArray()
	sub.AppendInt(vv...)
	a.values = append(a.values, sub)
}

// AppendFloatArray 追加 float64 数组: [1,2] => [1,2,[3,4.1,5]]
func (a *Array) AppendFloatArray(vv []float64) {
	sub := NewArray()
	sub.AppendFloat(vv...)
	a.values = append(a.values, sub)
}

// AppendBoolArray 追加 bool 数组: [1,2] => [1,2,[true,false]]
func (a *Array) AppendBoolArray(vv []bool) {
	sub := NewArray()
	sub.AppendBool(vv...)
	a.values = append(a.values, sub)
}

// AppendStringArray 追加 string 数组: [1,2] => [1,2,["A","b"]]
// a.AppendStringArray([]string{"A","b"})
func (a *Array) AppendStringArray(vv []string) {
	sub := NewArray()
	sub.AppendString(vv...)
	a.values = append(a.values, sub)
}

// AppendMapArray 追加 map 数组: [1,2] => [1,2,[{"A":1},{"b":true}]]
func (a *Array) AppendMapArray(vv []*Map) {
	sub := NewArray()
	sub.AppendMap(vv...)
	a.values = append(a.values, sub)
}

// AppendArrayArray 追加 array 数组: [1,2] => [1,2,[[3],[4],[{"b":true}]]]
func (a *Array) AppendArrayArray(vv []*Array) {
	sub := NewArray()
	sub.AppendArray(vv...)
	a.values = append(a.values, sub)
}

// AppendRawStringArray 追加原生 JSON 字符串数组: [1,2] => [1,2,["x",[4],[{"b":true}]]]
// a.AppendRawStringArray([]string{`"x"`, `[4]`, `[{"b":true}]`})
func (a *Array) AppendRawStringArray(ss []string) {
	sub := NewArray()
	sub.AppendRawString(ss...)
	a.values = append(a.values, sub)
}

// AppendRawBytesArray 追加原生 JSON 数据数组
func (a *Array) AppendRawBytesArray(vv [][]byte) {
	sub := NewArray()
	sub.AppendRawBytes(vv...)
	a.values = append(a.values, sub)
}

// Map 对象类(字典) JSON 数据
type Map struct {
	keys   []string
	values []Value
}

// NewMap 创建对象类(字典) JSON 数据集, 无添加值时, 结果中至少附加一个空对象: {}
func NewMap() *Map {
	return &Map{
		keys:   make([]string, 0, 8),
		values: make([]Value, 0, 8),
	}
}

func (m *Map) Serialize(buf []byte) []byte {
	if cap(buf) == 0 {
		buf = bytespool.Make(m.Size())
	}

	buf = append(buf, '{')
	count := len(m.keys)
	for i, key := range m.keys {
		buf = append(buf, key...)
		buf = append(buf, ':')
		buf = m.values[i].Serialize(buf)
		if i != count-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, '}')
}

func (m *Map) Size() int {
	size := 0
	for i, key := range m.keys {
		// +1 for :
		size += len(key) + 1
		size += m.values[i].Size()
	}

	// +2 for {}
	size += 2

	if len(m.keys) > 1 {
		// for ,
		size += len(m.keys) - 1
	}
	return size
}

func (m *Map) put(key string, value Value) {
	m.keys = append(m.keys, EscapeString(key))
	m.values = append(m.values, value)
}

func (m *Map) PutRawString(key, s string) {
	m.put(key, RawString(s))
}

func (m *Map) PutRawBytes(key string, b []byte) {
	m.put(key, RawBytes(b))
}

func (m *Map) PutUint(key string, u uint64) {
	m.put(key, V(strconv.FormatUint(u, 10)))
}

func (m *Map) PutInt(key string, i int64) {
	m.put(key, V(strconv.FormatInt(i, 10)))
}

func (m *Map) PutFloat(key string, f float64) {
	m.put(key, V(strconv.FormatFloat(f, 'f', -1, 64)))
}

func (m *Map) PutBool(key string, b bool) {
	m.put(key, V(strconv.FormatBool(b)))
}

func (m *Map) PutString(key, s string) {
	m.put(key, V(EscapeString(s)))
}

func (m *Map) PutRawStringArray(key string, ss []string) {
	a := NewArray()
	a.AppendRawString(ss...)
	m.put(key, a)
}

func (m *Map) PutRawBytesArray(key string, bs [][]byte) {
	a := NewArray()
	a.AppendRawBytes(bs...)
	m.put(key, a)
}

// PutUintArray 添加 uint64 数组数据项: {"A":[1,2]}
func (m *Map) PutUintArray(key string, u []uint64) {
	a := NewArray()
	a.AppendUint(u...)
	m.put(key, a)
}

func (m *Map) PutIntArray(key string, i []int64) {
	a := NewArray()
	a.AppendInt(i...)
	m.put(key, a)
}

func (m *Map) PutFloatArray(key string, f []float64) {
	a := NewArray()
	a.AppendFloat(f...)
	m.put(key, a)
}

func (m *Map) PutBoolArray(key string, b []bool) {
	a := NewArray()
	a.AppendBool(b...)
	m.put(key, a)
}

func (m *Map) PutStringArray(key string, s []string) {
	a := NewArray()
	a.AppendString(s...)
	m.put(key, a)
}

// PutArray 添加值为数组的数据项: {"A":[1,true,"x"]}
func (m *Map) PutArray(key string, oa *Array) {
	m.put(key, oa)
}

// PutMap 添加值为对象(字典)的数据项, map 嵌套: {"A":{"sub":1}}
func (m *Map) PutMap(key string, om *Map) {
	m.put(key, om)
}
