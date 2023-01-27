package conv

import (
	"math"
	"strconv"
)

type (
	booler   interface{ Bool() bool }
	floater  interface{ Float64() float64 }
	inter    interface{ Int64() int64 }
	uinter   interface{ Uint64() uint64 }
	stringer interface{ String() string }
)

// /////////////////////////////////////////
// Bool
// /////////////////////////////////////////

// Ttoi converts bool to int64
func Ttoi(t bool) int64 {
	if t {
		return 1
	}
	return 0
}

// Ttou converts bool to uint64
func Ttou(t bool) uint64 {
	if t {
		return 1
	}
	return 0
}

// Ttof converts bool to float64
func Ttof(t bool) float64 {
	if t {
		return 1
	}
	return 0
}

// Ttoa converts bool to string
func Ttoa(t bool) string {
	if t {
		return "true"
	}
	return "false"
}

// Ttov converts bool to any
func Ttov(t bool) interface{} {
	return t
}

// /////////////////////////////////////////
// Float64
// /////////////////////////////////////////

const maxUint64 = uint64(18446744073709551615)
const minInt64 = int64(-9223372036854775808)
const maxInt64 = int64(9223372036854775807)

var maxUint64Float = math.Nextafter(math.MaxUint64, 0)
var maxInt64Float = math.Nextafter(math.MaxInt64, 0)
var minInt64Float = math.Nextafter(math.MinInt64, 1)

// Ftot converts float64 to bool
func Ftot(f float64) bool {
	return f < 0 || f > 0
}

// Ftoi converts float64 to int64
func Ftoi(f float64) int64 {
	if math.IsNaN(f) {
		return 0
	}
	if f < -9007199254740991 || f > 9007199254740991 {
		// The number is outside the range for correct binary
		// representation of floating point as an integer value.
		// https://tc39.es/ecma262/#sec-number.min_safe_integer
		// https://tc39.es/ecma262/#sec-number.max_safe_integer
		if f < 0 {
			f = math.Ceil(f)
			if f < minInt64Float {
				return minInt64
			}
		} else {
			f = math.Floor(f)
			if f > maxInt64Float {
				return maxInt64
			}
		}
	}
	return int64(f)
}

// Ftou converts float64 to uint64
func Ftou(f float64) uint64 {
	if math.IsNaN(f) {
		return 0
	}
	if f < 0 {
		return 0
	}
	if f > 9007199254740991 {
		// Outside range: See Ftoi for description.
		f = math.Floor(f)
		if f > maxUint64Float {
			return maxUint64
		}
	}
	return uint64(f)
}

