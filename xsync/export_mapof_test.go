//go:build go1.18
// +build go1.18

package xsync

type (
	BucketOfPadded = bucketOfPadded
)

func MakeHasher[T comparable]() func(T, uint64) uint64 {
	return makeHasher[T]()
}

func CollectMapOfStats[K comparable, V any](m *MapOf[K, V]) MapStats {
	return MapStats{m.stats()}
}

func NewMapOfWithHasher[K comparable, V any](
	hasher func(K, uint64) uint64,
	options ...func(*MapConfig),
) *MapOf[K, V] {
	return newMapOf[K, V](hasher, options...)
}
