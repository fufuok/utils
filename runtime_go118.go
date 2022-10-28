//go:build !go1.19
// +build !go1.19

package utils

import (
	_ "unsafe"
)

func FastRand64() uint64 {
	return (uint64(FastRand()) << 32) | uint64(FastRand())
}

func FastRandu() uint {
	if PtrSize == 8 {
		return uint(FastRand64())
	}
	return uint(FastRand())
}
