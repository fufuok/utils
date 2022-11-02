package xsync

import (
	"hash/maphash"
)

func HashUint64[K IntegerConstraint](seed maphash.Seed, k K) uint64 {
	return hashUint64(seed, k)
}
