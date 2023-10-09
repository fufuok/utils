package utils

import (
	"bytes"
	"runtime"
	"strconv"
	"sync/atomic"
)

var runtimeID uint64

// ID 运行时自增 ID (每次程序启动从 1 开始)
func ID() uint64 {
	return atomic.AddUint64(&runtimeID, 1)
}

// GoroutineID 获取 Goroutine ID
func GoroutineID() (uint64, error) {
	b := make([]byte, 64)
	n := runtime.Stack(b, false)
	b = b[:n]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return strconv.ParseUint(string(b), 10, 64)
}
