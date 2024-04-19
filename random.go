package utils

import (
	"math"
)

const (
	decBytes      = "0123456789"
	hexBytes      = "0123456789abcdef"
	alphaBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 31 / letterIdxBits   // # of letter indices fitting in 31 bits
)

var (
	// Rand goroutine-safe, use Rand.xxx instead of rand.xxx
	Rand = NewRand()
	Seed = FastRand()
)

// RandInt (>=)min - (<)max
func RandInt(min, max int) int {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return FastIntn(max-min) + min
}

// RandUint32 (>=)min - (<)max
func RandUint32(min, max uint32) uint32 {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return FastRandn(max-min) + min
}

// FastIntn this is similar to rand.Intn, but faster.
// A non-negative pseudo-random number in the half-open interval [0,n).
// Return 0 if n <= 0.
func FastIntn(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= math.MaxUint32 {
		return int(FastRandn(uint32(n)))
	}
	return int(Rand.Int63n(int64(n)))
}

// RandString a random string, which may contain uppercase letters, lowercase letters and numbers.
// Ref: stackoverflow.icza
func RandString(n int) string {
	return B2S(RandBytes(n))
}

// RandHexString 指定长度的随机 hex 字符串
func RandHexString(n int) string {
	return B2S(RandHexBytes(n))
}

// RandAlphaString 指定长度的随机字母字符串
func RandAlphaString(n int) string {
	return B2S(RandAlphaBytes(n))
}

// RandDecString 指定长度的随机数字字符串
func RandDecString(n int) string {
	return B2S(RandDecBytes(n))
}

// RandBytes random bytes, but faster.
func RandBytes(n int) []byte {
	return RandBytesLetters(n, letterBytes)
}

// RandAlphaBytes generates random alpha bytes.
func RandAlphaBytes(n int) []byte {
	return RandBytesLetters(n, alphaBytes)
}

// RandHexBytes generates random hexadecimal bytes.
func RandHexBytes(n int) []byte {
	return RandBytesLetters(n, hexBytes)
}

// RandDecBytes 指定长度的随机数字切片
func RandDecBytes(n int) []byte {
	return RandBytesLetters(n, decBytes)
}

// RandBytesLetters 生成指定长度的字符切片
func RandBytesLetters(n int, letters string) []byte {
	if n < 1 || len(letters) < 2 {
		return nil
	}
	b := make([]byte, n)
	for i, cache, remain := n-1, FastRand(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = FastRand(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
