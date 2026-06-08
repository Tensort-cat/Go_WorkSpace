package main

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1

	left := 0
	cnt := 0
	for right := 0; right < len(nums); right++ {
		cnt += nums[right]
		if cnt >= target {
			res = min(res, right-left+1)
			for left <= right && cnt >= target {
				cnt -= nums[left]
				left++
				if cnt >= target {
					res = min(res, right-left+1)
				}
			}
		}

	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}
