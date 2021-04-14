package utils

import (
	"testing"
)

func TestDesCBCEnStringHex(t *testing.T) {
	actual := DesCBCEnStringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCEnStringHex("", "12345678")
	AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnStringHex("Fufu 中　文加密/解密~&#123a", "12345678")
	AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33", actual)
}

func TestDesCBCEnPKCS7StringHex(t *testing.T) {
	actual := DesCBCEnPKCS7StringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7StringHex("", "12345678")
	AssertEqual(t, "4431cc0267954866", actual)

	actual = DesCBCEnPKCS7StringHex("Fufu 中　文加密/解密~&#123a", "12345678")
	AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0", actual)
}

func TestDesCBCEnHex(t *testing.T) {
	actual := DesCBCEnHex([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = DesCBCEnHex(nil, []byte("12345678"))
	AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnHex([]byte(""), []byte("12345678"))
	AssertEqual(t, "96d0028878d58c89", actual)

	actual = DesCBCEnHex([]byte("Fufu 中　文加密/解密~&#123a"), []byte("12345678"))
	AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33", actual)
}

func TestDesCBCEnPKCS7Hex(t *testing.T) {
	actual := DesCBCEnPKCS7Hex([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7Hex([]byte(""), []byte("12345678"))
	AssertEqual(t, "4431cc0267954866", actual)

	actual = DesCBCEnPKCS7Hex([]byte("Fufu 中　文加密/解密~&#123a"), []byte("12345678"))
	AssertEqual(t,
		"6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0", actual)
}

func TestDesCBCDeStringHex(t *testing.T) {
	actual := DesCBCDeStringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCDeStringHex("96d0028878d58c89", "12345678")
	AssertEqual(t, "", actual)

	actual = DesCBCDeStringHex("6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77da8134c5db36d30b33",
		"12345678")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestDesCBCDePKCS7StringHex(t *testing.T) {
	actual := DesCBCDePKCS7StringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringHex("4431cc0267954866", "12345678")
	AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringHex("6470f9113753aa52ce510d7c2d0d64c22c91648e351359dd95d59f27865e77daf05ea2f6175eebc0",
		"12345678")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestDesCBCEnStringB64(t *testing.T) {
	actual := DesCBCEnStringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCEnStringB64("", "12345678")
	AssertEqual(t, "ltACiHjVjIk=", actual)

	actual = DesCBCEnStringB64("Fufu 中　文加密/解密~&#123a", "12345678")
	AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", actual)
}

func TestDesCBCEnPKCS7StringB64(t *testing.T) {
	actual := DesCBCEnPKCS7StringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7StringB64("", "12345678")
	AssertEqual(t, "RDHMAmeVSGY=", actual)

	actual = DesCBCEnPKCS7StringB64("Fufu 中　文加密/解密~&#123a", "12345678")
	AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==", actual)
}

func TestDesCBCEnB64(t *testing.T) {
	actual := DesCBCEnB64([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = DesCBCEnB64(nil, []byte("12345678"))
	AssertEqual(t, "ltACiHjVjIk=", actual)

	actual = DesCBCEnB64([]byte("Fufu 中　文加密/解密~&#123a"), []byte("12345678"))
	AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", actual)
}

func TestDesCBCEnPKCS7B64(t *testing.T) {
	actual := DesCBCEnPKCS7B64([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = DesCBCEnPKCS7B64(nil, []byte("12345678"))
	AssertEqual(t, "RDHMAmeVSGY=", actual)

	actual = DesCBCEnPKCS7B64([]byte("Fufu 中　文加密/解密~&#123a"), []byte("12345678"))
	AssertEqual(t,
		"ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==", actual)
}

func TestDesCBCDeStringB64(t *testing.T) {
	actual := DesCBCDeStringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCDeStringB64("ltACiHjVjIk=", "12345678")
	AssertEqual(t, "", actual)

	actual = DesCBCDeStringB64("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==", "12345678")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestDesCBCDePKCS7StringB644(t *testing.T) {
	actual := DesCBCDePKCS7StringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB64("RDHMAmeVSGY=", "12345678")
	AssertEqual(t, "", actual)

	actual = DesCBCDePKCS7StringB64("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==",
		"12345678")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestDesCBCEncrypt(t *testing.T) {
	b := []byte("Fufu 中　文加密/解密~&#123a")

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
		AssertEqual(t, v.expected, B64Encode(B2S(actual)))
	}

	actual := DesCBCEncrypt(false, []byte(""), []byte("abcdefgh"))
	AssertEqual(t, "Ko1p3p1f3/k=", B64Encode(B2S(actual)))

	actual = DesCBCEncrypt(false, []byte("1"), []byte("abcdefgh"))
	AssertEqual(t, "BBV4AdvFQhg=", B64Encode(B2S(actual)))

	// 指定向量加密
	actual = DesCBCEncrypt(false, b, []byte("12345678"), []byte("abcdefgh"))
	AssertEqual(t, "qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3m5/6K/HDHFyg==", B64Encode(B2S(actual)))
}

func TestDesCBCEncryptPKCS7(t *testing.T) {
	b := []byte("Fufu 中　文加密/解密~&#123a")

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
		AssertEqual(t, v.expected, B64Encode(B2S(actual)))
	}

	// 指定向量加密
	actual := DesCBCEncrypt(true, b, []byte("12345678"), []byte("abcdefgh"))
	AssertEqual(t, "qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3n1QubNNBPFLA==", B64Encode(B2S(actual)))
}

func TestDesCBCDecrypt(t *testing.T) {
	expected := "Fufu 中　文加密/解密~&#123a"

	AssertEqual(t, *new([]byte), DesCBCDecrypt(false, nil, nil))
	AssertEqual(t, *new([]byte), DesCBCDecrypt(false, []byte("1"), nil))
	AssertEqual(t, *new([]byte), DesCBCDecrypt(false, []byte("1"), []byte("2")))
	AssertEqual(t, []byte(expected),
		DesCBCDecrypt(false,
			S2B(B64Decode("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9qBNMXbNtMLMw==")), []byte("12345678")))

	// 指定向量加密
	actual := DesCBCDecrypt(
		false,
		S2B(B64Decode("qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3m5/6K/HDHFyg==")),
		[]byte("12345678"),
		[]byte("abcdefgh"),
	)
	AssertEqual(t, expected, B2S(actual))
}

func TestDesCBCDecryptPKCS7(t *testing.T) {
	expected := "Fufu 中　文加密/解密~&#123a"

	AssertEqual(t, *new([]byte), DesCBCDecrypt(true, nil, nil))
	AssertEqual(t, *new([]byte), DesCBCDecrypt(true, []byte("1"), nil))
	AssertEqual(t, *new([]byte), DesCBCDecrypt(true, []byte("1"), []byte("2")))
	AssertEqual(t, []byte(expected),
		DesCBCDecrypt(true,
			S2B(B64Decode("ZHD5ETdTqlLOUQ18LQ1kwiyRZI41E1ndldWfJ4Zed9rwXqL2F17rwA==")), []byte("12345678")))

	// 指定向量加密
	actual := DesCBCDecrypt(
		true,
		S2B(B64Decode("qGrKD8THQqjHNafLu/KDK6cqWoTKRxExsh4k8TQQp3n1QubNNBPFLA==")),
		[]byte("12345678"),
		[]byte("abcdefgh"),
	)
	AssertEqual(t, expected, B2S(actual))
}
