// 请实现一个函数 makeCounter，它返回一个闭包，每次调用闭包时，计数器加 1 并返回当前值
package main

import "fmt"

// TODO: 实现 makeCounter
func makeCounter() func() int {
	// 在这里写代码
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()
	fmt.Println(counter()) // 输出 1
	fmt.Println(counter()) // 输出 2
	fmt.Println(counter()) // 输出 3
}
