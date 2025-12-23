package main

// func ParseCIDR(s string) (IP, *IPNet, error)

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr) // ip地址
	fmt.Println(ipv4Net)  // 网络号
}
