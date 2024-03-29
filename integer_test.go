package utils

import (
	"math"
	"math/big"
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestMinInt(t *testing.T) {
	assert.Equal(t, -1, MinInt(1, -1))
	assert.Equal(t, 0, MinInt(0, 1))
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 1, MaxInt(1, -1))
	assert.Equal(t, 1, MaxInt(0, 1))
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, 1, GetInt("1"))
	assert.Equal(t, 1, GetInt("1", 2))
	assert.Equal(t, 1, GetInt(nil, 1))
	assert.Equal(t, 1, GetInt(0, 1))
	assert.Equal(t, 1, GetInt(-1, 1))
	assert.Equal(t, 0, GetInt(-1, 0))
	assert.Equal(t, -1, GetInt(-1))
}

func TestSearchInt(t *testing.T) {
	val := []int{1, 2, 3}
	assert.Equal(t, 0, SearchInt(val, 1))
	assert.Equal(t, 1, SearchInt(val, 2))
	assert.Equal(t, 2, SearchInt(val, 3))
	assert.Equal(t, -1, SearchInt(val, 4))
}

func TestInInts(t *testing.T) {
	val := []int{1, 2, 3}
	assert.Equal(t, true, InInts(val, 1))
	assert.Equal(t, false, InInts(val, 4))
}

// Ref: dustin/go-humanize
func TestComma(t *testing.T) {
	for _, v := range []struct {
		title    string
		actual   string
		expected string
	}{
		{"0", Comma(0), "0"},
		{"10", Comma(10), "10"},
		{"100", Comma(100), "100"},
		{"1,000", Comma(1000), "1,000"},
		{"10,000", Comma(10000), "10,000"},
		{"100,000", Comma(100000), "100,000"},
		{"10,000,000", Comma(10000000), "10,000,000"},
		{"10,100,000", Comma(10100000), "10,100,000"},
		{"10,010,000", Comma(10010000), "10,010,000"},
		{"10,001,000", Comma(10001000), "10,001,000"},
		{"123,456,789", Comma(123456789), "123,456,789"},
		{"maxint", Comma(9.223372e+18), "9,223,372,000,000,000,000"},
		{"math.maxint", Comma(math.MaxInt64), "9,223,372,036,854,775,807"},
		{"math.minint", Comma(math.MinInt64), "-9,223,372,036,854,775,808"},
		{"minint", Comma(-9.223372e+18), "-9,223,372,000,000,000,000"},
		{"-123,456,789", Comma(-123456789), "-123,456,789"},
		{"-10,100,000", Comma(-10100000), "-10,100,000"},
		{"-10,010,000", Comma(-10010000), "-10,010,000"},
		{"-10,001,000", Comma(-10001000), "-10,001,000"},
		{"-10,000,000", Comma(-10000000), "-10,000,000"},
		{"-100,000", Comma(-100000), "-100,000"},
		{"-10,000", Comma(-10000), "-10,000"},
		{"-1,000", Comma(-1000), "-1,000"},
		{"-100", Comma(-100), "-100"},
		{"-10", Comma(-10), "-10"},

		{"123,456,789", Commai(123456789), "123,456,789"},
	} {
		assert.Equal(t, v.expected, v.actual, v.title)
	}
}

// Ref: dustin/go-humanize
func TestBigComma(t *testing.T) {
	for _, v := range []struct {
		title    string
		actual   string
		expected string
	}{
		{"0", BigComma(big.NewInt(0)), "0"},
		{"10", BigComma(big.NewInt(10)), "10"},
		{"100", BigComma(big.NewInt(100)), "100"},
		{"1,000", BigComma(big.NewInt(1000)), "1,000"},
		{"10,000", BigComma(big.NewInt(10000)), "10,000"},
		{"100,000", BigComma(big.NewInt(100000)), "100,000"},
		{"10,000,000", BigComma(big.NewInt(10000000)), "10,000,000"},
		{"10,100,000", BigComma(big.NewInt(10100000)), "10,100,000"},
		{"10,010,000", BigComma(big.NewInt(10010000)), "10,010,000"},
		{"10,001,000", BigComma(big.NewInt(10001000)), "10,001,000"},
		{"123,456,789", BigComma(big.NewInt(123456789)), "123,456,789"},
		{"maxint", BigComma(big.NewInt(9.223372e+18)), "9,223,372,000,000,000,000"},
		{"minint", BigComma(big.NewInt(-9.223372e+18)), "-9,223,372,000,000,000,000"},
		{"-123,456,789", BigComma(big.NewInt(-123456789)), "-123,456,789"},
		{"-10,100,000", BigComma(big.NewInt(-10100000)), "-10,100,000"},
		{"-10,010,000", BigComma(big.NewInt(-10010000)), "-10,010,000"},
		{"-10,001,000", BigComma(big.NewInt(-10001000)), "-10,001,000"},
		{"-10,000,000", BigComma(big.NewInt(-10000000)), "-10,000,000"},
		{"-100,000", BigComma(big.NewInt(-100000)), "-100,000"},
		{"-10,000", BigComma(big.NewInt(-10000)), "-10,000"},
		{"-1,000", BigComma(big.NewInt(-1000)), "-1,000"},
		{"-100", BigComma(big.NewInt(-100)), "-100"},
		{"-10", BigComma(big.NewInt(-10)), "-10"},
	} {
		assert.Equal(t, v.expected, v.actual, v.title)
	}
}

func TestCommau(t *testing.T) {
	var u0 uint64
	var u1 uint64 = 1111111111
	assert.Equal(t, Commau(u0), "0")
	assert.Equal(t, Commau(u1), "1,111,111,111")
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma(1234567890)
	}
}

func BenchmarkBigComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigComma(big.NewInt(1234567890))
	}
}

func TestParseInts(t *testing.T) {
	var empty []int
	for _, v := range []struct {
		s    string
		want []int
		ok   bool
	}{
		{"", empty, false},
		{"3,1,3", []int{1, 3}, true},
		{"3  \t\n,1- 5, 3\r-4", []int{1, 2, 3, 4, 5}, true},
	} {
		got, err := ParseInts(v.s)
		assert.Equal(t, v.want, got)
		assert.Equal(t, v.ok, err == nil)
	}
}
