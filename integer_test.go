package utils

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	AssertEqual(t, -1, MinInt(1, -1))
}

func TestMaxInt(t *testing.T) {
	AssertEqual(t, 1, MaxInt(1, -1))
}

func TestGetInt(t *testing.T) {
	AssertEqual(t, 1, GetInt("1"))
	AssertEqual(t, 1, GetInt("1", 2))
	AssertEqual(t, 1, GetInt(nil, 1))
	AssertEqual(t, 1, GetInt(0, 1))
	AssertEqual(t, 1, GetInt(-1, 1))
	AssertEqual(t, 0, GetInt(-1, 0))
	AssertEqual(t, -1, GetInt(-1))
}

func TestSearchInt(t *testing.T) {
	val := []int{1, 2, 3}
	AssertEqual(t, 0, SearchInt(val, 1))
	AssertEqual(t, 1, SearchInt(val, 2))
	AssertEqual(t, 2, SearchInt(val, 3))
	AssertEqual(t, -1, SearchInt(val, 4))
}

func TestInInts(t *testing.T) {
	val := []int{1, 2, 3}
	AssertEqual(t, true, InInts(val, 1))
	AssertEqual(t, false, InInts(val, 4))
}
