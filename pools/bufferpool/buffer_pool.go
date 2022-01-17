package bufferpool

import (
	"bytes"
	"sync"
)

var (
	// 8 MiB
	defaultMaxSize = 8 << 20

	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
)

// SetMaxSize 设置回收时允许的最大字节
// smallBufferSize is an initial allocation minimal capacity.
// const smallBufferSize = 64
func SetMaxSize(size int) bool {
	// 64 <= size <= 2GiB
	if size >= 64 && size <= 2<<30 {
		defaultMaxSize = size
		return true
	}
	return false
}

func New(bs []byte) *bytes.Buffer {
	buf := Get()
	buf.Write(bs)
	return buf
}

func NewByte(c byte) *bytes.Buffer {
	buf := Get()
	buf.WriteByte(c)
	return buf
}

func NewString(s string) *bytes.Buffer {
	buf := Get()
	buf.WriteString(s)
	return buf
}

func NewRune(r rune) *bytes.Buffer {
	buf := Get()
	buf.WriteRune(r)
	return buf
}

func Get() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func Put(buf *bytes.Buffer) {
	Release(buf)
}

func Release(buf *bytes.Buffer) bool {
	if buf.Cap() > defaultMaxSize {
		return false
	}
	buf.Reset()
	bufferPool.Put(buf)
	return true
}
