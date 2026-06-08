package main

import "fmt"

func longestConsecutive(nums []int) int {

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	// 遍历 set 去重后的元素
	for num := range set {
		// 不是起点
		if set[num-1] {
			continue
		}

		cur := num
		length := 1

		for set[cur+1] {
			cur++
			length++
		}
		res = max(res, length)
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(longestConsecutive(nums))
}
