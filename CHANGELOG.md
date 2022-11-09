# Go-Utils

## v0.9.0

**2022-11-09**

- 增加带超时时间的互斥锁: `utils.NewTryMutex()`
- 调整 `Hash` 函数到子目录: `xhash`
- 优化测试助手函数
- 同步 `xsync v2.4.0`, 增加 Map 初始容量参数

## v0.8.3

**2022-11-03**

- 同步优化 `xsync.Map`

## v0.8.2

**2022-10-31**

- 同步优化 `xsync.Map` `xsync.Counter`
  - 重大变化, 现在需要: `xsync.NewCounter()` 初始化后使用

## v0.8.1

**2022-10-28**

- 增加 `FastRand64` `FastRandu`
- 同步 `xsync` `xid` 修改

## v0.8.0

**2022-10-27**

- 同步 `xsync v2.0.1`, 增加通用助手函数:
  - `func NewHashMapOf[K comparable, V any](hasher ...func(maphash.Seed, K) uint64) HashMapOf[K, V]`
- 增加 `AssertNotEqual` `AssertNotEqualf` 用于测试
- 增加 `GenSeedHasher64` `GenHasher64` `GenHasher`, 按数据类型生成 Hash 函数, 基于 xxHash
- 同步 `generic` 优化, 增加 `sclices.Replace`

## v0.7.18

**2022-10-23**

- 同步 `xsync` 修改, 增加任意可比较类型为键的 Map:
  - `func NewMapOf[V any]() *MapOf[string, V]`
  - `func NewIntegerMapOf[K IntegerConstraint, V any]() *MapOf[K, V]`
  - `func NewTypedMapOf[K comparable, V any](hasher func(K) uint64) *MapOf[K, V]`

## v0.7.17

**2022-10-17**

- 新增 `ParseIP` `ParseHostPort` 解析 IP 和端口助手函数

## v0.7.16

**2022-10-10**

