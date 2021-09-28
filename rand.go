package utils

import (
	crand "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandInt (>=)m - (<)n 间随机整数
func RandInt(min int, max int) int {
	x := max - min
	if x <= 0 {
		return 0
	}
	return random.Intn(x) + min
}

// RandString 随机字符串, 大小写字母数字
// Ref: stackoverflow.icza
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, random.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = random.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			// fmt.Println(idx)
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return B2S(b)
}

// RandHex 随机 Hex 编码字符串
func RandHex(nHalf int) string {
	return hex.EncodeToString(RandBytes(nHalf))
}

// RandBytes 随机 bytes
func RandBytes(n int) []byte {
	if n < 1 {
		return nil
	}

	b := make([]byte, n)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}

	return b
}
