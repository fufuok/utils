//go:build go1.18
// +build go1.18

package xslices

// Deduplication removes repeatable elements from s.
func Deduplication[E comparable](s []E) []E {
	m := make(map[E]struct{}, len(s))
	j := 0
	for i := range s {
		if _, ok := m[s[i]]; ok {
			continue
		}
		m[s[i]] = struct{}{}
		if i != j {
			s[j] = s[i]
		}
		j++
	}
	return s[:j]
}
