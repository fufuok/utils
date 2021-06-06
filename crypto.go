package utils

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

const bufferSize = 65536

func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256(S2B(s)))
}

func Sha256(b []byte) []byte {
	return Hash(b, sha256.New())
}

func Sha512Hex(s string) string {
	return hex.EncodeToString(Sha512(S2B(s)))
}

func Sha512(b []byte) []byte {
	return Hash(b, sha512.New())
}

func Sha1Hex(s string) string {
	return hex.EncodeToString(Sha1(S2B(s)))
}

func Sha1(b []byte) []byte {
	return Hash(b, sha1.New())
}

func HmacSHA256Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA256(S2B(s), S2B(key)))
}

func HmacSHA256(b, key []byte) []byte {
	return Hmac(b, key, sha256.New)
}

func HmacSHA512Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA512(S2B(s), S2B(key)))
}

func HmacSHA512(b, key []byte) []byte {
	return Hmac(b, key, sha512.New)
}

func HmacSHA1Hex(s, key string) string {
	return hex.EncodeToString(HmacSHA1(S2B(s), S2B(key)))
}

func HmacSHA1(b, key []byte) []byte {
	return Hmac(b, key, sha1.New)
}

// MD5Hex 字符串 MD5
func MD5Hex(s string) string {
	b := md5.Sum(S2B(s))
	return hex.EncodeToString(b[:])
}

func MD5(b []byte) []byte {
	return Hash(b, nil)
}

func Hmac(b []byte, key []byte, h func() hash.Hash) []byte {
	if h == nil {
		h = md5.New
	}
	mac := hmac.New(h, key)
	mac.Write(b)

	return mac.Sum(nil)
}

func Hash(b []byte, h hash.Hash) []byte {
	if h == nil {
		h = md5.New()
	}
	h.Reset()
	h.Write(b)

	return h.Sum(nil)
}

// MD5Sum 文件 MD5
func MD5Sum(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	return MD5Reader(file)
}

// MD5Reader 计算 MD5
func MD5Reader(r io.Reader) (string, error) {
	h := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(r); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		h.Write(buf[:n])
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

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
		value = AesCBCEnStringB58(value, S2B(MD5Hex(secret)))
	}

	return value
}

// Decrypt 解密
func Decrypt(value, secret string) string {
	if secret != "" {
		value = AesCBCDeStringB58(value, S2B(MD5Hex(secret)))
	}

	return value
}
