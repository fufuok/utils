<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# bufferpool

```go
import "github.com/fufuok/utils/pools/bufferpool"
```

## Index

- [func Get() *bytes.Buffer](<#func-get>)
- [func New(bs []byte) *bytes.Buffer](<#func-new>)
- [func NewByte(c byte) *bytes.Buffer](<#func-newbyte>)
- [func NewRune(r rune) *bytes.Buffer](<#func-newrune>)
- [func NewString(s string) *bytes.Buffer](<#func-newstring>)
- [func Put(buf *bytes.Buffer)](<#func-put>)
- [func Release(buf *bytes.Buffer) bool](<#func-release>)
- [func SetMaxSize(size int) bool](<#func-setmaxsize>)


## func Get

```go
func Get() *bytes.Buffer
```

## func New

```go
func New(bs []byte) *bytes.Buffer
```

## func NewByte

```go
func NewByte(c byte) *bytes.Buffer
```

## func NewRune

```go
func NewRune(r rune) *bytes.Buffer
```

## func NewString

```go
func NewString(s string) *bytes.Buffer
```

## func Put

```go
func Put(buf *bytes.Buffer)
```

## func Release

```go
func Release(buf *bytes.Buffer) bool
```

## func SetMaxSize

```go
func SetMaxSize(size int) bool
```

SetMaxSize 设置回收时允许的最大字节 smallBufferSize is an initial allocation minimal capacity. const smallBufferSize = 64



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
