/*
TCP 地址支持 tcp4，tcp6，签名如下
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
*/

package main

import (
	"fmt"
	"net"
)

func main() {
	tcp4Addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:2020")
	if err != nil {
		panic(err)
	}
	fmt.Println(tcp4Addr)
	tcp6Addr, err := net.ResolveTCPAddr("tcp6", "[::1]:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(tcp6Addr)
}
