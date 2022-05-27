// Package jsongen forked from darjun/json-gen
package jsongen

import (
	"encoding/json"
	"strconv"
	"unsafe"
)

// Value 表示将要序列化到`json`字符串中的值
type Value interface {
	// Serialize 将值序列化为字符串，追加到`buf`后，返回新的`buf`
	Serialize(buf []byte) []byte
	// Size 返回值在最终的`json`串中占有多少字节
	Size() int
}

// QuotedValue 表示需要用"包裹起来的值，例如字符串
type QuotedValue string

// Serialize 将`q`序列化为字符串，追加到`buf`后，返回新的`buf`
func (q QuotedValue) Serialize(buf []byte) []byte {
	buf = append(buf, '"')
	buf = append(buf, q...)
	return append(buf, '"')
}

// Size 返回`q`在最终的`json`串中占有多少字节
func (q QuotedValue) Size() int {
	return len(q) + 2
}

// UnquotedValue 表示不需要用"包裹起来的值，例如整数，浮点数等
type UnquotedValue string

// Serialize 将`u`序列化为字符串，追加到`buf`后，返回新的`buf`
func (u UnquotedValue) Serialize(buf []byte) []byte {
	return append(buf, u...)
}

// Size 返回`u`在最终的`json`串中占有多少字节
func (u UnquotedValue) Size() int {
	return len(u)
}

type RawBytes []byte

func (b RawBytes) Serialize(buf []byte) []byte {
	return append(buf, b...)
}

func (b RawBytes) Size() int {
	return len(b)
}

type RawString string

func (s RawString) Serialize(buf []byte) []byte {
	return append(buf, s...)
}

func (s RawString) Size() int {
	return len(s)
}

// Array 表示一个`json`数组
type Array []Value

// NewArray 创建一个`json`数组，返回其指针
func NewArray() *Array {
	a := Array(make([]Value, 0, 1))
	return &a
}

// Serialize 将`a`序列化为字符串，追加到`buf`后，返回新的`buf`
func (a Array) Serialize(buf []byte) []byte {
	if len(buf) == 0 {
		buf = make([]byte, 0, a.Size())
	}

	buf = append(buf, '[')
	count := len(a)
	for i, e := range a {
		buf = e.Serialize(buf)
		if i != count-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, ']')
}

// Size 返回`a`在最终的`json`串中占有多少字节
func (a Array) Size() int {
	size := 0
	for _, e := range a {
		size += e.Size()
	}

	// for []
	size += 2

	if len(a) > 1 {
		// for ,
		size += len(a) - 1
	}
	return size
}

func (a *Array) AppendRawString(s string) {
	*a = append(*a, RawString(s))
}

func (a *Array) AppendRawBytes(b []byte) {
	*a = append(*a, RawBytes(b))
}

func (a *Array) AppendRawStringArray(ss []string) {
	value := make([]Value, 0, len(ss))
	for _, v := range ss {
		value = append(value, RawString(v))
	}
	*a = append(*a, Array(value))
}

func (a *Array) AppendRawBytesArray(bs [][]byte) {
	value := make([]Value, 0, len(bs))
	for _, v := range bs {
		value = append(value, RawBytes(v))
	}
	*a = append(*a, Array(value))
}

// AppendUint 将`uint64`类型的值`u`追加到数组`a`后
func (a *Array) AppendUint(u uint64) {
	value := strconv.FormatUint(u, 10)
	*a = append(*a, UnquotedValue(value))
}

// AppendInt 将`int64`类型的值`i`追加到数组`a`后
func (a *Array) AppendInt(i int64) {
	value := strconv.FormatInt(i, 10)
	*a = append(*a, UnquotedValue(value))
}

// AppendFloat 将`float64`类型的值`f`追加到数组`a`后
func (a *Array) AppendFloat(f float64) {
	value := strconv.FormatFloat(f, 'g', 10, 64)
	*a = append(*a, UnquotedValue(value))
}

