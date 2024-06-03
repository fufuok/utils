//go:build go1.18
// +build go1.18

package xslices

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestToString(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	assert.Equal(t, "12345", ToString(xs, ""))
	assert.Equal(t, "1.2.3.4.5", ToString(xs, "."))
	assert.Equal(t, "1--2--3--4--5", ToString(xs, "--"))
	assert.Equal(t, "", ToString([]int{}, "."))
	assert.Equal(t, "0", ToString([]float64{0.0}, "."))
	assert.Equal(t, "0,1.2", ToString([]float64{0.0, 1.2}, ","))
}

func TestAverage(t *testing.T) {
	xs := []uint32{1, 2, 3, 4}
	assert.Equal(t, 2.5, Average(xs))
	assert.Equal(t, float64(0), Average([]float32{0.0}))
	assert.Equal(t, 0.6, Average([]float64{0.0, 1.2}))

	fs := []float64{0.0, 0.5, 0.5}
	assert.Equal(t, 0.33, Average(fs, 2))
	fs = []float64{0.0, 1, 1}
	assert.Equal(t, 0.667, Average(fs, 3))
}
