package main

import "slices"

func nextPermutation(nums []int) {
	n := len(nums)

	// 第一步：从右向左找到第一个小于右侧相邻数字的数 nums[i]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果找到了，进入第二步；否则跳过第二步，反转整个数组
	if i >= 0 {
		// 第二步：从右向左找到 nums[i] 右边最小的大于 nums[i] 的数 nums[j]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 第三步：反转 nums[i+1:]（如果上面跳过第二步，此时 i = -1）
	slices.Reverse(nums[i+1:])
}
