# Go-Utils

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
