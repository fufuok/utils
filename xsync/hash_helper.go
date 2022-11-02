package xsync

import (
	"hash/maphash"
	"unsafe"
)

// HashSeedString calculates a hash of s with the given seed.
func HashSeedString(seed maphash.Seed, s string) uint64 {
	return hashString(seed, s)
}

// HashSeedUint64 calculates a hash of n with the given seed.
func HashSeedUint64(seed maphash.Seed, n uint64) uint64 {
	// Java's Long standard hash function.
	n = n ^ (n >> 32)
	nseed := *(*uint64)(unsafe.Pointer(&seed))
	// 64-bit variation of boost's hash_combine.
	nseed ^= n + 0x9e3779b97f4a7c15 + (nseed << 12) + (nseed >> 4)
	return nseed
}
