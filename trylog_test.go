package utils

import (
	"sync"
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
	ok = m.TryLock(100 * time.Millisecond)
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
