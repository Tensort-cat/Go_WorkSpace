package main

func findPeakElement(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return idx
}
