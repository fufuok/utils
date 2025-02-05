package utils

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTrylock(t *testing.T) {
	m := NewTryMutex()
	ok := m.TryLock(20 * time.Millisecond)
	if !ok {
		t.Error("it should be the lock succeeded but failed")
	}

	ok = m.TryLock(20 * time.Millisecond)
	if ok {
		t.Error("It should be the lock failed but it succeeded")
	}

	m.Unlock()

	ok = m.TryLock()
	if !ok {
		t.Fatal("it should be the lock succeeded but failed")
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		m.Unlock()
	}()

	ts := time.Now()
	ok = m.TryLock(120 * time.Millisecond)
	took := time.Since(ts)
	if !ok {
		t.Fatal("it should be the lock succeeded but failed")
	}
	if took < 50*time.Millisecond {
		t.Fatalf("expected time to acquire lock < 100ms, got: %s", took)
	}
}

func BenchmarkTrylock(b *testing.B) {
	b.Run("trylock", func(b *testing.B) {
		lock := NewTryMutex()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					defer lock.Unlock()
					lock.TryLock(20 * time.Millisecond)
				}()
			}
		})
	})
	b.Run("lock", func(b *testing.B) {
		lock := NewTryMutex()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					defer lock.Unlock()
					lock.Lock()
				}()
			}
		})
	})
	b.Run("mutex", func(b *testing.B) {
		var lock sync.Mutex
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					defer lock.Unlock()
					lock.Lock()
				}()
			}
		})
	})
}

/*
Ref: https://github.com/panjf2000/ants/blob/dev/pkg/sync/spinlock_test.go
Benchmark result for three types of locks:
	goos: darwin
	goarch: arm64
	pkg: github.com/panjf2000/ants/v2/pkg/sync
	BenchmarkMutex-10              	10452573	       111.1 ns/op	       0 B/op	       0 allocs/op
	BenchmarkSpinLock-10           	58953211	        18.01 ns/op	       0 B/op	       0 allocs/op
	BenchmarkBackOffSpinLock-10    	100000000	        10.81 ns/op	       0 B/op	       0 allocs/op
*/

type originSpinLock uint32

func (sl *originSpinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *originSpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func NewOriginSpinLock() sync.Locker {
	return new(originSpinLock)
}

func BenchmarkSpinMutex(b *testing.B) {
	m := sync.Mutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			//nolint:staticcheck
			m.Unlock()
		}
	})
}

func BenchmarkSpinLock(b *testing.B) {
	spin := NewOriginSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			//nolint:staticcheck
			spin.Unlock()
		}
	})
}

func BenchmarkBackOffSpinLock(b *testing.B) {
	spin := NewSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			//nolint:staticcheck
			spin.Unlock()
		}
	})
}

// # go test -run=^$ -benchmem -benchtime=1s -bench=Spin
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// BenchmarkSpinMutex-16           23156875                54.36 ns/op            0 B/op          0 allocs/op
// BenchmarkSpinLock-16            215412934                5.564 ns/op           0 B/op          0 allocs/op
// BenchmarkBackOffSpinLock-16     246957524                4.877 ns/op           0 B/op          0 allocs/op
