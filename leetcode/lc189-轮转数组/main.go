package main

import "slices"

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	point := l - k
	// 反转前 l - k 个
	slices.Reverse(nums[:point])
	// 反转后 k 个
	slices.Reverse(nums[point:])
	// 最后整体反转
	slices.Reverse(nums)
}

func main() {

}
