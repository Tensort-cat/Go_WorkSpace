package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}
	stack := make([]int, 100)
	for x > 0 {
		e := x % 10
		x /= 10
		stack = append(stack, e)
	}

	var end int
	for i, v := range stack {
		if v != 0 {
			end = i
			break
		}
	}

	t := 1
	for i := len(stack) - 1; i >= end; i-- {
		res += stack[i] * t
		t *= 10
	}

	if flag {
		res = -res
	}
	if res < -int(math.Pow(2, 31)) || res >= int(math.Pow(2, 31)) {
		return 0
	}

	return res
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(reverse(x))
}
