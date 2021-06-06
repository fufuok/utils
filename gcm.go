package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/fufuok/utils/base58"
)

const gcmStandardNonceSize = 12

// AesGCMEnStringHex 加密
func AesGCMEnStringHex(s string, key []byte) (string, []byte) {
	return AesGCMEnHex(S2B(s), key)
}

// AesGCMEnHex 加密
func AesGCMEnHex(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return hex.EncodeToString(res), nonce
}

// AesGCMDeStringHex 解密
func AesGCMDeStringHex(s string, key, nonce []byte) string {
	return B2S(AesGCMDeHex(s, key, nonce))
}

// AesGCMDeHex 解密
func AesGCMDeHex(s string, key, nonce []byte) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		return AesGCMDecrypt(data, key, nonce)
	}

	return nil
}

// AesGCMEnStringB58 加密
func AesGCMEnStringB58(s string, key []byte) (string, []byte) {
	return AesGCMEnB58(S2B(s), key)
}

// AesGCMEnB58 加密
func AesGCMEnB58(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return base58.Encode(res), nonce
}

// AesGCMDeStringB58 解密
func AesGCMDeStringB58(s string, key, nonce []byte) string {
	return B2S(AesGCMDeB58(s, key, nonce))
}

// AesGCMDeB58 解密
func AesGCMDeB58(s string, key, nonce []byte) []byte {
	return AesGCMDecrypt(base58.Decode(s), key, nonce)
}

// AesGCMEnStringB64 加密
func AesGCMEnStringB64(s string, key []byte) (string, []byte) {
	return AesGCMEnB64(S2B(s), key)
}

// AesGCMEnB64 加密
func AesGCMEnB64(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return B64UrlEncode(res), nonce
}

// AesGCMDeStringB64 解密
func AesGCMDeStringB64(s string, key, nonce []byte) string {
	return B2S(AesGCMDeB64(s, key, nonce))
}

// AesGCMDeB64 解密
func AesGCMDeB64(s string, key, nonce []byte) []byte {
	return AesGCMDecrypt(B64UrlDecode(s), key, nonce)
}

// AesGCMEncrypt AES-GCM 加密
func AesGCMEncrypt(plaintext, key []byte) (ciphertext []byte, nonce []byte) {
	ciphertext, nonce, _ = AesGCMEncryptWithNonce(plaintext, key, nil, nil)
	return
}

// AesGCMDecrypt AES-GCM 解密
func AesGCMDecrypt(ciphertext, key, nonce []byte) (plaintext []byte) {
	plaintext, _ = AesGCMDecryptWithNonce(ciphertext, key, nonce, nil)
	return
}

// AesGCMEncryptWithNonce AES-GCM 加密, (Galois/Counter Mode (GCM))
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

// AesGCMDecryptWithNonce AES-GCM 解密, (Galois/Counter Mode (GCM))
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
