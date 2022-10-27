//go:build go1.19
// +build go1.19

package xsync

import (
	"hash/maphash"
)

// hashString calculates a hash of s with the given seed.
func hashString(seed maphash.Seed, s string) uint64 {
	return maphash.String(seed, s)
}
