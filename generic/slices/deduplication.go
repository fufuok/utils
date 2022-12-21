//go:build go1.18
// +build go1.18

package slices

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

// Filter removes any elements from s for which pred(element) is false.
func Filter[E any, S ~[]E](s S, pred func(E) bool) S {
	j := 0
	for i := range s {
		if pred(s[i]) {
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}
