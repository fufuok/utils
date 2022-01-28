// Package jsongen forked from darjun/json-gen
package jsongen

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnquotedValue(t *testing.T) {
	testCases := []struct {
		value    UnquotedValue
		expected string
	}{
		{"1234", "1234"},
		{"12.34", "12.34"},
	}

	for _, c := range testCases {
		if string(c.value.Serialize(nil)) != c.expected {
			t.Errorf("actual(%s) != expected(%s)", string(c.value.Serialize(nil)), c.expected)
		}
	}
}

func TestQuotedValue(t *testing.T) {
	testCases := []struct {
		value    QuotedValue
		expected string
	}{
		{QuotedValue("string"), `"string"`},
		{QuotedValue(`string with \`), `"string with \"`},
		{QuotedValue(`string with "`), `"string with ""`},
	}

	for _, c := range testCases {
		if string(c.value.Serialize(nil)) != c.expected {
			t.Errorf("actual(%s) != expected(%s)", string(c.value.Serialize(nil)), c.expected)
		}
	}
}

func array1() (*Array, string) {
	a1 := NewArray()
	a1.AppendUint(123)
	a1.AppendInt(-45)
	a1.AppendFloat(12.34)
	a1.AppendBool(true)
	a1.AppendString("test string")
	a1.AppendString(`string with \`)
	a1.AppendString(`string with "`)
	expected1 := `[123,-45,12.34,true,"test string","string with \\","string with \""]`

	return a1, expected1
}

func array2() (*Array, string) {
	a2 := NewArray()
	a2.AppendUintArray([]uint64{123, 456, 789})
	a2.AppendIntArray([]int64{-12, -45, -78})
	a2.AppendFloatArray([]float64{12.34, -56.78, 9.0})
	a2.AppendBoolArray([]bool{true, false, true})
	a2.AppendStringArray([]string{"test string", `string with \`, `string with "`})
	expected2 := `[[123,456,789],[-12,-45,-78],[12.34,-56.78,9],[true,false,true],["test string","string with \\","string with \""]]`

	return a2, expected2
}

func array3() (*Array, string) {
	a3 := NewArray()
	m1 := NewMap()
	m1.PutUint("uintkey", 123)
	m1.PutInt("intkey", -456)
	m1.PutFloat("floatkey", 12.34)
	m1.PutBool("boolkey", true)
	m1.PutString("stringkey", "test string")
	a3.AppendMap(m1)

	m2 := NewMap()
	m2.PutUint("uintkey", 456)
	m2.PutInt("intkey", -789)
	m2.PutFloat("floatkey", 56.78)
	m2.PutBool("boolkey", false)
	m2.PutString("stringkey", `string with \`)
	a3.AppendMap(m2)
	expected3 := `[{"uintkey":123,"intkey":-456,"floatkey":12.34,"boolkey":true,"stringkey":"test string"},{"uintkey":456,"intkey":-789,"floatkey":56.78,"boolkey":false,"stringkey":"string with \\"}]`
	return a3, expected3
}

func array4() (*Array, string) {
	a4 := NewArray()

	a1, expected1 := array1()
	a2, expected2 := array2()
	a3, expected3 := array3()
	a4.AppendArray(*a1)
	a4.AppendArray(*a2)
	a4.AppendArray(*a3)
	expected4 := fmt.Sprintf("[%s,%s,%s]", expected1, expected2, expected3)

	return a4, expected4
}

func TestArrayValue(t *testing.T) {
	a1, expected1 := array1()

	a2, expected2 := array2()

	a3, expected3 := array3()

	a4, expected4 := array4()

	testCases := []struct {
		name     string
		value    *Array
		expected string
	}{
		{"basic", a1, expected1},
		{"primitive array", a2, expected2},
		{"map array", a3, expected3},
		{"nested general array", a4, expected4},
	}

	for _, c := range testCases {
		data := c.value.Serialize(nil)
		if string(data) != c.expected {
			t.Errorf("array name:%s actual:%s != expected:%s", c.name, string(data), c.expected)
		}

		if len(data) != c.value.Size() {
			t.Errorf("array name:%s buf size error, actual:%d, expected:%d", c.name, len(data), c.value.Size())
		}

		var obj []interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			t.Errorf("array name:%s unmarshal error:%v", c.name, err)
		} else {
			t.Logf("array name:%s unmarshal: %v", c.name, obj)
		}
	}
}

func map1() (*Map, string) {
	m1 := NewMap()
	m1.PutUint("uintkey", 123)
	m1.PutInt("intkey", -45)
	m1.PutFloat("floatkey", 12.34)
	m1.PutBool("boolkey", true)
	m1.PutString("stringkey1", "teststring")
	m1.PutString("stringkey2", `string with \`)
	m1.PutString("stringkey3", `string with "`)
	expected1 := `{"uintkey":123,"intkey":-45,"floatkey":12.34,"boolkey":true,"stringkey1":"teststring","stringkey2":"string with \\","stringkey3":"string with \""}`

	return m1, expected1
}

func map2() (*Map, string) {
	m2 := NewMap()
	m2.PutUintArray("uintarray", []uint64{123, 456, 789})
	m2.PutIntArray("intarray", []int64{-23, -45, -89})
	m2.PutFloatArray("floatarray", []float64{12.34, -56.78, 90})
	m2.PutBoolArray("boolarray", []bool{true, false, true})
	m2.PutStringArray("stringarray", []string{"test string", `string with \`, `string with "`})
	expected2 := `{"uintarray":[123,456,789],"intarray":[-23,-45,-89],"floatarray":[12.34,-56.78,90],"boolarray":[true,false,true],"stringarray":["test string","string with \\","string with \""]}`

	return m2, expected2
}

func map3() (*Map, string) {
	m3 := NewMap()

	a1, expected1 := array1()
	a2, expected2 := array2()
	a3, expected3 := array3()
	a4, expected4 := array4()

	m3.PutArray("array1", a1)
	m3.PutArray("array2", a2)
	m3.PutArray("array3", a3)
	m3.PutArray("array4", a4)

	expected := fmt.Sprintf(`{"array1":%s,"array2":%s,"array3":%s,"array4":%s}`, expected1, expected2, expected3, expected4)

	return m3, expected
}

func map4() (*Map, string) {
	m4 := NewMap()

	m1, expected1 := map1()
	m2, expected2 := map2()
	m3, expected3 := map3()

	m4.PutMap("map1", m1)
	m4.PutMap("map2", m2)
	m4.PutMap("map3", m3)

	expected := fmt.Sprintf(`{"map1":%s,"map2":%s,"map3":%s}`, expected1, expected2, expected3)

	return m4, expected
}

func TestMapValue(t *testing.T) {
	m1, expected1 := map1()
	m2, expected2 := map2()
	m3, expected3 := map3()
	m4, expected4 := map4()

	testCases := []struct {
		name     string
		value    *Map
		expected string
	}{
		{"basic", m1, expected1},
		{"primitive array", m2, expected2},
		{"general array", m3, expected3},
		{"nested map", m4, expected4},
	}

	for _, c := range testCases {
		data := c.value.Serialize(nil)
		if string(data) != c.expected {
			t.Errorf("map name:%s actual:%s != expected:%s", c.name, string(data), c.expected)
		}

		if len(data) != c.value.Size() {
			t.Errorf("map name:%s buf size error, actual:%d, expected:%d", c.name, len(data), c.value.Size())
		}

		var obj map[string]interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			t.Errorf("map name:%s unmarshal error:%v", c.name, err)
		} else {
			t.Logf("map name:%s unmarshal: %v", c.name, obj)
		}
	}
}
