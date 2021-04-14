package utils

import (
	"testing"
)

// Ref: gofiber/utils
func TestAssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}

func TestAssertPanics(t *testing.T) {
	t.Parallel()
	var a []int
	AssertPanics(nil, "should panic when index out of range", func() {
		_ = a[1]
	})
	AssertPanics(t, "should panic when index out of range", func() {
		_ = a[1]
	})
}
