//go:build go1.18
// +build go1.18

package xsync_test

import (
	"hash/maphash"
	"strconv"
	"testing"

	. "github.com/fufuok/utils/xsync"
)

func TestHashMapOf_StructKey_CustomHasher(t *testing.T) {
	const num = 200
	type location struct {
		lon float32
		lat float32
	}
	hasher := func(_ maphash.Seed, l location) uint64 {
		return uint64(6*l.lon + 7*l.lat)
	}
	m := NewHashMapOf[location, int](hasher)

	for i := 0; i < num; i++ {
		m.Store(location{float32(i), float32(-i)}, i)
	}
	for i := 0; i < num; i++ {
		v, ok := m.Load(location{float32(i), float32(-i)})
		if !ok {
			t.Errorf("value not found: %d", i)
		}
		if v != i {
			t.Errorf("values do not match, %d: %v", i, v)
		}
	}
}

func TestHashMapOf_StructKey_GenHasher(t *testing.T) {
	const num = 200
	type location struct {
		lon float32
		lat float32
	}
	// Warning: panic: unsupported key type xsync_test.location of kind struct
	// m := NewHashMapOf[location, int]()
	m := NewHashMapOf[*location, int]()
	keys := make([]*location, 0, num)

	for i := 0; i < num; i++ {
		key := &location{float32(i), float32(-i)}
		m.Store(key, i)
		keys = append(keys, key)
	}
	i := 0
	for _, k := range keys {
		v, ok := m.Load(k)
		if !ok {
			t.Errorf("value not found: %d", i)
		}
		if v != i {
			t.Errorf("values do not match, %d: %v", i, v)
		}
		i++
	}
	for i := 0; i < num; i++ {
		v, ok := m.Load(keys[i])
		if !ok {
			t.Errorf("value not found: %d", i)
		}
		if v != i {
			t.Errorf("values do not match, %d: %v", i, v)
		}
	}
	key := &location{float32(1), float32(-1)}
	v, ok := m.Load(key)
	if ok {
		t.Errorf("find value, expect or not: %v", v)
	}
}

func TestHashMapOf_UniqueValuePointers_Int(t *testing.T) {
	EnableAssertions()
	m := NewHashMapOf[string, int]()
	v := 42
	m.Store("foo", v)
	m.Store("foo", v)
	DisableAssertions()
}

func TestHashMapOf_UniqueValuePointers_Struct(t *testing.T) {
	type foo struct{}
	EnableAssertions()
	m := NewHashMapOf[string, foo]()
	v := foo{}
	m.Store("foo", v)
	m.Store("foo", v)
	DisableAssertions()
}

func TestHashMapOf_UniqueValuePointers_Pointer(t *testing.T) {
	type foo struct{}
	EnableAssertions()
	m := NewHashMapOf[string, *foo]()
	v := &foo{}
	m.Store("foo", v)
	m.Store("foo", v)
	DisableAssertions()
}

func TestHashMapOf_UniqueValuePointers_Slice(t *testing.T) {
	EnableAssertions()
	m := NewHashMapOf[string, []int]()
	v := make([]int, 13)
	m.Store("foo", v)
	m.Store("foo", v)
	DisableAssertions()
}

func TestHashMapOf_UniqueValuePointers_String(t *testing.T) {
	EnableAssertions()
	m := NewHashMapOf[string, string]()
	v := "bar"
	m.Store("foo", v)
	m.Store("foo", v)
	DisableAssertions()
}

func TestHashMapOf_UniqueValuePointers_Nil(t *testing.T) {
	EnableAssertions()
	m := NewHashMapOf[string, *struct{}]()
	m.Store("foo", nil)
	m.Store("foo", nil)
	DisableAssertions()
}

func TestHashMapOf_MissingEntry(t *testing.T) {
	m := NewHashMapOf[string, string]()
	v, ok := m.Load("foo")
	if ok {
		t.Errorf("value was not expected: %v", v)
	}
	if deleted, loaded := m.LoadAndDelete("foo"); loaded {
		t.Errorf("value was not expected %v", deleted)
	}
	if actual, loaded := m.LoadOrStore("foo", "bar"); loaded {
		t.Errorf("value was not expected %v", actual)
	}
}

func TestHashMapOf_EmptyStringKey(t *testing.T) {
	m := NewHashMapOf[string, string]()
	m.Store("", "foobar")
	v, ok := m.Load("")
	if !ok {
		t.Error("value was expected")
	}
	if v != "foobar" {
		t.Errorf("value does not match: %v", v)
	}
}

func TestHashMapOf_Store_NilValue(t *testing.T) {
	m := NewHashMapOf[string, *struct{}]()
	m.Store("foo", nil)
	v, ok := m.Load("foo")
	if !ok {
		t.Error("nil value was expected")
	}
	if v != nil {
		t.Errorf("value was not nil: %v", v)
	}
}

