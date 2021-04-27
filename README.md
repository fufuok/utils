# Go-Utils

常用的助手函数

若有直接引用的, 会在函数定义或目录 README 中注明来源, 保留 LICENSE, 感谢之!

## 安装

```shell
go get github.com/fufuok/utils
```

## 目录

```go
package utils // import "github.com/fufuok/utils"

func AddString(s ...string) string
func AddStringBytes(s ...string) []byte
func AesCBCDeB58(s, key string) []byte
func AesCBCDeB64(s, key string) []byte
func AesCBCDeHex(s, key string) []byte
func AesCBCDePKCS7B58(s, key string) []byte
func AesCBCDePKCS7B64(s, key string) []byte
func AesCBCDePKCS7Hex(s, key string) []byte
func AesCBCDePKCS7StringB58(s, key string) string
func AesCBCDePKCS7StringB64(s, key string) string
func AesCBCDePKCS7StringHex(s, key string) string
func AesCBCDeStringB58(s, key string) string
func AesCBCDeStringB64(s, key string) string
func AesCBCDeStringHex(s, key string) string
func AesCBCDecrypt(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) (plaintext []byte)
func AesCBCDecryptE(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) ([]byte, error)
func AesCBCEnB58(b, key []byte) string
func AesCBCEnB64(b, key []byte) string
func AesCBCEnHex(b, key []byte) string
func AesCBCEnPKCS7B58(b, key []byte) string
func AesCBCEnPKCS7B64(b, key []byte) string
func AesCBCEnPKCS7Hex(b, key []byte) string
func AesCBCEnPKCS7StringB58(s, key string) string
func AesCBCEnPKCS7StringB64(s, key string) string
func AesCBCEnPKCS7StringHex(s, key string) string
func AesCBCEnStringB58(s, key string) string
func AesCBCEnStringB64(s, key string) string
func AesCBCEnStringHex(s, key string) string
func AesCBCEncrypt(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) (ciphertext []byte)
func AesCBCEncryptE(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) ([]byte, error)
func AesGCMDeB58(s, key, nonce string) []byte
func AesGCMDeB64(s, key, nonce string) []byte
func AesGCMDeHex(s, key, nonce string) []byte
func AesGCMDeStringB58(s, key, nonce string) string
func AesGCMDeStringB64(s, key, nonce string) string
func AesGCMDeStringHex(s, key, nonce string) string
func AesGCMDecrypt(ciphertext, key, nonce []byte) (plaintext []byte)
func AesGCMDecryptWithNonce(ciphertext, key, nonce, additionalData []byte) ([]byte, error)
func AesGCMEnB58(b, key []byte) (string, string)
func AesGCMEnB64(b, key []byte) (string, string)
func AesGCMEnHex(b, key []byte) (string, string)
func AesGCMEnStringB58(s, key string) (string, string)
func AesGCMEnStringB64(s, key string) (string, string)
func AesGCMEnStringHex(s, key string) (string, string)
func AesGCMEncrypt(plaintext, key []byte) (ciphertext []byte, nonce []byte)
func AesGCMEncryptWithNonce(plaintext, key, nonce, additionalData []byte) ([]byte, []byte, error)
func AssertEqual(t testing.TB, expected, actual interface{}, description ...string)
func AssertPanics(t *testing.T, title string, f func())
func B2S(b []byte) string
func B64Decode(s string) []byte
func B64Encode(b []byte) string
func B64UrlDecode(s string) []byte
func B64UrlEncode(b []byte) string
func CallPath() string
func CopyB2S(b []byte) string
func CopyBytes(b []byte) []byte
func CopyS2B(s string) []byte
func CopyString(s string) string
func Decrypt(value, secret string) string
func DesCBCDeB58(s, key string) []byte
func DesCBCDeB64(s, key string) []byte
func DesCBCDeHex(s, key string) []byte
func DesCBCDePKCS7B58(s, key string) []byte
func DesCBCDePKCS7B64(s, key string) []byte
func DesCBCDePKCS7Hex(s, key string) []byte
func DesCBCDePKCS7StringB58(s, key string) string
func DesCBCDePKCS7StringB64(s, key string) string
func DesCBCDePKCS7StringHex(s, key string) string
func DesCBCDeStringB58(s, key string) string
func DesCBCDeStringB64(s, key string) string
func DesCBCDeStringHex(s, key string) string
func DesCBCDecrypt(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) (plaintext []byte)
func DesCBCDecryptE(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) ([]byte, error)
func DesCBCEnB58(b, key []byte) string
func DesCBCEnB64(b, key []byte) string
func DesCBCEnHex(b, key []byte) string
func DesCBCEnPKCS7B58(b, key []byte) string
func DesCBCEnPKCS7B64(b, key []byte) string
func DesCBCEnPKCS7Hex(b, key []byte) string
func DesCBCEnPKCS7StringB58(s, key string) string
func DesCBCEnPKCS7StringB64(s, key string) string
func DesCBCEnPKCS7StringHex(s, key string) string
func DesCBCEnStringB58(s, key string) string
func DesCBCEnStringB64(s, key string) string
func DesCBCEnStringHex(s, key string) string
func DesCBCEncrypt(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) (ciphertext []byte)
func DesCBCEncryptE(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) ([]byte, error)
func EncodeUUID(id []byte) []byte
func Encrypt(value, secret string) string
func Executable(evalSymlinks ...bool) string
func ExecutableDir(evalSymlinks ...bool) string
func GetBytes(v interface{}, defaultVal ...[]byte) []byte
func GetInt(v interface{}, defaultInt ...int) int
func GetSafeB2S(b []byte, defaultVal ...string) string
func GetSafeBytes(b []byte, defaultVal ...[]byte) []byte
func GetSafeS2B(s string, defaultVal ...[]byte) []byte
func GetSafeString(s string, defaultVal ...string) string
func GetString(v interface{}, defaultVal ...string) string
func GetenvDecrypt(key string, secret string) string
func Hash(b []byte, h hash.Hash) []byte
func Hmac(b []byte, key []byte, h func() hash.Hash) []byte
func HmacSHA1(b, key []byte) []byte
func HmacSHA1Hex(s, key string) string
func HmacSHA256(b, key []byte) []byte
func HmacSHA256Hex(s, key string) string
func HmacSHA512(b, key []byte) []byte
func HmacSHA512Hex(s, key string) string
func InInts(slice []int, n int) bool
func InStrings(ss []string, s string) bool
func IsDir(s string) bool
func IsExist(s string) bool
func IsFile(s string) bool
func JoinBytes(b ...[]byte) []byte
func MD5(b []byte) []byte
func MD5Hex(s string) string
func MD5Reader(r io.Reader) (string, error)
func MD5Sum(filename string) (string, error)
func MaxInt(a, b int) int
func MinInt(a, b int) int
func MustBool(v interface{}) bool
func MustInt(v interface{}) int
func MustJSON(v interface{}) []byte
func MustJSONString(v interface{}) string
func MustString(v interface{}) string
func Padding(b []byte, bSize int, pkcs7 bool) []byte
func RandBytes(n int) []byte
func RandHex(nHalf int) string
func RandInt(min int, max int) int
func RandString(n int) string
func Round(val float64, precision int) float64
func RunPath() string
func S2B(s string) []byte
func SearchInt(slice []int, n int) int
func SearchString(ss []string, s string) int
func SetenvEncrypt(key, value, secret string) (string, error)
func Sha1(b []byte) []byte
func Sha1Hex(s string) string
func Sha256(b []byte) []byte
func Sha256Hex(s string) string
func Sha512(b []byte) []byte
func Sha512Hex(s string) string
func Str2Bytes(s string) (b []byte)
func String2Bytes(s string) (b []byte)
func StringToBytes(s string) (b []byte)
func UUID() []byte
func UUIDShort() string
func UUIDSimple() string
func UUIDString() string
func UnPadding(b []byte, pkcs7 bool) []byte
func WaitNextMinute()
func WeightedChoiceIndex(choices []TChoice) int
func WeightedChoiceMap(choices map[interface{}]int) interface{}
func WeightedChoiceWeightsIndex(weights []int) int
type TChoice struct{ ... }
    func WeightedChoice(choices ...TChoice) TChoice
```

