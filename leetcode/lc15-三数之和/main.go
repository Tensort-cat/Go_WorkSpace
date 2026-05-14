package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := make([][]int, 0)
	// 第一个元素大于0就不可能有三元组和为0
	if nums[0] > 0 {
		return res
	}
	for i := 0; i < l; i++ {
		// 对 a 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, l-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 对 b 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 对 c 去重
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
