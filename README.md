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
```

...





*ff*