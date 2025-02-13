# Go-Utils

Go 常用的助手函数, 0 依赖, 性能至上.

若有直接引用的, 会在函数定义或目录 README 中注明来源, 并保留 LICENSE, 感谢之!

## 安装

```shell
go get github.com/fufuok/utils
```

## 目录

### 常用函数

<details>
  <summary>DOC</summary>

```go
package utils // import "github.com/fufuok/utils"

const Byte = 1 << (iota * 10) ...
const IByte = 1 ...
const IPv4Min = 0 ...
const PtrSize = 4 << (^uintptr(0) >> 63)
var BigByte ...
var BigSIByte ...
var Rand = NewRand() ...
var ErrInvalidHostPort = errors.New("invalid Host or Port")
var StackTraceBufferSize = 4 << 10
func B2S(b []byte) string
func B64Decode(s string) []byte
func B64Encode(b []byte) string
func B64UrlDecode(s string) []byte
func B64UrlEncode(b []byte) string
func BeginOfDay(t time.Time) time.Time
func BeginOfHour(t time.Time) time.Time
func BeginOfLastMonth(t time.Time) time.Time
func BeginOfLastWeek(t time.Time) time.Time
func BeginOfMinute(t time.Time) time.Time
func BeginOfMonth(t time.Time) time.Time
func BeginOfNextMonth(t time.Time) time.Time
func BeginOfNextWeek(t time.Time) time.Time
func BeginOfSecond(t time.Time) time.Time
func BeginOfTomorrow(t time.Time) time.Time
func BeginOfWeek(t time.Time) time.Time
func BeginOfYear(t time.Time) time.Time
func BeginOfYesterday(t time.Time) time.Time
func BigComma(b *big.Int) string
func BigCommaf(v *big.Float) string
func Bigoom(n, b *big.Int) (float64, int)
func CPUTicks() int64
func CallPath() string
func Comma(v int64) string
func Commaf(v float64) string
func Commai(v int) string
func Commau(v uint64) string
func CopyB2S(b []byte) string
func CopyBytes(b []byte) []byte
func CopyS2B(s string) []byte
func CopyString(s string) string
func CutBytes(s, sep []byte) (before, after []byte, found bool)
func CutString(s, sep string) (before, after string, found bool)
func DaysInMonth(year int, m time.Month) int
func DaysInYear(year int) int
func EncodeUUID(id []byte) []byte
func EndOfDay(t time.Time) time.Time
func EndOfHour(t time.Time) time.Time
func EndOfLastMonth(t time.Time) time.Time
func EndOfLastWeek(t time.Time) time.Time
func EndOfMinute(t time.Time) time.Time
func EndOfMonth(t time.Time) time.Time
func EndOfNextMonth(t time.Time) time.Time
func EndOfNextWeek(t time.Time) time.Time
func EndOfSecond(t time.Time) time.Time
func EndOfTomorrow(t time.Time) time.Time
func EndOfWeek(t time.Time) time.Time
func EndOfYear(t time.Time) time.Time
func EndOfYesterday(t time.Time) time.Time
func EqualFold(b, s string) bool
func EqualFoldBytes(b, s []byte) bool
func Executable(evalSymlinks ...bool) string
func ExecutableDir(evalSymlinks ...bool) string
func FastIntn(n int) int
func FastRand() uint32
func FastRand64() uint64
func FastRandn(n uint32) uint32
func FastRandu() uint
func GetBytes(v interface{}, defaultVal ...[]byte) []byte
func GetIPPort(addr net.Addr) (ip net.IP, port int, err error)
func GetInt(v interface{}, defaultInt ...int) int
func GetMonthDays(t time.Time) int
func GetNotInternalIPv4(ip, defaultIP net.IP, flag ...bool) net.IP
func GetNotInternalIPv4String(ip, defaultIP string, flag ...bool) string
func GetSafeB2S(b []byte, defaultVal ...string) string
func GetSafeBytes(b []byte, defaultVal ...[]byte) []byte
func GetSafeS2B(s string, defaultVal ...[]byte) []byte
func GetSafeString(s string, defaultVal ...string) string
func GetString(v interface{}, defaultVal ...string) string
func GoroutineID() (uint64, error)
func Gzip(data []byte) ([]byte, error)
func GzipLevel(data []byte, level int) (dst []byte, err error)
func HumanBaseBytes(v uint64, base float64, sizes []string) string
func HumanBigBytes(s *big.Int) string
func HumanBigIBytes(s *big.Int) string
func HumanBigKbps(s *big.Int) string
func HumanBytes(v uint64) string
func HumanGBMB(v uint64) string
func HumanIBytes(v uint64) string
func HumanIntBytes(v int) string
func HumanIntIBytes(v int) string
func HumanIntKbps(v int) string
func HumanKbps(v uint64) string
func ID() uint64
func IPv42Long(ip net.IP) int
func IPv42LongLittle(ip net.IP) int
func IPv4String2Long(ip string) int
func IPv4String2LongLittle(ip string) int
func IPv62Int(ip net.IP) *big.Int
func IPv6String2Int(ip string) *big.Int
func InIPNet(ip net.IP, ipNets map[*net.IPNet]struct{}) bool
func InIPNetString(ip string, ipNets map[*net.IPNet]struct{}) bool
func InInts(slice []int, n int) bool
func InStrings(ss []string, s string) bool
func InitCSTLocation() (name string, loc *time.Location, cst *time.Location, ok bool)
func InitLocation(name string) (*time.Location, bool)
func Int2IPv6(ipInt *big.Int) net.IP
func Int2IPv6String(n *big.Int) string
func IsIP(ip string) bool
func IsIPv4(s string) bool
func IsIPv6(s string) bool
func IsInternalIPv4(ip net.IP) bool
func IsInternalIPv4String(ip string) bool
func IsLeapYear(year int) bool
func IsNumeric(s string) bool
func IsPrivateIP(ip net.IP) bool
func IsPrivateIPString(ip string) bool
func JoinBytes(b ...[]byte) []byte
func JoinString(s ...string) string
func JoinStringBytes(s ...string) []byte
func LeftPad(s, pad string, n int) string
func LeftPadBytes(b, pad []byte, n int) []byte
func Logn(n, b float64) float64
func Long2IPv4(n int) net.IP
func Long2IPv4String(n int) string
func LongLittle2IPv4(n int) net.IP
func LongLittle2IPv4String(n int) string
func MaxInt(a, b int) int
func MinInt(a, b int) int
func MustBool(v interface{}) bool
func MustInt(v interface{}) int
func MustJSON(v interface{}) []byte
func MustJSONIndent(v interface{}) []byte
func MustJSONIndentString(v interface{}) string
func MustJSONString(v interface{}) string
func MustParseHumanBigBytes(s string, defaultVal ...*big.Int) *big.Int
func MustParseHumanBytes(s string, defaultVal ...uint64) uint64
func MustString(v interface{}, timeLayout ...string) string
func NanoTime() int64
func NewRand(seed ...int64) *rand.Rand
func Pad(s, pad string, n int) string
func PadBytes(s, pad []byte, n int) []byte
func ParseHostPort(s string) (net.IP, uint16, bool, error)
func ParseHumanBigBytes(s string) (*big.Int, error)
func ParseHumanBytes(s string) (uint64, error)
func ParseIP(s string) (net.IP, bool)
func ParseIPv4(ip string) net.IP
func ParseIPv6(ip string) net.IP
func ParseIPx(s string) (net.IP, bool)
func ParseIPxWithNumeric(s string) (net.IP, bool)
func ParseInts(s string) ([]int, error)
func RandAlphaBytes(n int) []byte
func RandAlphaString(n int) string
func RandBytes(n int) []byte
func RandBytesLetters(n int, letters string) []byte
func RandDecBytes(n int) []byte
func RandDecString(n int) string
func RandHexBytes(n int) []byte
func RandHexString(n int) string
func RandInt(min, max int) int
func RandString(n int) string
func RandUint32(min, max uint32) uint32
func Recover(cb ...RecoveryCallback)
func RemoveString(ss []string, s string) ([]string, bool)
func ReplaceHost(a, b string) string
func Reverse(s string) string
func RightPad(s, pad string, n int) string
func RightPadBytes(b, pad []byte, n int) []byte
func Round(v float64, precision int) float64
func RunPath() string
func RuneReverse(s string) string
func RuneSubString(s string, length int, suffix string) string
func S2B(s string) []byte
func SafeGo(fn func(), cb ...RecoveryCallback)
func SafeGoCommonFunc(args interface{}, fn func(args interface{}), cb ...RecoveryCallback)
func SafeGoWithContext(ctx context.Context, fn func(ctx context.Context), cb ...RecoveryCallback)
func SearchInt(slice []int, n int) int
func SearchString(ss []string, s string) int
func Sleep(ctx context.Context, interval time.Duration) error
func SplitHostPort(hostPort string) (host, port string)
func SumInt(v ...int) int
func ToLower(b string) string
func ToLowerBytes(b []byte) []byte
func ToUpper(b string) string
func ToUpperBytes(b []byte) []byte
func Trim(s string, cutset byte) string
func TrimBytes(b []byte, cutset byte) []byte
func TrimLeft(s string, cutset byte) string
func TrimLeftBytes(b []byte, cutset byte) []byte
func TrimRight(s string, cutset byte) string
func TrimRightBytes(b []byte, cutset byte) []byte
func TrimSlice(ss []string) []string
func UUID() []byte
func UUIDShort() string
func UUIDSimple() string
func UUIDString() string
func Ungzip(data []byte) (src []byte, err error)
func Unzip(data []byte) (src []byte, err error)
func ValidOptionalPort(port string) bool
func WaitNextMinute(t ...time.Time)
func WaitNextMinuteWithTime(t ...time.Time) (now time.Time)
func WaitNextSecond(t ...time.Time)
func WaitNextSecondWithTime(t ...time.Time) (now time.Time)
func WaitSignal(sig ...os.Signal) os.Signal
func WaitUntilMinute(m int, t ...time.Time)
func WaitUntilSecond(s int, t ...time.Time)
func Zip(data []byte) ([]byte, error)
func ZipLevel(data []byte, level int) (dst []byte, err error)
type Bool struct{ ... }
    func NewBool(val bool) *Bool
    func NewFalse() *Bool
    func NewTrue() *Bool
type NoCmp [0]func()
type NoCopy struct{}
type RecoveryCallback func(err interface{}, trace []byte)
type TryMutex struct{ ... }
    func NewTryMutex() *TryMutex
```
</details>

