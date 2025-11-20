package main

import "fmt"

// 交换两个值
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// 判断切片是否包含元素
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// 去重函数, 返回去重后的切片
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// 使用示例
func main() {
	// Swap 示例
	a, b := 10, 20
	a, b = Swap(a, b)
	fmt.Printf("a=%d, b=%d\n", a, b) // 输出: a=20, b=10

	// Contains 示例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(numbers, 3)) // 输出: true

	// Unique 示例
	duplicates := []int{1, 2, 2, 3, 4, 4, 5}
	unique := Unique(duplicates)
	fmt.Println(unique) // 输出: [1 2 3 4 5]
}
