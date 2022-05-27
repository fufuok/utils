package jsongen

import (
	"encoding/json"
	"testing"
)

func BenchmarkGenJsonStandard(b *testing.B) {
	b.StopTimer()
	m := make(map[string]interface{})

	var raw map[string]interface{}
	_ = json.Unmarshal([]byte(jsStr), &raw)

	m1 := make(map[string]interface{})
	m1["uintkey"] = 123
	m1["intkey"] = -45
	m1["floatkey"] = 12.34
	m1["boolkey"] = true
	m1["stringkey1"] = "teststring"
	m1["stringkey2"] = `string with \`
	m1["stringkey3"] = `string with "`
	m1["raw_string"] = raw
	m1["raw_bytes"] = raw
	m1["raw_sarr"] = []map[string]interface{}{raw, raw}
	m1["raw_barr"] = []map[string]interface{}{raw, raw}
	m["map1"] = m1

	m2 := make(map[string]interface{})
	m2["uintarray"] = []uint64{123, 456, 789}
	m2["intarray"] = []int64{-23, -45, -89}
	m2["floatarray"] = []float64{12.34, -56.78, 90}
	m2["boolarray"] = []bool{true, false, true}
	m2["stringarray"] = []string{"test string", `string with \`, `string with "`}
	m["map2"] = m2

	m3 := make(map[string]interface{})
	{
		a1 := []interface{}{123, -45, 12.34, true, "test string", `string with \`, `string with "`}
		a2 := []interface{}{[]uint64{123, 456, 789}, []int64{-12, -45, -78}, []float64{12.34, -56.78, 9.0}, []bool{true, false, true}}
		a3 := []interface{}{
			map[string]interface{}{
				"uintkey":   123,
				"intkey":    -456,
				"floatkey":  12.34,
				"boolkey":   true,
				"stringkey": "test string",
			},
			map[string]interface{}{
				"uintkey":   455,
				"intkey":    -789,
				"floatkey":  56.78,
				"boolkey":   false,
				"stringkey": `string with \`,
			},
		}
		m3["array1"] = a1
		m3["array2"] = a2
		m3["array3"] = a3
	}
	{
		a1 := []interface{}{123, -45, 12.34, true, "test string", `string with \`, `string with "`}
		a2 := []interface{}{[]uint64{123, 456, 789}, []int64{-12, -45, -78}, []float64{12.34, -56.78, 9.0}, []bool{true, false, true}}
		a3 := []interface{}{
			map[string]interface{}{
				"uintkey":   123,
				"intkey":    -456,
				"floatkey":  12.34,
				"boolkey":   true,
				"stringkey": "test string",
			},
			map[string]interface{}{
				"uintkey":   455,
				"intkey":    -789,
				"floatkey":  56.78,
				"boolkey":   false,
				"stringkey": `string with \`,
			},
		}
		m3["array4"] = []interface{}{a1, a2, a3}
	}
	m["map3"] = m3

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		json.Marshal(m)
	}
}

func BenchmarkGenJson(b *testing.B) {
	b.StopTimer()
	m, _ := map4()
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		m.Serialize(nil)
	}
}

// go test -run=^$ -benchmem -count=2 -bench=^BenchmarkGenJson
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/jsongen
// cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
// BenchmarkGenJsonStandard-4         40786             29439 ns/op           10257 B/op        195 allocs/op
// BenchmarkGenJsonStandard-4         40789             29188 ns/op           10257 B/op        195 allocs/op
// BenchmarkGenJson-4                588708              2026 ns/op            1792 B/op          1 allocs/op
// BenchmarkGenJson-4                585769              2026 ns/op            1792 B/op          1 allocs/op
