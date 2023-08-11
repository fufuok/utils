//go:build go1.18
// +build go1.18

package xslices

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestDeduplication(t *testing.T) {
	in := []int{7, -1, 0, 0, 1, 2, 1, -1, -1}
	want := []int{7, -1, 0, 1, 2}
	assert.Equal(t, want, Deduplication(in))

	in = []int{7, -1, 0, 1, 2}
	want = []int{7, -1, 0, 1, 2}
	assert.Equal(t, want, Deduplication(in))
}
