// 基本语法：value, ok := interfaceValue.(Type)
package main

import "fmt"

func main() {
	var i interface{} = "Hello, Go!"
	// var i any = "Hello, Go!" // 也可 以使用 any 类型，等价

	// 尝试将 i 断言为 string 类型
	s, ok := i.(string) // 如果不用ok接收，断言失败时会直接 panic
	if ok {
		fmt.Println("断言成功:", s)
	} else {
		fmt.Println("断言失败")
	}

	// 尝试将 i 断言为 int 类型
	n, ok := i.(int)
	if ok {
		fmt.Println("断言成功:", n)
	} else {
		fmt.Println("断言失败")
	}
}

/*
注：
类型断言只能用于接口类型：类型断言只能用于接口类型的变量，不能用于非接口类型的变量。
避免 panic：在使用不返回布尔值的类型断言时，务必确保类型断言不会失败，否则会引发 panic。
类型断言的性能：类型断言在运行时进行类型检查，因此可能会带来一定的性能开销。在性能敏感的场景中，应谨慎使用。
*/
