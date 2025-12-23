package main

// IP 地址支持解析 ipv4，ipv6，函数签名如下
// func ResolveIPAddr(network, address string) (*IPAddr, error)

import (
	"fmt"
	"net"
)

func main() {
	ipv4Addr, err := net.ResolveIPAddr("ip4", "192.168.2.1")
	if err != nil {
		panic(err)
	}
	fmt.Println(ipv4Addr)

	ipv6Addr, err := net.ResolveIPAddr("ip6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	if err != nil {
		panic(err)
	}
	fmt.Println(ipv6Addr)
}
