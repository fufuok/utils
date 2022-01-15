package utils

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
	"io"
	"os"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	bufferSize = 65536

	// FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset32 = 2166136261
	offset64 = 14695981039346656037
	prime32  = 16777619
	prime64  = 1099511628211
)

var Seed = FastRand()

func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256(S2B(s)))
}

func Sha256(b []byte) []byte {
	return Hash(b, sha256.New())
}

func Sha512Hex(s string) string {
	return hex.EncodeToString(Sha512(S2B(s)))
}

func Sha512(b []byte) []byte {
	return Hash(b, sha512.New())
}

func Sha1Hex(s string) string {
	return hex.EncodeToString(Sha1(S2B(s)))
}

func Sha1(b []byte) []byte {
	return Hash(b, sha1.New())
}

func HmacSHA256Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA256(S2B(s), S2B(key)))
}

func HmacSHA256(b, key []byte) []byte {
	return Hmac(b, key, sha256.New)
}

func HmacSHA512Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA512(S2B(s), S2B(key)))
}

func HmacSHA512(b, key []byte) []byte {
	return Hmac(b, key, sha512.New)
}

func HmacSHA1Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA1(S2B(s), S2B(key)))
}

func HmacSHA1(b, key []byte) []byte {
	return Hmac(b, key, sha1.New)
}

// MD5Hex 字符串 MD5
func MD5Hex(s string) string {
	b := md5.Sum(S2B(s))
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
	var h uint64 = offset64
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// Sum32 获取字符串的哈希值
func Sum32(s string) uint32 {
	var h uint32 = offset32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= prime32
	}
	return h
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
		d = 5381 + Seed + l
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
	return strconv.FormatUint(HashStringUint64(s...), 10)
}

func HashStringUint64(s ...string) uint64 {
	return Sum64(AddString(s...))
}

// HashBytes 合并 Bytes, 得到字符串哈希
func HashBytes(b ...[]byte) string {
	return strconv.FormatUint(HashBytesUint64(b...), 10)
}

func HashBytesUint64(b ...[]byte) uint64 {
	return Sum64(B2S(JoinBytes(b...)))
}
