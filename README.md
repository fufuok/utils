# Go-Utils

Go 常用的助手函数, 性能至上.

若有直接引用的, 会在函数定义或目录 README 中注明来源, 并保留 LICENSE, 感谢之!

## 安装

```shell
go get github.com/fufuok/utils
```

## 目录

```go
package utils // import "github.com/fufuok/utils"

const Byte = 1 << (iota * 10) ...
const IByte = 1 ...
var BigByte ...
var BigSIByte ...
var Rand = NewRand()
var Seed = FastRand()
func AddString(s ...string) string
func AddStringBytes(s ...string) []byte
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
func AssertEqual(t testing.TB, expected, actual interface{}, description ...string)
func AssertPanics(t *testing.T, title string, f func())
func B2S(b []byte) string
func B64Decode(s string) []byte
func B64Encode(b []byte) string
func B64UrlDecode(s string) []byte
func B64UrlEncode(b []byte) string
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
func Djb33(s string) uint32
func EncodeUUID(id []byte) []byte
func Encrypt(value, secret string) string
func EqualFold(b, s string) (equals bool)
func EqualFoldBytes(b, s []byte) (equals bool)
func Executable(evalSymlinks ...bool) string
func ExecutableDir(evalSymlinks ...bool) string
func FastIntn(n int) int
func FastRand() uint32
func FastRandBytes(n int) []byte
func FastRandn(n uint32) uint32
func FnvHash(s string) uint64
func FnvHash32(s string) uint32
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
func Get0Hour(t time.Time) time.Time
func Get0LastMonth(t time.Time) time.Time
func Get0LastWeek(t time.Time) time.Time
func Get0Minute(t time.Time) time.Time
func Get0Month(t time.Time) time.Time
func Get0NextMonth(t time.Time) time.Time
func Get0NextWeek(t time.Time) time.Time
func Get0Second(t time.Time) time.Time
func Get0Tomorrow(t time.Time) time.Time
func Get0Week(t time.Time) time.Time
func Get0Yesterday(t time.Time) time.Time
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
func GetenvDecrypt(key string, secret string) string
func Hash(b []byte, h hash.Hash) []byte
func Hmac(b []byte, key []byte, h func() hash.Hash) []byte
func HmacSHA1(b, key []byte) []byte
func HmacSHA1Hex(s, key string) string
func HmacSHA256(b, key []byte) []byte
func HmacSHA256Hex(s, key string) string
func HmacSHA512(b, key []byte) []byte
func HmacSHA512Hex(s, key string) string
func HumanBaseBytes(v uint64, base float64, sizes []string) string
func HumanBigBytes(s *big.Int) string
func HumanBigIBytes(s *big.Int) string
func HumanBigKbps(s *big.Int) string
func HumanBytes(v uint64) string
func HumanIBytes(v uint64) string
func HumanIntBytes(v int) string
func HumanIntIBytes(v int) string
func HumanIntKbps(v int) string
func HumanKbps(v uint64) string
func IPv42Long(ip net.IP) int
func IPv4String2Long(ip string) int
func InIPNet(ip net.IP, ipNets map[*net.IPNet]struct{}) bool
func InIPNetString(ip string, ipNets map[*net.IPNet]struct{}) bool
func InInts(slice []int, n int) bool
func InStrings(ss []string, s string) bool
func IsDir(s string) bool
func IsExist(s string) bool
func IsFile(s string) bool
func IsInternalIPv4(ip net.IP) bool
func IsInternalIPv4String(ip string) bool
func IsNil(i interface{}) bool
func IsPrivateIP(ip net.IP) bool
func IsPrivateIPString(ip string) bool
func JoinBytes(b ...[]byte) []byte
func Logn(n, b float64) float64
func Long2IPv4(n int) net.IP
func Long2IPv4String(n int) string
func MD5(b []byte) []byte
func MD5Hex(s string) string
func MD5Reader(r io.Reader) (string, error)
func MD5Sum(filename string) (string, error)
func MaxInt(a, b int) int
func MemHash(s string) uint64
func MemHash32(s string) uint32
func MemHashb(b []byte) uint64
func MemHashb32(b []byte) uint32
func MinInt(a, b int) int
func MustBool(v interface{}) bool
func MustInt(v interface{}) int
func MustJSON(v interface{}) []byte
func MustJSONIndent(v interface{}) []byte
func MustJSONIndentString(v interface{}) string
func MustJSONString(v interface{}) string
func MustString(v interface{}) string
func NanoTime() int64
func NewRand(seed ...int64) *rand.Rand
func Padding(b []byte, bSize int, pkcs7 bool) []byte
func ParseHumanBigBytes(s string) (*big.Int, error)
func ParseHumanBytes(s string) (uint64, error)
func ParsePrivateKey(privateKey []byte) (priv *rsa.PrivateKey, err error)
func ParsePublicKey(publicKey []byte) (pub *rsa.PublicKey, err error)
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error)
func RSAEncrypt(plaintext, publicKey []byte) ([]byte, error)
func RSASign(data, privateKey []byte) ([]byte, error)
func RSASignVerify(data, publicKey, sig []byte) error
func RandBytes(n int) []byte
func RandHex(nHalf int) string
func RandInt(min, max int) int
func RandString(n int) string
func RandUint32(min, max uint32) uint32
func RemoveString(ss []string, s string) ([]string, bool)
func ReplaceHost(a, b string) string
func Round(v float64, precision int) float64
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
func SplitHostPort(hostPort string) (host, port string)
func Str2Bytes(s string) (b []byte)
func StrToBytes(s string) []byte
func String2Bytes(s string) (b []byte)
func StringToBytes(s string) (b []byte)
func Sum32(s string) uint32
func Sum64(s string) uint64
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
func UUID() []byte
func UUIDShort() string
func UUIDSimple() string
func UUIDString() string
func UnPadding(b []byte, pkcs7 bool) []byte
func Unzip(data []byte) (src []byte, err error)
func ValidOptionalPort(port string) bool
func WaitNextMinute()
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
func Zip(data []byte) (dst []byte, err error)
```

