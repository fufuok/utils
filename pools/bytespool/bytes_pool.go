package bytespool

import (
	"math"
	"math/bits"
	"sync"
	"unsafe"
)

const (
	minCapacity       = 1 << 0
	smallBufferSize   = 64
	largeBufferSize   = math.MaxInt32
	defaultBufferSize = 8 << 20 // 8 MiB
)

var (
	defaultMaxSize = defaultBufferSize

	// 默认的各容量刻度字节数组池
	defaultCapacityPools CapacityPools
)

type CapacityPools struct {
	pools [32]sync.Pool
}

// SetMaxSize 设置回收时允许的最大字节
func SetMaxSize(size int) bool {
	if size >= smallBufferSize && size <= largeBufferSize {
		defaultMaxSize = size
		return true
	}
	return false
}

// Make 返回 len 为 0, cap >= 给定值的 []byte
func (p *CapacityPools) Make(capacity int) []byte {
	return p.Get(capacity)[:0]
}

// Get 完全同 New()
func (p *CapacityPools) Get(size int) []byte {
	return p.New(size)
}

// New 返回指定长度的 []byte
// 注:
//  1. 返回的 buf != nil
//  2. 由于复用底层数组, buf 可能残留有旧数据
func (p *CapacityPools) New(size int) (buf []byte) {
	if size < minCapacity {
		return p.Make(minCapacity)
	}
	if size > defaultMaxSize {
		return Bytes(size, size)
	}
	idx := getIndex(size)
	ptr, _ := p.pools[idx].Get().(*byte)
	if ptr == nil {
		return Bytes(size, 1<<idx)
	}
	
	// go1.20
	// return unsafe.Slice(ptr, bp.capacity)[:size]

	sh := (*bytesHeader)(unsafe.Pointer(&buf))
	sh.Data = ptr
	sh.Len = size
	sh.Cap = 1 << idx
	return
}

// NewBytes returns a byte slice of the specified content.
func (p *CapacityPools) NewBytes(bs []byte) []byte {
	buf := p.Make(len(bs))
	return append(buf, bs...)
}

// NewString returns a byte slice of the specified content.
func (p *CapacityPools) NewString(s string) []byte {
	buf := p.Make(len(s))
	return append(buf, s...)
}

// Append 与内置的 append 功能相同, 当底层数组需要重建时, 会回收原数组
func (p *CapacityPools) Append(buf []byte, elems ...byte) []byte {
	return p.AppendString(buf, *(*string)(unsafe.Pointer(&elems)))
}

func (p *CapacityPools) AppendString(buf []byte, elems string) []byte {
	n := len(buf)
	c := cap(buf)
	m := n + len(elems)
	if c < m && c <= defaultMaxSize {
		bbuf := p.Get(m)
		copy(bbuf, buf)
		copy(bbuf[n:], elems)
		p.Put(buf)
		return bbuf
	}
	return append(buf, elems...)
}

// Put 同 Release(), 不返回是否回收成功
func (p *CapacityPools) Put(buf []byte) {
	p.Release(buf)
}

func (p *CapacityPools) Release(buf []byte) bool {
	n := cap(buf)
	if n == 0 || n > defaultMaxSize {
		return false
	}
	idx := getIndex(n)
	if n != 1<<idx {
		idx--
	}

	// go1.20, store array pointer,
	// bp.pool.Put(unsafe.SliceData(buf))

	sh := (*bytesHeader)(unsafe.Pointer(&buf))
	p.pools[idx].Put(sh.Data)
	return true
}

func getIndex(n int) int {
	return bits.Len32(uint32(n) - 1)
}

func Make(size int) []byte {
	return defaultCapacityPools.Make(size)
}

func Get(size int) []byte {
	return defaultCapacityPools.Get(size)
}

func New(size int) []byte {
	return defaultCapacityPools.New(size)
}

func NewBytes(bs []byte) []byte {
	return defaultCapacityPools.NewBytes(bs)
}

func NewString(s string) []byte {
	return defaultCapacityPools.NewString(s)
}

func Append(buf []byte, elems ...byte) []byte {
	return defaultCapacityPools.Append(buf, elems...)
}

func AppendString(buf []byte, elems string) []byte {
	return defaultCapacityPools.AppendString(buf, elems)
}

func Release(buf []byte) bool {
	return defaultCapacityPools.Release(buf)
}

func Put(buf []byte) {
	defaultCapacityPools.Put(buf)
}
