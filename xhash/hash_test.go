package xhash

import (
	"encoding/hex"
	"hash/maphash"
	"math"
	"strconv"
	"testing"

	"github.com/fufuok/utils"
)

var (
	tmpS       = "Fufu 中　文加密/解密~&#123a"
	testString = "  Fufu 中　文\u2728->?\n*\U0001F63A   "
	testBytes  = []byte(testString)
	seed       = maphash.MakeSeed()
)

func TestSha256Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "ed3772cefd8991edac6d198df7b62c224b92038e2d435a9a1e2734211e5b5e0b",
		Sha256Hex(tmpS))
}

func TestSha512Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "c3b70022a04f57c1ad335256d2adb2aeec825c6641b2576b48f64abf1bb2c3210dff1087b9f"+
		"27261e4062779e64f29fc39d555164c5a2ea6eb3fd6d8b19ed1d1",
		Sha512Hex(tmpS))
}

func TestSha1Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "ad4ebd7a388c536ff4fcb494ffb36c047151f751",
		Sha1Hex(tmpS))
}

func TestHmacSHA256Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "6d502095be042aab03ac7ae36dd0ca504e54eb72569547dca4e16e5de605ae7c",
		HmacSHA256Hex(tmpS, "Fufu"))
}

func TestHmacSHA512Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "516825116f0c14c563f0fdd377729dd0a2024fb9f16f0dc81d83450e97114683ece5c8886"+
		"aaa912af970b1a40505333fb16e770e6b9df1826556e5fac680782b",
		HmacSHA512Hex(tmpS, "Fufu"))
}

func TestHmacSHA1Hex(t *testing.T) {
	t.Parallel()
	utils.AssertEqual(t, "e6720f969c4aa396324845bed13e3cbf550b0d6d",
		HmacSHA1Hex(tmpS, "Fufu"))
}

func TestMD5Hex(t *testing.T) {
	for _, v := range []struct {
		in, out string
	}{
		{"12345", "827ccb0eea8a706c4c34a16891f84e7b"},
		{testString, "8d47309acf79aa15378c82475c167865"},
		{"Fufu 中　文", "0ab5820207b25880bc0a1d09ed64f10c"},
	} {
		utils.AssertEqual(t, v.out, MD5Hex(v.in))
		utils.AssertEqual(t, v.out, MD5BytesHex([]byte(v.in)))
		utils.AssertEqual(t, v.out, hex.EncodeToString(MD5([]byte(v.in))))
	}
}

func TestMD5Sum(t *testing.T) {
	t.Parallel()
	res, _ := MD5Sum("LICENSE")
	expected := []string{
		// Real result
		"8fad15baa71cfe5901d9ac1bbec2c56c",
		// Result on github windows (LF would be replaced by CRLF, Maybe core.autocrlf is true)
		"cd5c4d3bd8efa894619c1f3eab8a9174",
	}
	utils.AssertEqual(t, true, utils.InStrings(expected, res))
	utils.AssertEqual(t, true, utils.InStrings(expected, MustMD5Sum("LICENSE")))
}

func TestHashString(t *testing.T) {
	for _, v := range []struct {
		in, out string
	}{
		{"", "14695981039346656037"},
		{"12345", "16534377278781491704"},
		{testString, "13467076781014605639"},
		{"Fufu 中　文", "1485575821508720008"},
	} {
		utils.AssertEqual(t, v.out, HashString(v.in))
		utils.AssertEqual(t, v.out, HashBytes([]byte(v.in)))
	}

	for _, v := range []struct {
		in  string
		out uint32
	}{
		{"", 2166136261},
		{"12345", 1136836824},
		{testString, 475021159},
		{"Fufu 中　文", 2300112168},
	} {
		utils.AssertEqual(t, v.out, HashString32(v.in))
		utils.AssertEqual(t, v.out, HashBytes32([]byte(v.in)))
	}
}

func TestHashUint(t *testing.T) {
	for _, v := range []struct {
		in, out uint64
	}{
		{0, 12161962213042174405},
		{1, 12161961113530546194},
		{offset64, 5063649278745802162},
		{prime64, 14714463944532698764},
		{math.MaxUint64, 10157053723145373757},
	} {
		utils.AssertEqual(t, v.out, HashUint64(v.in))
	}

	for _, v := range []struct {
		in, out uint32
	}{
		{0, 1268118805},
		{1, 1251341186},
		{offset32, 2870680790},
		{prime32, 2389395716},
		{math.MaxUint32, 3809873841},
	} {
		utils.AssertEqual(t, v.out, HashUint32(v.in))
	}
}

