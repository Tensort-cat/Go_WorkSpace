package main

func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}

	// 找到数据应该所在的行
	var nums []int
	for _, row := range matrix {
		if target >= row[0] && target <= row[len(row)-1] {
			nums = row
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
