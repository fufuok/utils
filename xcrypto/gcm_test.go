package xcrypto

import (
	"testing"

	"github.com/fufuok/utils"
)

func TestAesGCMEnStringHex(t *testing.T) {
	actual, _ := AesGCMEnStringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringHex("", tmpK)
	utils.AssertEqual(t, "", AesGCMDeStringHex(res, tmpK, nonce))

	res, nonce = AesGCMEnStringHex(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, AesGCMDeStringHex(res, tmpK, nonce))
}

func TestAesGCMEnStringB64(t *testing.T) {
	actual, _ := AesGCMEnStringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB64("", tmpK)
	utils.AssertEqual(t, "", AesGCMDeStringB64(res, tmpK, nonce))

	res, nonce = AesGCMEnStringB64(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, AesGCMDeStringB64(res, tmpK, nonce))
}

func TestAesGCMEnStringB58(t *testing.T) {
	actual, _ := AesGCMEnStringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res, nonce := AesGCMEnStringB58("", tmpK)
	utils.AssertEqual(t, "", AesGCMDeStringB58(res, tmpK, nonce))

	res, nonce = AesGCMEnStringB58(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, AesGCMDeStringB58(res, tmpK, nonce))
}

func TestGCMEnStringHex(t *testing.T) {
	actual := GCMEnStringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res := GCMEnStringHex("", tmpK)
	utils.AssertEqual(t, "", GCMDeStringHex(res, tmpK))

	res = GCMEnStringHex(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, GCMDeStringHex(res, tmpK))
}

func TestGCMEnStringB64(t *testing.T) {
	actual := GCMEnStringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res := GCMEnStringB64("", tmpK)
	utils.AssertEqual(t, "", GCMDeStringB64(res, tmpK))

	res = GCMEnStringB64(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, GCMDeStringB64(res, tmpK))
}

func TestGCMEnStringB58(t *testing.T) {
	actual := GCMEnStringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	res := GCMEnStringB58("", tmpK)
	utils.AssertEqual(t, "", GCMDeStringB58(res, tmpK))

	res = GCMEnStringB58(tmpS, tmpK)
	utils.AssertEqual(t, tmpS, GCMDeStringB58(res, tmpK))
}
