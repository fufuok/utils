# Go-Utils

常用的助手函数

若有直接引用的, 会在函数定义或 README 中注明来源, 保留 LICENSE, 感谢之!

## 安装

```shell
go get github.com/fufuok/utils
```

## 使用

```go
var s string
s = utils.GetString(nil)            // ""
s = utils.GetString(nil, "default") // "default"
s = utils.GetString([]byte("ff"))   // "ff"
s = utils.GetString(true)           // "true"

x := utils.AddString(s, "OK") // "trueOK"

b := utils.S2B(x)               // []byte("trueOK")
s = utils.B2S(b)                // "trueOK"
f := utils.CopyString(s)        // 不可变字符串
u := string(b)                  // 标准转换
b[0] = 'F'                      // 注意: 底层数组变化会引起 s 发生改变
fmt.Println(string(b), s, f, u) // "FrueOK" "FrueOK" "trueOK" "trueOK"

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
```

...

## 加解密小工具

见: `envtools`

## 获取内外网 IP 小工具

见: `myip`

或: https://github.com/fufuok/myip

## 编码解码 base58

见: `base58`

或: https://github.com/fufuok/basex

## JSON

`json` 使用 `gin` 类似的可选组织方式:

- `go build .` 默认使用 `json-iterator/go`
- `go build -tags=gojson.` 使用标准 JSON 库 `encoding/json`
- `go build -tags=go_json .` 使用 `goccy/go-json` (暂不成熟, 观望中)







*ff*