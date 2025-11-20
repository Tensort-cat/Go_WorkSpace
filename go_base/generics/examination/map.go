// 练习 2：实现泛型映射函数
// 编写一个 Map 函数，将切片中的每个元素转换为另一种类型
package main

import "fmt"

// 你的实现代码在这里
func Map[T any, U any](slice []T, mapper func(T) U) []U {
	// 实现映射逻辑
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = mapper(v)
	}
	return res
}

// 测试代码
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	strings := Map(numbers, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})
	fmt.Println(strings)
}
