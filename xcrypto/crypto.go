package xcrypto

import (
	"os"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/xhash"
)

// SetenvEncrypt 加密并设置环境变量(string)
func SetenvEncrypt(key, value, secret string) (string, error) {
	value = Encrypt(value, secret)
	if err := os.Setenv(key, value); err != nil {
		return "", err
	}

	return value, nil
}

// GetenvDecrypt 解密环境变量参数(string)
func GetenvDecrypt(key string, secret string) string {
	return Decrypt(os.Getenv(key), secret)
}

// Encrypt 加密 (密钥取 32 位 MD5, AES-CBC, base58)
func Encrypt(value, secret string) string {
	if secret != "" {
		value = AesCBCEnStringB58(value, utils.S2B(xhash.MD5Hex(secret)))
	}

	return value
}

// Decrypt 解密
func Decrypt(value, secret string) string {
	if secret != "" {
		value = AesCBCDeStringB58(value, utils.S2B(xhash.MD5Hex(secret)))
	}

	return value
}
