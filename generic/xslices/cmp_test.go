//go:build go1.21
// +build go1.21

package xslices

import (
	"math"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestMaxMin(t *testing.T) {
	xs := []int{-1, 0, 1, 2}
	assert.Equal(t, 2, Max(xs...))
	assert.Equal(t, -1, Min(xs...))
	xu := []uint{1, 0, 1, 2}
	assert.Equal(t, uint(2), Max(xu...))
	assert.Equal(t, uint(0), Min(xu...))

	xf := []float64{-3.14, 0, 1, 2.1}
	assert.Equal(t, 2.1, Max(xf...))
	assert.Equal(t, -3.14, Min(xf...))
	x := math.NaN()
	xf = append(xf, x)
	assert.True(t, math.IsNaN(Max(xf...)))
	assert.True(t, math.IsNaN(Min(xf...)))
}