func TestHashMapOf_LoadOrStore_NilValue(t *testing.T) {
	m := NewHashMapOf[string, *struct{}]()
	m.LoadOrStore("foo", nil)
	v, loaded := m.LoadOrStore("foo", nil)
	if !loaded {
		t.Error("nil value was expected")
	}
	if v != nil {
		t.Errorf("value was not nil: %v", v)
	}
}

func TestHashMapOf_LoadOrStore_NonNilValue(t *testing.T) {
	type foo struct{}
	m := NewHashMapOf[string, *foo]()
	newv := &foo{}
	v, loaded := m.LoadOrStore("foo", newv)
	if loaded {
		t.Error("no value was expected")
	}
	if v != newv {
		t.Errorf("value does not match: %v", v)
	}
	newv2 := &foo{}
	v, loaded = m.LoadOrStore("foo", newv2)
	if !loaded {
		t.Error("value was expected")
	}
	if v != newv {
		t.Errorf("value does not match: %v", v)
	}
}

func TestHashMapOf_LoadAndStore_NilValue(t *testing.T) {
	m := NewHashMapOf[string, *struct{}]()
	m.LoadAndStore("foo", nil)
	v, loaded := m.LoadAndStore("foo", nil)
	if !loaded {
		t.Error("nil value was expected")
	}
	if v != nil {
		t.Errorf("value was not nil: %v", v)
	}
	v, loaded = m.Load("foo")
	if !loaded {
		t.Error("nil value was expected")
	}
	if v != nil {
		t.Errorf("value was not nil: %v", v)
	}
}

func TestHashMapOf_LoadAndStore_NonNilValue(t *testing.T) {
	m := NewHashMapOf[string, int]()
	v1 := 1
	v, loaded := m.LoadAndStore("foo", v1)
	if loaded {
		t.Error("no value was expected")
	}
	if v != v1 {
		t.Errorf("value does not match: %v", v)
	}
	v2 := 2
	v, loaded = m.LoadAndStore("foo", v2)
	if !loaded {
		t.Error("value was expected")
	}
	if v != v1 {
		t.Errorf("value does not match: %v", v)
	}
	v, loaded = m.Load("foo")
	if !loaded {
		t.Error("value was expected")
	}
	if v != v2 {
		t.Errorf("value does not match: %v", v)
	}
}

func TestHashMapOf_Range(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	iters := 0
	met := make(map[string]int)
	m.Range(func(key string, value int) bool {
		if key != strconv.Itoa(value) {
			t.Errorf("got unexpected key/value for iteration %d: %v/%v", iters, key, value)
			return false
		}
		met[key] += 1
		iters++
		return true
	})
	if iters != numEntries {
		t.Errorf("got unexpected number of iterations: %d", iters)
	}
	for i := 0; i < numEntries; i++ {
		if c := met[strconv.Itoa(i)]; c != 1 {
			t.Errorf("range did not iterate correctly over %d: %d", i, c)
		}
	}
}

func TestHashMapOf_Range_FalseReturned(t *testing.T) {
	m := NewHashMapOf[string, int]()
	for i := 0; i < 100; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	iters := 0
	m.Range(func(key string, value int) bool {
		iters++
		return iters != 13
	})
	if iters != 13 {
		t.Errorf("got unexpected number of iterations: %d", iters)
	}
}

func TestHashMapOf_Range_NestedDelete(t *testing.T) {
	const numEntries = 256
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	m.Range(func(key string, value int) bool {
		m.Delete(key)
		return true
	})
	for i := 0; i < numEntries; i++ {
		if _, ok := m.Load(strconv.Itoa(i)); ok {
			t.Errorf("value found for %d", i)
		}
	}
}

func TestHashMapOf_SerialStore(t *testing.T) {
	const numEntries = 128
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	for i := 0; i < numEntries; i++ {
		v, ok := m.Load(strconv.Itoa(i))
		if !ok {
			t.Errorf("value not found for %d", i)
		}
		if v != i {
			t.Errorf("values do not match for %d: %v", i, v)
		}
	}
}

func TestHashMapOf_IntegerMapOfSerialStore(t *testing.T) {
	const numEntries = 128
	m := NewHashMapOf[int, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(i, i)
	}
	for i := 0; i < numEntries; i++ {
		v, ok := m.Load(i)
		if !ok {
			t.Errorf("value not found for %d", i)
		}
		if v != i {
			t.Errorf("values do not match for %d: %v", i, v)
		}
	}
}

func TestHashMapOf_IntegerMapOfSerialStore_WithHasher(t *testing.T) {
	const numEntries = 128
	m := NewHashMapOf[int, int](func(_ maphash.Seed, k int) uint64 {
		return uint64(k)
	})
	for i := 0; i < numEntries; i++ {
		m.Store(i, i)
	}
	for i := 0; i < numEntries; i++ {
		v, ok := m.Load(i)
		if !ok {
			t.Errorf("value not found for %d", i)
		}
		if v != i {
			t.Errorf("values do not match for %d: %v", i, v)
		}
	}
}

