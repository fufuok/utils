# Go-Utils

常用的辅助函数

若有直接引用的, 会在函数定义时注明来源, 感谢之!

## 安装

```shell
go get github.com/fufuok/utils
```

## 使用

```go
var s string
s = utils.GetString(nil)  // ""
s = utils.GetString(nil, "default")  // "default"
s = utils.GetString([]byte("ff"))  // "ff"
s = utils.GetString(true)  // "true"

x := utils.AddString(s, "OK")  // "trueOK"

b := utils.S2B(x)  // []byte("trueOK")
s = utils.B2S(b)  // "trueOK"
f := utils.CopyString(s)  // 不可变字符串
u := string(b)  // 标准转换
b[0] = 'F'  // 注意: 底层数组变化会引起 s 发生改变
fmt.Println(string(b), s, f, u)  // "FrueOK" "FrueOK" "trueOK" "trueOK"

_ = utils.AesCBCEnPKCS7StringHex("myData", "myKey")
```

...

## 加解密小工具

见: `envtools`







*ff*