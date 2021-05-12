# Go-Utils

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
  - `go build -tags=gojson.` 使用标准 JSON 库 `encoding/json`
  - `go build -tags=go_json .` 使用 `goccy/go-json` (暂不成熟, 观望中)

## v0.0.2

**2021-04-16**

- 改用 `json-iterator`

## v0.0.1

**2021-04-14**

- init
