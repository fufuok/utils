//go:build go1.20
// +build go1.20

package utils

import (
	"unsafe"
)

// S2B converts string to byte slice without a memory allocation.
// Ref: https://github.com/golang/go/issues/53003#issuecomment-1140276077
func S2B(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// B2S converts byte slice to string without a memory allocation.
// Slower: unsafe.String(unsafe.SliceData(b), len(b))
// strings.Clone(): unsafe.String(&b[0], len(b))
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
