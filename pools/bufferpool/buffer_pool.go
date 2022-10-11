package bufferpool

import (
	"bytes"
	"math"
	"sync"
)

const (
	smallBufferSize   = 64
	largeBufferSize   = math.MaxInt32
	defaultBufferSize = 8 << 20 // 8 MiB
)

var (
	defaultMaxSize = defaultBufferSize

	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
)

// SetMaxSize 设置回收时允许的最大字节
func SetMaxSize(size int) bool {
	if size >= smallBufferSize && size <= largeBufferSize {
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
