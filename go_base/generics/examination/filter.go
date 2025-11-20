// 练习1：编写一个 Filter 函数，根据条件过滤切片元素
package main

import "fmt"

// 你的实现代码在这里
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// 实现过滤逻辑
	var res []T
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// 测试代码
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	even := Filter(numbers, func(n int) bool { // 得到所有偶数
		return n%2 == 0
	})
	fmt.Println(even) // 应该输出: [2 4 6]
}
