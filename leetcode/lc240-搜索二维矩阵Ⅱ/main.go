package main

func searchMatrix(matrix [][]int, target int) bool {
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[0]))
	}
	res := false
	var dfs func(int, int)
	dfs = func(x, y int) {
		if matrix[x][y] == target {
			res = true
			return
		}

		if matrix[x][y] < target {
			// 往右走
			right := y + 1
			down := x + 1
			if right < len(matrix[0]) && !visited[x][right] {
				visited[x][right] = true
				dfs(x, right)
			}
			// 往下走
			if down < len(matrix) && !visited[down][y] {
				visited[down][y] = true
				dfs(down, y)
			}
		} else {
			return
		}
	}

	visited[0][0] = true
	dfs(0, 0)
	return res
}
