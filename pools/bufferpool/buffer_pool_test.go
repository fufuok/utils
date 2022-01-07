package bufferpool

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestBufferPool(t *testing.T) {
	for i := 0; i < 10; i++ {
		want := fmt.Sprintf("test%d", i)
		buf := Get()
		buf.WriteString("test")
		buf.WriteString(strconv.Itoa(i))
		if buf.String() != want {
			t.Fatalf("Unexpected result: %q, Expecting %q", buf.String(), want)
		}
		Put(buf)

		buf = New([]byte(want))
		if buf.String() != want {
			t.Fatalf("Unexpected result: %q, Expecting %q", buf.String(), want)
		}
		Put(buf)

		buf = NewString(want)
		if buf.String() != want {
			t.Fatalf("Unexpected result: %q, Expecting %q", buf.String(), want)
		}
		if !Release(buf) {
			t.Fatal("Unexpected result: false, Expecting true")
		}

		buf = NewByte(65)
		if buf.String() != "A" {
			t.Fatalf("Unexpected result: %s, Expecting 'A'", buf.String())
		}

		buf = NewRune(20013)
		if buf.String() != "中" {
			t.Fatalf("Unexpected result: '中', Expecting %q", buf.String())
		}
	}
}

func BenchmarkBufferPool(b *testing.B) {
	bs := []byte("bufferpool")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := New(bs)
			Put(buf)
		}
	})
}

func BenchmarkBufferNew(b *testing.B) {
	bs := []byte("bufferpool")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var buf bytes.Buffer
			buf.Write(bs)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils/pools/bufferpool
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkBufferPool-4           162477872                9.719 ns/op           0 B/op          0 allocs/op
// BenchmarkBufferNew-4             76593980                15.72 ns/op          64 B/op          1 allocs/op
// PASS
// ok      github.com/fufuok/utils/pools/bufferpool        4.410s