### 泛型方法集

具体使用见各目录下的文档或测试

- [`xslices`](generic/xslices) slices 扩展助手函数.
- [`deepcopy`](generic/deepcopy) DeepCopy makes deep copies of things: unexported field values are not copied.. [@smallnest](https://github.com/smallnest)
- [`atox`](generic/atox): a generic wrapper for the Parse* functions in the strconv package. [@josharian](https://github.com/josharian)
- 更多泛型方法可参考: [zyedidia/generic](https://github.com/zyedidia/generic)

### 有序字典

见: [orderedmap](orderedmap) ([@iancoleman](https://github.com/iancoleman/orderedmap))

A golang data type equivalent to python's collections.OrderedDict

支持 JSON 序列化, 键/值排序.

### 数据类型转换函数集

见: [conv](conv) ([@tidwall](https://github.com/tidwall/conv))

<details>
  <summary>DOC</summary>

```go
package conv // import "github.com/fufuok/utils/conv"

func Atof(a string) float64
func Atoi(a string) int64
func Atot(a string) bool
func Atou(a string) uint64
func Atov(a string) interface{}
func Ftoa(f float64) string
func Ftoi(f float64) int64
func Ftot(f float64) bool
func Ftou(f float64) uint64
func Ftov(f float64) interface{}
func Itoa(i int64) string
func Itof(i int64) float64
func Itot(i int64) bool
func Itou(i int64) uint64
func Itov(i int64) interface{}
func Ttoa(t bool) string
func Ttof(t bool) float64
func Ttoi(t bool) int64
func Ttou(t bool) uint64
func Ttov(t bool) interface{}
func Utoa(u uint64) string
func Utof(u uint64) float64
func Utoi(u uint64) int64
func Utot(u uint64) bool
func Utov(u uint64) interface{}
func Vtoa(v interface{}) string
func Vtof(v interface{}) float64
func Vtoi(v interface{}) int64
func Vtot(v interface{}) bool
func Vtou(v interface{}) uint64
type Value struct{ ... }
    func Any(v interface{}) Value
    func Bool(t bool) Value
    func Byte(x byte) Value
    func Bytes(b []byte) Value
    func CustomBits(x uint64) Value
    func Float32(x float32) Value
    func Float64(f float64) Value
    func Int(x int) Value
    func Int16(x int16) Value
    func Int32(x int32) Value
    func Int64(x int64) Value
    func Int8(x int8) Value
    func Nil() Value
    func String(s string) Value
    func StringWithTag(s string, tag uint16) Value
    func Uint(x uint) Value
    func Uint16(x uint16) Value
    func Uint32(x uint32) Value
    func Uint64(x uint64) Value
    func Uint8(x uint8) Value
```
</details>

### 加解密小工具

见: [envtools](envtools)

### 常用对称加解密函数

见: [xcrypto](xcrypto)

<details>
  <summary>DOC</summary>

```go
package xcrypto // import "github.com/fufuok/utils/xcrypto"

func AesCBCDeB58(s string, key []byte) []byte
func AesCBCDeB64(s string, key []byte) []byte
func AesCBCDeHex(s string, key []byte) []byte
func AesCBCDePKCS7B58(s string, key []byte) []byte
func AesCBCDePKCS7B64(s string, key []byte) []byte
func AesCBCDePKCS7Hex(s string, key []byte) []byte
func AesCBCDePKCS7StringB58(s string, key []byte) string
func AesCBCDePKCS7StringB64(s string, key []byte) string
func AesCBCDePKCS7StringHex(s string, key []byte) string
func AesCBCDeStringB58(s string, key []byte) string
func AesCBCDeStringB64(s string, key []byte) string
func AesCBCDeStringHex(s string, key []byte) string
func AesCBCDecrypt(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) (plaintext []byte)
func AesCBCDecryptE(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) ([]byte, error)
func AesCBCEnB58(b, key []byte) string
func AesCBCEnB64(b, key []byte) string
func AesCBCEnHex(b, key []byte) string
func AesCBCEnPKCS7B58(b, key []byte) string
func AesCBCEnPKCS7B64(b, key []byte) string
func AesCBCEnPKCS7Hex(b, key []byte) string
func AesCBCEnPKCS7StringB58(s string, key []byte) string
func AesCBCEnPKCS7StringB64(s string, key []byte) string
func AesCBCEnPKCS7StringHex(s string, key []byte) string
func AesCBCEnStringB58(s string, key []byte) string
func AesCBCEnStringB64(s string, key []byte) string
func AesCBCEnStringHex(s string, key []byte) string
func AesCBCEncrypt(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) (ciphertext []byte)
func AesCBCEncryptE(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) ([]byte, error)
func AesGCMDeB58(s string, key, nonce []byte) []byte
func AesGCMDeB64(s string, key, nonce []byte) []byte
func AesGCMDeHex(s string, key, nonce []byte) []byte
func AesGCMDeStringB58(s string, key, nonce []byte) string
func AesGCMDeStringB64(s string, key, nonce []byte) string
func AesGCMDeStringHex(s string, key, nonce []byte) string
func AesGCMDecrypt(ciphertext, key, nonce []byte) (plaintext []byte)
func AesGCMDecryptWithNonce(ciphertext, key, nonce, additionalData []byte) ([]byte, error)
func AesGCMEnB58(b, key []byte) (string, []byte)
func AesGCMEnB64(b, key []byte) (string, []byte)
func AesGCMEnHex(b, key []byte) (string, []byte)
func AesGCMEnStringB58(s string, key []byte) (string, []byte)
func AesGCMEnStringB64(s string, key []byte) (string, []byte)
func AesGCMEnStringHex(s string, key []byte) (string, []byte)
func AesGCMEncrypt(plaintext, key []byte) (ciphertext []byte, nonce []byte)
func AesGCMEncryptWithNonce(plaintext, key, nonce, additionalData []byte) ([]byte, []byte, error)
func Decrypt(value, secret string) string
func DesCBCDeB58(s string, key []byte) []byte
func DesCBCDeB64(s string, key []byte) []byte
func DesCBCDeHex(s string, key []byte) []byte
func DesCBCDePKCS7B58(s string, key []byte) []byte
func DesCBCDePKCS7B64(s string, key []byte) []byte
func DesCBCDePKCS7Hex(s string, key []byte) []byte
func DesCBCDePKCS7StringB58(s string, key []byte) string
func DesCBCDePKCS7StringB64(s string, key []byte) string
func DesCBCDePKCS7StringHex(s string, key []byte) string
func DesCBCDeStringB58(s string, key []byte) string
func DesCBCDeStringB64(s string, key []byte) string
func DesCBCDeStringHex(s string, key []byte) string
func DesCBCDecrypt(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) (plaintext []byte)
func DesCBCDecryptE(asPKCS7 bool, ciphertext, key []byte, ivs ...[]byte) ([]byte, error)
func DesCBCEnB58(b, key []byte) string
func DesCBCEnB64(b, key []byte) string
func DesCBCEnHex(b, key []byte) string
func DesCBCEnPKCS7B58(b, key []byte) string
func DesCBCEnPKCS7B64(b, key []byte) string
func DesCBCEnPKCS7Hex(b, key []byte) string
func DesCBCEnPKCS7StringB58(s string, key []byte) string
func DesCBCEnPKCS7StringB64(s string, key []byte) string
func DesCBCEnPKCS7StringHex(s string, key []byte) string
func DesCBCEnStringB58(s string, key []byte) string
func DesCBCEnStringB64(s string, key []byte) string
func DesCBCEnStringHex(s string, key []byte) string
func DesCBCEncrypt(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) (ciphertext []byte)
func DesCBCEncryptE(asPKCS7 bool, plaintext, key []byte, ivs ...[]byte) ([]byte, error)
func Encrypt(value, secret string) string
func GCMDeB58(s string, key []byte) []byte
func GCMDeB64(s string, key []byte) []byte
func GCMDeHex(s string, key []byte) []byte
func GCMDeStringB58(s string, key []byte) string
func GCMDeStringB64(s string, key []byte) string
func GCMDeStringHex(s string, key []byte) string
func GCMDecrypt(encrypted, key []byte) ([]byte, error)
func GCMEnB58(b, key []byte) string
func GCMEnB64(b, key []byte) string
func GCMEnHex(b, key []byte) string
func GCMEnStringB58(s string, key []byte) string
func GCMEnStringB64(s string, key []byte) string
func GCMEnStringHex(s string, key []byte) string
func GCMEncrypt(plaintext, key []byte) ([]byte, error)
func GenRSAKey(bits int) (publicKey, privateKey []byte)
func GetenvDecrypt(key string, secret string) string
func Padding(b []byte, bSize int, pkcs7 bool) []byte
func ParsePrivateKey(privateKey []byte) (priv *rsa.PrivateKey, err error)
func ParsePublicKey(publicKey []byte) (pub *rsa.PublicKey, err error)
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error)
func RSAEncrypt(plaintext, publicKey []byte) ([]byte, error)
func RSASign(data, privateKey []byte) ([]byte, error)
func RSASignVerify(data, publicKey, sig []byte) error
func SetenvEncrypt(key, value, secret string) (string, error)
func UnPadding(b []byte, pkcs7 bool) []byte
func XOR(src, key []byte) []byte
func XORDeB58(s string, key []byte) []byte
func XORDeB64(s string, key []byte) []byte
func XORDeHex(s string, key []byte) []byte
func XORDeStringB58(s string, key []byte) string
func XORDeStringB64(s string, key []byte) string
func XORDeStringHex(s string, key []byte) string
func XORE(src, key []byte) ([]byte, error)
func XOREnB58(b, key []byte) string
func XOREnB64(b, key []byte) string
func XOREnHex(b, key []byte) string
func XOREnStringB58(s string, key []byte) string
func XOREnStringB64(s string, key []byte) string
func XOREnStringHex(s string, key []byte) string
```
</details>

### 获取内外网 IP 小工具

见: [myip](myip)

或: https://github.com/fufuok/myip

<details>
  <summary>DOC</summary>

```go
package myip // import "github.com/fufuok/utils/myip"

func ExternalIP(v ...string) string
func ExternalIPAny(retries ...int) string
func ExternalIPv4() string
func ExternalIPv6() string
func InterfaceAddrs(v ...string) (map[string][]net.IP, error)
func InternalIP(dstAddr, network string) string
func InternalIPv4() string
func InternalIPv6() string
func LocalIP() string
func LocalIPv4s() (ips []string)
```
</details>

### 编码解码 base62

见: [base62](base62) ([@jxskiss](https://github.com/jxskiss/base62))

或: https://github.com/fufuok/basex

<details>
  <summary>DOC</summary>

```go
package base62 // import "github.com/fufuok/utils/base62"

var StdEncoding = NewEncoding(encodeStd)
func AppendInt(dst []byte, num int64) []byte
func AppendUint(dst []byte, num uint64) []byte
func Decode(src []byte) ([]byte, error)
func DecodeString(src string) ([]byte, error)
func DecodeToBuf(dst []byte, src []byte) ([]byte, error)
func Encode(src []byte) []byte
func EncodeToBuf(dst []byte, src []byte) []byte
func EncodeToString(src []byte) string
func FormatInt(num int64) []byte
func FormatUint(num uint64) []byte
func ParseInt(src []byte) (int64, error)
func ParseUint(src []byte) (uint64, error)
type CorruptInputError int64
type Encoding struct{ ... }
    func NewEncoding(encoder string) *Encoding
```
</details>

<details>
  <summary>Usage</summary>

```go
// Basic usage.
Encode(src []byte) []byte
EncodeToString(src []byte) string
Decode(src []byte) ([]byte, error)
DecodeString(src string) ([]byte, error)
FormatInt(num int64) []byte
FormatUint(num uint64) []byte
ParseInt(src []byte) (int64, error)
ParseUint(src []byte) (uint64, error)

// Providing a dst buffer, you may reuse buffers to reduce memory allocation.
EncodeToBuf(dst []byte, src []byte) []byte
DecodeToBuf(dst []byte, src []byte) ([]byte, error)
AppendInt(dst []byte, num int64) []byte
AppendUint(dst []byte, num uint64) []byte

// Or you may use a custom encoding alphabet.
enc := NewEncoding("...my-62-byte-string-alphabet...")
enc.XXX()
```
</details>

### 编码解码 base58

见: [base58](base58)

或: https://github.com/fufuok/basex

<details>
  <summary>DOC</summary>

```go
package base58 // import "github.com/fufuok/utils/base58"

func CheckDecode(input string) (result []byte, version byte, err error)
func CheckEncode(input []byte, version byte) string
func Decode(b string) []byte
func Encode(b []byte) string
```
</details>

### 高性能 Hash 方法及生成器

见: [xhash](xhash)

<details>
  <summary>DOC</summary>

```go
package xhash // import "github.com/fufuok/utils/xhash"

func AddBytes32(h uint32, b []byte) uint32
func AddBytes64(h uint64, b []byte) uint64
func AddString32(h uint32, s string) uint32
func AddString64(h uint64, s string) uint64
func AddUint32(h, u uint32) uint32
func AddUint64(h uint64, u uint64) uint64
func Djb33(s string) uint32
func FnvHash(s string) uint64
func FnvHash32(s string) uint32
func GenHasher[K comparable]() func(K) uintptr
func GenHasher64[K comparable]() func(K) uint64
func GenSeedHasher64[K comparable]() func(maphash.Seed, K) uint64
func Hash(b []byte, h hash.Hash) []byte
func HashBytes(b ...[]byte) string
func HashBytes32(b ...[]byte) uint32
func HashBytes64(b ...[]byte) uint64
func HashSeedString(seed maphash.Seed, s string) uint64
func HashSeedUint64(seed maphash.Seed, n uint64) uint64
func HashString(s ...string) string
func HashString32(s ...string) uint32
func HashString64(s ...string) uint64
func HashUint32(u uint32) uint32
func HashUint64(u uint64) uint64
func Hmac(b []byte, key []byte, h func() hash.Hash) []byte
func HmacSHA1(b, key []byte) []byte
func HmacSHA1Hex(s, key string) string
func HmacSHA256(b, key []byte) []byte
func HmacSHA256Hex(s, key string) string
func HmacSHA512(b, key []byte) []byte
func HmacSHA512Hex(s, key string) string
func MD5(b []byte) []byte
func MD5BytesHex(bs []byte) string
func MD5Hex(s string) string
func MD5Reader(r io.Reader) (string, error)
func MD5Sum(filename string) (string, error)
func MemHash(s string) uint64
func MemHash32(s string) uint32
func MemHashb(b []byte) uint64
func MemHashb32(b []byte) uint32
func MustMD5Sum(filename string) string
func Sha1(b []byte) []byte
func Sha1Hex(s string) string
func Sha256(b []byte) []byte
func Sha256Hex(s string) string
func Sha512(b []byte) []byte
func Sha512Hex(s string) string
func Sum32(s string) uint32
func Sum64(s string) uint64
func SumBytes32(bs []byte) uint32
func SumBytes64(bs []byte) uint64
type Hashable interface{ ... }
```
</details>

### 可排序全局唯一 ID 生成器

比 UUID 更快, 更短

- 要使用 UUIDv4 可以使用 `utils.UUID()`

- 要使用程序运行时自增 ID 可以使用 `utils.ID()`

见: [xid](xid) ([@rs](https://github.com/rs/xid))

或: http://github.com/fufuok/xid

<details>
  <summary>DOC</summary>

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
</details>

### 自守护进程和后台运行

见: [xdaemon](xdaemon) ([@zh-five](https://github.com/zh-five/xdaemon))

或: https://github.com/fufuok/xdaemon

<details>
  <summary>DOC</summary>

```go
package xdaemon // import "github.com/fufuok/utils/xdaemon"

const EnvName = "XW_DAEMON_IDX"
func Background(logFile string, isExit bool) (*exec.Cmd, error)
func NewSysProcAttr() *syscall.SysProcAttr
type Daemon struct{ ... }
    func NewDaemon(logFile string) *Daemon
```
</details>

### 高性能并发安全同步扩展库

见: [xsync](xsync) ([@puzpuzpuz](https://github.com/puzpuzpuz/xsync))

或: https://github.com/fufuok/xsync

**官方版本: `v3.0.0` 已统一了调用方法并内置了 hasher 生成器, 不再需要之前的改动, 直接使用官方原版就好**

<details>
  <summary>DOC</summary>

```go
package xsync // import "github.com/fufuok/utils/xsync"

type Counter struct { ... }
    func (c *Counter) Add(delta int64)
    func (c *Counter) Dec()
    func (c *Counter) Inc()
    func (c *Counter) Reset()
    func (c *Counter) Value() int64

type IntegerConstraint interface {
        ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type MPMCQueue struct { ... }
    func NewMPMCQueue(capacity int) *MPMCQueue
    func (q *MPMCQueue) Dequeue() interface{}
    func (q *MPMCQueue) Enqueue(item interface{})
    func (q *MPMCQueue) TryDequeue() (item interface{}, ok bool)
    func (q *MPMCQueue) TryEnqueue(item interface{}) bool

type Map struct { ... }
    func NewMap() *Map
    func (m *Map) Delete(key string)
    func (m *Map) Load(key string) (value interface{}, ok bool)
    func (m *Map) LoadAndDelete(key string) (value interface{}, loaded bool)
    func (m *Map) LoadAndStore(key string, value interface{}) (actual interface{}, loaded bool)
    func (m *Map) LoadOrCompute(key string, valueFn func() interface{}) (actual interface{}, loaded bool)
    func (m *Map) LoadOrStore(key string, value interface{}) (actual interface{}, loaded bool)
    func (m *Map) Range(f func(key string, value interface{}) bool)
    func (m *Map) Size() int
    func (m *Map) Store(key string, value interface{})

type MapOf[K comparable, V any] struct { ... }
    func NewIntegerMapOf[K IntegerConstraint, V any]() *MapOf[K, V]
    func NewMapOf[V any]() *MapOf[string, V]
    func NewTypedMapOf[K comparable, V any](hasher func(K) uint64) *MapOf[K, V]
    func (m *MapOf[K, V]) Delete(key K)
    func (m *MapOf[K, V]) Load(key K) (value V, ok bool)
    func (m *MapOf[K, V]) LoadAndDelete(key K) (value V, loaded bool)
    func (m *MapOf[K, V]) LoadAndStore(key K, value V) (actual V, loaded bool)
    func (m *MapOf[K, V]) LoadOrCompute(key K, valueFn func() V) (actual V, loaded bool)
    func (m *MapOf[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool)
    func (m *MapOf[K, V]) Range(f func(key K, value V) bool)
    func (m *MapOf[K, V]) Size() int
    func (m *MapOf[K, V]) Store(key K, value V)

type RBMutex struct { ... }
    func (m *RBMutex) Lock()
    func (m *RBMutex) RLock() *RToken
    func (m *RBMutex) RUnlock(t *RToken)
    func (m *RBMutex) Unlock()

type RToken struct { ... }
```
</details>

### 高效的 JSON 字符串操作库集

1. **只有 `1` 次内存分配的 JSON 字符串生成器**

见: [jsongen](xjson/jsongen) ([@darjun](https://github.com/darjun/json-gen))

或: [https://github.com/fufuok/jsongen](https://github.com/fufuok/jsongen)

<details>
  <summary>DOC</summary>

```go
package jsongen // import "github.com/fufuok/utils/xjson/jsongen"

Package jsongen forked from darjun/json-gen

type Array []Value
    func NewArray() *Array
type Map struct{ ... }
    func NewMap() *Map
type QuotedValue string
type RawBytes []byte
type RawString string
type UnquotedValue string
type Value interface{ ... }
    func EscapeString(s string) Value
```
</details>

2. **超高效的 JSON 字符串解析和字段搜索**

见: [gjson](xjson/gjson) ([@tidwall](https://github.com/tidwall/gjson))

<details>
  <summary>DOC</summary>

```go
package gjson // import "github.com/fufuok/utils/xjson/gjson"

Package gjson provides searching for json strings.

var DisableModifiers = false
func AddModifier(name string, fn func(json, arg string) string)
func AppendJSONString(dst []byte, s string) []byte
func ForEachLine(json string, iterator func(line Result) bool)
func ModifierExists(name string, fn func(json, arg string) string) bool
func Valid(json string) bool
func ValidBytes(json []byte) bool
type Result struct{ ... }
    func Get(json, path string) Result
    func GetBytes(json []byte, path string) Result
    func GetMany(json string, path ...string) []Result
    func GetManyBytes(json []byte, path ...string) []Result
    func Parse(json string) Result
    func ParseBytes(json []byte) Result
    func (t Result) Array() []Result
    func (t Result) Bool() bool
    func (t Result) Exists() bool
    func (t Result) Float() float64
    func (t Result) ForEach(iterator func(key, value Result) bool)
    func (t Result) Get(path string) Result
    func (t Result) Int() int64
    func (t Result) IsArray() bool
    func (t Result) IsBool() bool
    func (t Result) IsObject() bool
    func (t Result) Less(token Result, caseSensitive bool) bool
    func (t Result) Map() map[string]Result
    func (t Result) Path(json string) string
    func (t Result) Paths(json string) []string
    func (t Result) String() string
    func (t Result) Time() time.Time
    func (t Result) Uint() uint64
    func (t Result) Value() interface{}
type Type int
    const Null Type = iota ...
```
</details>

3. **JSON 字符串字段修改和删除**

见: [sjson](xjson/sjson) ([@tidwall](https://github.com/tidwall/sjson))

<details>
  <summary>DOC</summary>

```go
package sjson // import "github.com/fufuok/utils/xjson/sjson"

Package sjson provides setting json values.

func Delete(json, path string) (string, error)
func DeleteBytes(json []byte, path string) ([]byte, error)
func Set(json, path string, value interface{}) (string, error)
func SetBytes(json []byte, path string, value interface{}) ([]byte, error)
func SetBytesOptions(json []byte, path string, value interface{}, opts *Options) ([]byte, error)
func SetOptions(json, path string, value interface{}, opts *Options) (string, error)
func SetRaw(json, path, value string) (string, error)
func SetRawBytes(json []byte, path string, value []byte) ([]byte, error)
func SetRawBytesOptions(json []byte, path string, value []byte, opts *Options) ([]byte, error)
func SetRawOptions(json, path, value string, opts *Options) (string, error)
type Options struct{ ... }
```
</details>

4. **JSON 字符串格式化和校验**

见: [pretty](xjson/pretty) ([@tidwall](https://github.com/tidwall/pretty))

<details>
  <summary>DOC</summary>

```go
package pretty // import "github.com/fufuok/utils/xjson/pretty"

var DefaultOptions = &Options{ ... }
func Color(src []byte, style *Style) []byte
func Pretty(json []byte) []byte
func PrettyOptions(json []byte, opts *Options) []byte
func Spec(src []byte) []byte
func SpecInPlace(src []byte) []byte
func Ugly(json []byte) []byte
func UglyInPlace(json []byte) []byte
type Options struct{ ... }
type Style struct{ ... }
    var TerminalStyle *Style
```
</details>

5. 字符串模式匹配(`*?`通配符搜索)

见: [match](xjson/match) ([@tidwall](https://github.com/tidwall/match))

<details>
  <summary>DOC</summary>

```go
package match // import "github.com/fufuok/utils/xjson/match"

Package match provides a simple pattern matcher with unicode support.

func Allowable(pattern string) (min, max string)
func IsPattern(str string) bool
func Match(str, pattern string) bool
func MatchLimit(str, pattern string, maxcomp int) (matched, stopped bool)
```
</details>

### 常用的池

`[]byte` 更多功能的字节切片池化见: [github.com/fufuok/bytespool](https://github.com/fufuok/bytespool)

见: [pools](pools)

<details>
  <summary>DOC</summary>

```go
package bytespool // import "github.com/fufuok/utils/pools/bytespool"

func Append(buf []byte, elems ...byte) []byte
func AppendString(buf []byte, elems string) []byte
func Get(size int) []byte
func Make(size int) []byte
func New(size int) []byte
func Put(buf []byte)
func Release(buf []byte) bool
func SetMaxSize(size int) bool
type CapacityPools struct{ ... }

package bufferpool // import "github.com/fufuok/utils/pools/bufferpool"

func Get() *bytes.Buffer
func New(bs []byte) *bytes.Buffer
func NewByte(c byte) *bytes.Buffer
func NewRune(r rune) *bytes.Buffer
func NewString(s string) *bytes.Buffer
func Put(buf *bytes.Buffer)
func Release(buf *bytes.Buffer) bool
func SetMaxSize(size int) bool

package readerpool // import "github.com/fufuok/utils/pools/readerpool"

func New(b []byte) *bytes.Reader
func Release(r *bytes.Reader)

package timerpool // import "github.com/fufuok/utils/pools/timerpool"

func New(d time.Duration) *time.Timer
func Release(t *time.Timer)

package tickerpool // import "github.com/fufuok/utils/pools/tickerpool"

func New(d time.Duration) *time.Ticker
func Release(t *time.Ticker)
```
</details>

### 并发任务调度库

见: [sched](sched)

或: https://github.com/fufuok/sched

简洁, 高效, 并发限制, 复用 goroutine

<details>
  <summary>DOC</summary>

```go
package sched // import "github.com/fufuok/utils/sched"

type Option func(w *Pool)
    func Queues(limit int) Option
    func Workers(limit int) Option
type Pool struct{ ... }
    func New(opts ...Option) *Pool
```
</details>

### NTP 简单时间同步

独立项目, 见: https://github.com/fufuok/ntp

默认优选 NTP Host, 周期性返回时钟偏移值或当前时间, 也可指定 Host 单次请求

### 简单的日志文件滚动

示例: [xfile/examples](xfile/examples)

带文件写入 Buffer 的日志文件滚动器, 比如按日期滚动生成新文件

<details>
  <summary>DOC</summary>

```go
package ntp // import "github.com/fufuok/utils/ntp"

Package ntp provides an implementation of a Simple NTP (SNTP) client capable
of querying the current time from a remote NTP server. See RFC5905
(https://tools.ietf.org/html/rfc5905) for more details.

This approach grew out of a go-nuts post by Michael Hofmann:
https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/FlcdMU5fkLQ

const LeapNoWarning LeapIndicator = 0 ...
func ClockOffsetChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Duration
func Time(host string) (time.Time, error)
func TimeChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Time
func TimeV(host string, version int) (time.Time, error)
type HostResponse struct{ ... }
    func HostPreferred(hosts []string) *HostResponse
type LeapIndicator uint8
type QueryOptions struct{ ... }
type Response struct{ ... }
    func GetResponse(host string) *Response
    func Query(host string) (*Response, error)
    func QueryWithOptions(host string, opt QueryOptions) (*Response, error)
```
</details>

## 使用示例

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/fufuok/utils"
	"github.com/fufuok/utils/assert"
	"github.com/fufuok/utils/base58"
	"github.com/fufuok/utils/base62"
	"github.com/fufuok/utils/conv"
	"github.com/fufuok/utils/pools/bufferpool"
	"github.com/fufuok/utils/sched"
	"github.com/fufuok/utils/xcrypto"
	"github.com/fufuok/utils/xhash"
	"github.com/fufuok/utils/xid"
	"github.com/fufuok/utils/xjson/jsongen"
)

func main() {
	var s string
	s = utils.GetString(123.45)         // "123.45"
	s = utils.GetString(nil)            // ""
	s = utils.GetString(nil, "default") // "default"
	s = utils.GetString([]byte("ff"))   // "ff"
	s = utils.GetString(true)           // "true"

	s = conv.Any(123.45).String()       // "123.45"
	s = conv.Any(nil).String()          // ""
	s = conv.Any([]byte("ff")).String() // "ff"
	s = conv.Any(true).String()         // "true"

	s = conv.Float64(123.45).String()     // "123.45"
	s = conv.Bytes([]byte("ff")).String() // "ff"

	s = conv.Ftoa(123.45) // "123.45"
	s = conv.Vtoa(nil)    // ""
	s = conv.Ttoa(true)   // "true"

	x := utils.JoinString(s, "OK") // "trueOK"

	b := []byte("trueOK")
	s = utils.B2S(b[0:1])                                     // "t"
	safeS1 := utils.B2S([]byte(s[0:1]))                       // 转换为不可变字符串, CopyString 的实现
	safeS2 := utils.CopyString(s[0:1])                        // 不可变字符串, s 可以被 GC 回收
	safeS3 := utils.GetSafeString(s[0:1], "optional default") // 不可变字符串
	safeS4 := utils.GetSafeB2S(b[0:1], "optional default")    // 转换为不可变字符串
	safeS5 := string(b[0:1])                                  // 标准转换
	b[0] = 70                                                 // 注意: 底层数组变化会引起字符串 s 发生改变
	fmt.Println(s, safeS1, safeS2, safeS3, safeS4, safeS5)    // F t t t t t

	x = xcrypto.Encrypt("myData", "myKey")
	fmt.Println(x) // Csi64LeQLmVhuZTh1xkCKM
	x = xcrypto.Decrypt("Csi64LeQLmVhuZTh1xkCKM", "myKey")
	fmt.Println(x) // myData

	k := utils.RandBytes(16)
	x = xcrypto.AesCBCEnPKCS7StringHex("myData", k)
	fmt.Println(x) // 9dce01c049de7493ce2fae6a2707fad1
	x = xcrypto.AesCBCEnStringB58("myData", k)
	fmt.Println(x) // 6NHTJhaW5mfeUioFLrbpRX

	k = []byte("1234567812345678")
	x = xcrypto.GCMEnStringHex("myData", k)
	fmt.Println(x) // 6501bb4737c772b9a2956d87183d4793d44b0e3c233bf2c6435d70502913d5701a98
	x = xcrypto.GCMEnStringB58("myData", k)
	fmt.Println(x) // 5H9ftJvyBesQS7yNGoW7sBCqpwedpEhkjB7k4Z8QBi8GYJo

	x = xcrypto.GCMEnStringB64("myData", k)
	fmt.Println(xcrypto.GCMDeStringB64(x, k)) // myData

	x = utils.UUIDString()
	fmt.Println(x) // 04a49f17-8c37-44f7-a9c5-ea291c3736d7
	x = utils.UUIDSimple()
	fmt.Println(x) // 16123e98b35a4cea8e9cc127f379ff52
	x = utils.UUIDShort()
	fmt.Println(x) // Mw4hP7t9bnMMczU2AvyorU
	x = xid.NewString()
	fmt.Println(x) // c294bsnn5ek0ub0200fg

	x = base62.EncodeToString([]byte("Test data"))
	fmt.Println(x) // hRXYkBCdzVGV
	bs, _ := base62.DecodeString("hRXYkBCdzVGV")
	x = utils.B2S(bs)
	fmt.Println(x) // Test data

	x = base58.Encode([]byte("Test data"))
	fmt.Println(x) // 25JnwSn7XKfNQ
	x = utils.B2S(base58.Decode("25JnwSn7XKfNQ"))
	fmt.Println(x) // Test data

	whoami := utils.Executable(true)
	pwd := utils.ExecutableDir(true)
	fmt.Println(whoami, pwd)

	fmt.Println(utils.HumanBytes(1234567890))  // 1.2 GB
	fmt.Println(utils.HumanIBytes(1234567890)) // 1.1 GiB
	fmt.Println(utils.HumanKbps(1234567890))   // 1.2 Gbps
	fmt.Println(utils.Commaf(1234567890.123))  // 1,234,567,890.123

	now := time.Date(2020, 2, 18, 12, 13, 14, 123456789, time.UTC)
	fmt.Println(utils.BeginOfLastWeek(now).Format(time.RFC3339Nano)) // 2020-02-10T00:00:00Z

	fmt.Println(utils.GetNotInternalIPv4String("100.125.1.1", "", true))  // 100.125.1.1
	fmt.Println(utils.GetNotInternalIPv4String("100.125.1.1", "1.2.3.4")) // 1.2.3.4
	fmt.Println(utils.GetNotInternalIPv4String("192.168.1.1", "1.2.3.4")) // 1.2.3.4
	fmt.Println(utils.GetNotInternalIPv4String("119.118.7.6", "1.2.3.4")) // 119.118.7.6

	var nilN struct{}
	var nilY *struct{}
	var nilY2 *struct{}
	var nilY3 = (*struct{})(nil)
	fmt.Println(assert.IsNil(nilN), assert.IsNil(nilY))                     // false true
	fmt.Println(assert.IsNil(nilY2), assert.IsNil(nilY3))                   // true true
	fmt.Println(assert.DeepEqual(nil, nilY), assert.DeepEqual(nilY2, nilY)) // false true
	fmt.Println(assert.DeepEqual(nilY3, nilY), nilY3 == nil)                // true true
	fmt.Println(nilY == nil, nilY == nilY2, nilY == nilY3)                  // true true true
	fmt.Println(assert.IsEmpty(nilN), assert.IsEmpty(nilY))                 // true true

	public, private := xcrypto.GenRSAKey(1024)
	fmt.Println(string(public))
	fmt.Println(string(private))

	fmt.Println(utils.IsPrivateIPString("FC00::"))         // true
	fmt.Println(utils.IsPrivateIPString("172.17.0.0"))     // true
	fmt.Println(utils.IsInternalIPv4String("100.125.1.1")) // true

	fmt.Println(utils.ToLower("TesT"))                                             // test
	fmt.Println(utils.EqualFold(utils.Trim("/TesT/", '/'), utils.ToUpper("Test"))) // true

	host, port := utils.SplitHostPort("demo.com:77")
	fmt.Println(host) // demo.com
	fmt.Println(port) // 77

	fmt.Println(utils.Rand.Intn(10), utils.FastIntn(10))

	dec, _ := utils.Zip(utils.RandBytes(3000))
	src, _ := utils.Unzip(dec)
	fmt.Println(len(dec), len(src)) // 2288 3000

	type T struct {
		Name string `json:"name"`
	}
	t1 := T{"ff"}
	buf := bufferpool.Get()
	_ = json.NewEncoder(buf).Encode(&t1)
	fmt.Println("json:", buf.String()) // json: {"name":"ff"}
	bufferpool.Put(buf)

	var t2 T
	buf = bufferpool.Get()
	buf.WriteString(`{"name":"ff"}`)
	_ = json.NewDecoder(buf).Decode(&t2)
	fmt.Printf("struct: %+v\n", t2)       // struct: {Name:ff}
	fmt.Println("empty:", buf.Len() == 0) // empty: true
	bufferpool.Put(buf)

	js := jsongen.NewMap()
	js.PutString("s", `a"b"\c`)
	js.PutFloat("f", 3.14)
	js.PutBool("b", false)
	jsArr := jsongen.NewArray()
	jsArr.AppendInt(7)
	jsArr.AppendStringArray([]string{"A", "B"})
	js.PutArray("sub", jsArr)
	jsBytes := js.Serialize(nil)
	fmt.Printf("%s\n", jsBytes) // {"s":"a\"b\"\\c","f":3.14,"b":false,"sub":[7,["A","B"]]}

	fmt.Println(utils.ID(), utils.ID()) // 1, 2

	fmt.Println(utils.CutString("test@fufuok.com", "@")) // test fufuok.com true

	fmt.Println(xhash.Sum64("test"))  // 18007334074686647077
	fmt.Println(xhash.MD5Hex("test")) // 098f6bcd4621d373cade4e832627b4f6

	lock := utils.NewTryMutex()
	lock.Lock()
	ok := lock.TryLock(20 * time.Millisecond)
	fmt.Println(ok) // false
	lock.Unlock()
	ok = lock.TryLock(20 * time.Millisecond)
	fmt.Println(ok) // true
	lock.Unlock()

	// 等待下一秒 0 毫秒 (近似)
	now = utils.WaitNextSecondWithTime()
	fmt.Println("hour:minute:second.00*ms", now)

	count := int64(0)
	bus := sched.New() // 默认并发数: runtime.NumCPU()
	for i := 0; i < 30; i++ {
		bus.Add(1)
		bus.RunWithArgs(func(n ...interface{}) {
			atomic.AddInt64(&count, int64(n[0].(int)))
		}, i)
	}
	bus.Wait()
	fmt.Println("count:", atomic.LoadInt64(&count)) // count: 435

	// 继续下一批任务
	bus.Add(1)
	bus.Run(func() {
		fmt.Println("is running:", bus.IsRunning(), bus.Running()) // is running: true 1
	})
	bus.Wait()
	bus.Release()

	// 指定并发数
	bus = sched.New(sched.Workers(2))
	bus.Add(5)
	for i := 0; i < 5; i++ {
		bus.Run(func() {
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		})
	}
	bus.WaitAndRelease()
	fmt.Println("is running:", bus.IsRunning()) // is running: false

	// 原子操作的安全布尔值 (与 go1.19 一致)
	var atomicBool utils.Bool
	atomicBool.StoreTrue()
	fmt.Println("is running:", atomicBool.Load()) // is running: true
	atomicBool.Toggle()
	fmt.Println("is running:", atomicBool.String()) // is running: false

	woo := func() {
		panic("woo...")
	}
	utils.SafeGo(woo)
	utils.SafeGo(woo, func(err interface{}, trace []byte) {
		fmt.Printf("SafeGo, ERR: %v, Traceback: \n%s\n", err, trace)
	})
	fmt.Println("Ctrl+c to exit")
	utils.WaitSignal()
	fmt.Println("Exited.")
}
```

## 🔑 JetBrains OS licenses

The `utils` package was developed using `GoLand` IDE under the ***Free JetBrains Open Source License*** granted by JetBrains s.r.o. I would like to take this opportunity to express my gratitude.

<a href="https://www.jetbrains.com/?from=FF.PyAdmin" target="_blank"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jetbrains.svg" alt="JetBrains logo."></a>




*ff*