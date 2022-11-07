package xhash

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/fnv"
	"hash/maphash"
	"io"
	"os"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/fufuok/utils"
)

const (
	bufferSize = 65536

	// FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset32 = 2166136261
	offset64 = 14695981039346656037
	prime32  = 16777619
	prime64  = 1099511628211
)

func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256(utils.S2B(s)))
}

func Sha256(b []byte) []byte {
	return Hash(b, sha256.New())
}

func Sha512Hex(s string) string {
	return hex.EncodeToString(Sha512(utils.S2B(s)))
}

func Sha512(b []byte) []byte {
	return Hash(b, sha512.New())
}

func Sha1Hex(s string) string {
	return hex.EncodeToString(Sha1(utils.S2B(s)))
}

func Sha1(b []byte) []byte {
	return Hash(b, sha1.New())
}

func HmacSHA256Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA256(utils.S2B(s), utils.S2B(key)))
}

func HmacSHA256(b, key []byte) []byte {
	return Hmac(b, key, sha256.New)
}

func HmacSHA512Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA512(utils.S2B(s), utils.S2B(key)))
}

func HmacSHA512(b, key []byte) []byte {
	return Hmac(b, key, sha512.New)
}

func HmacSHA1Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA1(utils.S2B(s), utils.S2B(key)))
}

func HmacSHA1(b, key []byte) []byte {
	return Hmac(b, key, sha1.New)
}

// MD5Hex 字符串 MD5
func MD5Hex(s string) string {
	b := md5.Sum(utils.S2B(s))
	return hex.EncodeToString(b[:])
}

func MD5BytesHex(bs []byte) string {
	b := md5.Sum(bs)
	return hex.EncodeToString(b[:])
}

func MD5(b []byte) []byte {
	return Hash(b, nil)
}

func Hmac(b []byte, key []byte, h func() hash.Hash) []byte {
	if h == nil {
		h = md5.New
	}
	mac := hmac.New(h, key)
	mac.Write(b)

	return mac.Sum(nil)
}

func Hash(b []byte, h hash.Hash) []byte {
	if h == nil {
		h = md5.New()
	}
	h.Reset()
	h.Write(b)

	return h.Sum(nil)
}

func MustMD5Sum(filename string) string {
	s, _ := MD5Sum(filename)
	return s
}

// MD5Sum 文件 MD5
func MD5Sum(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	return MD5Reader(file)
}