### 加解密小工具

见: `envtools`

### 获取内外网 IP 小工具

见: `myip`

或: https://github.com/fufuok/myip

```go
package myip // import "github.com/fufuok/utils/myip"

func ExternalIP(v ...string) string
func ExternalIPv4() string
func ExternalIPv6() string
func InternalIP(dstAddr, network string) string
func InternalIPv4() string
func InternalIPv6() string
func LocalIP() string
```

### 编码解码 base58

见: `base58`

或: https://github.com/fufuok/basex

```go
package base58 // import "github.com/fufuok/utils/base58"

func CheckDecode(input string) (result []byte, version byte, err error)
func CheckEncode(input []byte, version byte) string
func Decode(b string) []byte
func Encode(b []byte) string
```

### 可排序全局唯一ID生成器

比 UUID 更快, 更短

见: `xid`

或: http://github.com/fufuok/xid

```go
package xid // import "github.com/fufuok/utils/xid"

func NewBytes() []byte
func NewString() string
func Sort(ids []ID)
type ID [rawLen]byte
    func FromBytes(b []byte) (ID, error)
    func FromString(id string) (ID, error)
    func New() ID
    func NewWithTime(t time.Time) ID
    func NilID() ID
```

## JSON

`json` 使用 `gin` 类似的可选组织方式:

