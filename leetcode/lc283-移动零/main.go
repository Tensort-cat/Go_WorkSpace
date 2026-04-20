package main

import "fmt"

// https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < l; j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}

func main() {
	var elem int
	var nums []int
	for {
		fmt.Scan(&elem)
		if elem == -999999 {
			break
		}
		nums = append(nums, elem)
	}
	moveZeroes(nums)
	fmt.Println(nums)
}
