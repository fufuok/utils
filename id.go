package utils

import (
	"sync/atomic"
)

var runtimeID uint64

// ID 运行时自增 ID (每次程序启动从 1 开始)
func ID() uint64 {
	return atomic.AddUint64(&runtimeID, 1)
}
