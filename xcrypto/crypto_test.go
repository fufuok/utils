package xcrypto

import (
	"os"
	"testing"

	"github.com/fufuok/utils/assert"
)

var (
	tmpS = "Fufu 中　文加密/解密~&#123a"
	tmpK = []byte("0123456789012345")
	tmpB = []byte("Fufu 中　文\u2728->?\n*\U0001F63A0123456789012345")
)

func TestEncrypt(t *testing.T) {
	t.Parallel()
	assert.Equal(t, tmpS, Encrypt(tmpS, ""))
	assert.Equal(t, "8xMVBFnucwTQEmD8ZfPYWFnmXHwT3hDBN2m5vWKhRXhFXNmrJ188SMhfGTsbSPCv1R",
		Encrypt(tmpS, "\u1234\x33 Fufu@中文"))
	assert.Equal(t, "7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5",
		Encrypt(tmpS, "0123456789012345"))
}

func TestDecrypt(t *testing.T) {
	t.Parallel()
	assert.Equal(t, tmpS,
		Decrypt("7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5",
			"0123456789012345"))
}

func TestGetenvDecrypt(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("GO_TEST_1", tmpS)
	_ = os.Setenv("GO_TEST_2", "8xMVBFnucwTQEmD8ZfPYWFnmXHwT3hDBN2m5vWKhRXhFXNmrJ188SMhfGTsbSPCv1R")
	encrypt, _ := SetenvEncrypt("GO_TEST_3", "Fufu", "1")
	res1 := GetenvDecrypt("GO_TEST_1", "")
	res2 := GetenvDecrypt("GO_TEST_2", "\u1234\x33 Fufu@中文")
	res3 := GetenvDecrypt("GO_TEST_3", "1")
	assert.Equal(t, tmpS, res1)
	assert.Equal(t, tmpS, res2)
	assert.Equal(t, "Fufu", res3)
	assert.Equal(t, "Fufu", Decrypt(encrypt, "1"))
}
