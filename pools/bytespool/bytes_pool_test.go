package bytespool

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"testing"
)

func TestCapacityPools(t *testing.T) {
	maxSize := 2048
	SetMaxSize(maxSize)
	tests := []struct {
		size        int
		scaleSize   int
		bytesLength int
		releaseOK   bool
	}{
		{-1, minCapacity, 0, true},
		{0, minCapacity, 0, true},
		{minCapacity, minCapacity, minCapacity, true},
		{2, 2, 2, true},
		{64, 64, 64, true},
		{128, 128, 128, true},
		{2000, 2048, 2000, true},
		{2047, 2048, 2047, true},
		{maxSize, maxSize, maxSize, true},
		{4096, 0, 4096, false},
		{5000, 0, 5000, false},
	}
	for _, v := range tests {
		t.Run(fmt.Sprintf("bytes.Get(%d)", v.size), func(t *testing.T) {
			buf := Make(v.size)
			if buf == nil {
				t.Fatal("expect  buf != nil")
			}
			if len(buf) != 0 {
				t.Fatalf("expect buffer len is 0, but got %d", len(buf))
			}
			if cap(buf) < v.scaleSize {
				t.Fatalf("expect buffer cap >= %d, but got %d", v.scaleSize, cap(buf))
			}

			buf = Get(v.size)
			if len(buf) != v.bytesLength {
				t.Fatalf("expect buffer len is %d, but got %d", v.bytesLength, len(buf))
			}
			if cap(buf) < v.scaleSize {
				t.Fatalf("expect buffer cap >= %d, but got %d", v.scaleSize, cap(buf))
			}

			ok := Release(buf)
			if ok != v.releaseOK {
				t.Fatalf("expect to release the buffer result is %v, but got %v", v.releaseOK, ok)
			}
		})
	}
	SetMaxSize(defaultBufferSize)
}

func TestCapacityPools_Default(t *testing.T) {
	buf := Make(defaultMaxSize + 1)
	if len(buf) != 0 {
		t.Fatalf("expect buffer len is 0, but got %d", len(buf))
	}
	if cap(buf) <= defaultMaxSize {
		t.Fatalf("expect buffer cap > %d, but got %d", defaultMaxSize, cap(buf))
	}
	if Release(buf) {
		t.Fatal("expect to release the buffer failure, but not")
	}

	buf = Make(defaultMaxSize)
	if len(buf) != 0 {
		t.Fatalf("expect buffer len is 0, but got %d", len(buf))
	}
	if cap(buf) != defaultMaxSize {
		t.Fatalf("expect buffer cap is %d, but got %d", defaultMaxSize, cap(buf))
	}

	len0 := make([]byte, 0, 8)
	if !Release(len0) {
		t.Fatal("expect to release the buffer successfully, but not")
	}

	var cap0 []byte
	if Release(cap0) {
		t.Fatal("expect to release the buffer failure, but not")
	}

	abc := []byte("abc")
	buf = append(buf, abc...)

	// Disable GC to test re-acquire the same data
	gc := debug.SetGCPercent(-1)

	if !Release(buf) {
		t.Fatal("expect to release the buffer successfully, but not")
	}

	newBuf := Get(defaultMaxSize)
	if fmt.Sprintf("%p", newBuf) != fmt.Sprintf("%p", buf) {
		t.Fatal("expect the newBuf is the buf, but not")
	}
	if !bytes.Equal(abc, (newBuf)[:3]) {
		t.Fatal("expect that newBuf may contain old data, but not")
	}

	if !Release(newBuf) {
		t.Fatal("expect to release the buffer successfully, but not")
	}

	buf8 := Get(8)
	copy(buf8, "12345678")
	if string(buf8) != "12345678" {
		t.Fatal("expect copy result is 123456789, but not")
	}

	buf8 = append(buf8, '9')

	Put(buf8)

	buf16 := Get(16)
	if &buf8[0] != &buf16[0] {
		t.Fatal("expect buf8 and buf16 to be the same array")
	}
	if string(buf16[:9]) != "123456789" {
		t.Fatal("expect the buf8 is the buf16, but not")
	}

	// Re-enable GC
	debug.SetGCPercent(gc)

	SetMaxSize(smallBufferSize)

	buf = Make(3)
	if len(buf) != 0 {
		t.Fatalf("expect buffer len is 0, but got %d", len(buf))
	}
	if cap(buf) != 4 {
		t.Fatalf("expect buffer cap is 4, but got %d", cap(buf))
	}
	buf = Make(smallBufferSize + 3)
	if len(buf) != 0 {
		t.Fatalf("expect buffer len is 0, but got %d", len(buf))
	}
	if cap(buf) != smallBufferSize+3 {
		t.Fatalf("expect buffer cap is smallBufferSize+3, but got %d", cap(buf))
	}
	if Release(buf) {
		t.Fatal("expect to release the buffer failure, but not")
	}
	buf = append(buf, '1')
	if Release(buf) {
		t.Fatal("expect to release the buffer failure, but not")
	}

	SetMaxSize(defaultBufferSize)
}