// AppendBool 将`bool`类型的值`b`追加到数组`a`后
func (a *Array) AppendBool(b bool) {
	value := strconv.FormatBool(b)
	*a = append(*a, UnquotedValue(value))
}

// AppendString 将`string`类型的值`s`追加到数组`a`后
func (a *Array) AppendString(value string) {
	*a = append(*a, EscapeString(value))
}

// AppendMap 将`Map`类型的值`m`追加到数组`a`后
func (a *Array) AppendMap(m *Map) {
	*a = append(*a, m)
}

// AppendUintArray 将`uint64`数组`u`追加到数组`a`后
func (a *Array) AppendUintArray(u []uint64) {
	value := make([]Value, 0, len(u))
	for _, v := range u {
		value = append(value, UnquotedValue(strconv.FormatUint(v, 10)))
	}
	*a = append(*a, Array(value))
}

// AppendIntArray 将`int64`数组`i`追加到数组`a`后
func (a *Array) AppendIntArray(i []int64) {
	value := make([]Value, 0, len(i))
	for _, v := range i {
		value = append(value, UnquotedValue(strconv.FormatInt(v, 10)))
	}
	*a = append(*a, Array(value))
}

// AppendFloatArray 将`float64`数组`f`追加到数组`a`后
func (a *Array) AppendFloatArray(f []float64) {
	value := make([]Value, 0, len(f))
	for _, v := range f {
		value = append(value, UnquotedValue(strconv.FormatFloat(v, 'g', 10, 64)))
	}
	*a = append(*a, Array(value))
}

// AppendBoolArray 将`bool`数组`b`追加到数组`a`后
func (a *Array) AppendBoolArray(b []bool) {
	value := make([]Value, 0, len(b))
	for _, v := range b {
		value = append(value, UnquotedValue(strconv.FormatBool(v)))
	}
	*a = append(*a, Array(value))
}

// AppendStringArray 将`string`数组`s`追加到数组`a`后
func (a *Array) AppendStringArray(s []string) {
	value := make([]Value, 0, len(s))
	for _, v := range s {
		value = append(value, EscapeString(v))
	}
	*a = append(*a, Array(value))
}

// AppendMapArray 将`Map`数组`m`追加到数组`a`后
func (a *Array) AppendMapArray(m []Map) {
	value := make([]Value, 0, len(m))
	for _, v := range m {
		value = append(value, v)
	}
	*a = append(*a, Array(value))
}

// AppendArray 将`json`数组`oa`追加到数组`a`后
func (a *Array) AppendArray(oa Array) {
	*a = append(*a, oa)
}

// Map 表示一个`json`映射
type Map struct {
	keys   []string
	values []Value
}

// Serialize 将`m`序列化为字符串，追加到`buf`后，返回新的`buf`
func (m Map) Serialize(buf []byte) []byte {
	if len(buf) == 0 {
		buf = make([]byte, 0, m.Size())
	}

	buf = append(buf, '{')
	count := len(m.keys)
	for i, key := range m.keys {
		buf = append(buf, '"')
		buf = append(buf, key...)
		buf = append(buf, '"')
		buf = append(buf, ':')
		buf = m.values[i].Serialize(buf)
		if i != count-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, '}')
}

