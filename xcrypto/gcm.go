package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/base58"
)

const gcmStandardNonceSize = 12

// AesGCMEnStringHex 加密
func AesGCMEnStringHex(s string, key []byte) (string, []byte) {
	return AesGCMEnHex(utils.S2B(s), key)
}

// AesGCMEnHex 加密
func AesGCMEnHex(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return hex.EncodeToString(res), nonce
}

// AesGCMDeStringHex 解密
func AesGCMDeStringHex(s string, key, nonce []byte) string {
	return utils.B2S(AesGCMDeHex(s, key, nonce))
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
	return AesGCMEnB58(utils.S2B(s), key)
}

// AesGCMEnB58 加密
func AesGCMEnB58(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return base58.Encode(res), nonce
}

// AesGCMDeStringB58 解密
func AesGCMDeStringB58(s string, key, nonce []byte) string {
	return utils.B2S(AesGCMDeB58(s, key, nonce))
}

// AesGCMDeB58 解密
func AesGCMDeB58(s string, key, nonce []byte) []byte {
	return AesGCMDecrypt(base58.Decode(s), key, nonce)
}

// AesGCMEnStringB64 加密
func AesGCMEnStringB64(s string, key []byte) (string, []byte) {
	return AesGCMEnB64(utils.S2B(s), key)
}

// AesGCMEnB64 加密
func AesGCMEnB64(b, key []byte) (string, []byte) {
	res, nonce := AesGCMEncrypt(b, key)
	return utils.B64UrlEncode(res), nonce
}

// AesGCMDeStringB64 解密
func AesGCMDeStringB64(s string, key, nonce []byte) string {
	return utils.B2S(AesGCMDeB64(s, key, nonce))
}

// AesGCMDeB64 解密
func AesGCMDeB64(s string, key, nonce []byte) []byte {
	return AesGCMDecrypt(utils.B64UrlDecode(s), key, nonce)
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
		nonce = utils.S2B(utils.RandString(gcmStandardNonceSize))
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

// GCMEnStringHex 加密
func GCMEnStringHex(s string, key []byte) string {
	return GCMEnHex(utils.S2B(s), key)
}

// GCMEnHex 加密
func GCMEnHex(b, key []byte) string {
	if res, err := GCMEncrypt(b, key); err == nil {
		return hex.EncodeToString(res)
	}

	return ""
}

// GCMDeStringHex 解密
func GCMDeStringHex(s string, key []byte) string {
	return utils.B2S(GCMDeHex(s, key))
}

// GCMDeHex 解密
func GCMDeHex(s string, key []byte) []byte {
	if data, err := hex.DecodeString(s); err == nil {
		if res, err := GCMDecrypt(data, key); err == nil {
			return res
		}
	}

	return nil
}

// GCMEnStringB58 加密
func GCMEnStringB58(s string, key []byte) string {
	return GCMEnB58(utils.S2B(s), key)
}

// GCMEnB58 加密
func GCMEnB58(b, key []byte) string {
	if res, err := GCMEncrypt(b, key); err == nil {
		return base58.Encode(res)
	}

	return ""
}

// GCMDeStringB58 解密
func GCMDeStringB58(s string, key []byte) string {
	return utils.B2S(GCMDeB58(s, key))
}

// GCMDeB58 解密
func GCMDeB58(s string, key []byte) []byte {
	if res, err := GCMDecrypt(base58.Decode(s), key); err == nil {
		return res
	}

	return nil
}

// GCMEnStringB64 加密
func GCMEnStringB64(s string, key []byte) string {
	return GCMEnB64(utils.S2B(s), key)
}

// GCMEnB64 加密
func GCMEnB64(b, key []byte) string {
	if res, err := GCMEncrypt(b, key); err == nil {
		return utils.B64UrlEncode(res)
	}

	return ""
}

// GCMDeStringB64 解密
func GCMDeStringB64(s string, key []byte) string {
	return utils.B2S(GCMDeB64(s, key))
}

// GCMDeB64 解密
func GCMDeB64(s string, key []byte) []byte {
	if res, err := GCMDecrypt(utils.B64UrlDecode(s), key); err == nil {
		return res
	}

	return nil
}

// GCMEncrypt AES-GCM 加密
func GCMEncrypt(plaintext, key []byte) ([]byte, error) {
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

	nonce := utils.RandBytes(gcm.NonceSize())
	res := gcm.Seal(nonce, nonce, plaintext, nil)

	return res, nil
}

// GCMDecrypt AES-GCM 解密
func GCMDecrypt(encrypted, key []byte) ([]byte, error) {
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

	nonceSize := gcm.NonceSize()

	if len(encrypted) < nonceSize {
		return nil, errors.New("encrypted value is not valid")
	}

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
