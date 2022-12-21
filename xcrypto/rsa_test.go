package xcrypto

import (
	"testing"

	"github.com/fufuok/utils/assert"
)

func TestRSABase(t *testing.T) {
	pub, priv := GenRSAKey(2048)
	encrypted, err := RSAEncrypt(tmpB, pub)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, _ := RSADecrypt(encrypted, priv)
	assert.Equal(t, tmpB, decrypted)

	sig, err := RSASign(tmpB, priv)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nil, RSASignVerify(tmpB, pub, sig))

	privateKey := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDJIaZyj/Ikk8zEK6SkoPV2UXkhv7LkXuIyAnWivREd88WNaoI/
gn749FaFVEu8uqZCnZoJbN/UkTi8pzEfJaPgMxx3/zPR4/JHpRAbsNUIBScYRU4s
holUJp3zj4nSJ23Vv3js4vKlz+EUF1HuUvzGCHvamPPpoW3klbQf0q83CwIDAQAB
AoGASy/UHBFpcHj72/2An7Y37FIKmS4jPrGllxJKTVqmMm81U2cUZ49tzbzxNwhL
A7M2gcKOmaVkiv76mbNabo1Qk84mhNr9Z4KytTnz9ticpfk7zY2CRfyHOWeNleUM
Ji9pmmBRJG5sWG4ZAH0IV5hAJOktd8qEaSPbbg9LAuhtfzECQQDyApjlXP/VlXyS
mRdWvJfvoA9LHM10MPOXqzQzec70FrLyLeMXg35cYlMCGkDZ7Tl24u8NGNsV2IU/
pSKBKUmzAkEA1MIbW3rG5SMlCNwjVBJJlvdztvlAvTYyqA7qF1DBLOKF1O/JRizd
Vt6kbiTJ+65Y1z/Gy8BvcJYy4yVMAiKBSQJBAKk5/bYtAMxWOoS8PmCtgcTS9L6+
RkBgVoWQ9vCj1X5DPSAxzCFOFpb9PjQzLXP1+P/UEfrjjZdKD2sAyw7sUxcCQQCv
qWwpBZ/+RBwpyogou8iiqsCRjA5Vqs/8TgQdKAG263iQLULDe/tr4/tjLWDPOk4D
upaKV+Iq1PhC7uJoyNBxAkBXBlLAi2OA+p6UNzlJ1I8Fz3NSYkyEB7FBkHhKM5P+
CSZCbtGXzOwSx9syHTrBk79KPrTMcHeFmJXtBvskCJcn
-----END RSA PRIVATE KEY-----`
	publicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJIaZyj/Ikk8zEK6SkoPV2UXkh
v7LkXuIyAnWivREd88WNaoI/gn749FaFVEu8uqZCnZoJbN/UkTi8pzEfJaPgMxx3
/zPR4/JHpRAbsNUIBScYRU4sholUJp3zj4nSJ23Vv3js4vKlz+EUF1HuUvzGCHva
mPPpoW3klbQf0q83CwIDAQAB
-----END PUBLIC KEY-----`
	chipertext, err := RSAEncrypt(tmpB, []byte(publicKey))
	if err != nil {
		t.Fatal(err)
	}
	plaintext, _ := RSADecrypt(chipertext, []byte(privateKey))
	assert.Equal(t, tmpB, plaintext)
}
