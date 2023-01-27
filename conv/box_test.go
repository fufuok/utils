// Copyright 2023 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package conv

import (
	"fmt"
	"math"
	"sync"
	"testing"
	"time"
)

type Jello struct {
	Neat int
	Feet int
}

type Pudding struct {
	Neat int
	Feet int
}

func (p Pudding) Float64() float64 {
	return float64(p.Neat) * float64(p.Feet)
}
func (p Pudding) Uint64() uint64 {
	return uint64(p.Neat) * uint64(p.Feet)
}
func (p Pudding) Int64() int64 {
	return int64(p.Neat) * int64(p.Feet)
}
func (p Pudding) Bool() bool {
	return true
}

func (p Pudding) String() string {
	return fmt.Sprintf("Yum{%d %d}", p.Neat, p.Feet)
}

func TestValue(t *testing.T) {
	assert(t, Nil().Int() == 0)
	assert(t, Nil().IsNil() == true)
	assert(t, Nil().IsCustomBits() == false)
	assert(t, CustomBits(0).IsNil() == false)
	assert(t, CustomBits(0).Int() == 0)
	assert(t, CustomBits(1).Int() == 1)
	assert(t, CustomBits(1).Uint64() == 1)
	assert(t, Bool(true).Int() == 1)
	assert(t, Bool(false).Int() == 0)
	assert(t, Int(0).Bool() == false)
	assert(t, Int(1).Bool() == true)
	assert(t, Float64(math.NaN()).Bool() == false)
	assert(t, CustomBits(1).Bool() == true)
	assert(t, CustomBits(0).Bool() == false)
	assert(t, Int(1).Int() == 1)
	assert(t, Float64(1.0).Int() == 1)
	assert(t, Float64(math.NaN()).Int() == 0)
	assert(t, Float64(math.Inf(+1)).Int() == math.MaxInt64)
	assert(t, Float64(math.Inf(-1)).Int() == math.MinInt64)
	assert(t, Float64(math.Inf(+1)).Uint() == math.MaxUint64)
	assert(t, Float64(math.Inf(-1)).Uint() == 0)
	assert(t, Uint64(99).Int() == 99)
	assert(t, String("hello world").String() == "hello world")
	assert(t, String("hello world").Int() == 0)
	assert(t, String("hello world").IsNil() == false)
	assert(t, string(String("hello world").Bytes()) == "hello world")
	assert(t, Bytes([]byte("hello world")).String() == "hello world")
	assert(t, Bytes([]byte("hello world")).Int() == 0)
	assert(t, Bytes([]byte("hello world")).IsNil() == false)
	assert(t, string(Bytes([]byte("hello world")).Bytes()) == "hello world")
	forceIfaceStrs = true
	assert(t, String("hello world").String() == "hello world")
	assert(t, String("hello world").Int() == 0)
	assert(t, String("hello world").IsNil() == false)
	assert(t, string(String("hello world").Bytes()) == "hello world")
	assert(t, Bytes([]byte("hello world")).String() == "hello world")
	assert(t, Bytes([]byte("hello world")).Int() == 0)
	assert(t, Bytes([]byte("hello world")).IsNil() == false)
	assert(t, string(Bytes([]byte("hello world")).Bytes()) == "hello world")
	forceIfaceStrs = false
	assert(t, Any(Jello{1, 2}).IsNil() == false)
	assert(t, Any(Jello{1, 2}).String() == "{1 2}")
	assert(t, Any(Pudding{1, 2}).String() == "Yum{1 2}")
	assert(t, string(Any(Pudding{1, 2}).Bytes()) == "Yum{1 2}")
	forceIfacePtrs = true
	assert(t, Any(Jello{1, 2}).IsNil() == false)
	assert(t, Any(Jello{1, 2}).String() == "{1 2}")
	assert(t, Any(Jello{1, 2}).Any().(Jello).Feet == 2)
	assert(t, Any(Pudding{1, 2}).String() == "Yum{1 2}")
	assert(t, string(Any(Pudding{1, 2}).Bytes()) == "Yum{1 2}")
	forceIfacePtrs = false
	assert(t, Any(nil).IsNil())
	assert(t, Any("hello").String() == "hello")
	assert(t, Any([]byte("hello")).String() == "hello")
	assert(t, Any(true).Bool() == true)
	assert(t, Any(false).Bool() == false)
	assert(t, Any(int8(-1)).Int8() == -1)
	assert(t, Any(int16(-2)).Int16() == -2)
	assert(t, Any(int32(-3)).Int32() == -3)
	assert(t, Any(int64(-4)).Int64() == -4)
	assert(t, Any(uint8(1)).Int8() == 1)
	assert(t, Any(uint16(2)).Int16() == 2)
	assert(t, Any(uint32(3)).Int32() == 3)
	assert(t, Any(uint64(4)).Int64() == 4)
	assert(t, Any(int(1)).Int8() == 1)
	assert(t, Any(uint(2)).Int16() == 2)
	assert(t, Any(uintptr(3)).Int32() == 3)
	assert(t, Any(float32(4)).Float32() == 4)
	assert(t, Any(float64(5)).Float64() == 5)
	assert(t, Int(123).String() == "123")
	assert(t, string(Int(123).Bytes()) == "123")
	assert(t, Int(123).Any().(int64) == 123)
	assert(t, Any(Jello{1, 2}).Any().(Jello).Neat == 1)

	assert(t, CustomBits(99).String() == "99")
	assert(t, Bool(true).String() == "true")
	assert(t, Bool(false).String() == "false")
	assert(t, Uint64(99).String() == "99")
	assert(t, Int64(-99).String() == "-99")
	assert(t, Float64(-998).String() == "-998")
	assert(t, Nil().String() == "")

	assert(t, Any(CustomBits(99).Any()).String() == "99")
	assert(t, Any(Bool(true).Any()).String() == "true")
	assert(t, Any(Bool(false).Any()).String() == "false")
	assert(t, Any(Uint64(99).Any()).String() == "99")
	assert(t, Any(Int64(-99).Any()).String() == "-99")
	assert(t, Any(Float64(-998).Any()).String() == "-998")
	assert(t, Any(Nil().Any()).String() == "")

	assert(t, Int(99).Float64() == 99.0)
	assert(t, Nil().Float64() == 0)
	assert(t, CustomBits(1).Float64() == 1)
	assert(t, Bool(true).Float64() == 1)
	assert(t, Bool(false).Float64() == 0)
	assert(t, Uint64(98).Float64() == 98)
	assert(t, Int64(-98).Float64() == -98)
	assert(t, Float64(99).Float64() == 99.0)
	assert(t, Any("-99").Float64() == -99)
	assert(t, Any([]byte("-99")).Float64() == -99)
	assert(t, Any(interface{}(nil)).Float64() == 0)
	assert(t, Any(nil).Float64() == 0)
	assert(t, math.IsNaN(Any("hello").Float64()))
	assert(t, math.IsNaN(Any(Jello{10, 20}).Float64()))
	assert(t, Any(Pudding{10, 20}).Float64() == 200)

	assert(t, Int(99).Uint64() == 99.0)
	assert(t, Nil().Uint64() == 0)
	assert(t, CustomBits(1).Uint64() == 1)
	assert(t, Bool(true).Uint64() == 1)
	assert(t, Bool(false).Uint64() == 0)
	assert(t, Uint64(98).Uint64() == 98)
	assert(t, Int64(980).Uint64() == 980)
	assert(t, Float64(99).Uint64() == 99)
	assert(t, Any("990").Uint64() == 990)
	assert(t, Any([]byte("990")).Uint64() == 990)
	assert(t, Any(interface{}(nil)).Uint64() == 0)
	assert(t, Any(nil).Uint64() == 0)
	assert(t, Any("hello").Uint64() == 0)
	assert(t, Any(Jello{10, 20}).Uint64() == 0)
	assert(t, Any(Pudding{10, 20}).Uint64() == 200)

	assert(t, Int(99).Int64() == 99.0)
	assert(t, Nil().Int64() == 0)
	assert(t, CustomBits(1).Int64() == 1)
	assert(t, Bool(true).Int64() == 1)
	assert(t, Bool(false).Int64() == 0)
	assert(t, Uint64(98).Int64() == 98)
	assert(t, Int64(-98).Int64() == -98)
	assert(t, Float64(99).Int64() == 99.0)
	assert(t, Any("-99").Int64() == -99)
	assert(t, Any([]byte("-99")).Int64() == -99)
	assert(t, Any(interface{}(nil)).Int64() == 0)
	assert(t, Any(nil).Int64() == 0)
	assert(t, Any("hello").Int64() == 0)
	assert(t, Any(Jello{10, 20}).Int64() == 0)
	assert(t, Any(Pudding{10, 20}).Int64() == 200)

	assert(t, Int(99).Bool() == true)
	assert(t, Nil().Bool() == false)
	assert(t, CustomBits(1).Bool() == true)
	assert(t, Bool(true).Bool() == true)
	assert(t, Bool(false).Bool() == false)
	assert(t, Uint64(98).Bool() == true)
	assert(t, Int64(-98).Bool() == true)
	assert(t, Float64(99).Bool() == true)
	assert(t, Any("-99").Bool() == false)
	assert(t, Any("true").Bool() == true)
	assert(t, Any([]byte("-99")).Bool() == false)
	assert(t, Any([]byte("true")).Bool() == true)
	assert(t, Any(interface{}(nil)).Bool() == false)
	assert(t, Any(nil).Bool() == false)
	assert(t, Any("hello").Bool() == false)
	assert(t, Any(Jello{10, 20}).Bool() == false)
	assert(t, Any(Pudding{10, 20}).Bool() == true)

	assert(t, Any(nil).IsString() == false)
	assert(t, Any(123).IsString() == false)
	assert(t, Any("hello").IsString() == true)
	assert(t, Any([]byte("hello")).IsString() == false)
	assert(t, StringWithTag("hello", 10).IsString() == true)
	assert(t, StringWithTag("hello", 10).Tag() == 10)
	forceIfaceStrs = true
	assert(t, StringWithTag("hello", 10).IsString() == true)
	assert(t, StringWithTag("hello", 10).Tag() == 10)
	assert(t, Any("hello").IsString() == true)
	assert(t, Any([]byte("hello")).IsString() == false)
	forceIfaceStrs = false

	assert(t, Any(nil).IsBytes() == false)
	assert(t, Any(123).IsBytes() == false)
	assert(t, Any("hello").IsBytes() == false)
	assert(t, Any([]byte("hello")).IsBytes() == true)
	forceIfaceStrs = true
	assert(t, Any("hello").IsBytes() == false)
	assert(t, Any([]byte("hello")).IsBytes() == true)
	forceIfaceStrs = false

	assert(t, Int8(-10).Int8() == -10)
	assert(t, Int(500).Int8() == -12)
	assert(t, Int16(-10).Int16() == -10)
	assert(t, Int32(-10).Int32() == -10)
	assert(t, Int64(-10).Int64() == -10)
	assert(t, Int64(-10).Float32() == -10)
	assert(t, Float32(10.1239123).Float32() == 10.1239123)

	assert(t, Uint8(10).Uint8() == 10)
	assert(t, Uint(500).Uint8() == 500&0xFF)
	assert(t, Uint16(10).Uint16() == 10)
	assert(t, Uint32(11).Uint32() == 11)
	assert(t, Uint64(12).Uint64() == 12)
	assert(t, Uint64(12).Uint() == 12)

	assert(t, Uint64(10).IsUint() == true)
	assert(t, Uint8(10).IsUint() == true)
	assert(t, Int64(10).IsUint() == false)

	assert(t, Int64(10).IsInt() == true)
	assert(t, Int8(10).IsInt() == true)
	assert(t, Uint64(10).IsInt() == false)

	assert(t, Float64(10).IsFloat() == true)
	assert(t, Float32(10).IsFloat() == true)
	assert(t, Uint64(10).IsFloat() == false)

	assert(t, Bool(true).IsBool() == true)
	assert(t, Bool(false).IsBool() == true)
	assert(t, Uint64(10).IsBool() == false)

	assert(t, Byte(10).Byte() == 10)
	assert(t, Uint64(257).Byte() == 1)

	assert(t, String("hello").IsNumber() == false)
	assert(t, Int(10).IsNumber() == true)
	assert(t, Uint(10).IsNumber() == true)
	assert(t, Float64(10).IsNumber() == true)
	assert(t, Any(10).IsNumber() == true)

	assert(t, Uint64(10).Tag() == 0)
	assert(t, Bytes(nil).Tag() == 0)
	assert(t, Bytes([]byte{}).Tag() == 0)

	assert(t, String("hello").Tag() == 0)
	assert(t, StringWithTag("hello", 999).Tag() == 999)
	assert(t, StringWithTag("hello", 999).String() == "hello")
	forceIfaceStrs = true
	assert(t, String("hello").Tag() == 0)
	assert(t, StringWithTag("hello", 999).Tag() == 999)
	assert(t, StringWithTag("hello", 999).String() == "hello")
	forceIfaceStrs = false

}

