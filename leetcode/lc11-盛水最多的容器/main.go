package main

import "fmt"

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
		} else {
			res = max(res, height[j]*(j-i))
			j--
		}
	}

	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
