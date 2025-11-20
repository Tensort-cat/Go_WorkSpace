package main

/* 切片（slice）是对数组（array）的一个 动态视图（动态窗口）。
简单理解就是：切片是一个轻量级的结构体，内部指向底层数组。
它不是独立存储数据的容器，而是对底层数组的一段“引用”。
*/

import "fmt"

func main() {
	// 声明一个未指定大小的数组来定义切片
	// var identifier []type

	// 创建方式1：从数组中切
	s := [5]int{1, 2, 3, 4, 5}
	s2 := s[1:3]
	fmt.Println("s2 =", s2) // [2 3]
	// 改变s2的元素值，会影响到s
	s2[0] = 100
	fmt.Println("改变s1后的s =", s) // [1 100 3 4 5]

	// 创建方式2：使用make函数创建切片
	s3 := make([]int, 3, 5) // 底层数组: [0, 0, 0, _, _]
	fmt.Println("s3 =", s3) // [0 0 0]

	// 创建方式3：使用字面量语法创建切片
	s4 := []int{1, 2, 3}    // 这会隐式创建一个底层数组 [1, 2, 3] 并返回它的切片
	fmt.Println("s4 =", s4) // [1 2 3]

	// 遍历切片
	for index, value := range s4 {
		fmt.Println("s4[", index, "] =", value)
	}

}
