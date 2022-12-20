//go:build go1.18
// +build go1.18

package slices

func Deduplication[T comparable](ss []T) []T {
	m := make(map[T]struct{}, len(ss))
	j := 0
	for i := range ss {
		if _, ok := m[ss[i]]; ok {
			continue
		}
		m[ss[i]] = struct{}{}
		if i != j {
			ss[j] = ss[i]
		}
		j++
	}
	return ss[:j]
}