### 加解密小工具

见: [envtools](envtools)

### 获取内外网 IP 小工具

见: [myip](myip)

或: https://github.com/fufuok/myip

```go
package myip // import "github.com/fufuok/utils/myip"

func ExternalIP(v ...string) string
func ExternalIPAny(retries ...int) string
func ExternalIPv4() string
func ExternalIPv6() string
func InternalIP(dstAddr, network string) string
func InternalIPv4() string
func InternalIPv6() string
func LocalIP() string
func LocalIPv4s() (ips []string)
```

### 编码解码 base58

见: [base58](base58)

或: https://github.com/fufuok/basex

```go
package base58 // import "github.com/fufuok/utils/base58"

func CheckDecode(input string) (result []byte, version byte, err error)
func CheckEncode(input []byte, version byte) string
func Decode(b string) []byte
func Encode(b []byte) string
```

### 可排序全局唯一 ID 生成器

比 UUID 更快, 更短

见: [xid](xid)

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

### 自守护进程和后台运行

见: [xdaemon](xdaemon)

或: https://github.com/fufuok/xdaemon

```go
package xdaemon // import "github.com/fufuok/utils/xdaemon"

const EnvName = "XW_DAEMON_IDX"
func Background(logFile string, isExit bool) (*exec.Cmd, error)
func NewSysProcAttr() *syscall.SysProcAttr
type Daemon struct{ ... }
    func NewDaemon(logFile string) *Daemon
```

### 高性能并发安全同步扩展库

见: [xsync](xsync)

或: https://github.com/fufuok/xsync

```go
package xsync // import "github.com/fufuok/utils/xsync"

type Counter struct{ ... }
    func (c *Counter) Add(delta int64)
    func (c *Counter) Dec()
    func (c *Counter) Inc()
    func (c *Counter) Reset()
    func (c *Counter) Value() int64
type MPMCQueue struct{ ... }
    func NewMPMCQueue(capacity int) *MPMCQueue
    func (q *MPMCQueue) Dequeue() interface{}
    func (q *MPMCQueue) Enqueue(item interface{})
    func (q *MPMCQueue) TryDequeue() (item interface{}, ok bool)
    func (q *MPMCQueue) TryEnqueue(item interface{}) bool
type Map struct{ ... }
    func NewMap() *Map
    func (m *Map) Delete(key string)
    func (m *Map) Load(key string) (value interface{}, ok bool)
    func (m *Map) LoadAndDelete(key string) (value interface{}, loaded bool)
    func (m *Map) LoadOrStore(key string, value interface{}) (actual interface{}, loaded bool)
    func (m *Map) Range(f func(key string, value interface{}) bool)
    func (m *Map) Store(key string, value interface{})
type RBMutex struct{ ... }
    func (m *RBMutex) Lock()
    func (m *RBMutex) RLock() *RToken
    func (m *RBMutex) RUnlock(t *RToken)
    func (m *RBMutex) Unlock()
type RToken struct{ ... }
```

