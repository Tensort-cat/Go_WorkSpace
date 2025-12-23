package main

import (
	"fmt"
	"net"
)

func main() {
	// 解析域名的 IP 地址
	addrs, err := net.LookupHost("github.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(addrs)

	// 查询记录值
	mxs, err := net.LookupMX("github.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(mxs) // [0xc0000081c8 0xc000008198 0xc000008180 0xc000008168 0xc0000081b0]
}
