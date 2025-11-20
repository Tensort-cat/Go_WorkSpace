// var arrayName [size]dataType
package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums) // [0 0 0 0 0]

	words := [3]string{"apple", "banana", "orange"}
	fmt.Println(words) // [apple banana orange]

	// 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度
	balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance, len(balance)) // [1000 2 3.4 7 50]

	// 将索引为 1 和 3 的元素初始化
	balance2 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance2) // [0 2 0 7 0]
}
