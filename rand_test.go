package utils

import (
	"fmt"
	"testing"
)

func TestWeightedChoice(t *testing.T) {
	choices := make([]TChoice, 10)
	count := make(map[TChoice]int)

	for i := 0; i < 10; i++ {
		c := TChoice{
			Item:   i,
			Weight: i * i,
		}
		choices[i] = c
		count[c] = 0
	}

	for i := 0; i < 100000; i++ {
		c := WeightedChoice(choices...)
		count[c] += 1
	}

	for i, c := range choices[0:9] {
		next := choices[i+1]
		AssertEqual(t, true, count[c] < count[next])
	}

	AssertEqual(t, nil, WeightedChoice().Item)
}

func TestWeightedChoiceWeightsIndex(t *testing.T) {
	count := make(map[interface{}]int)
	items := make([]interface{}, 10)
	weights := make([]int, 10)
	for i := 0; i < 10; i++ {
		items[i] = fmt.Sprintf("Item.%d", i)
		weights[i] = i * i
	}

	for i := 0; i < 100000; i++ {
		idx := WeightedChoiceWeightsIndex(weights)
		count[idx] += 1
		_ = items[idx]
	}

	for i := range items[0:9] {
		AssertEqual(t, true, count[i] < count[i+1])
	}

	AssertEqual(t, -1, WeightedChoiceWeightsIndex(nil))
	AssertEqual(t, -1, WeightedChoiceWeightsIndex([]int{-1, 0}))
}

func TestWeightedChoiceMap(t *testing.T) {
	choices := make(map[interface{}]int)
	count := make(map[interface{}]int)
	items := make([]interface{}, 10)

	for i := 0; i < 10; i++ {
		item := fmt.Sprintf("Item.%d", i)
		choices[item] = i * i
		count[item] = 0
		items[i] = item
	}

	for i := 0; i < 100000; i++ {
		c := WeightedChoiceMap(choices)
		count[c] += 1
	}

	for i, c := range items[0:9] {
		next := items[i+1]
		AssertEqual(t, true, count[c] < count[next])
	}

	AssertEqual(t, nil, WeightedChoiceMap(nil))
	AssertEqual(t, nil, WeightedChoiceMap(map[interface{}]int{"A": 0}))
	AssertEqual(t, nil, WeightedChoiceMap(map[interface{}]int{"A": -1}))
}

func TestRandInt(t *testing.T) {
	AssertEqual(t, true, RandInt(1, 2) == 1)
	AssertEqual(t, true, RandInt(-1, 0) == -1)
	AssertEqual(t, true, RandInt(0, 5) >= 0)
	AssertEqual(t, true, RandInt(0, 5) < 5)
}

func TestRandString(t *testing.T) {
	AssertEqual(t, true, len(RandString(16)) == 16)
}

func TestRandHex(t *testing.T) {
	AssertEqual(t, true, len(RandHex(16)) == 32)
}

func BenchmarkWeightedChoice(b *testing.B) {
	choices := make([]TChoice, 20)
	for i := 0; i < 20; i++ {
		c := TChoice{
			Item:   i,
			Weight: i,
		}
		choices[i] = c
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeightedChoice(choices...)
	}
}

func BenchmarkWeightedChoiceWeightsIndex(b *testing.B) {
	choices := make([]interface{}, 20)
	weights := make([]int, 20)
	for i := 0; i < 20; i++ {
		choices[i] = fmt.Sprintf("Item.%d", i)
		weights[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = choices[WeightedChoiceWeightsIndex(weights)]
	}
}

func BenchmarkWeightedChoiceMap(b *testing.B) {
	choices := make(map[interface{}]int)
	for i := 0; i < 20; i++ {
		choices[&i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WeightedChoiceMap(choices)
	}
}

// BenchmarkWeightedChoice-8               	16799522	        74.7 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoice-8               	15464216	        79.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoice-8               	16557890	        71.8 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceWeightsIndex-8   	18043592	        75.4 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceWeightsIndex-8   	14749660	        71.2 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceWeightsIndex-8   	15543214	        81.1 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceMap-8            	10280749	       129 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceMap-8            	10496216	       126 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWeightedChoiceMap-8            	10238050	       115 ns/op	       0 B/op	       0 allocs/op
