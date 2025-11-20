package main

import "fmt"

func main() {
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

	if ptr == nil {
		fmt.Println("ptr 是一个空指针")
	} else {
		fmt.Println("ptr 是一个非空指针")
	}
}
