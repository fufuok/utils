package utils

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"

	"github.com/fufuok/utils/base58"
)

// DES 加密, ZerosPadding
func DesCBCEnStringHex(s, key string) string {
	return hex.EncodeToString(DesCBCEncrypt(false, S2B(s), S2B(key)))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7StringHex(s, key string) string {
	return hex.EncodeToString(DesCBCEncrypt(true, S2B(s), S2B(key)))
}

// DES 加密, ZerosPadding
func DesCBCEnHex(b, key []byte) string {
	return hex.EncodeToString(DesCBCEncrypt(false, b, key))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7Hex(b, key []byte) string {
	return hex.EncodeToString(DesCBCEncrypt(true, b, key))
}

// DES 解密, ZerosPadding
func DesCBCDeStringHex(s, key string) string {
	return B2S(DesCBCDeHex(s, key))
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7StringHex(s, key string) string {
	return B2S(DesCBCDePKCS7Hex(s, key))
}

// DES 解密, ZerosPadding
func DesCBCDeHex(s, key string) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return DesCBCDecrypt(false, data, S2B(key))
	}

	return nil
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7Hex(s, key string) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return DesCBCDecrypt(true, data, S2B(key))
	}

	return nil
}

// DES 加密, ZerosPadding
func DesCBCEnStringB58(s, key string) string {
	return base58.Encode(DesCBCEncrypt(false, S2B(s), S2B(key)))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7StringB58(s, key string) string {
	return base58.Encode(DesCBCEncrypt(true, S2B(s), S2B(key)))
}

// DES 加密, ZerosPadding
func DesCBCEnB58(b, key []byte) string {
	return base58.Encode(DesCBCEncrypt(false, b, key))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7B58(b, key []byte) string {
	return base58.Encode(DesCBCEncrypt(true, b, key))
}

// DES 解密, ZerosPadding
func DesCBCDeStringB58(s, key string) string {
	return B2S(DesCBCDeB58(s, key))
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7StringB58(s, key string) string {
	return B2S(DesCBCDePKCS7B58(s, key))
}

// DES 解密, ZerosPadding
func DesCBCDeB58(s, key string) []byte {
	return DesCBCDecrypt(false, base58.Decode(s), S2B(key))
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7B58(s, key string) []byte {
	return DesCBCDecrypt(true, base58.Decode(s), S2B(key))
}

// DES 加密, ZerosPadding
func DesCBCEnStringB64(s, key string) string {
	return B64UrlEncode(DesCBCEncrypt(false, S2B(s), S2B(key)))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7StringB64(s, key string) string {
	return B64UrlEncode(DesCBCEncrypt(true, S2B(s), S2B(key)))
}

// DES 加密, ZerosPadding
func DesCBCEnB64(b, key []byte) string {
	return B64UrlEncode(DesCBCEncrypt(false, b, key))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7B64(b, key []byte) string {
	return B64UrlEncode(DesCBCEncrypt(true, b, key))
}

// DES 解密, ZerosPadding
func DesCBCDeStringB64(s, key string) string {
	return B2S(DesCBCDeB64(s, key))
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7StringB64(s, key string) string {
	return B2S(DesCBCDePKCS7B64(s, key))
}

// DES 解密, ZerosPadding
func DesCBCDeB64(s, key string) []byte {
	return DesCBCDecrypt(false, B64UrlDecode(s), S2B(key))
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7B64(s, key string) []byte {
	return DesCBCDecrypt(true, B64UrlDecode(s), S2B(key))
}

// AES-CBC 加密
func DesCBCEncrypt(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) (ciphertext []byte) {
	ciphertext, _ = DesCBCEncryptE(asPKCS7, plaintext, key, ivs...)
	return
}

// AES-CBC 解密
func DesCBCDecrypt(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) (plaintext []byte) {
	plaintext, _ = DesCBCDecryptE(asPKCS7, ciphertext, key, ivs...)
	return
}

// DES-CBC 加密, 密码分组链接模式 (Cipher Block Chaining (CBC))
// key 长度固定为 8
// asPKCS7: false (ZerosPadding), true (Pkcs7Padding)
func DesCBCEncryptE(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	bSize := block.BlockSize()
	plaintext = Padding(plaintext, bSize, asPKCS7)

	// 向量无效时自动为 key[:blockSize]
	var iv []byte
	if len(ivs) > 0 && len(ivs[0]) == bSize {
		iv = ivs[0]
	} else {
		iv = key[:bSize]
	}

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// DES-CBC 解密, 密码分组链接模式 (Cipher Block Chaining (CBC))
func DesCBCDecryptE(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
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

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = UnPadding(plaintext, asPKCS7)

	return plaintext, nil
}
