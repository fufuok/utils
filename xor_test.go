package utils

import (
	"bytes"
	"testing"
)

var (
	tmpS = "Fufu 中　文加密/解密~&#123a"
	tmpK = []byte("0123456789012345")
	tmpB = []byte("Fufu 中　文\u2728->?\n*\U0001F63A0123456789012345")
)

func TestXOREnStringHex(t *testing.T) {
	actual := XOREnStringHex("1", []byte("1"))
	AssertEqual(t, "51", actual)

	res := XOREnStringHex("", tmpK)
	AssertEqual(t, "", XORDeStringHex(res, tmpK))

	res = XOREnStringHex(tmpS, tmpK)
	AssertEqual(t, tmpS, XORDeStringHex(res, tmpK))
}

func TestXOREnStringB64(t *testing.T) {
	actual := XOREnStringB64("1", []byte("1"))
	AssertEqual(t, "UQ==", actual)

	res := XOREnStringB64("", tmpK)
	AssertEqual(t, "", XORDeStringB64(res, tmpK))

	res = XOREnStringB64(tmpS, tmpK)
	AssertEqual(t, tmpS, XORDeStringB64(res, tmpK))
}

func TestXOREnStringB58(t *testing.T) {
	actual := XOREnStringB58("1", []byte("1"))
	AssertEqual(t, "2Q", actual)

	res := XOREnStringB58("", tmpK)
	AssertEqual(t, "", XORDeStringB58(res, tmpK))

	res = XOREnStringB58(tmpS, tmpK)
	AssertEqual(t, tmpS, XORDeStringB58(res, tmpK))
}

func BenchmarkEnDeXOR(b *testing.B) {
	src := bytes.Repeat(tmpB, 10)
	key := tmpK
	// res := XOR(src, key)
	// fmt.Println(string(XOR(res, key)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := XOR(src, key)
		XOR(res, key)
	}
}

func BenchmarkEnDeAesCBC(b *testing.B) {
	src := bytes.Repeat(tmpB, 10)
	key := tmpK
	// res := AesCBCEncrypt(false, src, key)
	// fmt.Println(string(AesCBCDecrypt(false, res, key)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := AesCBCEncrypt(false, src, key)
		AesCBCDecrypt(false, res, key)
	}
}

func BenchmarkEnDeGCM(b *testing.B) {
	src := bytes.Repeat(tmpB, 10)
	key := tmpK
	// res, nonce := AesGCMEncrypt(src, key)
	// fmt.Println(string(AesGCMDecrypt(res, key, nonce)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, nonce := AesGCMEncrypt(src, key)
		AesGCMDecrypt(res, key, nonce)
	}
}

func BenchmarkEnDeDesCBC(b *testing.B) {
	src := bytes.Repeat(tmpB, 10)
	key := []byte("12345678")
	// res := DesCBCEncrypt(false, src, key)
	// fmt.Println(string(DesCBCDecrypt(false, res, key)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := DesCBCEncrypt(false, src, key)
		DesCBCDecrypt(false, res, key)
	}
}

func BenchmarkEnDeAesCBCB64(b *testing.B) {
	src := bytes.Repeat(tmpB, 10)
	key := tmpK
	// res := AesCBCEnB64(src, key)
	// fmt.Println(string(AesCBCDeB64(res, key)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := AesCBCEnB64(src, key)
		AesCBCDeB64(res, key)
	}
}

// BenchmarkEnDeXOR-12               209040              5701 ns/op            3200 B/op          4 allocs/op
// BenchmarkEnDeXOR-12               211089              5716 ns/op            3200 B/op          4 allocs/op
// BenchmarkEnDeAesCBC-12            601584              1933 ns/op            2928 B/op         18 allocs/op
// BenchmarkEnDeAesCBC-12            601578              1932 ns/op            2928 B/op         18 allocs/op
// BenchmarkEnDeGCM-12               925539              1259 ns/op            2448 B/op         13 allocs/op
// BenchmarkEnDeGCM-12               925489              1261 ns/op            2448 B/op         13 allocs/op
// BenchmarkEnDeDesCBC-12             72048             16426 ns/op            2248 B/op         12 allocs/op
// BenchmarkEnDeDesCBC-12             73365             16461 ns/op            2248 B/op         12 allocs/op
// BenchmarkEnDeAesCBCB64-12         376006              3225 ns/op            5104 B/op         22 allocs/op
// BenchmarkEnDeAesCBCB64-12         353877              3219 ns/op            5104 B/op         22 allocs/op
