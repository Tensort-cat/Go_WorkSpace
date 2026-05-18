package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"

	if strings.IndexByte(s, 'e') != -1 {
		fmt.Println("存在")
	}

	if strings.HasPrefix(s, "he") {
		fmt.Println("有前缀")
	}

	if strings.HasSuffix(s, "ll") {
		fmt.Println("有后缀")
	}
}
