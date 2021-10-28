package timerpool

import (
	"sync"
	"time"
)

var timerPool sync.Pool

func New(d time.Duration) *time.Timer {
	v := timerPool.Get()
	if v == nil {
		return time.NewTimer(d)
	}
	t := v.(*time.Timer)
	if t.Reset(d) {
		// active timer trapped to the pool?
		// t.Stop()
		return time.NewTimer(d)
	}
	return t
}

func Release(t *time.Timer) {
	if !t.Stop() {
		return
	}
	timerPool.Put(t)
}
