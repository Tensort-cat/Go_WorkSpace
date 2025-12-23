/*
UDP 地址支持 udp4，udp6，签名如下
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
*/

package main

import (
	"fmt"
	"net"
)

func main() {
	udp4Addr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:2020")
	if err != nil {
		panic(err)
	}
	fmt.Println(udp4Addr)
	udp6Addr, err := net.ResolveUDPAddr("udp6", "[::1]:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(udp6Addr)
}
