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
