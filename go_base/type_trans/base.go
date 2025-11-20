// Go 语言类型转换基本格式如下：
// type_name(expression)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int 类型转换为 float64 类型
	var i int = 10
	f := float64(i)
	fmt.Println(f)

	// 字符串转换为 int 类型
	var s string = "123"
	num, err := strconv.Atoi(s) // strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	// int 类型转换为字符串类型
	var n int = 123
	str := strconv.Itoa(n)
	fmt.Printf("str = %s, 类型为%T\n", str, str)

	// 字符串转换为浮点型类型
	var str1 string = "3.1415926"
	f1, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("f1 = %f, 类型为%T\n", f1, f1)

	// 浮点数转字符串
	var f2 float64 = 3.1415926
	str2 := strconv.FormatFloat(f2, 'f', 2, 64) // 第一个参数为浮点数，第二个参数为格式化类型，第三个参数为保留小数位数，第四个参数为float位数
	fmt.Printf("str2 = %s, 类型为%T\n", str2, str2)
}
