package utils

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
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
func DesCBCEnStringB64(s, key string) string {
	return base64.URLEncoding.EncodeToString(DesCBCEncrypt(false, S2B(s), S2B(key)))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7StringB64(s, key string) string {
	return base64.URLEncoding.EncodeToString(DesCBCEncrypt(true, S2B(s), S2B(key)))
}

// DES 加密, ZerosPadding
func DesCBCEnB64(b, key []byte) string {
	return base64.URLEncoding.EncodeToString(DesCBCEncrypt(false, b, key))
}

// DES 加密, Pkcs7Padding
func DesCBCEnPKCS7B64(b, key []byte) string {
	return base64.URLEncoding.EncodeToString(DesCBCEncrypt(true, b, key))
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
	if data, err := base64.URLEncoding.DecodeString(s); err == nil {
		return DesCBCDecrypt(false, data, S2B(key))
	}

	return nil
}

// DES 解密, Pkcs7Padding
func DesCBCDePKCS7B64(s, key string) []byte {
	if data, err := base64.URLEncoding.DecodeString(s); err == nil {
		return DesCBCDecrypt(true, data, S2B(key))
	}

	return nil
}

// DES-CBC 加密, 密码分组链接模式 (Cipher Block Chaining (CBC))
// key 长度固定为 8
// asPKCS7: false (ZerosPadding), true (Pkcs7Padding)
func DesCBCEncrypt(asPKCS7 bool, b, key []byte, ivs ...[]byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

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

// DES-CBC 解密, 密码分组链接模式 (Cipher Block Chaining (CBC))
func DesCBCDecrypt(asPKCS7 bool, b, key []byte, ivs ...[]byte) []byte {
	block, err := des.NewCipher(key)
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
