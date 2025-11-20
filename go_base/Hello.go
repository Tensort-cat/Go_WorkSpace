package main

import "fmt"

func main() {
	// 我是单行注释
	/* 	我是
		多行
	 	注释
	*/
	fmt.Println("Hello World!")
	fmt.Println("Hello Again!")

	// 字符串拼接
	fmt.Println("Hello" + " " + "World!")

	a := 10
	if a > 2 {
		fmt.Println("a is greater than 2")
	}

	var code = 123
	var date = "2021-01-01"
	var url = "Code = %d & endDate = %s"
	var target = fmt.Sprintf(url, code, date)
	fmt.Println(target) // 输出结果：Code = 123 & endDate = 2021-01-01
}
