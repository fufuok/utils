// 环境变量加密工具
// go run main.go -d=Fufu
// go run main.go -d="Fufu  777"
// go run main.go -d=Fufu -k=TestEnv
// go run main.go -k=TestEnv
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fufuok/utils"
)

var (
	// 项目基础密钥 (环境变量名)
	baseSecretKeyName = ""
	// 基础密钥
	baseSecretKey = ""

	// 环境变量名(可选)
	key string
	// 待加解密内容
	value string
)

func init() {
	// 基础密钥环境变量名称
	baseSecretKeyName = utils.GetenvDecrypt("ENV_TOOLS_NAME", "")
	if baseSecretKeyName == "" {
		log.Fatalln("基础密钥的名称不存在\n请设置: export ENV_TOOLS_NAME=你的项目基础密钥环境变量名称")
	}
	// 基础密钥
	baseSecretKey = utils.GetenvDecrypt(baseSecretKeyName, "")
	if baseSecretKey == "" {
		log.Fatalf("基础密钥不能为空\n请设置: export %s=你的项目基础密钥", baseSecretKeyName)
	}
}

func main() {
	// 参数
	flag.StringVar(&key, "k", "envname", "环境变量名")
	flag.StringVar(&value, "d", "", "待加密字符串")
	flag.Parse()

	if value != "" {
		// 加密
		result, err := utils.SetenvEncrypt(key, value, baseSecretKey)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("\nplaintext:\n\t%s\nciphertext:\n\t%s\nLinux:\n\texport %s=%s\nWindows:\n\tset %s=%s\n\n",
			value, result, key, result, key, result)
	}

	// 解密
	result := utils.GetenvDecrypt(key, baseSecretKey)
	fmt.Printf("\ntestGetenv: %s = %s\n\n", key, result)
}
