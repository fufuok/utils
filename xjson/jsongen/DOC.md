<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# jsongen

```go
import "github.com/fufuok/utils/xjson/jsongen"
```

Package jsongen forked from darjun/json\-gen

## Index

- [type Array](<#type-array>)
  - [func NewArray() *Array](<#func-newarray>)
  - [func (a *Array) AppendArray(oa Array)](<#func-array-appendarray>)
  - [func (a *Array) AppendBool(b bool)](<#func-array-appendbool>)
  - [func (a *Array) AppendBoolArray(b []bool)](<#func-array-appendboolarray>)
  - [func (a *Array) AppendFloat(f float64)](<#func-array-appendfloat>)
  - [func (a *Array) AppendFloatArray(f []float64)](<#func-array-appendfloatarray>)
  - [func (a *Array) AppendInt(i int64)](<#func-array-appendint>)
  - [func (a *Array) AppendIntArray(i []int64)](<#func-array-appendintarray>)
  - [func (a *Array) AppendMap(m *Map)](<#func-array-appendmap>)
  - [func (a *Array) AppendMapArray(m []Map)](<#func-array-appendmaparray>)
  - [func (a *Array) AppendRawBytes(b []byte)](<#func-array-appendrawbytes>)
  - [func (a *Array) AppendRawBytesArray(bs [][]byte)](<#func-array-appendrawbytesarray>)
  - [func (a *Array) AppendRawString(s string)](<#func-array-appendrawstring>)
  - [func (a *Array) AppendRawStringArray(ss []string)](<#func-array-appendrawstringarray>)
  - [func (a *Array) AppendString(value string)](<#func-array-appendstring>)
  - [func (a *Array) AppendStringArray(s []string)](<#func-array-appendstringarray>)
  - [func (a *Array) AppendUint(u uint64)](<#func-array-appenduint>)
  - [func (a *Array) AppendUintArray(u []uint64)](<#func-array-appenduintarray>)
  - [func (a Array) Serialize(buf []byte) []byte](<#func-array-serialize>)
  - [func (a Array) Size() int](<#func-array-size>)
- [type Map](<#type-map>)
  - [func NewMap() *Map](<#func-newmap>)
  - [func (m *Map) PutArray(key string, a *Array)](<#func-map-putarray>)
  - [func (m *Map) PutBool(key string, b bool)](<#func-map-putbool>)
  - [func (m *Map) PutBoolArray(key string, b []bool)](<#func-map-putboolarray>)
  - [func (m *Map) PutFloat(key string, f float64)](<#func-map-putfloat>)
  - [func (m *Map) PutFloatArray(key string, f []float64)](<#func-map-putfloatarray>)
  - [func (m *Map) PutInt(key string, i int64)](<#func-map-putint>)
  - [func (m *Map) PutIntArray(key string, i []int64)](<#func-map-putintarray>)
  - [func (m *Map) PutMap(key string, om *Map)](<#func-map-putmap>)
  - [func (m *Map) PutRawBytes(key string, b []byte)](<#func-map-putrawbytes>)
  - [func (m *Map) PutRawBytesArray(key string, bs [][]byte)](<#func-map-putrawbytesarray>)
  - [func (m *Map) PutRawString(key, s string)](<#func-map-putrawstring>)
  - [func (m *Map) PutRawStringArray(key string, ss []string)](<#func-map-putrawstringarray>)
  - [func (m *Map) PutString(key, value string)](<#func-map-putstring>)
  - [func (m *Map) PutStringArray(key string, s []string)](<#func-map-putstringarray>)
  - [func (m *Map) PutUint(key string, u uint64)](<#func-map-putuint>)
  - [func (m *Map) PutUintArray(key string, u []uint64)](<#func-map-putuintarray>)
  - [func (m Map) Serialize(buf []byte) []byte](<#func-map-serialize>)
  - [func (m Map) Size() int](<#func-map-size>)
- [type QuotedValue](<#type-quotedvalue>)
  - [func (q QuotedValue) Serialize(buf []byte) []byte](<#func-quotedvalue-serialize>)
  - [func (q QuotedValue) Size() int](<#func-quotedvalue-size>)
- [type RawBytes](<#type-rawbytes>)
  - [func (b RawBytes) Serialize(buf []byte) []byte](<#func-rawbytes-serialize>)
  - [func (b RawBytes) Size() int](<#func-rawbytes-size>)
- [type RawString](<#type-rawstring>)
  - [func (s RawString) Serialize(buf []byte) []byte](<#func-rawstring-serialize>)
  - [func (s RawString) Size() int](<#func-rawstring-size>)
- [type UnquotedValue](<#type-unquotedvalue>)
  - [func (u UnquotedValue) Serialize(buf []byte) []byte](<#func-unquotedvalue-serialize>)
  - [func (u UnquotedValue) Size() int](<#func-unquotedvalue-size>)
- [type Value](<#type-value>)
  - [func EscapeString(s string) Value](<#func-escapestring>)


## type Array

Array 表示一个\`json\`数组

```go
type Array []Value
```

### func NewArray

```go
func NewArray() *Array
```

NewArray 创建一个\`json\`数组，返回其指针

### func \(\*Array\) AppendArray

```go
func (a *Array) AppendArray(oa Array)
```

AppendArray 将\`json\`数组\`oa\`追加到数组\`a\`后

### func \(\*Array\) AppendBool

```go
func (a *Array) AppendBool(b bool)
```

AppendBool 将\`bool\`类型的值\`b\`追加到数组\`a\`后

### func \(\*Array\) AppendBoolArray

```go
func (a *Array) AppendBoolArray(b []bool)
```

AppendBoolArray 将\`bool\`数组\`b\`追加到数组\`a\`后

### func \(\*Array\) AppendFloat

```go
func (a *Array) AppendFloat(f float64)
```

AppendFloat 将\`float64\`类型的值\`f\`追加到数组\`a\`后

### func \(\*Array\) AppendFloatArray

```go
func (a *Array) AppendFloatArray(f []float64)
```

AppendFloatArray 将\`float64\`数组\`f\`追加到数组\`a\`后

### func \(\*Array\) AppendInt

```go
func (a *Array) AppendInt(i int64)
```

AppendInt 将\`int64\`类型的值\`i\`追加到数组\`a\`后

### func \(\*Array\) AppendIntArray

```go
func (a *Array) AppendIntArray(i []int64)
```

AppendIntArray 将\`int64\`数组\`i\`追加到数组\`a\`后

### func \(\*Array\) AppendMap

```go
func (a *Array) AppendMap(m *Map)
```

AppendMap 将\`Map\`类型的值\`m\`追加到数组\`a\`后

### func \(\*Array\) AppendMapArray

```go
func (a *Array) AppendMapArray(m []Map)
```

AppendMapArray 将\`Map\`数组\`m\`追加到数组\`a\`后

### func \(\*Array\) AppendRawBytes

```go
func (a *Array) AppendRawBytes(b []byte)
```

### func \(\*Array\) AppendRawBytesArray

```go
func (a *Array) AppendRawBytesArray(bs [][]byte)
```

### func \(\*Array\) AppendRawString

```go
func (a *Array) AppendRawString(s string)
```

### func \(\*Array\) AppendRawStringArray

```go
func (a *Array) AppendRawStringArray(ss []string)
```

### func \(\*Array\) AppendString

```go
func (a *Array) AppendString(value string)
```

AppendString 将\`string\`类型的值\`s\`追加到数组\`a\`后

### func \(\*Array\) AppendStringArray

```go
func (a *Array) AppendStringArray(s []string)
```

AppendStringArray 将\`string\`数组\`s\`追加到数组\`a\`后

### func \(\*Array\) AppendUint

```go
func (a *Array) AppendUint(u uint64)
```

AppendUint 将\`uint64\`类型的值\`u\`追加到数组\`a\`后

### func \(\*Array\) AppendUintArray

```go
func (a *Array) AppendUintArray(u []uint64)
```

AppendUintArray 将\`uint64\`数组\`u\`追加到数组\`a\`后

### func \(Array\) Serialize

```go
func (a Array) Serialize(buf []byte) []byte
```

Serialize 将\`a\`序列化为字符串，追加到\`buf\`后，返回新的\`buf\`

### func \(Array\) Size

```go
func (a Array) Size() int
```

Size 返回\`a\`在最终的\`json\`串中占有多少字节

## type Map

Map 表示一个\`json\`映射

```go
type Map struct {
    // contains filtered or unexported fields
}
```

### func NewMap

```go
func NewMap() *Map
```

NewMap 创建一个\`json\`映射返回其指针

### func \(\*Map\) PutArray

```go
func (m *Map) PutArray(key string, a *Array)
```

PutArray 将\`json\`数组\`a\`与键\`key\`关联

### func \(\*Map\) PutBool

```go
func (m *Map) PutBool(key string, b bool)
```

PutBool 将\`bool\`类型的值\`b\`与键\`key\`关联

### func \(\*Map\) PutBoolArray

```go
func (m *Map) PutBoolArray(key string, b []bool)
```

PutBoolArray 将\`bool\`数组类型的值\`b\`与键\`key\`关联

### func \(\*Map\) PutFloat

```go
func (m *Map) PutFloat(key string, f float64)
```

PutFloat 将\`float64\`类型的值\`f\`与键\`key\`关联

### func \(\*Map\) PutFloatArray

```go
func (m *Map) PutFloatArray(key string, f []float64)
```

PutFloatArray 将\`float64\`数组类型的值\`f\`与键\`key\`关联

### func \(\*Map\) PutInt

```go
func (m *Map) PutInt(key string, i int64)
```

PutInt 将\`int64\`类型的值\`i\`与键\`key\`关联

### func \(\*Map\) PutIntArray

```go
func (m *Map) PutIntArray(key string, i []int64)
```

PutIntArray 将\`int64\`数组类型的值\`i\`与键\`key\`关联

### func \(\*Map\) PutMap

```go
func (m *Map) PutMap(key string, om *Map)
```

PutMap 将\`json\`映射\`om\`与键\`key\`关联

### func \(\*Map\) PutRawBytes

```go
func (m *Map) PutRawBytes(key string, b []byte)
```

### func \(\*Map\) PutRawBytesArray

```go
func (m *Map) PutRawBytesArray(key string, bs [][]byte)
```

### func \(\*Map\) PutRawString

```go
func (m *Map) PutRawString(key, s string)
```

### func \(\*Map\) PutRawStringArray

```go
func (m *Map) PutRawStringArray(key string, ss []string)
```

### func \(\*Map\) PutString

```go
func (m *Map) PutString(key, value string)
```

PutString 将\`string\`类型的值\`value\`与键\`key\`关联

### func \(\*Map\) PutStringArray

```go
func (m *Map) PutStringArray(key string, s []string)
```

PutStringArray 将\`string\`数组类型的值\`s\`与键\`key\`关联

### func \(\*Map\) PutUint

```go
func (m *Map) PutUint(key string, u uint64)
```

PutUint 将\`uint64\`类型的值\`u\`与键\`key\`关联

### func \(\*Map\) PutUintArray

```go
func (m *Map) PutUintArray(key string, u []uint64)
```

PutUintArray 将\`uint64\`数组类型的值\`u\`与键\`key\`关联

### func \(Map\) Serialize

```go
func (m Map) Serialize(buf []byte) []byte
```

Serialize 将\`m\`序列化为字符串，追加到\`buf\`后，返回新的\`buf\`

### func \(Map\) Size

```go
func (m Map) Size() int
```

Size 返回\`m\`在最终的\`json\`串中占有多少字节

## type QuotedValue

QuotedValue 表示需要用"包裹起来的值，例如字符串

```go
type QuotedValue string
```

### func \(QuotedValue\) Serialize

```go
func (q QuotedValue) Serialize(buf []byte) []byte
```

Serialize 将\`q\`序列化为字符串，追加到\`buf\`后，返回新的\`buf\`

### func \(QuotedValue\) Size

```go
func (q QuotedValue) Size() int
```

Size 返回\`q\`在最终的\`json\`串中占有多少字节

## type RawBytes

```go
type RawBytes []byte
```

### func \(RawBytes\) Serialize

```go
func (b RawBytes) Serialize(buf []byte) []byte
```

### func \(RawBytes\) Size

```go
func (b RawBytes) Size() int
```

## type RawString

```go
type RawString string
```

### func \(RawString\) Serialize

```go
func (s RawString) Serialize(buf []byte) []byte
```

### func \(RawString\) Size

```go
func (s RawString) Size() int
```

## type UnquotedValue

UnquotedValue 表示不需要用"包裹起来的值，例如整数，浮点数等

```go
type UnquotedValue string
```

### func \(UnquotedValue\) Serialize

```go
func (u UnquotedValue) Serialize(buf []byte) []byte
```

Serialize 将\`u\`序列化为字符串，追加到\`buf\`后，返回新的\`buf\`

### func \(UnquotedValue\) Size

```go
func (u UnquotedValue) Size() int
```

Size 返回\`u\`在最终的\`json\`串中占有多少字节

## type Value

Value 表示将要序列化到\`json\`字符串中的值

```go
type Value interface {
    // Serialize 将值序列化为字符串，追加到`buf`后，返回新的`buf`
    Serialize(buf []byte) []byte
    // Size 返回值在最终的`json`串中占有多少字节
    Size() int
}
```

### func EscapeString

```go
func EscapeString(s string) Value
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
