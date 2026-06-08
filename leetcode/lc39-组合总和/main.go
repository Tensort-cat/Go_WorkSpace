package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sum := 0
	path := []int{}

	var dfs func(int)
	dfs = func(start int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		} else if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	dfs(0)
	return res
}
