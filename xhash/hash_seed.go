//go:build !go1.19
// +build !go1.19

package xhash

import (
	"hash/maphash"
	"reflect"
	"unsafe"
)

// hashString calculates a hash of s with the given seed.
func hashString(seed maphash.Seed, s string) uint64 {
	seed64 := *(*uint64)(unsafe.Pointer(&seed))
	if s == "" {
		return seed64
	}
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return uint64(memhash(unsafe.Pointer(strh.Data), uintptr(seed64), uintptr(strh.Len)))
}
