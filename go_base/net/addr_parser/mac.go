package main

// func ParseMAC(s string) (hw HardwareAddr, err error)

import (
	"fmt"
	"net"
)

func main() {
	hw, err := net.ParseMAC("00:1A:2B:3C:4D:5E")
	if err != nil {
		panic(err)
	}
	fmt.Println(hw)
	fmt.Printf("%T", hw)
}