- 新增 `pools/bytespool`, 精简版 [bytespool](https://github.com/fufuok/bytespool)
- 新增 `SafeGo` `Recover` `WaitSignal` 安全 goroutine 助手, IPv4 小端转换函数
- 增加 `generic/bimap` `generic/set` 包
- 同步 `xid` 包修改: 修正 nilID, 优化检查逻辑
- 同步 `generic/slices` 修改

## v0.7.15

**2022-10-08**

- 优化 `zlib` `gzip`

## v0.7.14

**2022-08-25**

- 代码整理

## v0.7.13

**2022-08-13**

- 同步 `xsync` 包修改: 修复 MapOf 结构值指针唯一性

## v0.7.12

**2022-08-08**

- 增加 `base62`, 来自 `jxskiss/base62`

## v0.7.11

**2022-08-03**

- 同步 `xsync` 包修改: 修复 32 位编译

## v0.7.10

**2022-07-28**

- 优化 `WaitNextMinute` 可选传入时间参数
- 同步官方扩展库 `sync` `exp`
- 优化 `Map.LoadAndStore` 为原子操作

## v0.7.9

**2022-07-19**

- 新增 NTP 简单时间同步包

## v0.7.8

**2022-07-17**

- 同步 `generic` 库更新, 为 `hashset` `mapset` 增加 `Of` 方法
- 更新 `DOC.md` 文档

## v0.7.7

**2022-07-08**

- 更新 `xsync`, 暴露 `m.Size`

## v0.7.6

**2022-07-05**

- !调整压缩助手函数: `Zip` `ZipLevel` `Unzip` `Gzip` `GzipLevel` `Ungzip`

## v0.7.5

**2022-06-28**

- 增加 `Map.LoadAndStore` 返回旧值, 存入新值 (旧值不存在时, `ok` 为 `false`, 返回的值是新值)

## v0.7.2

**2022-06-22**

- 增加 `IsIP` `IsIPv4` `IsIPv6` `ParseIPv4` `ParseIPv6` `myip.InterfaceAddrs`

## v0.7.1

**2022-06-16**

- 补全 `tickerpool`

## v0.7.0

**2022-06-12**

- 增加 `generic` 包, 包含常用的泛型实现, 大多来自: `zyedidia/generic`, 感谢!
- 移动 `deepcopy` `maps` `slices` `constraints` 到 `generic`
- 优化文本 Hash 算法

## v0.6.1

**2022-06-10**

- 调整 `sched` 并发调度实现机制

## v0.6.0

**2022-05-27**

- 增加 `xjson` 包, 包含常用到的 JSON 字符串操作库
  - `github.com/fufuok/utils/xjson/jsongen` JSON 字符串生成
  - `github.com/fufuok/utils/xjson/gjson` JSON 字符串搜索
  - `github.com/fufuok/utils/xjson/sjson` JSON 字符串字段修改
  - `github.com/fufuok/utils/xjson/pretty` JSON 字符串校验和格式化
  - `github.com/fufuok/utils/xjson/match` 字符串通配符匹配

## v0.5.3

**2022-05-20**

- 修正 `jsongen` 多行文本转义问题

## v0.5.2

**2022-05-15**

- 修正 `EqualFold`
- 增加 `AssertEqualf`
- 增加 2 个特殊的类型: `NoCopy` `NoCmp`
- 优化 `jsongen`: 增加插入原始 JSON 的方法 `RawString` `RawBytes`
- `MustString` 增加更多类型转换: `time.Time` `reflect.Value` `fmt.Stringer`

## v0.5.1

**2022-05-11**

- 增加原子操作的安全布尔值

## v0.5.0

**2022-05-09**

- 增加 `sched` `deepcopy` 来自 `poly.red`, 感谢
- 增加运行时自增 ID
- 增加 `CutString` `CutBytes`, 同 go1.18
- 增加 `golang.org/x/exp` 的 `maps` `slices` `constraints` 到本地
- 加解密函数迁入 `xcrypto` 包, Hash 签名函数不迁移

## v0.4.6

**2022-05-07**

- 增加 `golang.org/x/sync` 到本地

## v0.4.5

**2022-04-26**

- 修正获取上/下月时间的计算误差
- 增加新的 `myip` 接口地址

## v0.4.4

**2022-04-09**

- 增加 `MustMD5Sum`

## v0.4.3

**2022-04-07**

- 增加 `MD5BytesHex`
- 升级 `xsync` 泛型版本

## v0.4.2

**2022-03-16**

- 增加 `HumanGBMB`

## v0.4.1

**2022-03-11**

- 同步 `xid` 更新

## v0.4.0

**2022-02-20**

- 重命名 `time` 相关助手函数, 并增加了一些

## v0.3.15

**2022-02-16**

- 增加 `Pad` `PadBytes` 用指定字符串填充原字符串到指定长度
- 修正 `Trim` 裁剪全部内容的问题

## v0.3.14

**2022-01-26**

- 增加 `jsongen` 包, 高效的 JSON 字符串生成器

## v0.3.13

**2022-01-21**

- 增加 `MustParseHumanBytes` `MustParseHumanBigBytes`

## v0.3.12

**2022-01-18**

- 增加 `bufferpool.SetMaxSize()` 以设置回收到池的最大字节容量值

## v0.3.11

**2022-01-14**

- 增加字符串哈希助手函数, 生成数字字符串, 比 Md5 快(会有重复, 注意使用场景, 比如简单的 Token 校验)
  - `HashString` `HashStringUint64` `HashBytes` `HashBytesUint64`

## v0.3.10

**2022-01-06**

- 增加 `bufferpool.NewByte()` `bufferpool.NewRune()`

## v0.3.9

**2022-01-05**

- 增加 bytes.Buffer Pool

## v0.3.8

**2021-11-09**

- 更新 `myip` 接口地址

## v0.3.6

**2021-10-30**

- 增加了几个简单的池: `readerpool` `timerpool` `tickerpool`

## v0.3.4

**2021-10-23**

- 增加 gzip 助手函数: `Zip` `Unzip`

## v0.3.3

**2021-10-20**

- 优化随机数, 并发安全且性能好. 参考新包: [github.com/fufuok/random](https://github.com/fufuok/random)

## v0.3.2

**2021-09-30**

- 增加高性能的 `Trim` 类函数和大小写转换函数
- 增加 `ReplaceHost`, 替换 URL 主机名
- 增加 `sync` 扩展包, 提供: `Counter` `Map` `MPMCQueue` `RBMutex`, 来自: `puzpuzpuz/xsync`

## v0.3.0

**2021-09-28**

- 清理代码, 移除负载均衡类函数, 新包: [github.com/fufuok/balancer: Goroutine-safe, High-performance general load balancing algorithm library.](https://github.com/fufuok/balancer)

## v0.2.2

**2021-08-27**

- 补充 `InIPNetString` 测试用例

## v0.2.1

**2021-08-26**

- 补齐 `AES-GCM` 加解密助手函数
- 增加 `IsPrivateIP` `IsPrivateIPString`, 判断是否为私有 IP (RFC 1918, RFC 4193)

## v0.2.0

**2021-08-12**

- 移出 `json`, `timewheel` 包, 消除依赖, 以下包代替:
  - [github.com/fufuok/internal/json](https://github.com/fufuok/internal)
  - [github.com/fufuok/timewheel](https://github.com/fufuok/timewheel)
- `xid` 修正 JSON 解码无效 ID 崩溃问题

## v0.1.16

**2021-08-02**

- 增加从 `net.Addr` 获取 IP 和 端口的方法 `GetIPPort`
- 增加后台运行和守护包 `xdaemon`

## v0.1.15

**2021-06-30**

- 修正 `MustJSON` 传参方式

## v0.1.14

**2021-06-26**

- 增加 IP 包含关系方法 `InIPNet` `InIPNetString`
- 升级 `goccy/go-json@v0.7.2`

## v0.1.12

**2021-06-17**

- 修正时间轮 Ticker 使用小于时间刻度时, 只会执行一次

## v0.1.11

**2021-06-12**

- 增加 `HumanBigBytes` `HumanBigKbps` `HumanIntKbps` `HumanKbps`

## v0.1.10

**2021-06-06**

- 调整对称加密密钥参数始终为 `[]byte`
- 增加 `XOR` `RSA`
- 规范注释

## v0.1.9

**2021-06-04**

- 增加时间轮

## v0.1.8

**2021-05-30**

- 变更 `-tags=gojson` 为 `-tags=go_json`, 与 `gin` 一致
- 增加 `IsNil` 判断对象是否为 nil

## v0.1.7

**2021-05-21**

- 增加 `MustJSONIndent` `MustJSONIndentString`
- 启用可选 JSON 库: `goccy/go-json@v0.5.1`, 编译参数: `go build -tags=go_json .`
- 升级 `json-iterator/go@v1.1.11`, 默认使用: `ConfigCompatibleWithStandardLibrary`

## v0.1.6

**2021-05-20**

- 增加 `myip.ExternalIPAny` 重试指定次数取得外网地址

## v0.1.5

**2021-05-19**

- 增加时间点助手函数: `Get0Hour` `Get0LastWeek` `GetMonthDays` 等
- 增加内网 IPv4 判断和取值, IP 与数值转换

## v0.1.4

**2021-05-12**

- 增加 `Commai` `HumanIntBytes` `HumanIntIBytes` 方便对 `int` 操作

## v0.1.3

**2021-05-12**

- 增加数字转逗号分隔千分位字符串系列函数, 支持 `int64` `uint64` `big.Int`, 如: `Commaf(234.123)` -> `1,234.123`
- 增加数字转容量单位字符串系列函数, 如: `HumanBytes(234567890)` -> `1.2 GB`
- 增加 `myip.LocalIPv4s()` 取本机网卡所有 IPv4 地址

## v0.1.2

**2021-04-27**

- 整合 `xid.NewString` 可排序全局ID生成器, 增加助手函数
- 增加 `Executable` `ExecutableDir` 运行时函数

## v0.1.1

**2021-04-25**

- 增加 `GetSafeB2S` `GetSafeS2B` `GetSafeBytes` `GetSafeString` 等安全函数, 用于基于 `fasthttp` 的应用
- 优化部分转换函数, 增加更多测试用例

## v0.1.0

**2021-04-23**

- 增加 `AES-GCM` 系列助手函数
- 增加加权随机函数 `WeightedChoice` `WeightedChoiceIndex` `WeightedChoiceWeightsIndex` `WeightedChoiceMap`
- 增加 UUID 相关助手函数
- 调整 `Base64` 助手函数参数类型, 增加 `base58` 加解密
- 整合 `MyIP` 获取服务器内外网 IP
- 加密小工具加密方式改为: `AesCBCEnStringB58(value, MD5Hex(secret))`
- `json` 使用 `gin` 类似的可选组织方式:
    - `go build .` 默认使用 `json-iterator/go`
    - `go build -tags=stdjson.` 使用标准 JSON 库 `encoding/json`
    - `go build -tags=go_json .` 使用 `goccy/go-json`

## v0.0.2

**2021-04-16**

- 改用 `json-iterator`

## v0.0.1

**2021-04-14**

- init