// MD5Reader 计算 MD5
func MD5Reader(r io.Reader) (string, error) {
	h := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(r); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		h.Write(buf[:n])
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sum64 获取字符串的哈希值
func Sum64(s string) uint64 {
	return AddString64(offset64, s)
}

// SumBytes64 获取 bytes 的哈希值
func SumBytes64(bs []byte) uint64 {
	return AddBytes64(offset64, bs)
}

// Sum32 获取字符串的哈希值
func Sum32(s string) uint32 {
	return AddString32(offset32, s)
}

// SumBytes32 获取 bytes 的哈希值
func SumBytes32(bs []byte) uint32 {
	return AddBytes32(offset32, bs)
}

// FnvHash 获取字符串的哈希值
func FnvHash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// FnvHash32 获取字符串的哈希值
func FnvHash32(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// MemHashb 使用内置的 memhash 获取哈希值
func MemHashb(b []byte) uint64 {
	h := (*reflect.StringHeader)(unsafe.Pointer(&b))
	return uint64(memhash(unsafe.Pointer(h.Data), offset64, uintptr(h.Len)))
}

// MemHash 使用内置的 memhash 获取字符串哈希值
func MemHash(s string) uint64 {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return uint64(memhash(unsafe.Pointer(h.Data), offset64, uintptr(h.Len)))
}

// MemHashb32 使用内置的 memhash 获取哈希值
func MemHashb32(b []byte) uint32 {
	h := (*reflect.StringHeader)(unsafe.Pointer(&b))
	return uint32(memhash(unsafe.Pointer(h.Data), offset32, uintptr(h.Len)))
}

// MemHash32 使用内置的 memhash 获取字符串哈希值
func MemHash32(s string) uint32 {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return uint32(memhash(unsafe.Pointer(h.Data), offset32, uintptr(h.Len)))
}

// Djb33 比 FnvHash32 更快的获取字符串哈希值
// djb2 with better shuffling. 5x faster than FNV with the hash.Hash overhead.
// Ref: patrickmn/go-cache
func Djb33(s string) uint32 {
	var (
		l = uint32(len(s))
		d = 5381 + utils.Seed + l
		i = uint32(0)
	)
	// Why is all this 5x faster than a for loop?
	if l >= 4 {
		for i < l-4 {
			d = (d * 33) ^ uint32(s[i])
			d = (d * 33) ^ uint32(s[i+1])
			d = (d * 33) ^ uint32(s[i+2])
			d = (d * 33) ^ uint32(s[i+3])
			i += 4
		}
	}
	switch l - i {
	case 1:
	case 2:
		d = (d * 33) ^ uint32(s[i])
	case 3:
		d = (d * 33) ^ uint32(s[i])
		d = (d * 33) ^ uint32(s[i+1])
	case 4:
		d = (d * 33) ^ uint32(s[i])
		d = (d * 33) ^ uint32(s[i+1])
		d = (d * 33) ^ uint32(s[i+2])
	}
	return d ^ (d >> 16)
}

// HashString 合并一串文本, 得到字符串哈希
func HashString(s ...string) string {
	return strconv.FormatUint(HashString64(s...), 10)
}

func HashString64(s ...string) uint64 {
	return Sum64(utils.AddString(s...))
}

func HashString32(s ...string) uint32 {
	return Sum32(utils.AddString(s...))
}

// HashBytes 合并 Bytes, 得到字符串哈希
func HashBytes(b ...[]byte) string {
	return strconv.FormatUint(HashBytes64(b...), 10)
}

func HashBytes64(b ...[]byte) uint64 {
	return SumBytes64(utils.JoinBytes(b...))
}

func HashBytes32(b ...[]byte) uint32 {
	return SumBytes32(utils.JoinBytes(b...))
}

// HashUint64 returns the hash of u.
// Ref: segmentio/fasthash
func HashUint64(u uint64) uint64 {
	return AddUint64(offset64, u)
}

// HashUint32 returns the hash of u.
// Ref: segmentio/fasthash
func HashUint32(u uint32) uint32 {
	return AddUint32(offset32, u)
}

// AddString64 adds the hash of s to the precomputed hash value h.
// Ref: segmentio/fasthash
func AddString64(h uint64, s string) uint64 {
	/*
		This is an unrolled version of this algorithm:
		for _, c := range s {
			h = (h ^ uint64(c)) * prime64
		}
		It seems to be ~1.5x faster than the simple loop in BenchmarkHash64:
		- BenchmarkHash64/hash_function-4   30000000   56.1 ns/op   642.15 MB/s   0 B/op   0 allocs/op
		- BenchmarkHash64/hash_function-4   50000000   38.6 ns/op   932.35 MB/s   0 B/op   0 allocs/op
	*/
	for len(s) >= 8 {
		h = (h ^ uint64(s[0])) * prime64
		h = (h ^ uint64(s[1])) * prime64
		h = (h ^ uint64(s[2])) * prime64
		h = (h ^ uint64(s[3])) * prime64
		h = (h ^ uint64(s[4])) * prime64
		h = (h ^ uint64(s[5])) * prime64
		h = (h ^ uint64(s[6])) * prime64
		h = (h ^ uint64(s[7])) * prime64
		s = s[8:]
	}

	if len(s) >= 4 {
		h = (h ^ uint64(s[0])) * prime64
		h = (h ^ uint64(s[1])) * prime64
		h = (h ^ uint64(s[2])) * prime64
		h = (h ^ uint64(s[3])) * prime64
		s = s[4:]
	}

	if len(s) >= 2 {
		h = (h ^ uint64(s[0])) * prime64
		h = (h ^ uint64(s[1])) * prime64
		s = s[2:]
	}

	if len(s) > 0 {
		h = (h ^ uint64(s[0])) * prime64
	}

	return h
}

// AddBytes64 adds the hash of b to the precomputed hash value h.
// Ref: segmentio/fasthash
func AddBytes64(h uint64, b []byte) uint64 {
	for len(b) >= 8 {
		h = (h ^ uint64(b[0])) * prime64
		h = (h ^ uint64(b[1])) * prime64
		h = (h ^ uint64(b[2])) * prime64
		h = (h ^ uint64(b[3])) * prime64
		h = (h ^ uint64(b[4])) * prime64
		h = (h ^ uint64(b[5])) * prime64
		h = (h ^ uint64(b[6])) * prime64
		h = (h ^ uint64(b[7])) * prime64
		b = b[8:]
	}

	if len(b) >= 4 {
		h = (h ^ uint64(b[0])) * prime64
		h = (h ^ uint64(b[1])) * prime64
		h = (h ^ uint64(b[2])) * prime64
		h = (h ^ uint64(b[3])) * prime64
		b = b[4:]
	}

	if len(b) >= 2 {
		h = (h ^ uint64(b[0])) * prime64
		h = (h ^ uint64(b[1])) * prime64
		b = b[2:]
	}

	if len(b) > 0 {
		h = (h ^ uint64(b[0])) * prime64
	}

	return h
}

// AddUint64 adds the hash value of the 8 bytes of u to h.
// Ref: segmentio/fasthash
func AddUint64(h uint64, u uint64) uint64 {
	h = (h ^ ((u >> 56) & 0xFF)) * prime64
	h = (h ^ ((u >> 48) & 0xFF)) * prime64
	h = (h ^ ((u >> 40) & 0xFF)) * prime64
	h = (h ^ ((u >> 32) & 0xFF)) * prime64
	h = (h ^ ((u >> 24) & 0xFF)) * prime64
	h = (h ^ ((u >> 16) & 0xFF)) * prime64
	h = (h ^ ((u >> 8) & 0xFF)) * prime64
	h = (h ^ ((u >> 0) & 0xFF)) * prime64
	return h
}

// AddString32 adds the hash of s to the precomputed hash value h.
// Ref: segmentio/fasthash
func AddString32(h uint32, s string) uint32 {
	for len(s) >= 8 {
		h = (h ^ uint32(s[0])) * prime32
		h = (h ^ uint32(s[1])) * prime32
		h = (h ^ uint32(s[2])) * prime32
		h = (h ^ uint32(s[3])) * prime32
		h = (h ^ uint32(s[4])) * prime32
		h = (h ^ uint32(s[5])) * prime32
		h = (h ^ uint32(s[6])) * prime32
		h = (h ^ uint32(s[7])) * prime32
		s = s[8:]
	}

	if len(s) >= 4 {
		h = (h ^ uint32(s[0])) * prime32
		h = (h ^ uint32(s[1])) * prime32
		h = (h ^ uint32(s[2])) * prime32
		h = (h ^ uint32(s[3])) * prime32
		s = s[4:]
	}

	if len(s) >= 2 {
		h = (h ^ uint32(s[0])) * prime32
		h = (h ^ uint32(s[1])) * prime32
		s = s[2:]
	}

	if len(s) > 0 {
		h = (h ^ uint32(s[0])) * prime32
	}

	return h
}

// AddBytes32 adds the hash of b to the precomputed hash value h.
// Ref: segmentio/fasthash
func AddBytes32(h uint32, b []byte) uint32 {
	for len(b) >= 8 {
		h = (h ^ uint32(b[0])) * prime32
		h = (h ^ uint32(b[1])) * prime32
		h = (h ^ uint32(b[2])) * prime32
		h = (h ^ uint32(b[3])) * prime32
		h = (h ^ uint32(b[4])) * prime32
		h = (h ^ uint32(b[5])) * prime32
		h = (h ^ uint32(b[6])) * prime32
		h = (h ^ uint32(b[7])) * prime32
		b = b[8:]
	}

	if len(b) >= 4 {
		h = (h ^ uint32(b[0])) * prime32
		h = (h ^ uint32(b[1])) * prime32
		h = (h ^ uint32(b[2])) * prime32
		h = (h ^ uint32(b[3])) * prime32
		b = b[4:]
	}

	if len(b) >= 2 {
		h = (h ^ uint32(b[0])) * prime32
		h = (h ^ uint32(b[1])) * prime32
		b = b[2:]
	}

	if len(b) > 0 {
		h = (h ^ uint32(b[0])) * prime32
	}

	return h
}

// AddUint32 adds the hash value of the 8 bytes of u to h.
// Ref: segmentio/fasthash
func AddUint32(h, u uint32) uint32 {
	h = (h ^ ((u >> 24) & 0xFF)) * prime32
	h = (h ^ ((u >> 16) & 0xFF)) * prime32
	h = (h ^ ((u >> 8) & 0xFF)) * prime32
	h = (h ^ ((u >> 0) & 0xFF)) * prime32
	return h
}

// HashSeedString calculates a hash of s with the given seed.
func HashSeedString(seed maphash.Seed, s string) uint64 {
	return hashString(seed, s)
}

// HashSeedUint64 calculates a hash of n with the given seed.
func HashSeedUint64(seed maphash.Seed, n uint64) uint64 {
	// Java's Long standard hash function.
	n = n ^ (n >> 32)
	nseed := *(*uint64)(unsafe.Pointer(&seed))
	// 64-bit variation of boost's hash_combine.
	nseed ^= n + 0x9e3779b97f4a7c15 + (nseed << 12) + (nseed >> 4)
	return nseed
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr
