//go:build !go1.20
// +build !go1.20

package utils

import (
	"unsafe"
)

// S2B StringToBytes converts string to byte slice without a memory allocation.
// Ref: gin
func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// B2S BytesToString
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
