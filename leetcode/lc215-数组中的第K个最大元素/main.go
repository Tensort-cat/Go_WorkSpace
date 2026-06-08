package main

import (
	"fmt"
	"math/rand"
)

// 在子数组 [left, right] 中随机选择一个基准元素 pivot
// 根据 pivot 重新排列子数组 [left, right]
// 重新排列后，<= pivot 的元素都在 pivot 的左侧，>= pivot 的元素都在 pivot 的右侧
// 返回 pivot 在重新排列后的 nums 中的下标
// 特别地，如果子数组的所有元素都等于 pivot，我们会返回子数组的中心下标，避免退化
func partition(nums []int, left int, right int) int {
	// 1. 选择一个基准元素 pivot
	i := left + rand.Intn(right-left+1)
	pivot := nums[i]
	// 把 pivot 与子数组第一个元素交换，避免 pivot 干扰后续划分，从而简化实现逻辑
	nums[i], nums[left] = nums[left], nums[i]

	// 2. 相向双指针遍历子数组 [left + 1, right]
	// 循环不变量：在循环过程中，子数组的数据分布始终如下图
	// [ pivot | <=pivot | 尚未遍历 | >=pivot ]
	//   ^                 ^     ^         ^
	//   left              i     j         right

	i, j := left+1, right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		// 此时 nums[i] >= pivot

		for i <= j && nums[j] > pivot {
			j--
		}
		// 此时 nums[j] <= pivot

		if i >= j {
			break
		}

		// 维持循环不变量
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 循环结束后
	// [ pivot | <=pivot | >=pivot ]
	//   ^             ^   ^     ^
	//   left          j   i     right

	// 3. 把 pivot 与 nums[j] 交换，完成划分（partition）
	// 为什么与 j 交换？
	// 如果与 i 交换，可能会出现 i = right + 1 的情况，已经下标越界了，无法交换
	// 另一个原因是如果 nums[i] > pivot，交换会导致一个大于 pivot 的数出现在子数组最左边，不是有效划分
	// 与 j 交换，即使 j = left，交换也不会出错
	nums[left], nums[j] = nums[j], nums[left]

	// 交换后
	// [ <=pivot | pivot | >=pivot ]
	//               ^
	//               j

	// 返回 pivot 的下标
	return j
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	targetIndex := n - k  // 第 k 大元素在升序数组中的下标是 n - k
	left, right := 0, n-1 // 闭区间
	for {
		i := partition(nums, left, right)
		if i == targetIndex {
			return nums[i]
		} else if i < targetIndex {
			left = i + 1
		} else {
			right = i - 1
		}
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(findKthLargest(nums, k))
}
