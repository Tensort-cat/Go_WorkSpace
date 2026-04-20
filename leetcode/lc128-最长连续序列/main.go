package main

import "fmt"

func longestConsecutive(nums []int) int {
	res := 0
	l := len(nums)
	if l == 0 {
		return res
	}
	m := make(map[int]struct{}, 0)
	for _, num := range nums {
		m[num] = struct{}{}
	}

	var exists bool
	for x, _ := range m {
		if _, exists := m[x-1]; exists {
			continue
		}

		// x 是序列的起点
		y := x + 1
		_, exists = m[y]
		for exists {
			y++
			_, exists = m[y]
		}
		res = max(res, y-x)
	}

	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
