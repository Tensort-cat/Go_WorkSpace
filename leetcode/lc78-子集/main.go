package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(int)
	dfs = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}

	fmt.Printf("subsets(nums): %v\n", subsets(nums))
}
