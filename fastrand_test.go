package utils

import (
	"testing"
)

func TestNewRand(t *testing.T) {
	rd := NewRand(1)
	AssertEqual(t, int64(5577006791947779410), rd.Int63())

	rd = NewRand()
	for i := 0; i < 1000; i++ {
		AssertEqual(t, 0, rd.Intn(1))
		AssertEqual(t, int64(0), rd.Int63n(1))
		AssertEqual(t, 0, Rand.Intn(1))
		AssertEqual(t, int64(0), Rand.Int63n(1))
	}
}
