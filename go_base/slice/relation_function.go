package main

import "fmt"

func main() {
	// len() 和 cap() 函数
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	// nil切片
	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("nilSlice is nil")
	}

	// append() 和 copy() 函数
	var numbers2 []int
	printSlice(numbers2)

	/* 允许对空切片进行 append() 操作 */
	numbers2 = append(numbers2, 0)
	printSlice(numbers2)

	/* 向切片添加一个元素 */
	numbers2 = append(numbers2, 1)
	printSlice(numbers2)

	/* 向切片添加多个元素 */
	numbers2 = append(numbers2, 2, 3, 4)
	printSlice(numbers2)

	// 创建切片numbers3是numbers2的两倍容量
	numbers3 := make([]int, len(numbers2), cap(numbers2)*2)
	printSlice(numbers3)

	// 拷贝numbers2到numbers3
	copy(numbers3, numbers2) // copy(dst, src) int, 返回拷贝的元素个数
	printSlice(numbers3)

	// nil切片 与 空切片 区别
	nullSlice := []int{}          // 空切片
	fmt.Println(nullSlice == nil) // false
	fmt.Println(nilSlice == nil)  // true

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
