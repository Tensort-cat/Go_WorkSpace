package main

func subarraySum(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		cnt := 0
		for j := i; j < len(nums); j++ {
			cnt += nums[j]
			if cnt == k {
				res++
			}
		}
	}
	return res
}
