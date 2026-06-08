package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "he.llo"

	if strings.IndexByte(s, 'e') != -1 {
		fmt.Println("存在")
	}

	fmt.Printf("strings.LastIndex(s, \"lo\"): %v\n", strings.LastIndex(s, "lo"))
	if strings.HasPrefix(s, "he") {
		fmt.Println("有前缀")
	}

	if strings.HasSuffix(s, "ll") {
		fmt.Println("有后缀")
	}

	sp := strings.Split(s, ".")
	fmt.Printf("sp: %v\n", sp)

	num := "0123"
	val, _ := strconv.Atoi(num)
	fmt.Printf("val: %v\n", val)

	fmt.Printf("strings.Count(num, \"1\"): %v\n", strings.Count(num, "1"))
}
