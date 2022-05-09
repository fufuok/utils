package xcrypto

import (
	"crypto/rc4"
	"encoding/hex"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/base58"
)

// XOREnStringHex 加密
func XOREnStringHex(s string, key []byte) string {
	return XOREnHex(utils.S2B(s), key)
}

// XOREnHex 加密
func XOREnHex(b, key []byte) string {
	return hex.EncodeToString(XOR(b, key))
}

// XORDeStringHex 解密
func XORDeStringHex(s string, key []byte) string {
	return utils.B2S(XORDeHex(s, key))
}

// XORDeHex 解密
func XORDeHex(s string, key []byte) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return XOR(data, key)
	}

	return nil
}

// XOREnStringB58 加密
func XOREnStringB58(s string, key []byte) string {
	return XOREnB58(utils.S2B(s), key)
}

// XOREnB58 加密
func XOREnB58(b, key []byte) string {
	return base58.Encode(XOR(b, key))
}

// XORDeStringB58 解密
func XORDeStringB58(s string, key []byte) string {
	return utils.B2S(XORDeB58(s, key))
}

// XORDeB58 解密
func XORDeB58(s string, key []byte) []byte {
	return XOR(base58.Decode(s), key)
}

// XOREnStringB64 加密
func XOREnStringB64(s string, key []byte) string {
	return XOREnB64(utils.S2B(s), key)
}

// XOREnB64 加密
func XOREnB64(b, key []byte) string {
	return utils.B64UrlEncode(XOR(b, key))
}

// XORDeStringB64 解密
func XORDeStringB64(s string, key []byte) string {
	return utils.B2S(XORDeB64(s, key))
}

// XORDeB64 解密
func XORDeB64(s string, key []byte) []byte {
	return XOR(utils.B64UrlDecode(s), key)
}

// XOR 异或加解密
func XOR(src, key []byte) []byte {
	dst, _ := XORE(src, key)
	return dst
}

// XORE RC4 加密算法(异或运算), 简单加解密, 不够安全
// key 长度是 1-256
func XORE(src, key []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return src, err
	}

	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)

	return dst, nil
}