func TestBytes(t *testing.T) {
	testBytes := func(t *testing.T, ncap int) {
		t.Helper()
		b := make([]byte, 0, ncap)
		b = append(b, "hello world"...)
		if ncap < cap(b) {
			ncap = cap(b)
		}
		assert(t, len(b) == 11 && cap(b) == ncap)
		v := Bytes(b)
		b2 := v.Bytes()
		assert(t, len(b2) == 11 && cap(b2) == ncap)
	}

	testBytes(t, 0)
	testBytes(t, 1)
	testBytes(t, 0xFF)
	testBytes(t, 0xFFF)
	testBytes(t, 0xFFFF)
	testBytes(t, 0xFFFFF)
	testBytes(t, 0x7FFFFF)
	testBytes(t, 0x7FFFFF+1)
	testBytes(t, 0xFFFFFF)
}

func TestUnits(t *testing.T) {
	assert(t, Float64(-98).toFloat64() == -98)
	assert(t, Uint64(98).toUint64() == 98)
	assert(t, Int64(-98).toInt64() == -98)
	assert(t, Bool(true).toBool() == true)
	assert(t, Bool(false).toBool() == false)
}

func TestPLocks(t *testing.T) {
	// Tests the psave() with plock/punlock using multiple goroutines.
	// Best if used with -race
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			start := time.Now()
			for time.Since(start) < time.Second/10 {
				Any(&Jello{1, 2})
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkIfaceInt(b *testing.B) {
	gen := func(b *testing.B, reset bool) []interface{} {
		arr := make([]interface{}, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = i
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ReportAllocs()
		b.ResetTimer()
		var res int
		for i := 0; i < b.N; i++ {
			res += arr[i].(int)
		}
	})
}

func BenchmarkBoxInt(b *testing.B) {
	gen := func(b *testing.B, reset bool) []Value {
		arr := make([]Value, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = Int(i)
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ReportAllocs()
		b.ResetTimer()
		var res int
		for i := 0; i < b.N; i++ {
			res += arr[i].Int()
		}
	})
}

func BenchmarkIfaceString(b *testing.B) {
	gen := func(b *testing.B, reset bool) []interface{} {
		strs := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			strs[i] = fmt.Sprint(i)
		}
		arr := make([]interface{}, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = strs[i]
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ResetTimer()
		b.ReportAllocs()
		var n int
		for i := 0; i < b.N; i++ {
			s := arr[i].(string)
			n += int(s[0]) + int(s[len(s)-1])
		}
	})
}

func BenchmarkBoxString(b *testing.B) {
	gen := func(b *testing.B, reset bool) []Value {
		strs := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			strs[i] = fmt.Sprint(i)
		}
		arr := make([]Value, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = String(strs[i])
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ResetTimer()
		b.ReportAllocs()
		var n int
		for i := 0; i < b.N; i++ {
			s := arr[i].String()
			n += int(s[0]) + int(s[len(s)-1])
		}
	})
}

