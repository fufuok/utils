package utils

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestNewRand(t *testing.T) {
	rd := NewRand(1)
	assert.Equal(t, int64(5577006791947779410), rd.Int63())

	rd = NewRand()
	for i := 1; i < 1000; i++ {
		assert.Equal(t, true, rd.Intn(i) < i)
		assert.Equal(t, true, rd.Int63n(int64(i)) < int64(i))
		assert.Equal(t, true, Rand.Intn(i) < i)
		assert.Equal(t, true, Rand.Int63n(int64(i)) < int64(i))
	}
}
