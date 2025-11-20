package main

import "fmt"

func main() {
	count := 10
	// 类似c语言的for循环
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	fmt.Println("===================================")

	// 类似c语言的while循环
	i := 0
	for i < count {
		fmt.Println(i)
		i++
	}

	// 无限循环
	/* for {
		fmt.Println("infinite loop")
		break
	} */

	// for-each循环
	fruits := []string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	numbers := [6]int{1, 2, 3, 5} // 未赋值的有初始值“零值”
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// for 循环的 range 格式可以省略 key 和 value
	for _, x := range numbers { // 只要 value
		fmt.Printf("x 的值 = %d\n", x)
	}

	for key := range numbers { // 只要 key
		fmt.Printf("key 的值 = %d\n", key)
	}

	// 两个都要
	for key, value := range numbers { // key 和 value
		fmt.Printf("key 的值 = %d, value 的值 = %d\n", key, value)
	}
}
