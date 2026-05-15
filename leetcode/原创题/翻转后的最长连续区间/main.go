package main

import "fmt"

func ans(nums []int, k int) int {
	res := 0
	l := len(nums)
	if l == 0 {
		return res
	}

	left, right := 0, 0
	for right < l {
		zeroCnt := 0
		// 统计区间中的0是否小于等于k
		for i := left; i <= right; i++ {
			if nums[i] == 0 {
				zeroCnt++
			}
		}
		if zeroCnt <= k {
			res = max(res, right-left+1)
		} else {
			// 将left向右移到第一个0的后面
			for nums[left] != 0 {
				left++
			}
			left++
		}
		right++
	}

	return res
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	nums := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(nums)
	fmt.Println(ans(nums, k))
}
