# 获取本机内网 / 公网 IP

## 使用

```go
package main

import (
	"fmt"

	"github.com/fufuok/utils/myip"
)

var (
	// 服务器 IP
	InternalIPv4  string
	ExternalIPv4  string
	ExternalIPAny string
)

func init() {
	// 推荐方式
	go func() {
		InternalIPv4 = myip.InternalIPv4()
	}()
	go func() {
		ExternalIPv4 = myip.ExternalIPv4()
	}()
	go func() {
		ExternalIPAny = myip.ExternalIPAny(5)
	}()
}

func main() {
	fmt.Println("MyIP(可能为空, 但不阻塞)", InternalIPv4, ExternalIPv4, ExternalIPAny)

	fmt.Println("获取外网地址 (IPv4):", myip.ExternalIPv4())
	fmt.Println("获取外网地址 (IPv6):", myip.ExternalIPv6())
	fmt.Println("获取外网地址 (出口公网地址, 优先获取 IPv6):", myip.ExternalIP())
	fmt.Println("获取外网地址 (出口公网地址 IPv4):", myip.ExternalIP("ipv4"))
	fmt.Println("获取外网地址 (出口公网地址 IPv6):", myip.ExternalIP("ipv6"))

	fmt.Println("获取内网地址 (IPv4):", myip.InternalIPv4())
	fmt.Println("获取内网地址 (临时 IPv6 地址):", myip.InternalIPv6())
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("", ""))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("1.1.1.1:53", "udp"))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("baidu.com:443", "tcp"))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("1.1.1.1", "ip4:icmp"))

	fmt.Println("获取本地地址 (第一个):", myip.LocalIP())
	fmt.Println("获取所有本地地址 (IPv4):", myip.LocalIPv4s())

	fmt.Println("MyIP:", InternalIPv4, ExternalIPv4, ExternalIPAny)
}
```





*ff*