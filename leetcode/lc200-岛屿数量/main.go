package main

func numIslands(grid [][]byte) int {
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	res := 0
	var dfs func(int, int)
	dfs = func(x int, y int) {
		for i := 0; i < 4; i++ {
			nextX, nextY := x+dir[i][0], y+dir[i][1]
			if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
				continue
			}
			if !visited[nextX][nextY] && grid[nextX][nextY] == '1' {
				visited[nextX][nextY] = true
				dfs(nextX, nextY)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				visited[i][j] = true
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
