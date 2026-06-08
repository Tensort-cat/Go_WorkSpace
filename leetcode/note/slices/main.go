package main

import (
	"fmt"
	"slices"
)

func main() {
	// 切片反转
	s1 := []int{1, 2, 3, 4, 5}
	slices.Reverse(s1)
	fmt.Println(s1)
	index := slices.Index(s1, 2)
	fmt.Println(index)
}
