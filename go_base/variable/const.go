package main

import "fmt"

func main() {

	// 常量：const 关键字声明的变量，在程序运行时，其值不能被修改。
	// const 常量名 类型 = 值
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()

	// 常量可以用作枚举
	const (
		RED   = 1
		GREEN = 2
		BLUE  = 3
	)
	fmt.Printf("颜色为 : %d\n", RED)

	// iota 特殊常量，可以被用来生成一系列的常量，常量的值从 0 开始，每次加 1
	// iota 可理解为 const 语句块中的行索引
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