func TestHashStringToInt(t *testing.T) {
	utils.AssertEqual(t, uint64(offset64), Sum64(""))
	utils.AssertEqual(t, uint32(offset32), Sum32(""))
	utils.AssertEqual(t, uint64(offset64), FnvHash(""))
	utils.AssertEqual(t, uint32(offset32), FnvHash32(""))
	utils.AssertEqual(t, uint64(13467076781014605639), Sum64(testString))
	utils.AssertEqual(t, uint32(475021159), Sum32(testString))
	utils.AssertEqual(t, uint64(13467076781014605639), FnvHash(testString))
	utils.AssertEqual(t, uint32(475021159), FnvHash32(testString))

	v := MemHash(testString)
	for i := 0; i < 100000; i++ {
		utils.AssertEqual(t, v, MemHash(testString))
	}

	v = MemHashb(testBytes)
	for i := 0; i < 100000; i++ {
		utils.AssertEqual(t, v, MemHashb(testBytes))
	}

	v32 := MemHash32(testString)
	for i := 0; i < 100000; i++ {
		utils.AssertEqual(t, v32, MemHash32(testString))
	}

	v32 = Djb33(testString)
	for i := 0; i < 100000; i++ {
		utils.AssertEqual(t, v32, Djb33(testString))
	}
}

func TestHashSeedString(t *testing.T) {
	const numEntries = 1000
	m := make(map[uint64]uint64)
	for i := 0; i < numEntries; i++ {
		k := HashSeedString(seed, strconv.Itoa(i))
		if _, ok := m[k]; ok {
			t.Fatalf("expect key %d to not exist", k)
		}
		m[k] = k
	}
	if len(m) != numEntries {
		t.Fatalf("expect count of %d, but got: %d", numEntries, len(m))
	}
}

func TestHashSeedUint64(t *testing.T) {
	const numEntries = 1000
	m := make(map[uint64]uint64)
	for i := 0; i < numEntries; i++ {
		k := HashSeedUint64(seed, uint64(i))
		if _, ok := m[k]; ok {
			t.Fatalf("expect key %d to not exist", k)
		}
		m[k] = k
	}
	if len(m) != numEntries {
		t.Fatalf("expect count of %d, but got: %d", numEntries, len(m))
	}
}

func BenchmarkHashString(b *testing.B) {
	str := utils.RandString(20)
	b.ResetTimer()
	b.Run("MD5Hex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MD5Hex(str)
		}
	})
	b.Run("HashString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = HashString(str)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkHashString
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkHashString/MD5Hex-8             5115880               226.1 ns/op            64 B/op          2 allocs/op
// BenchmarkHashString/MD5Hex-8             5377476               230.3 ns/op            64 B/op          2 allocs/op
// BenchmarkHashString/MD5Hex-8             5227404               226.4 ns/op            64 B/op          2 allocs/op
// BenchmarkHashString/HashString-8        11307057               102.8 ns/op            24 B/op          1 allocs/op
// BenchmarkHashString/HashString-8        12567037               96.83 ns/op            24 B/op          1 allocs/op
// BenchmarkHashString/HashString-8        12295094               101.1 ns/op            24 B/op          1 allocs/op

func BenchmarkHash(b *testing.B) {
	buf := utils.RandBytes(20)
	str := utils.RandString(20)
	b.ResetTimer()
	b.Run("Sum64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Sum64(str)
		}
	})
	b.Run("Sum32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Sum32(str)
		}
	})
	b.Run("FnvHash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FnvHash(str)
		}
	})
	b.Run("FnvHash32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FnvHash32(str)
		}
	})
	b.Run("MemHashb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MemHashb(buf)
		}
	})
	b.Run("MemHash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MemHash(str)
		}
	})
	b.Run("MemHashb32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MemHashb32(buf)
		}
	})
	b.Run("MemHash32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = MemHash32(str)
		}
	})
	b.Run("Djb33", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Djb33(str)
		}
	})
	b.Run("SeedString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = HashSeedString(seed, str)
		}
	})
}

// go test -run=^$ -benchmem -benchtime=1s -count=3 -bench=BenchmarkHash$
// goos: linux
// goarch: amd64
// pkg: github.com/fufuok/utils
// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkHash/Sum64-8           53414522                22.67 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Sum64-8           54474142                22.05 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Sum64-8           47609222                22.30 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Sum32-8           48003264                21.97 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Sum32-8           56251669                21.41 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Sum32-8           51219012                22.23 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash-8         28918519                37.55 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash-8         33735345                35.68 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash-8         36855715                37.35 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash32-8       35083820                32.14 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash32-8       47042220                28.03 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/FnvHash32-8       46833836                28.16 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb-8       153161814                9.470 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb-8       152563350                8.079 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb-8       142128412                7.964 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash-8        125632875                8.775 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash-8        150410337                8.372 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash-8        135773145                9.427 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb32-8     146760392                7.950 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb32-8     146271644                7.759 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHashb32-8     154832995                8.686 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash32-8      100000000                10.13 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash32-8      100000000                10.30 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/MemHash32-8      143592832                8.360 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Djb33-8           60568434                18.17 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Djb33-8           75151240                18.17 ns/op            0 B/op          0 allocs/op
// BenchmarkHash/Djb33-8           68428705                19.79 ns/op            0 B/op          0 allocs/op
