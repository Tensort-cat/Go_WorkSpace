package main

func longestMountain(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	left, right := 0, 1
	res := 0

	for right < len(arr) {
		asc := false
		desc := false
		for right < len(arr) && arr[right] > arr[right-1] {
			asc = true
			right++
		}

		for right < len(arr) && arr[right] < arr[right-1] {
			desc = true
			right++
		}
		if asc && desc {
			res = max(res, right-left)
		}

		// 去重
		for right < len(arr) && arr[right] == arr[right-1] {
			right++
		}
		left = right - 1
	}

	return res
}
