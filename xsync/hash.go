package xsync

import (
	"hash/maphash"
)

// HashSeedString calculates a hash of s with the given seed.
func HashSeedString(seed maphash.Seed, s string) uint64 {
	return hashString(seed, s)
}

// HashSeedUint64 calculates a hash of v with the given seed.
func HashSeedUint64(seed maphash.Seed, v uint64) uint64 {
	return hashUint64(seed, v)
}