// Size 返回`m`在最终的`json`串中占有多少字节
func (m Map) Size() int {
	size := 0
	for i, key := range m.keys {
		// +2 for ", +1 for :
		size += len(key) + 2 + 1
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
	m.keys = append(m.keys, key)
	m.values = append(m.values, value)
}

func (m *Map) PutRawString(key, s string) {
	m.put(key, RawString(s))
}

func (m *Map) PutRawBytes(key string, b []byte) {
	m.put(key, RawBytes(b))
}

func (m *Map) PutRawStringArray(key string, ss []string) {
	value := make([]Value, 0, len(ss))
	for _, v := range ss {
		value = append(value, RawString(v))
	}
	m.put(key, Array(value))
}

func (m *Map) PutRawBytesArray(key string, bs [][]byte) {
	value := make([]Value, 0, len(bs))
	for _, v := range bs {
		value = append(value, RawBytes(v))
	}
	m.put(key, Array(value))
}

// PutUint 将`uint64`类型的值`u`与键`key`关联
func (m *Map) PutUint(key string, u uint64) {
	value := strconv.FormatUint(u, 10)
	m.put(key, UnquotedValue(value))
}

// PutInt 将`int64`类型的值`i`与键`key`关联
func (m *Map) PutInt(key string, i int64) {
	value := strconv.FormatInt(i, 10)
	m.put(key, UnquotedValue(value))
}

// PutFloat 将`float64`类型的值`f`与键`key`关联
func (m *Map) PutFloat(key string, f float64) {
	value := strconv.FormatFloat(f, 'g', 10, 64)
	m.put(key, UnquotedValue(value))
}

// PutBool 将`bool`类型的值`b`与键`key`关联
func (m *Map) PutBool(key string, b bool) {
	value := strconv.FormatBool(b)
	m.put(key, UnquotedValue(value))
}

// PutString 将`string`类型的值`value`与键`key`关联
func (m *Map) PutString(key, value string) {
	m.put(key, EscapeString(value))
}

// PutUintArray 将`uint64`数组类型的值`u`与键`key`关联
func (m *Map) PutUintArray(key string, u []uint64) {
	value := make([]Value, 0, len(u))
	for _, v := range u {
		value = append(value, UnquotedValue(strconv.FormatUint(v, 10)))
	}
	m.put(key, Array(value))
}

// PutIntArray 将`int64`数组类型的值`i`与键`key`关联
func (m *Map) PutIntArray(key string, i []int64) {
	value := make([]Value, 0, len(i))
	for _, v := range i {
		value = append(value, UnquotedValue(strconv.FormatInt(v, 10)))
	}
	m.put(key, Array(value))
}

// PutFloatArray 将`float64`数组类型的值`f`与键`key`关联
func (m *Map) PutFloatArray(key string, f []float64) {
	value := make([]Value, 0, len(f))
	for _, v := range f {
		value = append(value, UnquotedValue(strconv.FormatFloat(v, 'g', 10, 64)))
	}
	m.put(key, Array(value))
}

// PutBoolArray 将`bool`数组类型的值`b`与键`key`关联
func (m *Map) PutBoolArray(key string, b []bool) {
	value := make([]Value, 0, len(b))
	for _, v := range b {
		value = append(value, UnquotedValue(strconv.FormatBool(v)))
	}
	m.put(key, Array(value))
}

// PutStringArray 将`string`数组类型的值`s`与键`key`关联
func (m *Map) PutStringArray(key string, s []string) {
	value := make([]Value, 0, len(s))
	for _, v := range s {
		value = append(value, EscapeString(v))
	}
	m.put(key, Array(value))
}

// PutArray 将`json`数组`a`与键`key`关联
func (m *Map) PutArray(key string, a *Array) {
	m.put(key, a)
}

// PutMap 将`json`映射`om`与键`key`关联
func (m *Map) PutMap(key string, om *Map) {
	m.put(key, om)
}

// NewMap 创建一个`json`映射返回其指针
func NewMap() *Map {
	return &Map{
		keys:   make([]string, 0, 8),
		values: make([]Value, 0, 8),
	}
}

func EscapeString(s string) Value {
	for i := 0; i < len(s); i++ {
		if s[i] == '"' || s[i] == '\\' || s[i] < ' ' || s[i] > 0x7f {
			b, _ := json.Marshal(s)
			return UnquotedValue(*(*string)(unsafe.Pointer(&b)))
		}
	}
	return QuotedValue(s)
}
