package main

import (
	"fmt"
	"math/rand"
)

func partition(nums []int, left, right int) int {
	// 随机选择基准元素
	randomIndex := left + rand.Intn(right-left+1)
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

	pivot := nums[left]

	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pos := partition(nums, low, high)
		quickSort(nums, low, pos-1)
		quickSort(nums, pos+1, high)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}
	fmt.Printf("nums: %v\n", nums)
	quickSort(nums, 0, n-1)
	fmt.Printf("nums after quick sort: %v\n", nums)
}
