package xcrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/fufuok/utils"
)

// GenRSAKey 生成 RSA 密钥对
// openssl genrsa -out rsa_private_key.pem 1024
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func GenRSAKey(bits int) (publicKey, privateKey []byte) {
	if bits < 64 {
		bits = 64
	}
	if bits > 8192 {
		bits = 8192
	}
	priv, _ := rsa.GenerateKey(rand.Reader, bits)

	x509PrivateKey := x509.MarshalPKCS1PrivateKey(priv)
	privateKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509PrivateKey,
	})

	pub := priv.PublicKey
	x509PublicKey, _ := x509.MarshalPKIXPublicKey(&pub)
	publicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509PublicKey,
	})

	return
}

// RSAEncrypt 公钥加密
func RSAEncrypt(plaintext, publicKey []byte) ([]byte, error) {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	// 加密明文
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// RSADecrypt 私钥解密
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	priv, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 解密密文
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// RSASign 私钥签名
func RSASign(data, privateKey []byte) ([]byte, error) {
	priv, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	hashed := utils.Sha256(data)

	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed, nil)
}

// RSASignVerify 公钥验证签名
func RSASignVerify(data, publicKey, sig []byte) error {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		return err
	}
	hashed := utils.Sha256(data)

	return rsa.VerifyPSS(pub, crypto.SHA256, hashed, sig, nil)
}

// ParsePrivateKey parses an RSA private key in PKCS #1, ASN.1 DER form.
func ParsePrivateKey(privateKey []byte) (priv *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(privateKey)
	priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)

	return
}

// ParsePublicKey parses a public key in PKIX, ASN.1 DER form.
func ParsePublicKey(publicKey []byte) (pub *rsa.PublicKey, err error) {
	var pubKey interface{}
	block, _ := pem.Decode(publicKey)
	pubKey, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid public key")
	}

	return
}
