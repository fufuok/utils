package xcrypto

import (
	"testing"

	"github.com/fufuok/utils"
)

func TestDesCBCEnStringHex(t *testing.T) {
	actual := DesCBCEnStringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnStringHex("", []byte("12345678"))
	utils.AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnStringHex(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33", actual)
}

func TestDesCBCEnPKCS7StringHex(t *testing.T) {
	actual := DesCBCEnPKCS7StringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7StringHex("", []byte("12345678"))
	utils.AssertEqual(t, "4431cc0267954866", actual)

	actual = DesCBCEnPKCS7StringHex(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0", actual)
}

func TestDesCBCEnHex(t *testing.T) {
	actual := DesCBCEnHex([]byte(("1")), []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnHex(nil, []byte(("12345678")))
	utils.AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnHex([]byte(""), []byte(("12345678")))
	utils.AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnHex([]byte(tmpS), []byte("12345678"))
	utils.AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33", actual)
}

func TestDesCBCEnPKCS7Hex(t *testing.T) {
	actual := DesCBCEnPKCS7Hex([]byte(("1")), []byte(("1")))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7Hex([]byte(""), []byte(("12345678")))
	utils.AssertEqual(t, "4431cc0267954866", actual)

	actual = DesCBCEnPKCS7Hex([]byte(tmpS), []byte(("12345678")))
	utils.AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0", actual)
}

func TestDesCBCDeStringHex(t *testing.T) {
	actual := DesCBCDeStringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringHex("96d0028878d58c89", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringHex("6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33",
		[]byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCDePKCS7StringHex(t *testing.T) {
	actual := DesCBCDePKCS7StringHex("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringHex("4431cc0267954866", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringHex("6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0",
		[]byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCEnStringB58(t *testing.T) {
	actual := DesCBCEnStringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnStringB58("", []byte("12345678"))
	utils.AssertEqual(t, "SE56JGzQwap", actual)

	actual = DesCBCEnStringB58(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"5zTdexqNF1BMJXSBTNocnSapeoor27RnosuA5QncNRGDkAzsCXkqsa2", actual)
}

func TestDesCBCEnPKCS7StringB58(t *testing.T) {
	actual := DesCBCEnPKCS7StringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7StringB58("", []byte("12345678"))
	utils.AssertEqual(t, "CQaC5oNT7z5", actual)

	actual = DesCBCEnPKCS7StringB58(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"5zTdexqNF1BMJXSBTNocnSapeoor27RnosuA5QncNRGE4mRYYqbJS6K", actual)
}

func TestDesCBCDeStringB58(t *testing.T) {
	actual := DesCBCDeStringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringB58("SE56JGzQwap", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringB58("5zTdexqNF1BMJXSBTNocnSapeoor27RnosuA5QncNRGDkAzsCXkqsa2", []byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCDePKCS7StringB58(t *testing.T) {
	actual := DesCBCDePKCS7StringB58("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB58("CQaC5oNT7z5", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB58("5zTdexqNF1BMJXSBTNocnSapeoor27RnosuA5QncNRGE4mRYYqbJS6K",
		[]byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCEnStringB64(t *testing.T) {
	actual := DesCBCEnStringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnStringB64("", []byte("12345678"))
	utils.AssertEqual(t, "ltACiHjVjIk=", actual)

	actual = DesCBCEnStringB64(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", actual)
}

func TestDesCBCEnPKCS7StringB64(t *testing.T) {
	actual := DesCBCEnPKCS7StringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7StringB64("", []byte("12345678"))
	utils.AssertEqual(t, "RDHMAmeVSGY=", actual)

	actual = DesCBCEnPKCS7StringB64(tmpS, []byte("12345678"))
	utils.AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==", actual)
}

func TestDesCBCEnB64(t *testing.T) {
	actual := DesCBCEnB64([]byte(("1")), []byte(("1")))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnB64(nil, []byte(("12345678")))
	utils.AssertEqual(t, "ltACiHjVjIk=", actual)

	actual = DesCBCEnB64([]byte(tmpS), []byte("12345678"))
	utils.AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", actual)
}

func TestDesCBCEnPKCS7B64(t *testing.T) {
	actual := DesCBCEnPKCS7B64([]byte(("1")), []byte(("1")))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7B64(nil, []byte(("12345678")))
	utils.AssertEqual(t, "RDHMAmeVSGY=", actual)

	actual = DesCBCEnPKCS7B64([]byte(tmpS), []byte("12345678"))
	utils.AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==", actual)
}

func TestDesCBCDeStringB64(t *testing.T) {
	actual := DesCBCDeStringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringB64("ltACiHjVjIk=", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDeStringB64("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", []byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCDePKCS7StringB64(t *testing.T) {
	actual := DesCBCDePKCS7StringB64("1", []byte("1"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB64("RDHMAmeVSGY=", []byte("12345678"))
	utils.AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB64("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==",
		[]byte("12345678"))
	utils.AssertEqual(t, tmpS, actual)
}

func TestDesCBCEncrypt(t *testing.T) {
	b := []byte(tmpS)

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{[]byte("12345678"), "ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw=="},
	} {
		actual := DesCBCEncrypt(false, b, v.key)
		utils.AssertEqual(t, v.expected, utils.B64Encode(actual))
	}

	actual := DesCBCEncrypt(false, []byte(""), []byte("abcdefgh"))
	utils.AssertEqual(t, "Ko1p3p1f3/k=", utils.B64Encode(actual))

	actual = DesCBCEncrypt(false, []byte(("1")), []byte("abcdefgh"))
	utils.AssertEqual(t, "BBV4AdvFQhg=", utils.B64Encode(actual))

	// 指定向量加密
	actual = DesCBCEncrypt(false, b, []byte(("12345678")), []byte("abcdefgh"))
	utils.AssertEqual(t, "qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3m5/6K/HDHFyg==", utils.B64Encode(actual))
}

func TestDesCBCEncryptPKCS7(t *testing.T) {
	b := []byte(tmpS)

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{[]byte("12345678"), "ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA=="},
	} {
		actual := DesCBCEncrypt(true, b, v.key)
		utils.AssertEqual(t, v.expected, utils.B64Encode(actual))
	}

	// 指定向量加密
	actual := DesCBCEncrypt(true, b, []byte(("12345678")), []byte("abcdefgh"))
	utils.AssertEqual(t, "qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3n1QubNNBPFLA==", utils.B64Encode(actual))
}

func TestDesCBCDecrypt(t *testing.T) {
	expected := tmpS

	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(false, nil, nil))
	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(false, []byte(("1")), nil))
	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(false, []byte(("1")), []byte("2")))
	utils.AssertEqual(t, []byte(expected),
		DesCBCDecrypt(false,
			utils.B64Decode("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw=="), []byte(("12345678"))))

	// 指定向量加密
	actual := DesCBCDecrypt(
		false,
		utils.B64Decode("qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3m5/6K/HDHFyg=="),
		[]byte(("12345678")),
		[]byte("abcdefgh"),
	)
	utils.AssertEqual(t, expected, utils.B2S(actual))
}

func TestDesCBCDecryptPKCS7(t *testing.T) {
	expected := tmpS

	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(true, nil, nil))
	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(true, []byte(("1")), nil))
	utils.AssertEqual(t, *new([]byte), DesCBCDecrypt(true, []byte(("1")), []byte("2")))
	utils.AssertEqual(t, []byte(expected),
		DesCBCDecrypt(true,
			utils.B64Decode("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA=="), []byte("12345678")))

	// 指定向量加密
	actual := DesCBCDecrypt(
		true,
		utils.B64Decode("qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3n1QubNNBPFLA=="),
		[]byte("12345678"),
		[]byte("abcdefgh"),
	)
	utils.AssertEqual(t, expected, utils.B2S(actual))
}
