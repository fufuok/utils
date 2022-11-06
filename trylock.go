package utils

import (
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
