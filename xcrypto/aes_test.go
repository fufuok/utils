package xcrypto

import (
	"testing"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/assert"
)

func TestAesCBCEnStringHex(t *testing.T) {
	actual := AesCBCEnStringHex("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnStringHex("", tmpK)
	assert.Equal(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnStringHex(tmpS, tmpK)
	assert.Equal(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608520cbae12ea453abc9d9e66ceb7d60be", actual)
}

func TestAesCBCEnPKCS7StringHex(t *testing.T) {
	actual := AesCBCEnPKCS7StringHex("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7StringHex("", tmpK)
	assert.Equal(t, "3c8a535cc4de40f7cf961e54b4e8661f", actual)

	actual = AesCBCEnPKCS7StringHex(tmpS, tmpK)
	assert.Equal(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c56088e05e1c12308135d889866051367833b", actual)
}

func TestAesCBCEnHex(t *testing.T) {
	actual := AesCBCEnHex([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnHex(nil, tmpK)
	assert.Equal(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnHex([]byte(""), tmpK)
	assert.Equal(t, "5f7df0bf103a8c4ae6faad9906ac3b2a", actual)

	actual = AesCBCEnHex([]byte(tmpS), tmpK)
	assert.Equal(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608520cbae12ea453abc9d9e66ceb7d60be", actual)
}

func TestAesCBCEnPKCS7Hex(t *testing.T) {
	actual := AesCBCEnPKCS7Hex([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7Hex([]byte(""), tmpK)
	assert.Equal(t, "3c8a535cc4de40f7cf961e54b4e8661f", actual)

	actual = AesCBCEnPKCS7Hex([]byte(tmpS), tmpK)
	assert.Equal(t,
		"b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c56088e05e1c12308135d889866051367833b", actual)
}

func TestAesCBCDeStringHex(t *testing.T) {
	actual := AesCBCDeStringHex("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringHex("5f7df0bf103a8c4ae6faad9906ac3b2a", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringHex("b94cd86a309b3391c50f38eb8091563df0f9b329a385ace279574278"+
		"0f8c5608520cbae12ea453abc9d9e66ceb7d60be", tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCDePKCS7StringHex(t *testing.T) {
	actual := AesCBCDePKCS7StringHex("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringHex("3c8a535cc4de40f7cf961e54b4e8661f", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringHex("b94cd86a309b3391c50f38eb8091563df0f9b329a385ace2795742780f8c5608"+
		"8e05e1c12308135d889866051367833b", tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCEnStringB58(t *testing.T) {
	actual := AesCBCEnStringB58("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnStringB58("", tmpK)
	assert.Equal(t, "CnvTeK2SG7tByrMkdSKhT7", actual)

	actual = AesCBCEnStringB58(tmpS, tmpK)
	assert.Equal(t,
		"7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6LrvouTcnhFSnmtghqSCSih", actual)

	actual = AesCBCEnStringB58(tmpS, []byte("d927ad81199aa7dcadfdb4e47b6dc694"))
	assert.Equal(t,
		"7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5", actual)
}

func TestAesCBCEnPKCS7StringB58(t *testing.T) {
	actual := AesCBCEnPKCS7StringB58("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7StringB58("", tmpK)
	assert.Equal(t, "8UbX7eoJQNJWbL1GqeM5LJ", actual)

	actual = AesCBCEnPKCS7StringB58(tmpS, tmpK)
	assert.Equal(t,
		"7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6UGTm6E1tXrpcrtXrg5NhBp", actual)
}

func TestAesCBCEnB58(t *testing.T) {
	actual := AesCBCEnB58([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnB58(nil, tmpK)
	assert.Equal(t, "CnvTeK2SG7tByrMkdSKhT7", actual)

	actual = AesCBCEnB58([]byte(tmpS), tmpK)
	assert.Equal(t,
		"7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6LrvouTcnhFSnmtghqSCSih", actual)

	actual = AesCBCEnB58([]byte(tmpS), []byte("d927ad81199aa7dcadfdb4e47b6dc694"))
	assert.Equal(t,
		"7nWaJ3TEGhKh7RCMJDCHpDrDH741KpMKhRLW2N9Gxb3Tdi48sgkXDHQZBAApVTBHx5", actual)
}

func TestAesCBCEnPKCS7B58(t *testing.T) {
	actual := AesCBCEnPKCS7B58([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7B58(nil, tmpK)
	assert.Equal(t, "8UbX7eoJQNJWbL1GqeM5LJ", actual)

	actual = AesCBCEnPKCS7B58([]byte(tmpS), tmpK)
	assert.Equal(t,
		"7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6UGTm6E1tXrpcrtXrg5NhBp", actual)
}

func TestAesCBCDeStringB58(t *testing.T) {
	actual := AesCBCDeStringB58("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringB58("CnvTeK2SG7tByrMkdSKhT7", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringB58("7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6LrvouTcnhFSnmtghqSCSih",
		tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCDePKCS7StringB58(t *testing.T) {
	actual := AesCBCDePKCS7StringB58("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringB58("8UbX7eoJQNJWbL1GqeM5LJ", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringB58("7oF2FkMyrFgs54hnnoymoAsXJnPgpvXpq8Tr8XFyDdw6UGTm6E1tXrpcrtXrg5NhBp",
		tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCEnStringB64(t *testing.T) {
	actual := AesCBCEnStringB64("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnStringB64("", tmpK)
	assert.Equal(t, "X33wvxA6jErm-q2ZBqw7Kg==", actual)

	actual = AesCBCEnStringB64(tmpS, tmpK)
	assert.Equal(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-", actual)
}

func TestAesCBCEnPKCS7StringB64(t *testing.T) {
	actual := AesCBCEnPKCS7StringB64("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7StringB64("", tmpK)
	assert.Equal(t, "PIpTXMTeQPfPlh5UtOhmHw==", actual)

	actual = AesCBCEnPKCS7StringB64(tmpS, tmpK)
	assert.Equal(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7", actual)
}

func TestAesCBCEnB64(t *testing.T) {
	actual := AesCBCEnB64([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnB64(nil, tmpK)
	assert.Equal(t, "X33wvxA6jErm-q2ZBqw7Kg==", actual)

	actual = AesCBCEnB64([]byte(tmpS), tmpK)
	assert.Equal(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-", actual)
}

func TestAesCBCEnPKCS7B64(t *testing.T) {
	actual := AesCBCEnPKCS7B64([]byte("1"), []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCEnPKCS7B64(nil, tmpK)
	assert.Equal(t, "PIpTXMTeQPfPlh5UtOhmHw==", actual)

	actual = AesCBCEnPKCS7B64([]byte(tmpS), tmpK)
	assert.Equal(t,
		"uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7", actual)
}

func TestAesCBCDeStringB64(t *testing.T) {
	actual := AesCBCDeStringB64("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringB64("X33wvxA6jErm-q2ZBqw7Kg==", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDeStringB64("uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVghSDLrhLqRTq8nZ5mzrfWC-",
		tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCDePKCS7StringB64(t *testing.T) {
	actual := AesCBCDePKCS7StringB64("1", []byte("1"))
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringB64("PIpTXMTeQPfPlh5UtOhmHw==", tmpK)
	assert.Equal(t, "", actual)

	actual = AesCBCDePKCS7StringB64("uUzYajCbM5HFDzjrgJFWPfD5symjhazieVdCeA-MVgiOBeHBIwgTXYiYZgUTZ4M7",
		tmpK)
	assert.Equal(t, tmpS, actual)
}

func TestAesCBCEncrypt(t *testing.T) {
	b := []byte(tmpS)

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{
			tmpK,
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
		assert.Equal(t, v.expected, utils.B64Encode(actual))
	}

	actual := AesCBCEncrypt(false, []byte(""), []byte("abcdefghabcdefgh"))
	assert.Equal(t, "/06zrVSl4UrsshCLDgplgA==", utils.B64Encode(actual))

	actual = AesCBCEncrypt(false, []byte("1"), []byte("abcdefghabcdefgh"))
	assert.Equal(t, "TzlqKoJCo9jWlqZ+gHvhvQ==", utils.B64Encode(actual))

	// 指定向量加密
	actual = AesCBCEncrypt(false, b, []byte("01234567890123456789012345678901"), []byte("abcdefghabcdefgh"))
	assert.Equal(t, "+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdpZn4m+l8HNpYkY1A30Y2xh", utils.B64Encode(actual))
}

func TestAesCBCEncryptPKCS7(t *testing.T) {
	b := []byte(tmpS)

	for _, v := range []struct {
		key      []byte
		expected string
	}{
		{nil, ""},
		{[]byte("1"), ""},
		{[]byte("01234567890123456"), ""},
		{
			tmpK,
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
		assert.Equal(t, v.expected, utils.B64Encode(actual))
	}

	// 指定向量加密
	actual := AesCBCEncrypt(true, b, []byte("01234567890123456789012345678901"), []byte("abcdefghabcdefgh"))
	assert.Equal(t, "+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdrGbnwtEDJRdtWYZKF2P0kb", utils.B64Encode(actual))
}

func TestAesCBCDecrypt(t *testing.T) {
	expected := tmpS

	for _, v := range []struct {
		key []byte
		res string
	}{
		{
			tmpK,
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
		actual := AesCBCDecrypt(false, utils.B64Decode(v.res), v.key)
		assert.Equal(t, expected, utils.B2S(actual))
	}

	assert.Equal(t, *new([]byte), AesCBCDecrypt(false, nil, nil))
	assert.Equal(t, *new([]byte), AesCBCDecrypt(false, []byte("1"), nil))
	assert.Equal(t, *new([]byte), AesCBCDecrypt(false, []byte("1"), []byte("2")))

	// 指定向量加密
	actual := AesCBCDecrypt(
		false,
		utils.B64Decode("+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdpZn4m+l8HNpYkY1A30Y2xh"),
		[]byte("01234567890123456789012345678901"),
		[]byte("abcdefghabcdefgh"),
	)
	assert.Equal(t, expected, utils.B2S(actual))
}

func TestAesCBCDecryptPKCS7(t *testing.T) {
	expected := tmpS

	for _, v := range []struct {
		key []byte
		res string
	}{
		{
			tmpK,
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
		actual := AesCBCDecrypt(true, utils.B64Decode(v.res), v.key)
		assert.Equal(t, expected, utils.B2S(actual))
	}

	assert.Equal(t, *new([]byte), AesCBCDecrypt(true, nil, nil))
	assert.Equal(t, *new([]byte), AesCBCDecrypt(true, []byte("1"), nil))
	assert.Equal(t, *new([]byte), AesCBCDecrypt(true, []byte("1"), []byte("2")))

	// 指定向量加密
	actual := AesCBCDecrypt(
		true,
		utils.B64Decode("+bmCWToTQ/EvZMbphPZ+cImUgpYrf/q78bflV70zIdrGbnwtEDJRdtWYZKF2P0kb"),
		[]byte("01234567890123456789012345678901"),
		[]byte("abcdefghabcdefgh"),
	)
	assert.Equal(t, expected, utils.B2S(actual))
}
