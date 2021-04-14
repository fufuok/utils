package utils

import (
	"testing"
)

func TestAesCBCEnStringHex(t *testing.T) {
	actual := AesCBCEnStringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCEnStringHex("", "0123456789012345")
	AssertEqual(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnStringHex("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608520cbae12ea453abc9d9e66ceb7d60be", actual)
}

func TestAesCBCEnPKCS7StringHex(t *testing.T) {
	actual := AesCBCEnPKCS7StringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCEnPKCS7StringHex("", "0123456789012345")
	AssertEqual(t, "3c8a535cc4de40f7cf961e54b4e8661f", actual)

	actual = AesCBCEnPKCS7StringHex("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c56088e05e1c12308135d889866051367833b", actual)
}

func TestAesCBCEnHex(t *testing.T) {
	actual := AesCBCEnHex([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = AesCBCEnHex(nil, []byte("0123456789012345"))
	AssertEqual(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnHex([]byte(""), []byte("0123456789012345"))
	AssertEqual(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnHex([]byte("Fufu 中　文加密/解密~&#123a"), []byte("0123456789012345"))
	AssertEqual(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608520cbae12ea453abc9d9e66ceb7d60be", actual)
}

func TestAesCBCEnPKCS7Hex(t *testing.T) {
	actual := AesCBCEnPKCS7Hex([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = AesCBCEnPKCS7Hex([]byte(""), []byte("0123456789012345"))
	AssertEqual(t, "3c8a535cc4de40f7cf961e54b4e8661f", actual)

	actual = AesCBCEnPKCS7Hex([]byte("Fufu 中　文加密/解密~&#123a"), []byte("0123456789012345"))
	AssertEqual(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c56088e05e1c12308135d889866051367833b", actual)
}

func TestAesCBCDeStringHex(t *testing.T) {
	actual := AesCBCDeStringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCDeStringHex("5f7df0bf103a8c4ae6faad9906ac3b2a", "0123456789012345")
	AssertEqual(t, "", actual)

	actual = AesCBCDeStringHex("b94cd86a309b3391c50f38eb8091563df0f9b329a385ace279574278"+
		"0f8c5608520cbae12ea453abc9d9e66ceb7d60be", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestAesCBCDePKCS7StringHex(t *testing.T) {
	actual := AesCBCDePKCS7StringHex("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCDePKCS7StringHex("3c8a535cc4de40f7cf961e54b4e8661f", "0123456789012345")
	AssertEqual(t, "", actual)

	actual = AesCBCDePKCS7StringHex("b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608"+
		"8e05e1c12308135d889866051367833b", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestAesCBCEnStringB64(t *testing.T) {
	actual := AesCBCEnStringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCEnStringB64("", "0123456789012345")
	AssertEqual(t, "X33wvxA6jErm-q2ZBqw7Kg==", actual)

	actual = AesCBCEnStringB64("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-", actual)
}

func TestAesCBCEnPKCS7StringB64(t *testing.T) {
	actual := AesCBCEnPKCS7StringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCEnPKCS7StringB64("", "0123456789012345")
	AssertEqual(t, "PIpTXMTeQPfPlh5UtOhmHw==", actual)

	actual = AesCBCEnPKCS7StringB64("Fufu 中　文加密/解密~&#123a", "0123456789012345")
	AssertEqual(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7", actual)
}

func TestAesCBCEnB64(t *testing.T) {
	actual := AesCBCEnB64([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = AesCBCEnB64(nil, []byte("0123456789012345"))
	AssertEqual(t, "X33wvxA6jErm-q2ZBqw7Kg==", actual)

	actual = AesCBCEnB64([]byte("Fufu 中　文加密/解密~&#123a"), []byte("0123456789012345"))
	AssertEqual(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-", actual)
}

func TestAesCBCEnPKCS7B64(t *testing.T) {
	actual := AesCBCEnPKCS7B64([]byte("1"), []byte("1"))
	AssertEqual(t, "", actual)

	actual = AesCBCEnPKCS7B64(nil, []byte("0123456789012345"))
	AssertEqual(t, "PIpTXMTeQPfPlh5UtOhmHw==", actual)

	actual = AesCBCEnPKCS7B64([]byte("Fufu 中　文加密/解密~&#123a"), []byte("0123456789012345"))
	AssertEqual(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7", actual)
}

func TestAesCBCDeStringB64(t *testing.T) {
	actual := AesCBCDeStringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCDeStringB64("X33wvxA6jErm-q2ZBqw7Kg==", "0123456789012345")
	AssertEqual(t, "", actual)

	actual = AesCBCDeStringB64("uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-", "0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestAesCBCDePKCS7StringB644(t *testing.T) {
	actual := AesCBCDePKCS7StringB64("1", "1")
	AssertEqual(t, "", actual)

	actual = AesCBCDePKCS7StringB64("PIpTXMTeQPfPlh5UtOhmHw==", "0123456789012345")
	AssertEqual(t, "", actual)

	actual = AesCBCDePKCS7StringB64("uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7",
		"0123456789012345")
	AssertEqual(t, "Fufu 中　文加密/解密~&#123a", actual)
}

func TestAesCBCEncrypt(t *testing.T) {
	b := []byte("Fufu 中　文加密/解密~&#123a")

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{
			[]byte("0123456789012345"),
			"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA+MVghSDLrhLqRTq8nZ5mzrfWC+",
		},
		{
			[]byte("012345678901234567891234"),
			"6qxrtFlWv1xcruY1VZSalQ/tD58cCWMD368cIt1Vp5GMNlv/ER6KB77H2+PqX+33",
		},
		{
			[]byte("01234567890123456789012345678901"),
			"4drGuwuCHASVh4nta3nml/rBWgkNcUyuTcbB1R5iFfyGkEsP6N9HuxAgg0n22vb6",
		},
	} {
		actual := AesCBCEncrypt(false, b, v.key)
		AssertEqual(t, v.expected, B64Encode(B2S(actual)))
	}

	actual := AesCBCEncrypt(false, []byte(""), []byte("abcdefghabcdefgh"))
	AssertEqual(t, "/06zrVSl4UrsshCLDgplgA==", B64Encode(B2S(actual)))

	actual = AesCBCEncrypt(false, []byte("1"), []byte("abcdefghabcdefgh"))
	AssertEqual(t, "TzlqKoJCo9jWlqZ+gHvhvQ==", B64Encode(B2S(actual)))

	// 指定向量加密
	actual = AesCBCEncrypt(false, b, []byte("01234567890123456789012345678901"), []byte("abcdefghabcdefgh"))
	AssertEqual(t, "+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdpZn4m+l8HNpYkY1A30Y2xh", B64Encode(B2S(actual)))
}

func TestAesCBCEncryptPKCS7(t *testing.T) {
	b := []byte("Fufu 中　文加密/解密~&#123a")

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{
			[]byte("0123456789012345"),
			"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA+MVgiOBeHBIwgTXYiYZgUTZ4M7",
		},
		{
			[]byte("012345678901234567891234"),
			"6qxrtFlWv1xcruY1VZSalQ/tD58cCWMD368cIt1Vp5Ega7Jeiz5gMWUgiZNOOElT",
		},
		{
			[]byte("01234567890123456789012345678901"),
			"4drGuwuCHASVh4nta3nml/rBWgkNcUyuTcbB1R5iFfwb6rQnIO47Hsu1fZrc9e1J",
		},
	} {
		actual := AesCBCEncrypt(true, b, v.key)
		AssertEqual(t, v.expected, B64Encode(B2S(actual)))
	}

	// 指定向量加密
	actual := AesCBCEncrypt(true, b, []byte("01234567890123456789012345678901"), []byte("abcdefghabcdefgh"))
	AssertEqual(t, "+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdrGbnwtEDJRdtWYZKF2P0kb", B64Encode(B2S(actual)))
}

func TestAesCBCDecrypt(t *testing.T) {
	expected := "Fufu 中　文加密/解密~&#123a"

	for _, v := range []struct {
		key []byte
		res string
	}{
		{
			[]byte("0123456789012345"),
			"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA+MVghSDLrhLqRTq8nZ5mzrfWC+",
		},
		{
			[]byte("012345678901234567891234"),
			"6qxrtFlWv1xcruY1VZSalQ/tD58cCWMD368cIt1Vp5GMNlv/ER6KB77H2+PqX+33",
		},
		{
			[]byte("01234567890123456789012345678901"),
			"4drGuwuCHASVh4nta3nml/rBWgkNcUyuTcbB1R5iFfyGkEsP6N9HuxAgg0n22vb6",
		},
	} {
		actual := AesCBCDecrypt(false, S2B(B64Decode(v.res)), v.key)
		AssertEqual(t, expected, B2S(actual))
	}

	AssertEqual(t, *new([]byte), AesCBCDecrypt(false, nil, nil))
	AssertEqual(t, *new([]byte), AesCBCDecrypt(false, []byte("1"), nil))
	AssertEqual(t, *new([]byte), AesCBCDecrypt(false, []byte("1"), []byte("2")))

	// 指定向量加密
	actual := AesCBCDecrypt(
		false,
		S2B(B64Decode("+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdpZn4m+l8HNpYkY1A30Y2xh")),
		[]byte("01234567890123456789012345678901"),
		[]byte("abcdefghabcdefgh"),
	)
	AssertEqual(t, expected, B2S(actual))
}

func TestAesCBCDecryptPKCS7(t *testing.T) {
	expected := "Fufu 中　文加密/解密~&#123a"

	for _, v := range []struct {
		key []byte
		res string
	}{
		{
			[]byte("0123456789012345"),
			"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA+MVgiOBeHBIwgTXYiYZgUTZ4M7",
		},
		{
			[]byte("012345678901234567891234"),
			"6qxrtFlWv1xcruY1VZSalQ/tD58cCWMD368cIt1Vp5Ega7Jeiz5gMWUgiZNOOElT",
		},
		{
			[]byte("01234567890123456789012345678901"),
			"4drGuwuCHASVh4nta3nml/rBWgkNcUyuTcbB1R5iFfwb6rQnIO47Hsu1fZrc9e1J",
		},
	} {
		actual := AesCBCDecrypt(true, S2B(B64Decode(v.res)), v.key)
		AssertEqual(t, expected, B2S(actual))
	}

	AssertEqual(t, *new([]byte), AesCBCDecrypt(true, nil, nil))
	AssertEqual(t, *new([]byte), AesCBCDecrypt(true, []byte("1"), nil))
	AssertEqual(t, *new([]byte), AesCBCDecrypt(true, []byte("1"), []byte("2")))

	// 指定向量加密
	actual := AesCBCDecrypt(
		true,
		S2B(B64Decode("+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdrGbnwtEDJRdtWYZKF2P0kb")),
		[]byte("01234567890123456789012345678901"),
		[]byte("abcdefghabcdefgh"),
	)
	AssertEqual(t, expected, B2S(actual))
}
