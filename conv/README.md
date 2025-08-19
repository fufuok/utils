# conv

*forked from tidwall/conv*

[![GoDoc](https://godoc.org/github.com/tidwall/conv?status.svg)](https://godoc.org/github.com/tidwall/conv)

A Go package for converting various primitive types.

Designed to follow [ECMA-262](https://en.wikipedia.org/wiki/ECMAScript),
therefore it differs slightly from how the built in `strconv` package works. 

Please see the [docs](https://godoc.org/github.com/tidwall/conv) for more details.

## Functions

```go
Ttoi(t bool) int64      // bool -> int64
Ttou(t bool) uint64     // bool -> uint64
Ttof(t bool) float64    // bool -> float64
Ttoa(t bool) string     // bool -> string
Ttov(t bool) any        // bool -> any

Ftot(f float64) bool    // float64 -> bool
Ftoi(f float64) int64   // float64 -> int64
Ftou(f float64) uint64  // float64 -> uint64
Ftoa(f float64) string  // float64 -> string
Ftov(f float64) any     // float64 -> any

Itot(i int64) bool      // int64 -> bool
Itof(i int64) float64   // int64 -> float64
Itou(i int64) uint64    // int64 -> uint64
Itoa(i int64) string    // int64 -> string
Itov(i int64) any       // int64 -> any

Utot(u uint64) bool     // uint64 -> bool
Utof(u uint64) float64  // uint64 -> float64
Utoi(u uint64) int64    // uint64 -> int64
Utoa(u uint64) string   // uint64 -> string
Utov(u uint64) any      // uint64 -> any

Atot(a string) bool     // string -> bool
Atof(a string) float64  // string -> float64
Atoi(a string) int64    // string -> int64
Atou(a string) uint64   // string -> uint64
Atov(a string) any      // string -> any

Vtot(v any) bool        // any -> bool
Vtof(v any) float64     // any -> float64
Vtoi(v any) int64       // any -> int64
Vtou(v any) uint64      // any -> uint64
Vtoa(v any) string      // any -> string
```
