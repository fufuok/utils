package utils

import (
	"testing"
)

func TestAesGCMEnStringHex(t *testing.T) {
	actual, _ := AesGCMEnStringHex("1", []byte("1"))
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringHex("", tmpK)
	AssertEqual(t, "", AesGCMDeStringHex(res, tmpK, nonce))

	res, nonce = AesGCMEnStringHex(tmpS, tmpK)
	AssertEqual(t, tmpS, AesGCMDeStringHex(res, tmpK, nonce))
}

func TestAesGCMEnStringB64(t *testing.T) {
	actual, _ := AesGCMEnStringB64("1", []byte("1"))
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB64("", tmpK)
	AssertEqual(t, "", AesGCMDeStringB64(res, tmpK, nonce))

	res, nonce = AesGCMEnStringB64(tmpS, tmpK)
	AssertEqual(t, tmpS, AesGCMDeStringB64(res, tmpK, nonce))
}

func TestAesGCMEnStringB58(t *testing.T) {
	actual, _ := AesGCMEnStringB58("1", []byte("1"))
	AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB58("", tmpK)
	AssertEqual(t, "", AesGCMDeStringB58(res, tmpK, nonce))

	res, nonce = AesGCMEnStringB58(tmpS, tmpK)
	AssertEqual(t, tmpS, AesGCMDeStringB58(res, tmpK, nonce))
}