func TestHashMapOf_SerialLoadOrStore(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	for i := 0; i < numEntries; i++ {
		if _, loaded := m.LoadOrStore(strconv.Itoa(i), i); !loaded {
			t.Errorf("value not found for %d", i)
		}
	}
}

func TestHashMapOf_SerialLoadOrCompute(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		v, loaded := m.LoadOrCompute(strconv.Itoa(i), func() int {
			return i
		})
		if loaded {
			t.Errorf("value not computed for %d", i)
		}
		if v != i {
			t.Errorf("values do not match for %d: %v", i, v)
		}
	}
	for i := 0; i < numEntries; i++ {
		v, loaded := m.LoadOrCompute(strconv.Itoa(i), func() int {
			return i
		})
		if !loaded {
			t.Errorf("value not loaded for %d", i)
		}
		if v != i {
			t.Errorf("values do not match for %d: %v", i, v)
		}
	}
}

func TestHashMapOf_SerialStoreThenDelete(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	for i := 0; i < numEntries; i++ {
		m.Delete(strconv.Itoa(i))
		if _, ok := m.Load(strconv.Itoa(i)); ok {
			t.Errorf("value was not expected for %d", i)
		}
	}
}

func TestHashMapOf_IntegerMapOfSerialStoreThenDelete(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[int32, int32]()
	for i := 0; i < numEntries; i++ {
		m.Store(int32(i), int32(i))
	}
	for i := 0; i < numEntries; i++ {
		m.Delete(int32(i))
		if _, ok := m.Load(int32(i)); ok {
			t.Errorf("value was not expected for %d", i)
		}
	}
}

func TestHashMapOf_SerialStoreThenLoadAndDelete(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	for i := 0; i < numEntries; i++ {
		if _, loaded := m.LoadAndDelete(strconv.Itoa(i)); !loaded {
			t.Errorf("value was not found for %d", i)
		}
		if _, ok := m.Load(strconv.Itoa(i)); ok {
			t.Errorf("value was not expected for %d", i)
		}
	}
}

func TestHashMapOf_IntegerMapOfSerialStoreThenLoadAndDelete(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[int, int]()
	for i := 0; i < numEntries; i++ {
		m.Store(i, i)
	}
	for i := 0; i < numEntries; i++ {
		if _, loaded := m.LoadAndDelete(i); !loaded {
			t.Errorf("value was not found for %d", i)
		}
		if _, ok := m.Load(i); ok {
			t.Errorf("value was not expected for %d", i)
		}
	}
}

func TestHashMapOf_Size(t *testing.T) {
	const numEntries = 1000
	m := NewHashMapOf[string, int]()
	size := m.Size()
	if size != 0 {
		t.Errorf("zero size expected: %d", size)
	}
	expectedSize := 0
	for i := 0; i < numEntries; i++ {
		m.Store(strconv.Itoa(i), i)
		expectedSize++
		size := m.Size()
		if size != expectedSize {
			t.Errorf("size of %d was expected, got: %d", expectedSize, size)
		}
	}
	for i := 0; i < numEntries; i++ {
		m.Delete(strconv.Itoa(i))
		expectedSize--
		size := m.Size()
		if size != expectedSize {
			t.Errorf("size of %d was expected, got: %d", expectedSize, size)
		}
	}
}

func BenchmarkHashMapOf_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := NewHashMapOf[string, int]()
			benchmarkMapOfStringKeys(b, func(k string) (int, bool) {
				return m.Load(k)
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkHashMapOf_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := NewHashMapOf[string, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(benchmarkKeyPrefix+strconv.Itoa(i), i)
			}
			benchmarkMapOfStringKeys(b, func(k string) (int, bool) {
				return m.Load(k)
			}, func(k string, v int) {
				m.Store(k, v)
			}, func(k string) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkIntegerHashMapOf_NoWarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		if bc.readPercentage == 100 {
			// This benchmark doesn't make sense without a warm-up.
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			m := NewHashMapOf[int, int]()
			benchmarkMapOfIntegerKeys(b, func(k int) (int, bool) {
				return m.Load(k)
			}, func(k int, v int) {
				m.Store(k, v)
			}, func(k int) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}

func BenchmarkIntegerHashMapOf_WarmUp(b *testing.B) {
	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			m := NewHashMapOf[int, int]()
			for i := 0; i < benchmarkNumEntries; i++ {
				m.Store(i, i)
			}
			benchmarkMapOfIntegerKeys(b, func(k int) (int, bool) {
				return m.Load(k)
			}, func(k int, v int) {
				m.Store(k, v)
			}, func(k int) {
				m.Delete(k)
			}, bc.readPercentage)
		})
	}
}
