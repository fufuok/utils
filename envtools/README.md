# 环境变量加密工具

用于项目中敏感配置项加解密. 比如各类 API secret

1. 项目 git 中不会出现明文信息
2. 运行环境中也不会见到明文信息, 也不能通过环境变量值解密

## 1. 前置

1. 先将这个项目用到的基础密钥名称设置到固定环境变量 `ENV_TOOLS_NAME`
2. 再把项目的基础加密密钥(明文, 也可以自定义密文, 见示例), 设置到环境变量, 名称即为项目中要用到的环境变量名称

```shell
# Linux
# 告知该加密工具, 项目的基础密钥环境变量名称是: XY_PROJECT_BASE_SECRET_KEY
export ENV_TOOLS_NAME=XY_PROJECT_BASE_SECRET_KEY
# 设置基础加密密钥(明文)到项目中用于获取密钥的环境变量名称: XY_PROJECT_BASE_SECRET_KEY
export XY_PROJECT_BASE_SECRET_KEY=myBASEkeyValue123
# Windows
set ENV_TOOLS_NAME=XY_PROJECT_BASE_SECRET_KEY
set XY_PROJECT_BASE_SECRET_KEY=myBASEkeyValue123
```

## 2. 加密

```shell
cd envtools
go build main.go
./main -d="待加密的字符串" -k="环境变量名"
```

```shell
# ./main -d="123.456" -k="XY_REDIS_AUTH"
plaintext:
	123.456
ciphertext:
	FH3Djy1UJiv2y5CrpDQzty
Linux:
	export XY_REDIS_AUTH=FH3Djy1UJiv2y5CrpDQzty
Windows:
	set XY_REDIS_AUTH=FH3Djy1UJiv2y5CrpDQzty


testGetenv: XY_REDIS_AUTH = 123.456
```

得到加密内容到环境中执行即可.

## 3. 解密

```shell
export XY_REDIS_AUTH=FH3Djy1UJiv2y5CrpDQzty
./main -k=XY_REDIS_AUTH
# testGetenv: XY_REDIS_AUTH = 123.456
```

## 4. 应用示例

[example/main.go](example/main.go)

```go
package main

import (
	"fmt"

	"github.com/fufuok/utils/xcrypto"
)

const (
	// BaseSecretKeyName 项目基础密钥 (环境变量名)
	BaseSecretKeyName = "FF_PROJECT_1_BASE_SECRET_KEY"

	// BaseSecretSalt 用于解密基础密钥值的密钥 (编译在程序中)
	BaseSecretSalt = "123"

	// RedisAuthKeyName Redis Auth 短语环境变量 Key
	RedisAuthKeyName = "PROJECT_1_REDIS_AUTH"
)

type tConfig struct {
	BaseSecret string
	RedisAuth  string
}

var Conf tConfig

func init() {
	// !!! 前置: 假如项目环境中已经执行了下面的配置
	// export FF_PROJECT_1_BASE_SECRET_KEY=EnUNZ1FkdnsvWXTukDe4FiwhLkw5eMmjGgAYNqYwB9zn
	// export PROJECT_1_REDIS_AUTH=FH3Djy1UJiv2y5CrpDQzty
	// 1. 项目 git 中不会出现明文信息
	// 2. 运行环境中也不会见到明文信息, 也不能通过环境变量值解密

	// 从环境变量中读取, 用程序中固化的密钥解密, 得到我们的基础密钥是: myBASEkeyValue123
	Conf.BaseSecret = xcrypto.GetenvDecrypt(BaseSecretKeyName, BaseSecretSalt)

	// 用 BaseSecret(或基于此的密钥) 解密其他项目配置
	Conf.RedisAuth = xcrypto.GetenvDecrypt(RedisAuthKeyName, Conf.BaseSecret)
}

func main() {
	// 业务中连接 Redis 就可以用 Conf.RedisAuth
	// Redis Auth: 123.456
	fmt.Println("Redis Auth:", Conf.RedisAuth)
}
```

## 5. 用户名密码编码

有时数据库连接账号密码包含有特殊字符, 可以先将用户名账号编码后, 再加密, 如:

账号密码为: `xy_monitor` `ABC$1^1#1.1>`

先编码:

```shell
# 有特殊字符, 注意用单引号, 对比输出的原始字符串是否与输入相符
./main -u=xy_monitor -p='ABC$1^1#1.1>'

url.UserPassword:
xy_monitor
ABC$1^1#1.1>
xy_monitor:ABC$1%5E1%231.1%3E
```

构造编码后的数据库连接:

`sqlserver://xy_monitor:ABC$1%5E1%231.1%3E@127.0.0.1:1433?database=dev_testdb`

加密该 DSN:

```shell
./main -k=TEST_DB_DSN -d='sqlserver://xy_monitor:ABC$1%5E1%231.1%3E@127.0.0.1:1433?database=dev_testdb'

plaintext:
        sqlserver://xy_monitor:ABC$1%5E1%231.1%3E@127.0.0.1:1433?database=dev_testdb
ciphertext:
        3n3nkZ5hbf3qjccsqUL8YSMtyCY67NcTbZgVodx34mMmpFpZ2G5AvVvUr2MYTeR8hkug4KQnoZHfxXvsXAA7HYaRpFm3vY5x44h7FXT4ghpkG8
Linux:
        export TEST_DB_DSN=3n3nkZ5hbf3qjccsqUL8YSMtyCY67NcTbZgVodx34mMmpFpZ2G5AvVvUr2MYTeR8hkug4KQnoZHfxXvsXAA7HYaRpFm3vY5x44h7FXT4ghpkG8
Windows:
        set TEST_DB_DSN=3n3nkZ5hbf3qjccsqUL8YSMtyCY67NcTbZgVodx34mMmpFpZ2G5AvVvUr2MYTeR8hkug4KQnoZHfxXvsXAA7HYaRpFm3vY5x44h7FXT4ghpkG8


testGetenv: TEST_DB_DSN = sqlserver://xy_monitor:ABC$1%5E1%231.1%3E@127.0.0.1:1433?database=dev_testdb
```





*ff*