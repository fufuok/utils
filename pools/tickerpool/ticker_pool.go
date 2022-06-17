package tickerpool

import (
	"sync"
	"time"
)

var tickerPool sync.Pool

func New(d time.Duration) *time.Ticker {
	v := tickerPool.Get()
	if v == nil {
		return time.NewTicker(d)
	}
	t := v.(*time.Ticker)
	t.Reset(d)
	return t
}

func Release(t *time.Ticker) {
	t.Stop()
	tickerPool.Put(t)
}