func TestNewBytesString(t *testing.T) {
	s := "Fufu 中文-123"
	bs := []byte(s)

	buf := NewString(s)
	if cap(buf) != 16 {
		t.Fatalf("expect buffer cap is 16, but got %d", cap(buf))
	}
	if string(buf) != s {
		t.Fatalf("expect buf to be equal to %s, but not", s)
	}

	buf = NewBytes(bs)
	if cap(buf) != 16 {
		t.Fatalf("expect buffer cap is 16, but got %d", cap(buf))
	}
	if !bytes.Equal(buf, bs) {
		t.Fatalf("expect buf to be equal to %s, but not", bs)
	}
}

func TestUnalignedCapacity(t *testing.T) {
	bs := make([]byte, 0, 7)
	bs = append(bs, "123"...)
	if !Release(bs) {
		t.Fatal("expect to release the buffer successfully, but not")
	}
	buf := Make(3)
	if cap(buf) != 4 {
		t.Fatalf("expect buffer cap is 4, but got %d", cap(buf))
	}
	if !Release(buf) {
		t.Fatal("expect to release the buffer successfully, but not")
	}
}

func TestAppend(t *testing.T) {
	// Disable GC to test re-acquire the same data
	gc := debug.SetGCPercent(-1)
	buf := Get(4)
	if len(buf) != 4 || cap(buf) != 4 {
		t.Fatalf("expect buf cap is 4, but got %d", cap(buf))
	}

	copy(buf, "1234")
	bbuf := Append(buf, '+')
	if len(bbuf) != 5 || cap(bbuf) != 8 {
		t.Fatalf("expect bbuf cap is 8, but got %d", cap(bbuf))
	}
	// Warning: you should stop using (buf)!
	if len(buf) != 4 || cap(buf) != 4 {
		t.Fatalf("expect buf cap is 4, but got %d", cap(buf))
	}

	if !bytes.EqualFold(bbuf, []byte("1234+")) || !bytes.EqualFold(buf, []byte("1234")) {
		t.Fatalf("expect bbuf is 1234+, buf is 1234")
	}

	if &bbuf[0] == &buf[0] {
		t.Fatal("expect bbuf and buf to not be the same array")
	}

	cbuf := AppendString(bbuf, "+2")
	if len(bbuf) != 7 && cap(bbuf) != 8 && len(cbuf) != 7 && cap(cbuf) != 8 {
		t.Fatalf("expect bbuf and cbuf cap is 8, but got %d", cap(bbuf))
	}

	if &cbuf[0] != &bbuf[0] {
		t.Fatal("expect bbuf and buf to be the same array")
	}
	// Re-enable GC
	debug.SetGCPercent(gc)
}
