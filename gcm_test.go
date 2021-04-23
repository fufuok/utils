package utils

import (
	"testing"
)

func TestAesGCMEnStringHex(t *testing.T) {
	actual, _ := AesGCMEnStringHex("1", "1")
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringHex("", "0123456789012345")
	AssertEqual(t, "", AesGCMDeStringHex(res, "0123456789012345", nonce))

	res, nonce = AesGCMEnStringHex("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", AesGCMDeStringHex(res, "0123456789012345", nonce))
}

func TestAesGCMEnStringB64(t *testing.T) {
	actual, _ := AesGCMEnStringB64("1", "1")
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB64("", "0123456789012345")
	AssertEqual(t, "", AesGCMDeStringB64(res, "0123456789012345", nonce))

	res, nonce = AesGCMEnStringB64("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", AesGCMDeStringB64(res, "0123456789012345", nonce))
}

func TestAesGCMEnStringB58(t *testing.T) {
	actual, _ := AesGCMEnStringB58("1", "1")
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB58("", "0123456789012345")
	AssertEqual(t, "", AesGCMDeStringB58(res, "0123456789012345", nonce))

	res, nonce = AesGCMEnStringB58("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", AesGCMDeStringB58(res, "0123456789012345", nonce))
}
