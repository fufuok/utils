//go:build go1.18
// +build go1.18

package xsync

const (
	EntriesPerMapOfBucket   = entriesPerMapOfBucket
	DefaultMinMapOfTableCap = defaultMinMapTableLen * entriesPerMapOfBucket
)

type (
	BucketOfPadded = bucketOfPadded
)

func DefaultHasher[T comparable]() func(T, uint64) uint64 {
	return defaultHasher[T]()
}
