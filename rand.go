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

var src = rand.NewSource(time.Now().UnixNano())

type TChoice struct {
	Item   interface{}
	Weight int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *TChoice) String() string {
	return MustJSONString(c)
}

// 加权随机, 返回选中项
func WeightedChoice(choices ...TChoice) TChoice {
	idx := WeightedChoiceIndex(choices)
	if idx == -1 {
		return TChoice{}
	}

	return choices[idx]
}

// 加权随机, 返回选中项的下标, -1 未选中任何项
func WeightedChoiceIndex(choices []TChoice) int {
	if len(choices) == 0 {
		return -1
	}

	m := 0
	for _, c := range choices {
		if c.Weight > 0 {
			m += c.Weight
		}
	}

	if m < 1 {
		return -1
	}

	n := rand.Intn(m)
	for i, c := range choices {
		if c.Weight <= 0 {
			continue
		}
		n -= c.Weight
		if n < 0 {
			return i
		}
	}

	return -1
}

// 加权随机, 参数为权重值列表, 返回选中项的下标, -1 未选中任何项
func WeightedChoiceWeightsIndex(weights []int) int {
	if len(weights) == 0 {
		return -1
	}

	m := 0
	for _, w := range weights {
		if w > 0 {
			m += w
		}
	}

	if m < 1 {
		return -1
	}

	n := rand.Intn(m)
	for i, w := range weights {
		if w <= 0 {
			continue
		}
		n -= w
		if n < 0 {
			return i
		}
	}

	return -1
}

// 加权随机, map[object]weight
func WeightedChoiceMap(choices map[interface{}]int) interface{} {
	if len(choices) == 0 {
		return nil
	}

	m := 0
	for _, w := range choices {
		if w > 0 {
			m += w
		}
	}

	if m < 1 {
		return nil
	}

	n := rand.Intn(m)
	for k, w := range choices {
		n -= w
		if n < 0 {
			return k
		}
	}

	return nil
}

// (>=)m - (<)n 间随机整数
func RandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

// 随机字符串, 大小写字母数字
// Ref: stackoverflow.icza
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
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

// 随机 Hex 编码字符串
func RandHex(nHalf int) string {
	return hex.EncodeToString(RandBytes(nHalf))
}

// 随机 bytes
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
