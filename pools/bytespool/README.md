# 💫 BytesPool

Reuse used byte slices to achieve zero allocation.

The existing byte slices are stored in groups according to the capacity length range, and suitable byte slice objects are automatically allocated according to the capacity length when used.

For more functions, please use: [https://github.com/fufuok/bytespool](https://github.com/fufuok/bytespool)

## ✨ Features

- Get byte slices always succeed without panic.
- Optional length of 0 or fixed-length byte slices.
- Automatic garbage collection of big-byte slices.
- High performance, See: [Benchmarks](#-benchmarks).

## ⚙️ Installation

```go
go get -u github.com/fufuok/utils
```

## 📚 Examples

### ⚡️ Quickstart

```go
package main

import (
	"fmt"

	"github.com/fufuok/utils/pools/bytespool"
)

func main() {
	// Get() is the same as New()
	bs := bytespool.Get(1024)
	// len: 1024, cap: 1024
	fmt.Printf("len: %d, cap: %d\n", len(bs), cap(bs))

	// Put() is the same as Release(), Put it back into the pool after use
	bytespool.Put(bs)

	// len: 0, capacity: 8 (Specified capacity)
	bs = bytespool.Make(8)
	bs = append(bs, "abc"...)
	// len: 3, cap: 8
	fmt.Printf("len: %d, cap: %d\n", len(bs), cap(bs))
	ok := bytespool.Release(bs)
	// true
	fmt.Println(ok)

	// len: 8, capacity: 8 (Fixed length)
	bs = bytespool.New(8)
	copy(bs, "12345678")
	// len: 8, cap: 8, value: 12345678
	fmt.Printf("len: %d, cap: %d, value: %s\n", len(bs), cap(bs), bs)
	bytespool.Release(bs)

	// Output:
	// len: 1024, cap: 1024
	// len: 3, cap: 8
	// true
	// len: 8, cap: 8, value: 12345678
}
```

### ⏳ Automated reuse

```go
package main

import (
	"fmt"

	"github.com/fufuok/utils/pools/bytespool"
)

func main() {
	// len: 0, cap: 4 (Specified capacity, automatically adapt to the capacity scale)
	bs3 := bytespool.Make(3)

	bs3 = append(bs3, "123"...)
	fmt.Printf("len: %d, cap: %d, %s\n", len(bs3), cap(bs3), bs3)

	bytespool.Release(bs3)

	// len: 4, cap: 4 (Fixed length)
	bs4 := bytespool.New(4)

	// Reuse of bs3
	fmt.Printf("same array: %v\n", &bs3[0] == &bs4[0])
	// Contain old data
	fmt.Printf("bs3: %s, bs4: %s\n", bs3, bs4[:3])

	copy(bs4, "xy")
	fmt.Printf("len: %d, cap: %d, %s\n", len(bs4), cap(bs4), bs4[:3])

	bytespool.Release(bs4)

	// Output:
	// len: 3, cap: 4, 123
	// same array: true
	// bs3: 123, bs4: 123
	// len: 4, cap: 4, xy3
}
```

### 🛠 SetMaxSize

```go
package main

import (
	"fmt"

	"github.com/fufuok/utils/pools/bytespool"
)

func main() {
	maxSize := 4096
	bytespool.SetMaxSize(maxSize)

	bs := bytespool.Make(10)
	fmt.Printf("len: %d, cap: %d\n", len(bs), cap(bs))
	bytespool.Release(bs)

	bs = bytespool.Make(maxSize)
	fmt.Printf("len: %d, cap: %d\n", len(bs), cap(bs))
	bytespool.Release(bs)

	bs = bytespool.New(maxSize + 1)
	fmt.Printf("len: %d, cap: %d\n", len(bs), cap(bs))
	ok := bytespool.Release(bs)
	fmt.Printf("Discard: %v\n", !ok)

	// Output:
	// len: 0, cap: 16
	// len: 0, cap: 4096
	// len: 4097, cap: 4097
	// Discard: true
}
```

## 🤖 Benchmarks

```go
go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=.
goos: linux
goarch: amd64
pkg: github.com/fufuok/utils/pools/bytespool
cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
BenchmarkCapacityPools/New-4            60133353                19.34 ns/op            0 B/op          0 allocs/op
BenchmarkCapacityPools/New-4            61206177                19.49 ns/op            0 B/op          0 allocs/op
BenchmarkCapacityPools/Make-4           61580971                19.58 ns/op            0 B/op          0 allocs/op
BenchmarkCapacityPools/Make-4           61389439                19.71 ns/op            0 B/op          0 allocs/op
BenchmarkCapacityPools/New.Parallel-4           240337632                5.041 ns/op           0 B/op          0 allocs/op
BenchmarkCapacityPools/New.Parallel-4           235125742                5.133 ns/op           0 B/op          0 allocs/op
BenchmarkCapacityPools/Make.Parallel-4          229302106                5.073 ns/op           0 B/op          0 allocs/op
BenchmarkCapacityPools/Make.Parallel-4          238298523                5.308 ns/op           0 B/op          0 allocs/op
```







*ff*