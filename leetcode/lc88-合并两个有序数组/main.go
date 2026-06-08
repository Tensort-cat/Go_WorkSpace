package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		tmp[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		tmp[k] = nums2[j]
		j++
		k++
	}

	for i := range m + n {
		nums1[i] = tmp[i]
	}
}
