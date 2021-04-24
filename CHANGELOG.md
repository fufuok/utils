# Go-Utils

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
