package bufferpool

import (
	"bytes"
	"sync"
)

// 8 MiB
const defaultMaxSize = 8 << 20

var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
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
