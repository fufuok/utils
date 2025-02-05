package utils

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fufuok/utils/pools/timerpool"
)

type TryMutex struct {
	lock chan struct{}
}

func NewTryMutex() *TryMutex {
	return &TryMutex{
		make(chan struct{}, 1),
	}
}

func (m *TryMutex) Lock() {
	m.lock <- struct{}{}
}

func (m *TryMutex) Unlock() {
	<-m.lock
}

// TryLock 实现可选等待时间尝试获取锁
func (m *TryMutex) TryLock(timeout ...time.Duration) bool {
	select {
	case m.lock <- struct{}{}:
		return true
	default:
		if len(timeout) == 0 || timeout[0] <= 0 {
			return false
		}
		timer := timerpool.New(timeout[0])
		defer timerpool.Release(timer)
		select {
		case m.lock <- struct{}{}:
			return true
		case <-timer.C:
			return false
		}

	}
}

// https://github.com/panjf2000/ants/blob/dev/pkg/sync/spinlock.go
type spinLock uint32

const maxBackoff = 16

func (sl *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		// Leverage the exponential backoff algorithm, see https://en.wikipedia.org/wiki/Exponential_backoff.
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() sync.Locker {
	return new(spinLock)
}
