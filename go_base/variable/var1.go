package main

import "fmt"

func main() {
	// var 变量名 类型 = 初始值
	var a int = 10
	var b string = "Hello World"
	fmt.Println(a, b)

	// 一行多个声明
	// var 变量名1, 变量名2, 变量名3 类型 = 初始值1, 初始值2, 初始值3
	var c, d int = 20, 30
	fmt.Println(c, d)

	// 不声明初始值则默认为零值
	var e int
	fmt.Println(e) // 0
	var f bool
	fmt.Println(f) // false

	// := 简短声明，根据初始值自动推断类型
	g := 40
	h := "Goodbye"
	fmt.Println(g, h)
}
