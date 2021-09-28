package utils

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	AssertEqual(t, true, RandInt(1, 2) == 1)
	AssertEqual(t, true, RandInt(-1, 0) == -1)
	AssertEqual(t, true, RandInt(0, 5) >= 0)
	AssertEqual(t, true, RandInt(0, 5) < 5)
	AssertEqual(t, 0, RandInt(2, 2))
	AssertEqual(t, 0, RandInt(3, 2))
}

func TestRandString(t *testing.T) {
	AssertEqual(t, true, len(RandString(16)) == 16)
}

func TestRandHex(t *testing.T) {
	AssertEqual(t, true, len(RandHex(16)) == 32)
}
