//go:build go1.18
// +build go1.18

package xslices

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
