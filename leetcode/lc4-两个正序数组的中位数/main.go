package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	target := n / 2

	tmp := make([]int, target+1)
	i, j, k := 0, 0, 0

	for i < len(nums1) && j < len(nums2) && k <= target {
		if nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) && k <= target {
		tmp[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) && k <= target {
		tmp[k] = nums2[j]
		j++
		k++
	}

	if n%2 == 1 { // 奇数
		return float64(tmp[target])
	}

	return float64((tmp[target] + tmp[target-1])) / 2
}
