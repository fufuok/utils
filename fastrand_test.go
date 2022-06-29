package utils

import (
	"testing"
)

func TestNewRand(t *testing.T) {
	rd := NewRand(1)
	AssertEqual(t, int64(5577006791947779410), rd.Int63())

	rd = NewRand()
	for i := 1; i < 1000; i++ {
		AssertEqual(t, true, rd.Intn(i) < i)
		AssertEqual(t, true, rd.Int63n(int64(i)) < int64(i))
		AssertEqual(t, true, Rand.Intn(i) < i)
		AssertEqual(t, true, Rand.Int63n(int64(i)) < int64(i))
	}
}
