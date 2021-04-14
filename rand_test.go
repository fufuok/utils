package utils

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	AssertEqual(t, true, RandInt(1, 2) == 1)
	AssertEqual(t, true, RandInt(-1, 0) == -1)
	AssertEqual(t, true, RandInt(0, 5) >= 0)
	AssertEqual(t, true, RandInt(0, 5) < 5)
}

func TestRandString(t *testing.T) {
	AssertEqual(t, true, len(RandString(16)) == 16)
}
