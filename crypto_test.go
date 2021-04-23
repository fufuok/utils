package utils

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestSha256Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "99b3ab61f5d8ccfa4a90ba4d70e4ab39bef28e5edcdb34f45f06d5f945aa6d19",
		Sha256Hex("解码/编码~ 顶替&#123a"))
}

func TestSha512Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "9391580d363cdaa1662f49f6e4d36ef50e6590f3b6a7a6ddba8d2e0e30951b968b533037"+
		"99c45859d8e8ae3f560fce94c9f96d043a7d936cebd0b7d2dd36dd72",
		Sha512Hex("解码/编码~ 顶替&#123a"))
}

func TestSha1Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "5318d394704c5c1fbe78aed167df4a1b2d50b53c",
		Sha1Hex("解码/编码~ 顶替&#123a"))
}

func TestHmacSHA256Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "6763dc663717bff585cab57ed83e4a1bae6f849eff89d9797e6fd90bb5f41ae6",
		HmacSHA256Hex("解码/编码~ 顶替&#123a", "Fufu"))
}

func TestHmacSHA512Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "7b583f85065e526111d0a10cdf508933e2df1a79f0b74f1285c3a0e2471abb6519fcfd313805"+
		"07c74808727d8d483efa7fd52e097b8f0ab29efff18b7ce2bc62",
		HmacSHA512Hex("解码/编码~ 顶替&#123a", "Fufu"))
}

func TestHmacSHA1Hex(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "8901f3e883769e9706d2a352bb937894d6a6c911",
		HmacSHA1Hex("解码/编码~ 顶替&#123a", "Fufu"))
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
	AssertEqual(t, "9da5118f5101a96cca4e261153587a90", res)
}

func TestEncrypt(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "解码/编码~ 顶替&#", Encrypt("解码/编码~ 顶替&#", ""))
	AssertEqual(t, "BxcwmUGXjHHq9XTcZnpbmdzFWUCx7SFscefaT9KepvZB",
		Encrypt("解码/编码~ 顶替&#", "\u1234\x33 Fufu@中文"))
	AssertEqual(t, "7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5",
		Encrypt("Fufu 中　文加密/解密~&#123a", "0123456789012345"))
}

func TestDecrypt(t *testing.T) {
	t.Parallel()
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a",
		Decrypt("7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5",
			"0123456789012345"))
}

func TestGetenvDecrypt(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("GO_TEST_1", "解码/编码~ 顶替&#")
	_ = os.Setenv("GO_TEST_2", "BxcwmUGXjHHq9XTcZnpbmdzFWUCx7SFscefaT9KepvZB")
	encrypt, _ := SetenvEncrypt("GO_TEST_3", "Fufu", "1")
	res1 := GetenvDecrypt("GO_TEST_1", "")
	res2 := GetenvDecrypt("GO_TEST_2", "\u1234\x33 Fufu@中文")
	res3 := GetenvDecrypt("GO_TEST_3", "1")
	AssertEqual(t, "解码/编码~ 顶替&#", res1)
	AssertEqual(t, "解码/编码~ 顶替&#", res2)
	AssertEqual(t, "Fufu", res3)
	AssertEqual(t, "Fufu", Decrypt(encrypt, "1"))
}
