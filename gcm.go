package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/fufuok/utils/base58"
)

const gcmStandardNonceSize = 12

// AES 加密
func AesGCMEnStringHex(s, key string) (string, string) {
	return AesGCMEnHex(S2B(s), S2B(key))
}

// AES 加密
func AesGCMEnHex(b, key []byte) (string, string) {
	res, nonce := AesGCMEncrypt(b, key)
	return hex.EncodeToString(res), B2S(nonce)
}

// AES 解密
func AesGCMDeStringHex(s, key, nonce string) string {
	return B2S(AesGCMDeHex(s, key, nonce))
}

// AES 解密
func AesGCMDeHex(s, key, nonce string) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return AesGCMDecrypt(data, S2B(key), S2B(nonce))
	}

	return nil
}

// AES 加密
func AesGCMEnStringB58(s, key string) (string, string) {
	return AesGCMEnB58(S2B(s), S2B(key))
}

// AES 加密
func AesGCMEnB58(b, key []byte) (string, string) {
	res, nonce := AesGCMEncrypt(b, key)
	return base58.Encode(res), B2S(nonce)
}

// AES 解密
func AesGCMDeStringB58(s, key, nonce string) string {
	return B2S(AesGCMDeB58(s, key, nonce))
}

// AES 解密
func AesGCMDeB58(s, key, nonce string) []byte {
	return AesGCMDecrypt(base58.Decode(s), S2B(key), S2B(nonce))
}

// AES 加密
func AesGCMEnStringB64(s, key string) (string, string) {
	return AesGCMEnB64(S2B(s), S2B(key))
}

// AES 加密
func AesGCMEnB64(b, key []byte) (string, string) {
	res, nonce := AesGCMEncrypt(b, key)
	return B64UrlEncode(res), B2S(nonce)
}

// AES 解密
func AesGCMDeStringB64(s, key, nonce string) string {
	return B2S(AesGCMDeB64(s, key, nonce))
}

// AES 解密
func AesGCMDeB64(s, key, nonce string) []byte {
	return AesGCMDecrypt(B64UrlDecode(s), S2B(key), S2B(nonce))
}

// AES-GCM 加密
func AesGCMEncrypt(plaintext, key []byte) (ciphertext []byte, nonce []byte) {
	ciphertext, nonce, _ = AesGCMEncryptWithNonce(plaintext, key, nil, nil)
	return
}

// AES-GCM 解密
func AesGCMDecrypt(ciphertext, key, nonce []byte) (plaintext []byte) {
	plaintext, _ = AesGCMDecryptWithNonce(ciphertext, key, nonce, nil)
	return
}

// AES-GCM 加密, (Galois/Counter Mode (GCM))
// key 长度分别是 16 (AES-128), 32 (AES-256)
func AesGCMEncryptWithNonce(plaintext, key, nonce, additionalData []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	gcm, err := cipher.NewGCMWithNonceSize(block, gcmStandardNonceSize)
	if err != nil {
		return nil, nil, err
	}

	if len(nonce) == 0 {
		nonce = S2B(RandString(gcmStandardNonceSize))
	}
	res := gcm.Seal(nil, nonce, plaintext, additionalData)

	return res, nonce, nil
}

// AES-GCM 解密, (Galois/Counter Mode (GCM))
func AesGCMDecryptWithNonce(ciphertext, key, nonce, additionalData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	res, err := gcm.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		return nil, err
	}

	return res, nil
}
