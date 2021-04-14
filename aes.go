package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// AES 加密, ZerosPadding
func AesCBCEnStringHex(s, key string) string {
	return hex.EncodeToString(AesCBCEncrypt(false, S2B(s), S2B(key)))
}

// AES 加密, Pkcs7Padding
func AesCBCEnPKCS7StringHex(s, key string) string {
	return hex.EncodeToString(AesCBCEncrypt(true, S2B(s), S2B(key)))
}

// AES 加密, ZerosPadding
func AesCBCEnHex(b, key []byte) string {
	return hex.EncodeToString(AesCBCEncrypt(false, b, key))
}

// AES 加密, Pkcs7Padding
func AesCBCEnPKCS7Hex(b, key []byte) string {
	return hex.EncodeToString(AesCBCEncrypt(true, b, key))
}

// AES 解密, ZerosPadding
func AesCBCDeStringHex(s, key string) string {
	return B2S(AesCBCDeHex(s, key))
}

// AES 解密, Pkcs7Padding
func AesCBCDePKCS7StringHex(s, key string) string {
	return B2S(AesCBCDePKCS7Hex(s, key))
}

// AES 解密, ZerosPadding
func AesCBCDeHex(s, key string) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return AesCBCDecrypt(false, data, S2B(key))
	}

	return nil
}

// AES 解密, Pkcs7Padding
func AesCBCDePKCS7Hex(s, key string) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return AesCBCDecrypt(true, data, S2B(key))
	}

	return nil
}

// AES 加密, ZerosPadding
func AesCBCEnStringB64(s, key string) string {
	return base64.URLEncoding.EncodeToString(AesCBCEncrypt(false, S2B(s), S2B(key)))
}

// AES 加密, Pkcs7Padding
func AesCBCEnPKCS7StringB64(s, key string) string {
	return base64.URLEncoding.EncodeToString(AesCBCEncrypt(true, S2B(s), S2B(key)))
}

// AES 加密, ZerosPadding
func AesCBCEnB64(b, key []byte) string {
	return base64.URLEncoding.EncodeToString(AesCBCEncrypt(false, b, key))
}

// AES 加密, Pkcs7Padding
func AesCBCEnPKCS7B64(b, key []byte) string {
	return base64.URLEncoding.EncodeToString(AesCBCEncrypt(true, b, key))
}

// AES 解密, ZerosPadding
func AesCBCDeStringB64(s, key string) string {
	return B2S(AesCBCDeB64(s, key))
}

// AES 解密, Pkcs7Padding
func AesCBCDePKCS7StringB64(s, key string) string {
	return B2S(AesCBCDePKCS7B64(s, key))
}

// AES 解密, ZerosPadding
func AesCBCDeB64(s, key string) []byte {
	if data, err := base64.URLEncoding.DecodeString(s); err == nil {
		return AesCBCDecrypt(false, data, S2B(key))
	}

	return nil
}

// AES 解密, Pkcs7Padding
func AesCBCDePKCS7B64(s, key string) []byte {
	if data, err := base64.URLEncoding.DecodeString(s); err == nil {
		return AesCBCDecrypt(true, data, S2B(key))
	}

	return nil
}

// AES-CBC 加密, 密码分组链接模式 (Cipher Block Chaining (CBC))
// key 长度分别是 16 (AES-128), 24 (AES-192?), 32 (AES-256?)
// asPKCS7: false (ZerosPadding), true (Pkcs7Padding)
func AesCBCEncrypt(asPKCS7 bool, b, key []byte, ivs ...[]byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	// go1.15.6: const BlockSize = 16
	bSize := block.BlockSize()
	b = Padding(b, bSize, asPKCS7)

	// 向量无效时自动为 key[:blockSize]
	var iv []byte
	if len(ivs) > 0 && len(ivs[0]) == bSize {
		iv = ivs[0]
	} else {
		iv = key[:bSize]
	}

	res := make([]byte, len(b))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(res, b)

	return res
}

// AES-CBC 解密, 密码分组链接模式 (Cipher Block Chaining (CBC))
func AesCBCDecrypt(asPKCS7 bool, b, key []byte, ivs ...[]byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	bSize := block.BlockSize()

	// 向量无效时自动为 key[:blockSize]
	var iv []byte
	if len(ivs) > 0 && len(ivs[0]) == bSize {
		iv = ivs[0]
	} else {
		iv = key[:bSize]
	}

	res := make([]byte, len(b))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(res, b)
	res = UnPadding(res, asPKCS7)

	return res
}

// 填充
func Padding(b []byte, bSize int, pkcs7 bool) []byte {
	if pkcs7 {
		n := bSize - len(b)%bSize
		pad := bytes.Repeat([]byte{byte(n)}, n)
		return append(b, pad...)
	}

	n := bSize - len(b)%bSize
	if n == 0 {
		return b
	} else {
		return append(b, bytes.Repeat([]byte{byte(0)}, n)...)
	}
}

// 去除填充
func UnPadding(b []byte, pkcs7 bool) []byte {
	if pkcs7 {
		l := len(b)
		n := int(b[l-1])
		return b[:(l - n)]
	}
	for i := len(b) - 1; ; i-- {
		if b[i] != 0 {
			return b[:i+1]
		}
	}
}
