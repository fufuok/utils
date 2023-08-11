//go:build go1.18
// +build go1.18

package xslices

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestMerge(t *testing.T) {
	s1 := make([]int, 0, 10)
	s2 := make([]int, 0, 10)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
	}
	for i := 5; i < 10; i++ {
		s2 = append(s2, i)
	}
	// s3 := append(s1, s2...) // warning
	s := Merge(s1, s2)
	s[0] = 11
	s[5] = 22
	assert.Equal(t, []int{0, 1, 2, 3, 4}, s1)
	assert.Equal(t, []int{5, 6, 7, 8, 9}, s2)
	assert.Equal(t, []int{11, 1, 2, 3, 4, 22, 6, 7, 8, 9}, s)

	assert.Equal(t, s1, Merge(s1))
	assert.Equal(t, s2, Merge(s2))

	var x []int
	assert.Nil(t, Merge(x))
	x = []int{}
	assert.Equal(t, x, Merge(x))
}
