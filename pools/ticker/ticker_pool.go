package tickerpool

import (
	"sync"
	"time"
)

var timerPool sync.Pool

func New(d time.Duration) *time.Ticker {
	v := timerPool.Get()
	if v == nil {
		return time.NewTicker(d)
	}
	return v.(*time.Ticker)
}

func Release(t *time.Ticker) {
	timerPool.Put(t)
}
