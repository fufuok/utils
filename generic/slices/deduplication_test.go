//go:build go1.18
// +build go1.18

package slices

import (
	"testing"
)

func TestDeduplication(t *testing.T) {
	in := []int{7, -1, 0, 0, 1, 2, 1, -1, -1}
	want := []int{7, -1, 0, 1, 2}
	got := Deduplication(in)
	if !Equal(want, got) {
		t.Fatalf("Deduplication(%v) = %v, want %v", in, got, want)
	}

	in = []int{7, -1, 0, 1, 2}
	want = []int{7, -1, 0, 1, 2}
	got = Deduplication(in)
	if !Equal(want, got) {
		t.Fatalf("Deduplication(%v) = %v, want %v", in, got, want)
	}
}
