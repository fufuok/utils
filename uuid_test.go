package utils

import (
	"testing"

	"github.com/fufuok/utils/xid"
)

func TestUUIDString(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 10000; i++ {
		id := UUIDString()
		if m[id] {
			t.Error("duplicated UUID:", id)
		}
		m[id] = true
		AssertEqual(t, uint8(4), UUID()[6]>>4)
		AssertEqual(t, uint8(0x80), UUID()[8]&0xc0)
	}
}

func BenchmarkUniqueUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUID()
	}
}

func BenchmarkUniqueUUIDString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUIDString()
	}
}

func BenchmarkUniqueUUIDSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUIDSimple()
	}
}

func BenchmarkUniqueXID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = xid.NewBytes()
	}
}

func BenchmarkUniqueXIDString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = xid.NewString()
	}
}

// BenchmarkUniqueUUID-8         	 1516308	      2409 ns/op	      16 B/op	       1 allocs/op
// BenchmarkUniqueUUID-8         	 1428595	      3121 ns/op	      16 B/op	       1 allocs/op
// BenchmarkUniqueUUID-8         	 1454421	      2641 ns/op	      16 B/op	       1 allocs/op
// BenchmarkUniqueUUIDString-8   	 1320424	      2716 ns/op	      64 B/op	       2 allocs/op
// BenchmarkUniqueUUIDString-8   	 1400548	      2770 ns/op	      64 B/op	       2 allocs/op
// BenchmarkUniqueUUIDString-8   	 1000000	      3083 ns/op	      64 B/op	       2 allocs/op
// BenchmarkUniqueUUIDSimple-8   	 1000000	      3235 ns/op	      80 B/op	       3 allocs/op
// BenchmarkUniqueUUIDSimple-8   	 1202796	      2658 ns/op	      80 B/op	       3 allocs/op
// BenchmarkUniqueUUIDSimple-8   	 1317488	      2608 ns/op	      80 B/op	       3 allocs/op
// BenchmarkUniqueXID-8          	 2616777	      1411 ns/op	       0 B/op	       0 allocs/op
// BenchmarkUniqueXID-8          	 2530496	      1378 ns/op	       0 B/op	       0 allocs/op
// BenchmarkUniqueXID-8          	 2729223	      1389 ns/op	       0 B/op	       0 allocs/op
// BenchmarkUniqueXIDString-8    	 2332801	      1525 ns/op	      32 B/op	       1 allocs/op
// BenchmarkUniqueXIDString-8    	 2385540	      1611 ns/op	      32 B/op	       1 allocs/op
// BenchmarkUniqueXIDString-8    	 2197720	      1516 ns/op	      32 B/op	       1 allocs/op