func BenchmarkIfaceBytes(b *testing.B) {
	gen := func(b *testing.B, reset bool) []interface{} {
		strs := make([][]byte, b.N)
		for i := 0; i < b.N; i++ {
			strs[i] = []byte(fmt.Sprint(i))
		}
		arr := make([]interface{}, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = strs[i]
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ResetTimer()
		b.ReportAllocs()
		var n int
		for i := 0; i < b.N; i++ {
			s := arr[i].([]byte)
			n += int(s[0]) + int(s[len(s)-1])
		}
	})
}

func BenchmarkBoxBytes(b *testing.B) {
	gen := func(b *testing.B, reset bool) []Value {
		strs := make([][]byte, b.N)
		for i := 0; i < b.N; i++ {
			strs[i] = []byte(fmt.Sprint(i))
		}
		arr := make([]Value, b.N)
		if reset {
			b.ResetTimer()
			b.ReportAllocs()
		}
		for i := 0; i < b.N; i++ {
			arr[i] = Bytes(strs[i])
		}
		return arr
	}
	b.Run("to", func(b *testing.B) {
		gen(b, true)
	})
	b.Run("from", func(b *testing.B) {
		arr := gen(b, false)
		b.ResetTimer()
		b.ReportAllocs()
		var n int
		for i := 0; i < b.N; i++ {
			s := arr[i].Bytes()
			n += int(s[0]) + int(s[len(s)-1])
		}
	})
}
