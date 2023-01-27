// Copyright 2023 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package conv

import (
	"math"
	"testing"
)

func assert(t *testing.T, cond bool) {
	t.Helper()
	if !cond {
		t.Fatal("assertion failed")
	}
}

type tbooler bool
type tfloater float64
type tstringer string
type tinter int64
type tuinter uint64

func (t tbooler) Bool() bool        { return bool(t) }
func (t tfloater) Float64() float64 { return float64(t) }
func (t tinter) Int64() int64       { return int64(t) }
func (t tuinter) Uint64() uint64    { return uint64(t) }
func (t tstringer) String() string  { return string(t) }

func TestConv(t *testing.T) {
	// Bool
	assert(t, Ttof(true) == 1)
	assert(t, Ttof(false) == 0)
	assert(t, Ttoi(true) == 1)
	assert(t, Ttoi(false) == 0)
	assert(t, Ttou(true) == 1)
	assert(t, Ttou(false) == 0)
	assert(t, Ttoa(true) == "true")
	assert(t, Ttoa(false) == "false")
	assert(t, Ttov(true) == true)
	assert(t, Ttov(false) == false)

	// Float
	assert(t, Ftot(0) == false)
	assert(t, Ftot(1) == true)
	assert(t, Ftot(-1) == true)
	assert(t, Ftot(math.Inf(-1)) == true)
	assert(t, Ftot(math.Inf(0)) == true)
	assert(t, Ftot(math.Inf(+1)) == true)
	assert(t, Ftot(math.NaN()) == false)
	assert(t, Ftoi(0) == 0)
	assert(t, Ftoi(math.MaxUint64) == math.MaxInt64)
	assert(t, Ftoi(math.MaxUint64) == math.MaxInt64)
	assert(t, Ftoi(math.MinInt64) == math.MinInt64)
	assert(t, Ftoi(math.NaN()) == 0)
	assert(t, Ftoi(math.Inf(-1)) == math.MinInt64)
	assert(t, Ftoi(math.Inf(0)) == math.MaxInt64)
	assert(t, Ftoi(math.Inf(+1)) == math.MaxInt64)
	assert(t, Ftou(0) == 0)
	assert(t, Ftou(math.MaxUint64) == math.MaxUint64)
	assert(t, Ftou(math.MinInt64) == 0)
	assert(t, Ftou(math.NaN()) == 0)
	assert(t, Ftou(math.Inf(-1)) == 0)
	assert(t, Ftou(math.Inf(0)) == math.MaxUint64)
	assert(t, Ftou(math.Inf(+1)) == math.MaxUint64)
	assert(t, Ftoa(0) == "0")
	assert(t, Ftoa(math.NaN()) == "NaN")
	assert(t, Ftoa(math.Inf(+1)) == "Infinity")
	assert(t, Ftoa(math.Inf(-1)) == "-Infinity")
	assert(t, Ftov(0) == 0.0)
	assert(t, math.IsNaN(Ftov(math.NaN()).(float64)))

	// Int
	assert(t, Itot(0) == false)
	assert(t, Itot(1) == true)
	assert(t, Itot(-1) == true)
	assert(t, Itof(0) == 0)
	assert(t, Itof(1) == 1)
	assert(t, Itof(-1) == -1)
	assert(t, Itou(0) == 0)
	assert(t, Itou(1) == 1)
	assert(t, Itou(-1) == 0)
	assert(t, Itoa(0) == "0")
	assert(t, Itoa(1) == "1")
	assert(t, Itoa(-1) == "-1")
	assert(t, Itov(-1) == int64(-1))

	// Uint
	assert(t, Utot(0) == false)
	assert(t, Utot(1) == true)
	assert(t, Utof(0) == 0)
	assert(t, Utof(1) == 1)
	assert(t, Utoi(0) == 0)
	assert(t, Utoi(1) == 1)
	assert(t, Utoi(math.MaxUint64) == math.MaxInt64)
	assert(t, Utoa(0) == "0")
	assert(t, Utoa(1) == "1")
	assert(t, Utov(1) == uint64(1))

	// String
	assert(t, Atot("") == false)
	assert(t, Atot("0") == true)
	assert(t, Atot("1") == true)
	assert(t, Atot("false") == true)
	assert(t, Atot("true") == true)
	assert(t, Atof("0") == 0)
	assert(t, math.IsNaN(Atof("")))
	assert(t, math.IsNaN(Atof("p1o1")))
	assert(t, math.IsNaN(Atof("NaN")))
	assert(t, !math.IsNaN(Atof("12312")))
	assert(t, math.IsNaN(Atof("+Inf")))
	assert(t, math.IsNaN(Atof("+inf")))
	assert(t, math.IsNaN(Atof("+infinity")))
	assert(t, math.IsNaN(Atof("Inf")))
	assert(t, math.IsNaN(Atof("inf")))
	assert(t, math.IsNaN(Atof("infinity")))
	assert(t, math.IsNaN(Atof("-Inf")))
	assert(t, math.IsNaN(Atof("-inf")))
	assert(t, math.IsNaN(Atof("-infinity")))
	assert(t, math.IsInf(Atof("+Infinity"), +1))
	assert(t, math.IsInf(Atof("Infinity"), +1))
	assert(t, math.IsInf(Atof("-Infinity"), -1))
	assert(t, Atoi("0") == 0)
	assert(t, Atoi("1") == 1)
	assert(t, Atoi("-1") == -1)
	assert(t, Atoi("op1") == 0)
	assert(t, Atoi("123129319238019283121231231") == math.MaxInt64)
	assert(t, Atou("0") == 0)
	assert(t, Atou("1") == 1)
	assert(t, Atou("op1") == 0)
	assert(t, Atou("123129319238019283121231231") == math.MaxUint64)
	assert(t, Atov("123") == "123")
	assert(t, Atov("") == "")

	// ///////////////////////////////////////////////////////
	// Any(bool)
	// any(bool) -> bool
	assert(t, Vtot(true) == true)
	assert(t, Vtot(false) == false)
	// any(float*) -> bool
	assert(t, Vtot(float64(0)) == false)
	assert(t, Vtot(float64(1)) == true)
	assert(t, Vtot(float32(0)) == false)
	assert(t, Vtot(float32(1)) == true)
	// any(int*) -> bool
	assert(t, Vtot(int(0)) == false)
	assert(t, Vtot(int(1)) == true)
	assert(t, Vtot(int8(0)) == false)
	assert(t, Vtot(int8(1)) == true)
	assert(t, Vtot(int16(0)) == false)
	assert(t, Vtot(int16(1)) == true)
	assert(t, Vtot(int32(0)) == false)
	assert(t, Vtot(int32(1)) == true)
	assert(t, Vtot(int64(0)) == false)
	assert(t, Vtot(int64(1)) == true)
	// any(uint*) -> bool
	assert(t, Vtot(uint(0)) == false)
	assert(t, Vtot(uint(1)) == true)
	assert(t, Vtot(uint8(0)) == false)
	assert(t, Vtot(uint8(1)) == true)
	assert(t, Vtot(uint16(0)) == false)
	assert(t, Vtot(uint16(1)) == true)
	assert(t, Vtot(uint32(0)) == false)
	assert(t, Vtot(uint32(1)) == true)
	assert(t, Vtot(uint64(0)) == false)
	assert(t, Vtot(uint64(1)) == true)
	// any(string) -> bool
	assert(t, Vtot("true") == true)
	assert(t, Vtot("false") == true)
	assert(t, Vtot("") == false)
	// any(booler) -> bool
	assert(t, Vtot(tbooler(true)) == true)
	assert(t, Vtot(tbooler(false)) == false)
	// any(floater) -> bool
	assert(t, Vtot(tfloater(math.NaN())) == false)
	assert(t, Vtot(tfloater(0)) == false)
	assert(t, Vtot(tfloater(1)) == true)
	// any(inter) -> bool
	assert(t, Vtot(tinter(0)) == false)
	assert(t, Vtot(tinter(1)) == true)
	assert(t, Vtot(tinter(-1)) == true)
	// any(uinter) -> bool
	assert(t, Vtot(tuinter(0)) == false)
	assert(t, Vtot(tuinter(1)) == true)
	// any(stringer) -> bool
	assert(t, Vtot(tstringer("true")) == true)
	assert(t, Vtot(tstringer("")) == false)
	assert(t, Vtot(tstringer("false")) == true)
	// any(fallback) -> bool
	assert(t, Vtot(nil) == false)

	// ///////////////////////////////////////////////////////
	// Any(Float)
	// ///////////////////////////////////////////////////////
	// any(bool) -> float64
	assert(t, Vtof(true) == 1)
	assert(t, Vtof(false) == 0)
	// any(int*) -> float64
	assert(t, Vtof(int(1)) == 1)
	assert(t, Vtof(int(0)) == 0)
	assert(t, Vtof(int8(1)) == 1)
	assert(t, Vtof(int8(0)) == 0)
	assert(t, Vtof(int16(1)) == 1)
	assert(t, Vtof(int16(0)) == 0)
	assert(t, Vtof(int32(1)) == 1)
	assert(t, Vtof(int32(0)) == 0)
	assert(t, Vtof(int64(1)) == 1)
	assert(t, Vtof(int64(0)) == 0)
	// any(uint*) -> float64
	assert(t, Vtof(uint(1)) == 1)
	assert(t, Vtof(uint(0)) == 0)
	assert(t, Vtof(uint8(1)) == 1)
	assert(t, Vtof(uint8(0)) == 0)
	assert(t, Vtof(uint16(1)) == 1)
	assert(t, Vtof(uint16(0)) == 0)
	assert(t, Vtof(uint32(1)) == 1)
	assert(t, Vtof(uint32(0)) == 0)
	assert(t, Vtof(uint64(1)) == 1)
	assert(t, Vtof(uint64(0)) == 0)
	// any(float*) -> float64
	assert(t, Vtof(float32(1)) == 1)
	assert(t, Vtof(float32(0)) == 0)
	assert(t, Vtof(float64(1)) == 1)
	assert(t, Vtof(float64(0)) == 0)
	// any(string) -> float64
	assert(t, Vtof(string("1")) == 1)
	assert(t, Vtof(string("0")) == 0)
	assert(t, math.IsNaN(Vtof(string("wqer9812039"))))
	// any(booler) -> float64
	assert(t, Vtof(tbooler(true)) == 1)
	assert(t, Vtof(tbooler(false)) == 0)
	// any(floater) -> float64
	assert(t, Vtof(tfloater(1)) == 1)
	assert(t, Vtof(tfloater(0)) == 0)
	// any(inter) -> float64
	assert(t, Vtof(tinter(1)) == 1)
	assert(t, Vtof(tinter(0)) == 0)
	assert(t, Vtof(tinter(-1)) == -1)
	// any(uinter) -> float64
	assert(t, Vtof(tuinter(1)) == 1)
	assert(t, Vtof(tuinter(0)) == 0)
	// any(stringer) -> float64
	assert(t, Vtof(tstringer("1")) == 1)
	assert(t, Vtof(tstringer("0")) == 0)
	assert(t, math.IsNaN(Vtof(tstringer("wqer9812039"))))
	// any(fallback) -> float64
	assert(t, math.IsNaN(Vtof(nil)))

	// ///////////////////////////////////////////////////////
	// Any(Int)
	// ///////////////////////////////////////////////////////
	// any(bool) -> int64
	assert(t, Vtoi(true) == 1)
	assert(t, Vtoi(false) == 0)
	// any(int*) -> int64
	assert(t, Vtoi(int(1)) == 1)
	assert(t, Vtoi(int(0)) == 0)
	assert(t, Vtoi(int8(1)) == 1)
	assert(t, Vtoi(int8(0)) == 0)
	assert(t, Vtoi(int16(1)) == 1)
	assert(t, Vtoi(int16(0)) == 0)
	assert(t, Vtoi(int32(1)) == 1)
	assert(t, Vtoi(int32(0)) == 0)
	assert(t, Vtoi(int64(1)) == 1)
	assert(t, Vtoi(int64(0)) == 0)
	// any(uint*) -> int64
	assert(t, Vtoi(uint(1)) == 1)
	assert(t, Vtoi(uint(0)) == 0)
	assert(t, Vtoi(uint8(1)) == 1)
	assert(t, Vtoi(uint8(0)) == 0)
	assert(t, Vtoi(uint16(1)) == 1)
	assert(t, Vtoi(uint16(0)) == 0)
	assert(t, Vtoi(uint32(1)) == 1)
	assert(t, Vtoi(uint32(0)) == 0)
	assert(t, Vtoi(uint64(1)) == 1)
	assert(t, Vtoi(uint64(0)) == 0)
	// any(float*) -> int64
	assert(t, Vtoi(float32(1)) == 1)
	assert(t, Vtoi(float32(0)) == 0)
	assert(t, Vtoi(float64(1)) == 1)
	assert(t, Vtoi(float64(0)) == 0)
	// any(string) -> int64
	assert(t, Vtoi(string("1")) == 1)
	assert(t, Vtoi(string("0")) == 0)
	assert(t, Vtoi(string("wqer9812039")) == 0)
	// any(booler) -> int64
	assert(t, Vtoi(tbooler(true)) == 1)
	assert(t, Vtoi(tbooler(false)) == 0)
	// any(floater) -> int64
	assert(t, Vtoi(tfloater(1)) == 1)
	assert(t, Vtoi(tfloater(0)) == 0)
	// any(inter) -> int64
	assert(t, Vtoi(tinter(1)) == 1)
	assert(t, Vtoi(tinter(0)) == 0)
	assert(t, Vtoi(tinter(-1)) == -1)
	// any(uinter) -> int64
	assert(t, Vtoi(tuinter(1)) == 1)
	assert(t, Vtoi(tuinter(0)) == 0)
	// any(stringer) -> int64
	assert(t, Vtoi(tstringer("1")) == 1)
	assert(t, Vtoi(tstringer("0")) == 0)
	assert(t, Vtoi(tstringer("wqer9812039")) == 0)
	// any(fallback) -> int64
	assert(t, Vtoi(nil) == 0)

	// ///////////////////////////////////////////////////////
	// Any(Uint)
	// ///////////////////////////////////////////////////////
	// any(bool) -> uint64
	assert(t, Vtou(true) == 1)
	assert(t, Vtou(false) == 0)
	// any(int*) -> uint64
	assert(t, Vtou(int(1)) == 1)
	assert(t, Vtou(int(0)) == 0)
	assert(t, Vtou(int8(1)) == 1)
	assert(t, Vtou(int8(0)) == 0)
	assert(t, Vtou(int16(1)) == 1)
	assert(t, Vtou(int16(0)) == 0)
	assert(t, Vtou(int32(1)) == 1)
	assert(t, Vtou(int32(0)) == 0)
	assert(t, Vtou(int64(1)) == 1)
	assert(t, Vtou(int64(0)) == 0)
	// any(uint*) -> uint64
	assert(t, Vtou(uint(1)) == 1)
	assert(t, Vtou(uint(0)) == 0)
	assert(t, Vtou(uint8(1)) == 1)
	assert(t, Vtou(uint8(0)) == 0)
	assert(t, Vtou(uint16(1)) == 1)
	assert(t, Vtou(uint16(0)) == 0)
	assert(t, Vtou(uint32(1)) == 1)
	assert(t, Vtou(uint32(0)) == 0)
	assert(t, Vtou(uint64(1)) == 1)
	assert(t, Vtou(uint64(0)) == 0)
	// any(float*) -> uint64
	assert(t, Vtou(float32(1)) == 1)
	assert(t, Vtou(float32(0)) == 0)
	assert(t, Vtou(float64(1)) == 1)
	assert(t, Vtou(float64(0)) == 0)
	// any(string) -> uint64
	assert(t, Vtou(string("1")) == 1)
	assert(t, Vtou(string("0")) == 0)
	assert(t, Vtou(string("wqer9812039")) == 0)
	// any(booler) -> uint64
	assert(t, Vtou(tbooler(true)) == 1)
	assert(t, Vtou(tbooler(false)) == 0)
	// any(floater) -> uint64
	assert(t, Vtou(tfloater(1)) == 1)
	assert(t, Vtou(tfloater(0)) == 0)
	// any(inter) -> uint64
	assert(t, Vtou(tinter(1)) == 1)
	assert(t, Vtou(tinter(0)) == 0)
	// any(uinter) -> uint64
	assert(t, Vtou(tuinter(1)) == 1)
	assert(t, Vtou(tuinter(0)) == 0)
	// any(stringer) -> uint64
	assert(t, Vtou(tstringer("1")) == 1)
	assert(t, Vtou(tstringer("0")) == 0)
	assert(t, Vtou(tstringer("wqer9812039")) == 0)
	// any(fallback) -> uint64
	assert(t, Vtou(nil) == 0)

	// ///////////////////////////////////////////////////////
	// Any(String)
	// ///////////////////////////////////////////////////////
	// any(bool) -> string
	assert(t, Vtoa(true) == "true")
	assert(t, Vtoa(false) == "false")
	// any(int*) -> string
	assert(t, Vtoa(int(1)) == "1")
	assert(t, Vtoa(int(0)) == "0")
	assert(t, Vtoa(int8(1)) == "1")
	assert(t, Vtoa(int8(0)) == "0")
	assert(t, Vtoa(int16(1)) == "1")
	assert(t, Vtoa(int16(0)) == "0")
	assert(t, Vtoa(int32(1)) == "1")
	assert(t, Vtoa(int32(0)) == "0")
	assert(t, Vtoa(int64(1)) == "1")
	assert(t, Vtoa(int64(0)) == "0")
	// any(uint*) -> string
	assert(t, Vtoa(uint(1)) == "1")
	assert(t, Vtoa(uint(0)) == "0")
	assert(t, Vtoa(uint8(1)) == "1")
	assert(t, Vtoa(uint8(0)) == "0")
	assert(t, Vtoa(uint16(1)) == "1")
	assert(t, Vtoa(uint16(0)) == "0")
	assert(t, Vtoa(uint32(1)) == "1")
	assert(t, Vtoa(uint32(0)) == "0")
	assert(t, Vtoa(uint64(1)) == "1")
	assert(t, Vtoa(uint64(0)) == "0")
	// any(float*) -> string
	assert(t, Vtoa(float32(1)) == "1")
	assert(t, Vtoa(float32(0)) == "0")
	assert(t, Vtoa(float64(1)) == "1")
	assert(t, Vtoa(float64(0)) == "0")
	assert(t, Vtoa(float64(math.NaN())) == "NaN")
	assert(t, Vtoa(float64(math.Inf(+1))) == "Infinity")
	assert(t, Vtoa(float64(math.Inf(-1))) == "-Infinity")
	// any(string) -> string
	assert(t, Vtoa(string("1")) == "1")
	// any(booler) -> string
	assert(t, Vtoa(tbooler(true)) == "true")
	assert(t, Vtoa(tbooler(false)) == "false")
	// any(floater) -> string
	assert(t, Vtoa(tfloater(1)) == "1")
	assert(t, Vtoa(tfloater(0)) == "0")
	assert(t, Vtoa(tfloater(math.NaN())) == "NaN")
	assert(t, Vtoa(tfloater(math.Inf(+1))) == "Infinity")
	assert(t, Vtoa(tfloater(math.Inf(-1))) == "-Infinity")
	// any(inter) -> string
	assert(t, Vtoa(tinter(1)) == "1")
	assert(t, Vtoa(tinter(0)) == "0")
	// any(uinter) -> string
	assert(t, Vtoa(tuinter(1)) == "1")
	assert(t, Vtoa(tuinter(0)) == "0")
	// any(stringer) -> string
	assert(t, Vtoa(tstringer("1")) == "1")
	assert(t, Vtoa(tstringer("0")) == "0")
	assert(t, Vtoa(tstringer("wqer9812039")) == "wqer9812039")
	// any(fallback) -> string
	assert(t, Vtoa(nil) == "")

}

func TestFloatConversions(t *testing.T) {
	assert(t, Ftoi(math.NaN()) == 0)
	assert(t, Ftoi(math.Inf(+1)) == math.MaxInt64)
	assert(t, Ftoi(math.Inf(-1)) == math.MinInt64)
	assert(t, Ftou(math.NaN()) == 0)
	assert(t, Ftou(math.Inf(+1)) == math.MaxUint64)
	assert(t, Ftou(math.Inf(-1)) == 0)
}
