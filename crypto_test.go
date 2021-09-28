package utils

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestSha256Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "ed3772cefd8991edac6d198df7b62c224b92038e2d435a9a1e2734211e5b5e0b",
		Sha256Hex(tmpS))
}

func TestSha512Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "c3b70022a04f57c1ad335256d2adb2aeec825c6641b2576b48f64abf1bb2c3210dff1087b9f"+
		"27261e4062779e64f29fc39d555164c5a2ea6eb3fd6d8b19ed1d1",
		Sha512Hex(tmpS))
}

func TestSha1Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "ad4ebd7a388c536ff4fcb494ffb36c047151f751",
		Sha1Hex(tmpS))
}

func TestHmacSHA256Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "6d502095be042aab03ac7ae36dd0ca504e54eb72569547dca4e16e5de605ae7c",
		HmacSHA256Hex(tmpS, "Fufu"))
}

func TestHmacSHA512Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "516825116f0c14c563f0fdd377729dd0a2024fb9f16f0dc81d83450e97114683ece5c8886"+
		"aaa912af970b1a40505333fb16e770e6b9df1826556e5fac680782b",
		HmacSHA512Hex(tmpS, "Fufu"))
}

func TestHmacSHA1Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "e6720f969c4aa396324845bed13e3cbf550b0d6d",
		HmacSHA1Hex(tmpS, "Fufu"))
}

func TestMD5Hex(t *testing.T) {
	for _, v := range []struct {
		in, out string
	}{
		{"12345", "827ccb0eea8a706c4c34a16891f84e7b"},
		{"Fufu 中　文\u2728->?\n*\U0001F63A", "4ed346c6efaec048fa503f329751e019"},
		{"Fufu 中　文", "0ab5820207b25880bc0a1d09ed64f10c"},
	} {
		AssertEqual(t, v.out, MD5Hex(v.in))
	}
}

func TestMD5(t *testing.T) {
	for _, v := range []struct {
		in  []byte
		out string
	}{
		{[]byte("12345"), "827ccb0eea8a706c4c34a16891f84e7b"},
		{[]byte("Fufu 中　文\u2728->?\n*\U0001F63A"), "4ed346c6efaec048fa503f329751e019"},
		{[]byte("Fufu 中　文"), "0ab5820207b25880bc0a1d09ed64f10c"},
	} {
		AssertEqual(t, v.out, hex.EncodeToString(MD5(v.in)))
	}
}

func TestMD5Sum(t *testing.T) {
	t.Parallel()
	res, _ := MD5Sum("LICENSE")
	AssertEqual(t, "8fad15baa71cfe5901d9ac1bbec2c56c", res)
}

func TestEncrypt(t *testing.T) {
	t.Parallel()
	AssertEqual(t, tmpS, Encrypt(tmpS, ""))
	AssertEqual(t, "8xMVBFnucwTQEmD8ZfPYWFnmXHwT3hDBN2m5vWKhRXhFXNmrJ188SMhfGTsbSPCv1R",
		Encrypt(tmpS, "\u1234\x33 Fufu@中文"))
	AssertEqual(t, "7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5",
		Encrypt(tmpS, "0123456789012345"))
}

func TestDecrypt(t *testing.T) {
	t.Parallel()
	AssertEqual(t, tmpS,
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
	AssertEqual(t, tmpS, res1)
	AssertEqual(t, tmpS, res2)
	AssertEqual(t, "Fufu", res3)
	AssertEqual(t, "Fufu", Decrypt(encrypt, "1"))
}