- `go build .` 默认使用 `json-iterator/go`
- `go build -tags=gojson.` 使用标准 JSON 库 `encoding/json`
- `go build -tags=go_json .` 使用 `goccy/go-json` (暂不成熟, 观望中)


## 使用

```go
var s string
s = utils.GetString(123.45)         // "123.45"
s = utils.GetString(nil)            // ""
s = utils.GetString(nil, "default") // "default"
s = utils.GetString([]byte("ff"))   // "ff"
s = utils.GetString(true)           // "true"

x := utils.AddString(s, "OK") // "trueOK"

b = []byte("trueOK")
s = utils.B2S(b[0:1])                                     // "t"
safeS1 := utils.B2S([]byte(s[0:1]))                       // 转换为不可变字符串, CopyString 的实现
safeS2 := utils.CopyString(s[0:1])                        // 不可变字符串, s 可以被 GC 回收
safeS3 := utils.GetSafeString(s[0:1], "optional default") // 不可变字符串
safeS4 := utils.GetSafeB2S(b[0:1], "optional default")    // 转换为不可变字符串
safeS5 := string(b[0:1])                                  // 标准转换
b[0] = 70                                                 // 注意: 底层数组变化会引起字符串 s 发生改变
fmt.Println(s, safeS1, safeS2, safeS3, safeS4, safeS5)    // F t t t t t

x = utils.AesCBCEnPKCS7StringHex("myData", "myKey")
fmt.Println(x)

x = utils.UUIDString()
fmt.Println(x)  // 04a49f17-8c37-44f7-a9c5-ea291c3736d7
x = utils.UUIDSimple()
fmt.Println(x)  // 16123e98b35a4cea8e9cc127f379ff52
x = utils.UUIDShort()
fmt.Println(x)  // Mw4hP7t9bnMMczU2AvyorU

x = base58.Encode([]byte("Test data"))
fmt.Println(x)  // 25JnwSn7XKfNQ
x = utils.B2S(base58.Decode("25JnwSn7XKfNQ"))
fmt.Println(x)  // Test data

choice :=utils.WeightedChoice([]utils.TChoice{
    {"A", 5},
    {"B", 3},
    {"C", 2},
    {"D", 0},
}...)
fmt.Println(choice.String())  // {"Item":"B","Weight":3}

items := []interface{}{"Item.1", "Item.2", "Item.3", "Item.4"}
weights := []int{1, 2, 3, 100}
idx := utils.WeightedChoiceWeightsIndex(weights)
fmt.Println(items[idx])  // Item.4

itemMap := map[interface{}]int{"Item.1": 1, "Item.2": 2, "Item.3": 3, "Item.4": 100}
item := utils.WeightedChoiceMap(itemMap)
fmt.Println(item)  // Item.4

whoami := utils.Executable(true)
whereami := utils.ExecutableDir(true)
```

...





*ff*