### 常用的池

`[]byte` 字节切片池化见: [github.com/fufuok/bytespool](https://github.com/fufuok/bytespool)

见: [pools](pools)

```go
package bufferpool // import "github.com/fufuok/utils/pools/bufferpool"

func Get() *bytes.Buffer
func New(bs []byte) *bytes.Buffer
func NewString(s string) *bytes.Buffer
func Put(buf *bytes.Buffer)
func Release(buf *bytes.Buffer) bool

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

## 使用

```go
var s string
s = utils.GetString(123.45)         // "123.45"
s = utils.GetString(nil)            // ""
s = utils.GetString(nil, "default") // "default"
s = utils.GetString([]byte("ff"))   // "ff"
s = utils.GetString(true)           // "true"

x := utils.AddString(s, "OK") // "trueOK"

b := []byte("trueOK")
s = utils.B2S(b[0:1])                                     // "t"
safeS1 := utils.B2S([]byte(s[0:1]))                       // 转换为不可变字符串, CopyString 的实现
safeS2 := utils.CopyString(s[0:1])                        // 不可变字符串, s 可以被 GC 回收
safeS3 := utils.GetSafeString(s[0:1], "optional default") // 不可变字符串
safeS4 := utils.GetSafeB2S(b[0:1], "optional default")    // 转换为不可变字符串
safeS5 := string(b[0:1])                                  // 标准转换
b[0] = 70                                                 // 注意: 底层数组变化会引起字符串 s 发生改变
fmt.Println(s, safeS1, safeS2, safeS3, safeS4, safeS5)    // F t t t t t

x = utils.Encrypt("myData", "myKey")
fmt.Println(x) // Csi64LeQLmVhuZTh1xkCKM
x = utils.Decrypt("Csi64LeQLmVhuZTh1xkCKM", "myKey")
fmt.Println(x) // myData

k := utils.RandBytes(16)
x = utils.AesCBCEnPKCS7StringHex("myData", k)
fmt.Println(x) // 9dce01c049de7493ce2fae6a2707fad1
x = utils.AesCBCEnStringB58("myData", k)
fmt.Println(x) // 6NHTJhaW5mfeUioFLrbpRX

k = []byte("1234567812345678")
x = utils.GCMEnStringHex("myData", k)
fmt.Println(x) // 6501bb4737c772b9a2956d87183d4793d44b0e3c233bf2c6435d70502913d5701a98
x = utils.GCMEnStringB58("myData", k)
fmt.Println(x) // 5H9ftJvyBesQS7yNGoW7sBCqpwedpEhkjB7k4Z8QBi8GYJo

x = utils.GCMEnStringB64("myData", k)
fmt.Println(utils.GCMDeStringB64(x, k)) // myData

x = utils.UUIDString()
fmt.Println(x) // 04a49f17-8c37-44f7-a9c5-ea291c3736d7
x = utils.UUIDSimple()
fmt.Println(x) // 16123e98b35a4cea8e9cc127f379ff52
x = utils.UUIDShort()
fmt.Println(x) // Mw4hP7t9bnMMczU2AvyorU
x = xid.NewString()
fmt.Println(x) // c294bsnn5ek0ub0200fg

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
fmt.Println(utils.Get0LastWeek(now).Format(time.RFC3339Nano)) // 2020-02-10T00:00:00Z

fmt.Println(utils.GetNotInternalIPv4String("100.125.1.1", "", true))  // 100.125.1.1
fmt.Println(utils.GetNotInternalIPv4String("100.125.1.1", "1.2.3.4")) // 1.2.3.4
fmt.Println(utils.GetNotInternalIPv4String("192.168.1.1", "1.2.3.4")) // 1.2.3.4
fmt.Println(utils.GetNotInternalIPv4String("119.118.7.6", "1.2.3.4")) // 119.118.7.6

var nilN struct{}
var nilY *struct{}
fmt.Println(utils.IsNil(nilN), utils.IsNil(nilY)) // false true

public, private := utils.GenRSAKey(1024)
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

dec, _ := utils.Zip(utils.FastRandBytes(3000))
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
fmt.Printf("struct: %+v\n", t2)      // struct: {Name:ff}
fmt.Println("empty:", buf.Len() == 0) // empty: true
bufferpool.Put(buf)
```







*ff*