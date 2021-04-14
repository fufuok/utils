package utils

import (
	"time"
)

// 下一分钟, 对齐时间, 0 秒
func WaitNextMinute() {
	t := time.Now().Add(time.Minute)
	t0 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	<-time.After(t0.Sub(time.Now()))
}