// Ftoa converts float64 to string
// Returns 'Infinity' or '-Infinity', not 'Inf' or '-Inf', for infinite numbers.
// Returns 'NaN' for not-a-number.
func Ftoa(f float64) string {
	switch {
	case math.IsNaN(f):
		return "NaN"
	case math.IsInf(f, +1):
		return "Infinity"
	case math.IsInf(f, -1):
		return "-Infinity"
	default:
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
}

// Ftov converts float64 to any
func Ftov(f float64) interface{} {
	return f
}

// /////////////////////////////////////////
// Int64
// /////////////////////////////////////////

// Itot converts int64 to bool
func Itot(i int64) bool {
	return i != 0
}

// Itof converts int64 to float64
func Itof(i int64) float64 {
	return float64(i)
}

// Itou converts int64 to uint64
func Itou(i int64) uint64 {
	if i < 0 {
		return 0
	}
	return uint64(i)
}

// Itoa converts int64 to string
func Itoa(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Itov converts int64 to any
func Itov(i int64) interface{} {
	return i
}

// /////////////////////////////////////////
// Uint64
// /////////////////////////////////////////

// Utot converts uint64 to bool
func Utot(u uint64) bool {
	return u != 0
}

// Utof converts uint64 to float64
func Utof(u uint64) float64 {
	return float64(u)
}

// Utoi converts uint64 to int64
func Utoi(u uint64) int64 {
	if u > math.MaxInt64 {
		return math.MaxInt64
	}
	return int64(u)
}

// Utoa converts uint64 to string
func Utoa(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// Utov converts uint64 to any
func Utov(u uint64) interface{} {
	return u
}

// /////////////////////////////////////////
// String
// /////////////////////////////////////////

func isnumch(c byte) bool {
	return (c >= '0' && c <= '9') || c == '.'
}
func parseFloat(a string) (float64, error) {
	if a == "" {
		return 0, strconv.ErrSyntax
	}
	if len(a) == 1 || isnumch(a[0]) || (a[0] == '-' && isnumch(a[1])) ||
		(a[0] == '+' && isnumch(a[1])) {
		return strconv.ParseFloat(a, 64)
	}
	switch a {
	case "+Infinity", "Infinity":
		return math.Inf(+1), nil
	case "-Infinity":
		return math.Inf(-1), nil
	case "NaN":
		return math.NaN(), nil
	default:
		return 0, strconv.ErrSyntax
	}
}

// Atot converts string to bool
// Always returns true unless string is empty.
func Atot(a string) bool {
	return a != ""
}

// Atof converts string to float64
// For infinte numbers use 'Infinity' or '-Infinity', not 'Inf' or '-Inf'.
// Returns NaN for invalid syntax
func Atof(a string) float64 {
	f, err := parseFloat(a)
	if err != nil {
		return math.NaN()
	}
	return f
}

// Atoi converts string to int64
// Returns 0 for invalid syntax
func Atoi(a string) int64 {
	x, err := strconv.ParseInt(a, 10, 64)
	if err == nil {
		return x
	}
	f, err := parseFloat(a)
	if err == nil {
		return Ftoi(f)
	}
	return 0
}

// Atou converts string to uint64
// Returns 0 for invalid syntax
func Atou(a string) uint64 {
	x, err := strconv.ParseUint(a, 10, 64)
	if err == nil {
		return x
	}
	f, err := parseFloat(a)
	if err == nil {
		return Ftou(f)
	}
	return 0
}

// Atov converts string to any
func Atov(a string) interface{} {
	return a
}

// /////////////////////////////////////////
// Any
// /////////////////////////////////////////

// Vtot converts any to bool
func Vtot(v interface{}) bool {
	switch v := v.(type) {
	case bool:
		return v
	case int:
		return Itot(int64(v))
	case int8:
		return Itot(int64(v))
	case int16:
		return Itot(int64(v))
	case int32:
		return Itot(int64(v))
	case int64:
		return Itot(v)
	case uint:
		return Utot(uint64(v))
	case uint8:
		return Utot(uint64(v))
	case uint16:
		return Utot(uint64(v))
	case uint32:
		return Utot(uint64(v))
	case uint64:
		return Utot(v)
	case float64:
		return Ftot(v)
	case float32:
		return Ftot(float64(v))
	case string:
		return Atot(v)
	default:
		// order matters (bool,int,uint,float,string)
		switch v := v.(type) {
		case booler:
			return v.Bool()
		case inter:
			return Itot(v.Int64())
		case uinter:
			return Utot(v.Uint64())
		case stringer:
			return Atot(v.String())
		case floater:
			return Ftot(v.Float64())
		default:
			return false
		}
	}
}

// Vtof converts any to float64
func Vtof(v interface{}) float64 {
	switch v := v.(type) {
	case bool:
		return Ttof(v)
	case int:
		return Itof(int64(v))
	case int8:
		return Itof(int64(v))
	case int16:
		return Itof(int64(v))
	case int32:
		return Itof(int64(v))
	case int64:
		return Itof(v)
	case uint:
		return Utof(uint64(v))
	case uint8:
		return Utof(uint64(v))
	case uint16:
		return Utof(uint64(v))
	case uint32:
		return Utof(uint64(v))
	case uint64:
		return Utof(v)
	case float64:
		return v
	case float32:
		return float64(v)
	case string:
		return Atof(v)
	default:
		// order matters (float,int,uint,string,bool)
		switch v := v.(type) {
		case floater:
			return v.Float64()
		case inter:
			return Itof(v.Int64())
		case uinter:
			return Utof(v.Uint64())
		case stringer:
			return Atof(v.String())
		case booler:
			return Ttof(v.Bool())
		default:
			return math.NaN()
		}
	}
}

// Vtoi converts any to int64
func Vtoi(v interface{}) int64 {
	switch v := v.(type) {
	case bool:
		return Ttoi(v)
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return Utoi(uint64(v))
	case uint8:
		return Utoi(uint64(v))
	case uint16:
		return Utoi(uint64(v))
	case uint32:
		return Utoi(uint64(v))
	case uint64:
		return Utoi(v)
	case float64:
		return Ftoi(v)
	case float32:
		return Ftoi(float64(v))
	case string:
		return Atoi(v)
	default:
		// order matters (int,uint,float,string,bool)
		switch v := v.(type) {
		case inter:
			return v.Int64()
		case uinter:
			return Utoi(v.Uint64())
		case floater:
			return Ftoi(v.Float64())
		case stringer:
			return Atoi(v.String())
		case booler:
			return Ttoi(v.Bool())
		default:
			return 0
		}
	}
}

// Vtou converts any to uint64
func Vtou(v interface{}) uint64 {
	switch v := v.(type) {
	case bool:
		return Ttou(v)
	case int:
		return Itou(int64(v))
	case int8:
		return Itou(int64(v))
	case int16:
		return Itou(int64(v))
	case int32:
		return Itou(int64(v))
	case int64:
		return Itou(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float64:
		return Ftou(v)
	case float32:
		return Ftou(float64(v))
	case string:
		return Atou(v)
	default:
		// order matters (uint,int,float,string,bool)
		switch v := v.(type) {
		case uinter:
			return v.Uint64()
		case inter:
			return Itou(v.Int64())
		case floater:
			return Ftou(v.Float64())
		case stringer:
			return Atou(v.String())
		case booler:
			return Ttou(v.Bool())
		default:
			return 0
		}
	}
}

// Vtoa converts any to string
func Vtoa(v interface{}) string {
	switch v := v.(type) {
	case bool:
		return Ttoa(v)
	case int:
		return Itoa(int64(v))
	case int8:
		return Itoa(int64(v))
	case int16:
		return Itoa(int64(v))
	case int32:
		return Itoa(int64(v))
	case int64:
		return Itoa(v)
	case uint:
		return Utoa(uint64(v))
	case uint8:
		return Utoa(uint64(v))
	case uint16:
		return Utoa(uint64(v))
	case uint32:
		return Utoa(uint64(v))
	case uint64:
		return Utoa(v)
	case float64:
		return Ftoa(v)
	case float32:
		return Ftoa(float64(v))
	case string:
		return v
	default:
		// order matters (string,int,uint,float,bool)
		switch v := v.(type) {
		case stringer:
			return v.String()
		case inter:
			return Itoa(v.Int64())
		case uinter:
			return Utoa(v.Uint64())
		case floater:
			return Ftoa(v.Float64())
		case booler:
			return Ttoa(v.Bool())
		default:
			return ""
		}
	}
}
