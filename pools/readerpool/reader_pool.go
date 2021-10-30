package readerpool

import (
	"bytes"
	"sync"
)

var readerPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewReader(nil)
	},
}

func New(b []byte) *bytes.Reader {
	r := readerPool.Get().(*bytes.Reader)
	r.Reset(b)
	return r
}

func Release(r *bytes.Reader) {
	r.Reset(nil)
	readerPool.Put(r)
}
